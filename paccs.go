package main

// Copyright (C) 2017  Jan Wollschläger <jmw.tau@gmail.com>
// This file is part of goccs.
//
// goccs is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

import (
  "io/ioutil"
  "math"
  "math/rand"
  "strings"
  "strconv"
)

type ParameterSet map[string]float64

// The parameters to use, i.e. radii of atoms in angstrom.
func Parametersforname(name string) ParameterSet {
    if name == "mobcal" {
      return map[string]float64{
        "H": 2.2, "C": 2.7, "N": 2.7, "O": 2.7,
      }
    }
    if name == "siu_guo_2010" || name == ""{
      return ParameterSet {  "H": 2.01, "C": 2.35, "N": 2.26, "O": 2.26, }
    }
    return nil
}

// a molecule consists of the list of atom labels, the x,y,z coordinates
// of the atoms, so it is essentially just represented as the list of its
// constituting atoms.
type Molecule struct {
    atom_labels []string
    xs []float64
    ys []float64
    zs []float64
}

func minSlice (slice []float64)float64{
     rslt := math.MaxFloat64
     for _,val := range(slice){
       if val < rslt {
         rslt = val
       }
     }
     return rslt
}

func maxSlice (slice []float64)float64{
     rslt := -(math.MaxFloat64-10) // no overflows
     for _,val := range(slice){
       if val > rslt {
         rslt = val
       }
     }
     return rslt
}

// Calculates the projection approximation collision cross section
// for a molecule by averaging over all rotamers.
func PACCS (mol Molecule, trialsperrotamer int, numrotamers int, parameters ParameterSet) float64 {
    var ccssum float64 = 0.0
    for count := 0; count < numrotamers; count++{
      mol = RotateMolecule(mol, 4 * math.Pi * rand.Float64(), 4 * math.Pi * rand.Float64(), 4 * math.Pi * rand.Float64())
      ccssum += PACCSRotamer(mol, trialsperrotamer, parameters)
    }
    return ccssum / float64(numrotamers)
}

// Calculates the projection approximation collision cross section
// for a single rotamer.
func PACCSRotamer (mol Molecule, trials int, parameters ParameterSet) float64 {
    padding := 5.0 // padding of 5 angstrom sufficient
    minx := minSlice(mol.xs) - padding; maxx := maxSlice(mol.xs) + padding;
    miny := minSlice(mol.ys) - padding; maxy := maxSlice(mol.ys) + padding;

    maxminx := maxx-minx; maxminy := maxy-miny;
    hits := 0
    for count := 0; count < trials; count++ {
      randx := rand.Float64() * maxminx + minx
      randy := rand.Float64() * maxminy + miny
      for idx,x := range mol.xs {
        y := mol.ys[idx]
        dx, dy := math.Abs(randx - x), math.Abs(randy - y)
        radius := parameters[mol.atom_labels[idx]]
        if dx * dx + dy * dy < radius * radius {
            hits += 1
            break // no double/multiple hits
        }
      }
    }
    return (float64(hits) / float64(trials)) * maxminx * maxminy
}

// Rotates a molecule <mol> by rotx, roty and rotx.
func RotateMolecule (mol Molecule, rotx float64, roty float64, rotz float64) Molecule {
    var cx float64 = 0.0; var cy float64 = 0.0;  var cz float64 = 0.0;
    var nxs []float64; var nys []float64; var nzs []float64;
    for idx, x := range mol.xs {
      cx += x; cy += mol.ys[idx]; cz += mol.zs[idx];
    }
    var mol_len = float64(len(mol.xs))
    cx = cx / mol_len; cy = cy / mol_len; cz = cz / mol_len;
    for idx, x := range mol.xs {
      x = x - cx; var y = mol.ys[idx] - cy; var z = mol.zs[idx] - cz;
      y, z = y * math.Cos(rotx) - z * math.Sin(rotx), y * math.Sin(rotx) + z * math.Cos(rotx)
      x, z = x * math.Cos(roty) + z * math.Sin(roty),
            -x * math.Sin(roty) + z * math.Cos(roty)
      x, y = x * math.Cos(rotz) - y * math.Sin(rotz),
            x * math.Sin(rotz) + y * math.Cos(rotz)
      x, y, z = x + cx, y + cy, z + cz
      nxs = append(nxs, x); nys = append(nys, y); nzs = append(nzs, z);
    }
    return Molecule { atom_labels: mol.atom_labels, xs:nxs, ys:nys, zs:nzs };
}

// Loads a molecule, i.e. atom-symbols, atom-x, atom-y, atom-z given
// a string of the xyz-molecule-file format.
func Loadxyzstring(xyzstring string) Molecule {
    xyzstring = strings.Replace(xyzstring,"\t"," ",-1)
    for strings.Contains(xyzstring, "  "){
      xyzstring = strings.Replace(xyzstring,"  "," ",1)
    }
    var atom_labels []string;
    var atoms_x []float64;
    var atoms_y []float64;
    var atoms_z []float64;
    for count, line := range strings.Split(strings.TrimSuffix(xyzstring, "\n"), "\n") {
      err := false
      if len(strings.Split(line, " ")) > 3{
        atom_label := strings.Split(line, " ")[0]
        atom_x, err_x := strconv.ParseFloat(strings.Split(line, " ")[1],64)
        atom_y, err_y := strconv.ParseFloat(strings.Split(line, " ")[2],64)
        atom_z, err_z := strconv.ParseFloat(strings.Split(line, " ")[3],64)
        err = (err_x != nil || err_y != nil || err_z != nil)
        if !err {
          atom_labels = append(atom_labels, atom_label)
          atoms_x = append(atoms_x, atom_x)
          atoms_y = append(atoms_y, atom_y)
          atoms_z = append(atoms_z, atom_z)
        }
      }else{err = true}
      if err {
        if count > 3 {
          // could not understand input, so panic
          panic("could not read line"+line)
        }else{
          // skip on first three lines, as ill-formated input is
          // expected here, and the information is not needed anyway.
          // note that this could lead to a silent error in case that
          // the comment is of a form like C 1 2 3, but the feature
          // of xyz file comments should not be used like that IN ANY CASE.
          continue
        }
      }else{

    }
    }
    return Molecule { atom_labels: atom_labels, xs: atoms_x, ys: atoms_y, zs: atoms_z }
}

// Loads a molecule from a xyzfile.
func Loadxyzfile(xyzfile string) Molecule {
    xyzbytes,err := ioutil.ReadFile(xyzfile)
    if err != nil {panic(err)}
    return Loadxyzstring(string(xyzbytes))
}






//

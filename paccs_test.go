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

import ("fmt" // for inspection
        "testing"
        "math"
      )

func TestPACCSRotamerSimple(t *testing.T) {
  mol := Loadxyzstring("C 0 0 1")
  ccs := PACCSRotamer(mol, 100000, Parametersforname("mobcal"))
  fmt.Println("ccs of single rotamer of mol")
  fmt.Println(mol)
  fmt.Println("is:")
  fmt.Println(ccs)
}

func TestPACCSSimple(t *testing.T) {
  mol := Loadxyzstring("C 0 0 1")
  ccs := PACCS(mol, 10000, 1000, Parametersforname("mobcal"))
  fmt.Println("rotationally averaged ccs of mol")
  fmt.Println(mol)
  fmt.Println("is:")
  fmt.Println(ccs)
}

func TestPACCSAlkanes(t *testing.T) {
  delta_ccs := 0.3
  mol := Loadxyzfile("xyz/methane.xyz")
  exp_ccs := 27.5 // calculated with mobcal
  ccs := PACCS(mol, 10000, 1000, Parametersforname("mobcal"))
  fmt.Println("-- ccs of methane")
  fmt.Print("\tis:  ");fmt.Print(ccs);fmt.Print("  should be:  ");fmt.Println(exp_ccs);
  if math.Abs(ccs - exp_ccs) > delta_ccs {
    t.Errorf("calculated ccs differs from mobcal reference values!")
  }

  mol = Loadxyzfile("xyz/ethane.xyz")
  exp_ccs = 35.81 // calculated with mobcal
  ccs = PACCS(mol, 10000, 1000, Parametersforname("mobcal"))
  fmt.Println("-- ccs of ethane")
  fmt.Print("\tis:  ");fmt.Print(ccs);fmt.Print("  should be:  ");fmt.Println(exp_ccs);
  if math.Abs(ccs - exp_ccs) > delta_ccs {
    t.Errorf("calculated ccs differs from mobcal reference values!")
  }

  mol = Loadxyzfile("xyz/propane.xyz")
  exp_ccs = 42.46 // calculated with mobcal
  ccs = PACCS(mol, 10000, 1000, Parametersforname("mobcal"))
  fmt.Println("-- ccs of propane")
  fmt.Print("\tis:  ");fmt.Print(ccs);fmt.Print("  should be:  ");fmt.Println(exp_ccs);
  if math.Abs(ccs - exp_ccs) > delta_ccs {
    t.Errorf("calculated ccs differs from mobcal reference values!")
  }

  mol = Loadxyzfile("xyz/butane.xyz")
  exp_ccs = 50.11 // calculated with mobcal
  ccs = PACCS(mol, 10000, 1000, Parametersforname("mobcal"))
  fmt.Println("-- ccs of butane")
  fmt.Print("\tis:  ");fmt.Print(ccs);fmt.Print("  should be:  ");fmt.Println(exp_ccs);
  if math.Abs(ccs - exp_ccs) > delta_ccs {
    t.Errorf("calculated ccs differs from mobcal reference values!")
  }
}

func TestRotateMolecule(t *testing.T) {
  floatDelta := 0.00001;
  mol := Loadxyzstring("C 0 0 1\nN 0 0 -1")
  rotmol := RotateMolecule(mol,math.Pi,0,0) // 90
  fmt.Println("rotating")
  fmt.Println(mol)
  fmt.Println("180 degrees around x axis results in")
  fmt.Println(rotmol)
  if (rotmol.xs[0] - mol.xs[0]) > floatDelta || (rotmol.ys[0] - mol.ys[0]) > floatDelta{
    t.Errorf("wrong effect of rotation")
  }
  if (rotmol.xs[1] - mol.xs[1]) > floatDelta || (rotmol.ys[1] - mol.ys[1]) > floatDelta{
    t.Errorf("wrong effect of rotation")
  }
  if rotmol.zs[0] + mol.zs[0] > floatDelta || rotmol.zs[1] + mol.zs[1] > floatDelta {
    t.Errorf("wrong effect of rotation")
  }
}


func TestParametersforname(t *testing.T) {
  rslt := Parametersforname("mobcal")
  if rslt == nil{
    t.Errorf("parametersforname('mobcal') should not be null!")
  }
  if rslt["H"] != 2.2 || rslt["C"] != 2.7 || rslt["N"] != 2.7 || rslt["O"] != 2.7 {
    t.Errorf("wrong mobcal parameters loaded")
  }
  rslt = Parametersforname("")
  if rslt == nil{
    t.Errorf("parametersforname('') should not be null!")
  }
  if rslt["H"] != 2.01 || rslt["C"] != 2.35 || rslt["N"] != 2.26 || rslt["O"] != 2.26 {
    t.Errorf("wrong default (siu et al) parameters loaded")
  }
}


func TestLoadxyzstring(t *testing.T) {
  fmt.Println("TestLoadxyzstring")

  var labels []string; var xs []float64; var ys []float64; var zs []float64;
  mol := Loadxyzstring("C 0 0 1\n")
  labels = mol.atom_labels; xs = mol.xs; ys = mol.ys; zs = mol.zs;
  fmt.Println("Loadxyzstring(C 0 0 1\\n):")
  fmt.Println(labels)
  fmt.Println(xs);fmt.Println(ys);fmt.Println(zs)
  if labels == nil{
    t.Errorf("Loadxyzstring('C 0 0 1\n') should not be null!")
  }
  if len(labels) != len([]string{"C"}){
    t.Errorf("len(Loadxyzstring('C 0 0 1\n')) should be 1!")
  }

  mol = Loadxyzstring("C 0 0 1\nN 0 0 2\n")
  labels = mol.atom_labels; xs = mol.xs; ys = mol.ys; zs = mol.zs;
  fmt.Println("Loadxyzstring(C 0 0 1\\nN 0 0 2\\n):")
  fmt.Println(labels)
  fmt.Println(xs);fmt.Println(ys);fmt.Println(zs)
  if labels == nil{
    t.Errorf("Loadxyzstring('C 0 0 1\nN 0 0 2\n') should not be null!")
  }
  if len(labels) != len([]string{"C", "N"}){
    t.Errorf("len(Loadxyzstring('C 0 0 1\nN 0 0 2\n')) should be 2!")
  }

  mol = Loadxyzstring("3\nA test xyzfile\nC 0 0 1\nN 0 0 2\nBr 11.1 -22.2 -33.3\n")
  labels = mol.atom_labels; xs = mol.xs; ys = mol.ys; zs = mol.zs;
  fmt.Println("Loadxyzstring('3\\nA test xyzfile\\nC 0 0 1\\nN 0 0 2\\nBr 11.1 -22.2 -33.3\\n'):")
  fmt.Println(labels)
  fmt.Println(xs);fmt.Println(ys);fmt.Println(zs)
  if labels == nil{
    t.Errorf("Loadxyzstring('3\\nA test xyzfile\\nC 0 0 1\\nN 0 0 2\\nBr 11.1 -22.2 -33.3\\n') should not be null!")
  }
  if len(labels) != len([]string{"C", "N", "Br"}){
    t.Errorf("len(Loadxyzstring('3\\nA test xyzfile\\nC 0 0 1\\nN 0 0 2\\nBr 11.1 -22.2 -33.3\\n')) should be 3!")
  }

}

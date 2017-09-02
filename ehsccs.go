package main

import (
  "math"
  "math/rand"
)

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

type vec3 struct {
    x float64
    y float64
    z float64
}

type line struct {
    origin vec3
    direction vec3
}

type sphere struct {
    radius float64
    center vec3
}

func vecPlus(fstVec vec3, sndVec vec3) vec3 {
    return vec3 { x: fstVec.x + sndVec.x, y: fstVec.y + sndVec.y, z: fstVec.z + sndVec.z }
}

func vecMult(vec vec3, scalar float64) vec3 {
  return vec3 { x: vec.x * scalar, y: vec.y * scalar, z: vec.z * scalar }
}

func vecMinus(fstVec vec3, sndVec vec3) vec3 {
  return vec3 { x: fstVec.x - sndVec.x, y: fstVec.y - sndVec.y, z: fstVec.z - sndVec.z }
}

func vecDist(fstVec vec3, sndVec vec3) float64 {
  return vecLen(vecMinus(fstVec, sndVec))
}

func vecLenSquare(vec vec3) float64 {
    return (vec.x*vec.x + vec.y*vec.y + vec.z*vec.z)
}

func vecLen(vec vec3) float64 {
    return math.Sqrt(vec.x*vec.x + vec.y*vec.y + vec.z*vec.z)
}

func toUnitVec(vec vec3) vec3 {
    return vecMult(vec, 1 / vecLen(vec))
}

func dotProduct(fstVec vec3, sndVec vec3) float64 {
  return fstVec.x * sndVec.x + fstVec.y * sndVec.y + fstVec.z * sndVec.z
}

// Calculates the exact hard sphere (EHS) collision cross section
// for a molecule by averaging over all rotamers.
func EHSCCS (mol Molecule, trialsperrotamer int, numrotamers int, parameters ParameterSet) float64 {
    var ccssum float64 = 0.0
    for count := 0; count < numrotamers; count++{
      mol = RotateMolecule(mol, 4 * math.Pi * rand.Float64(), 4 * math.Pi * rand.Float64(), 4 * math.Pi * rand.Float64())
      ccssum += EHSCCSRotamer(mol, trialsperrotamer, parameters)
    }
    return ccssum / float64(numrotamers)
}


// Calculates the exact hard sphere (EHS) collision cross section
// for a single rotamer.
func EHSCCSRotamer (mol Molecule, trials int, parameters ParameterSet) float64 {

    spheres := moleculeToSpheres(mol, parameters)

    padding := 5.0 // padding of 5 angstrom sufficient
    minx := minSlice(mol.xs) - padding; maxx := maxSlice(mol.xs) + padding;
    miny := minSlice(mol.ys) - padding; maxy := maxSlice(mol.ys) + padding;

    maxminx := maxx-minx; maxminy := maxy-miny;
    hits := 0.0
    for count := 0; count < trials; count++ {
      randx := rand.Float64() * maxminx + minx
      randy := rand.Float64() * maxminy + miny

      a := line{direction: vec3{x: 0, y: 0, z: 1}, origin: vec3{x: randx, y: randy, z: -100}}
      b := lineSpheresTrajectory(a, spheres)
      ab := dotProduct(a.direction,b.direction)
      //abs_a := vecLen(a.direction) // => unit vector anyway
      //abs_b := vecLen(b.direction) // => unit vector anyway
      hits += 1 - ab // / (abs_a * abs_b)
    }
    return (float64(hits) / float64(trials)) * maxminx * maxminy
}

func moleculeToSpheres (mol Molecule, parameters ParameterSet) []sphere {
    var spheres []sphere
    for idx,atm_lab := range mol.atom_labels {
      spheres = append(spheres, sphere{center: vec3{x: mol.xs[idx], y: mol.ys[idx], z: mol.zs[idx]}, radius: parameters[atm_lab]})
    }
    return spheres
}


func lineSpheresTrajectory(lne line, spheres []sphere) line {
    for true {
      nextIntsctLineScalar,nextIntsctSphere := nextLineSpheresIntersection(lne, spheres)
      if nextIntsctLineScalar == math.MaxFloat64 { // ???
        break
      }
      lne = reflectLineOnSphere(lne, nextIntsctSphere, nextIntsctLineScalar)
      pointOfCollision := vecPlus(vecMult(lne.direction, nextIntsctLineScalar),lne.origin)
      lne.origin = pointOfCollision // move ray to current position
    }
    return lne
}


func reflectLineOnSphere(lne line, sph sphere, intsctLineScalar float64) line {
    pointOfCollision := vecPlus(vecMult(lne.direction, intsctLineScalar),lne.origin)
    // see https://math.stackexchange.com/questions/2334939/reflection-of-line-on-a-sphere/2334963?noredirect=1#comment4807112_2334963
    // 2 * [(line-direction) * (point-of-collision - center-of-sphere)] *
    // (point-of-collision - center-of-sphere) - (line-direction)
    // == 2 * [v * (x-c)] * (x-c) -v
    x_c := vecMinus(pointOfCollision,sph.center)
    newDir := vecMinus(vecMult(x_c, 2 * dotProduct(lne.direction,x_c)), lne.direction)
    newDir = vecMult(newDir, 1.0 / vecLen(newDir)) // dont forget to normalize
    return line{direction: newDir, origin: lne.origin}
}


func nextLineSpheresIntersection(lne line, spheres []sphere) (float64,sphere) {
    var nextIntsctSphere sphere
    nextIntsctLineScalar := math.MaxFloat64
    for _,sph := range spheres {
      fstIntersection,sndIntersection,success := lineSphereIntersections(lne, sph)
      if !success { continue }
      intersections := filterAboveZero([]float64{fstIntersection,sndIntersection})
      for _,intersectionScalar := range intersections {
        if intersectionScalar < nextIntsctLineScalar {
          nextIntsctLineScalar = intersectionScalar
          nextIntsctSphere = sph
        }
      }
    }
    return nextIntsctLineScalar,nextIntsctSphere
}



// computes the intersections of the line lne with the sphere sph:
//  according to https://en.wikipedia.org/wiki/Line%E2%80%93sphere_intersection
//  a line given in the parametric form
//    x = o + d * L
//  will have intersections with the sphere
//    ||x - c||^2 = r^2
//  at the two intersections given by
//    d1,2 = -(L * (o - c)) +- ( (L * (o-c))^2 - ||o - c||^2 +r^2 )^1/2
//  In the function below, we will use the substitutions
//    loc := L * (o - c)
//  and
//    oc := ||o - c||^2
func lineSphereIntersections(lne line, sph sphere)(float64,float64,bool) {

    if math.Abs(vecLen(lne.direction)-1.0) > 0.00001 {
        panic("non-unit vector encountered for direction of line in lineSphereIntersections")
    }
    locVec := vec3{ x: (lne.origin.x - sph.center.x) * lne.direction.x,
              y: (lne.origin.y - sph.center.y) * lne.direction.y,
              z: (lne.origin.z - sph.center.z) * lne.direction.z,
              }
    ocVec := vecMinus(lne.origin, sph.center)
    loc := locVec.x + locVec.y + locVec.z
    oc := vecLenSquare(ocVec)

    radicant := loc * loc - oc + sph.radius * sph.radius
    if radicant <= 0.0 {
      return 0.0, 0.0, false
    }
    radicant = math.Sqrt(radicant)
    return -loc - radicant, -loc + radicant, true
}





























//

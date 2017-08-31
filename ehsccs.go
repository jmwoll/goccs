package main

import (
  "math"
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

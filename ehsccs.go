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


func lineSphereIntersections(lne line, sph sphere)(float64,float64,bool) {
    oc := vec3{ x: lne.origin.x - sph.center.x,
              y: lne.origin.y - sph.center.y,
              z: lne.origin.z - sph.center.z,
              }
    ocSum := oc.x + oc.y + oc.z
    ocAbsSq := oc.x * oc.x + oc.y * oc.y + oc.z * oc.z
    ocSq := ocSum * ocSum
    radicant := ocSq - ocAbsSq + sph.radius * sph.radius
    if radicant < 0 {
      return 0.0, 0.0, false
    }
    radicant = math.Sqrt(radicant)
    return -ocSum - radicant, -ocSum + radicant, true

}





























//

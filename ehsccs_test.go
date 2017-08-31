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
        "testing"
      )

// Given a line in parametric form and a sphere, lineSphereIntersections
// should return the parameters of the line sphere intersection or
// false, if there is no intersection.
func TestLineSphereIntersections(t *testing.T) {
    logTestName("TestLineSphereIntersections")

    // -- simple cases where intersections occur
    logTest("given a line (0,0,0) + t*(0,0,1)")
    lne := line{ origin: vec3{x: 0, y: 0, z: 0},
                 direction: vec3{x: 0, y: 0, z: 1},
                }
    logTest("and a sphere with radius 1 at (0,0,3)")
    sph := sphere{ center: vec3{x: 0, y: 0, z: 3},
                   radius: 1,
                 }
    i1, i2, _ := lineSphereIntersections(lne, sph)
    assertTrue(i1 == 2 && i2 == 4, "the intersections should be t = 2,4", t)
    assertTrue(vecDist(vecPlus(vecMult(lne.direction, i1),lne.origin), sph.center) <= sph.radius, "lne should intersect sphere", t)
    assertTrue(vecDist(vecPlus(vecMult(lne.direction, i2),lne.origin), sph.center) <= sph.radius, "lne should intersect sphere", t)


    logTest("given a line (0,0,0) + t*(0,1,0)")
    lne = line{ origin: vec3{x: 0, y: 0, z: 0},
                 direction: vec3{x: 0, y: 1, z: 0},
                }
    logTest("and a sphere with radius 2 at (0,6,0)")
    sph = sphere{ center: vec3{x: 0, y: 6, z: 0},
                   radius: 2,
                 }
    i1, i2, _ = lineSphereIntersections(lne, sph)
    assertTrue(i1 == 4 && i2 == 8, "the intersections should be t = 4,8", t)
    assertTrue(vecDist(vecPlus(vecMult(lne.direction, i1),lne.origin), sph.center) <= sph.radius, "lne should intersect sphere", t)
    assertTrue(vecDist(vecPlus(vecMult(lne.direction, i2),lne.origin), sph.center) <= sph.radius, "lne should intersect sphere", t)




    logTest("given a line (-6,0,0) + t*(1,0,0)")
    lne = line{ origin: vec3{x: -6, y: 0, z: 0},
                 direction: vec3{x: 1, y: 0, z: 0},
              }
    logTest("and a sphere with radius 0.5 at (6,0,0)")
    sph = sphere{ center: vec3{x: 6, y: 0, z: 0},
                   radius: 0.5,
                }
    i1, i2, _ = lineSphereIntersections(lne, sph)
    logTestInline("i1,2");logTest(i1);logTest(i2)
    assertTrue(i1 == 11.5 && i2 == 12.5, "the intersections should be t = 11.5,12.5", t)
    assertTrue(vecDist(vecPlus(vecMult(lne.direction, i1),lne.origin), sph.center) <= sph.radius, "lne should intersect sphere", t)
    assertTrue(vecDist(vecPlus(vecMult(lne.direction, i2),lne.origin), sph.center) <= sph.radius, "lne should intersect sphere", t)


    logTest("given a line (1,1,1) + t*(0,0,1)")
    lne = line{ origin: vec3{x: -1, y: -1, z: 1},
                 direction: vec3{x: 0, y: 0, z: 1},
                }
    // normalize line
    lne.direction = toUnitVec(lne.direction)
    logTest("and a sphere with radius 1 at (0,0,1000)")
    sph = sphere{ center: vec3{x: -1, y: -1, z: 1000},
                   radius: 1,
                 }

    i1, i2, _ = lineSphereIntersections(lne, sph)
    logTestInline("i1,2");logTest(i1);logTest(i2)
    assertTrue(i1 == 998 && i2 == 1000, "the intersections should be t = 998,1000", t)
    assertTrue(vecDist(vecPlus(vecMult(lne.direction, i1),lne.origin), sph.center) <= sph.radius, "lne should intersect sphere", t)
    assertTrue(vecDist(vecPlus(vecMult(lne.direction, i2),lne.origin), sph.center) <= sph.radius, "lne should intersect sphere", t)


    logTest("given a line (1,1,1) + t*(0,1,1)")
    lne = line{ origin: vec3{x: -1, y: 0, z: 0},
                 direction: vec3{x: 0, y: 1, z: 1},
                }
    // normalize line
    lne.direction = toUnitVec(lne.direction)
    logTest("after normalization:");logTest(lne);
    logTest("and a sphere with radius 1 at (0,1000,1000)")
    sph = sphere{ center: vec3{x: -1, y: 1000, z: 1000},
                   radius: 1,
                 }

    i1, i2, _ = lineSphereIntersections(lne, sph)
    logTestInline("i1,2");logTest(i1);logTest(i2)
    assertTrue(vecDist(vecPlus(vecMult(lne.direction, i1),lne.origin), sph.center) <= sph.radius, "lne should intersect sphere", t)
    assertTrue(vecDist(vecPlus(vecMult(lne.direction, i2),lne.origin), sph.center) <= sph.radius, "lne should intersect sphere", t)


    // -- simple cases where no intersections occur
    logTest("given a line (1,1,1) + t*(1,0,0)")
    lne = line{ origin: vec3{x: 1, y: 1, z: 1},
                 direction: vec3{x: 1, y: 0, z: 0},
                }
    // normalize line
    lne.direction = toUnitVec(lne.direction)
    logTest("and a sphere with radius 1 at (0,0,1000)")
    sph = sphere{ center: vec3{x: 0, y: 0, z: 1000},
                   radius: 1,
                 }

    i1, i2, hit := lineSphereIntersections(lne, sph)
    logTestInline("i1,2");logTest(i1);logTest(i2)
    assertTrue(i1 == 0 && i2 == 0, "there should be no intersections (intersections should be 0,0)", t)
    logTestInline("and flag should be false");logTest(hit);
    assertTrue(hit == false, "when no intersections are found, the hit flag should be false", t);

    logTest("given a line (1,1,1) + t*(0,1,0)")
    lne = line{ origin: vec3{x: 1, y: 1, z: 1},
                 direction: vec3{x: 0, y: 1, z: 0},
                }
    // normalize line
    lne.direction = toUnitVec(lne.direction)
    logTest("and a sphere with radius 1 at (0,0,1000)")
    sph = sphere{ center: vec3{x: 0, y: 0, z: 1000},
                   radius: 1,
                 }

    i1, i2, hit = lineSphereIntersections(lne, sph)
    logTestInline("i1,2");logTest(i1);logTest(i2)
    assertTrue(i1 == 0 && i2 == 0, "there should be no intersections (intersections should be 0,0)", t)
    logTestInline("and flag should be false");logTest(hit);
    assertTrue(hit == false, "when no intersections are found, the hit flag should be false", t);


    logTest("given a line (1,1,1) + t*(0,0,1)")
    lne = line{ origin: vec3{x: 1, y: 1, z: 1},
                 direction: vec3{x: 0, y: 0, z: 1},
                }
    // normalize line
    lne.direction = toUnitVec(lne.direction)
    logTest("and a sphere with radius 1 at (0,0,1000)")
    sph = sphere{ center: vec3{x: 0, y: 0, z: 1000},
                   radius: 1,
                 }

    i1, i2, _ = lineSphereIntersections(lne, sph)
    logTestInline("i1,2");logTest(i1);logTest(i2)
    assertTrue(i1 == 0 && i2 == 0, "there should be no intersections (intersections should be 0,0)", t)

}































//

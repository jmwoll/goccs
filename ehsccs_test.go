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
        "fmt"
        "testing"
        "math"
      )


func TestEHSCCSRotamerSimple(t *testing.T) {
    logTestName("TestEHSCCSRotamerSimple")
    mol := Loadxyzstring("C 0 0 1")
    ccs := EHSCCSRotamer(mol, 1000, EHSParametersforname("mobcal"))
    logTest("EHS CCS:");logTest(ccs);
    assertTrue(ccs > 0, "CCS values always non-zero and positive", t)
}

func TestEHSCCSMethane(t *testing.T) {
    logTestName("TestEHSCCSRotamerMethane")
    mol := Loadxyzfile("xyz/methane.xyz")
    logTest(mol)
    ccs := EHSCCS(mol, 10000, 3000, EHSParametersforname("mobcal"))
    logTest("EHS CCS:");logTest(ccs);
    assertTrue(ccs > 0, "CCS values always non-zero and positive", t)
    refccs := 27.602
    assertTrue(math.Abs(ccs - refccs) < 0.5, "CCS values differs from reference", t)
}

func TestEHSCCSButane(t *testing.T) {
    logTestName("TestEHSCCSRotamerButane")
    mol := Loadxyzfile("xyz/butane.xyz")
    ccs := EHSCCS(mol, 10000, 3000, EHSParametersforname("mobcal"))
    logTest("EHS CCS:");logTest(ccs);
    assertTrue(ccs > 0, "CCS values always non-zero and positive", t)
    refccs := 52.101
    assertTrue(math.Abs(ccs - refccs) < 0.5, "CCS values differs from reference", t)
}

func TestEHSCCSPentane(t *testing.T) {
    logTestName("TestEHSCCSRotamerPentane")
    mol := Loadxyzfile("xyz/pentane.xyz")
    ccs := EHSCCS(mol, 10000, 3000, EHSParametersforname("mobcal"))
    logTest("EHS CCS:");logTest(ccs);
    assertTrue(ccs > 0, "CCS values always non-zero and positive", t)
    refccs := 59.653
    assertTrue(math.Abs(ccs - refccs) < 0.5, "CCS values differs from reference", t)
}

func TestEHSCCSOctabenzone(t *testing.T) {
    logTestName("TestEHSCCSOctabenzone")
    mol := Loadxyzfile("xyz/octabenzone.xyz")
    ccs := EHSCCS(mol, 10000, 3000, EHSParametersforname("mobcal"))
    logTest("EHS CCS:");logTest(ccs);
    assertTrue(ccs > 0, "CCS values always non-zero and positive", t)
    refccs := 157.8
    assertTrue(math.Abs(ccs - refccs) < 0.5, "CCS values differs from reference", t)
}

func TestEHSCCSAbacavir(t *testing.T) {
    logTestName("TestEHSCCSAbacavir")
    mol := Loadxyzfile("xyz/abacavir.xyz")
    ccs := EHSCCS(mol, 10000, 3000, EHSParametersforname("mobcal"))
    logTest("EHS CCS:");logTest(ccs);
    assertTrue(ccs > 0, "CCS values always non-zero and positive", t)
    refccs := 126.28
    assertTrue(math.Abs(ccs - refccs) < 0.5, "CCS values differs from reference", t)
}

func TestEHSCCSPaclitaxel(t *testing.T) {
    logTestName("TestEHSCCSPaclitaxel")
    mol := Loadxyzfile("xyz/paclitaxel.xyz")
    ccs := EHSCCS(mol, 10000, 3000, EHSParametersforname("mobcal"))
    logTest("EHS CCS:");logTest(ccs);
    assertTrue(ccs > 0, "CCS values always non-zero and positive", t)
    refccs := 231.0
    assertTrue(math.Abs(ccs - refccs) < 0.5, "CCS values differs from reference", t)
}

func TestEHSCCSCiclosporin(t *testing.T) {
    logTestName("TestEHSCCSCiclosporin")
    mol := Loadxyzfile("xyz/ciclosporin.xyz")
    ccs := EHSCCS(mol, 10000, 3000, EHSParametersforname("mobcal"))
    logTest("EHS CCS:");logTest(ccs);
    assertTrue(ccs > 0, "CCS values always non-zero and positive", t)
    refccs := 306.0
    assertTrue(math.Abs(ccs - refccs) < 0.5, "CCS values differs from reference", t)
}


func TestReflectLineOnsphere(t *testing.T) {
    logTestName("TestReflectLineOnsphere (1.)")
    lne := line{ origin: vec3{x:0,y:0,z:0}, direction: vec3{x:1,y:0,z:0}}
    sph := sphere{ center: vec3{x:5,y:0,z:0}, radius: 1.0}
    logTest(fmt.Sprintf("given a line %v",lne))
    logTest(fmt.Sprintf("and a sphere %v",sph))
    intsct_1, intsct_2, _ := lineSphereIntersections(lne,sph)
    intscts := []float64{intsct_1,intsct_2}
    logTest("line sphere intersections are:");logTest(intscts);
    intsct := minSlice(filterAboveZero(intscts))
    logTest("smallest intersection (where hit happens)")
    logTest(intsct)
    logTest("reflected line")
    rslt := reflectLineOnSphere(lne,sph,intsct)
    logTest(rslt)
    assertTrue(rslt.direction.x == -1,fmt.Sprintf("the reflected line should be %v (and was %v)",
      line{origin:vec3{x:0,y:0,z:0},direction:vec3{x:-1,y:0,z:0}},rslt),t)

    logTestName("TestReflectLineOnsphere (2.)")
    lne = line{ origin: vec3{x:0,y:0,z:0}, direction: vec3{x:0,y:1,z:0}}
    sph = sphere{ center: vec3{x:0,y:5,z:0}, radius: 1.0}
    logTest(fmt.Sprintf("given a line %v",lne))
    logTest(fmt.Sprintf("and a sphere %v",sph))
    intsct_1, intsct_2, _ = lineSphereIntersections(lne,sph)
    intscts = []float64{intsct_1,intsct_2}
    logTest("line sphere intersections are:");logTest(intscts);
    intsct = minSlice(filterAboveZero(intscts))
    logTest("smallest intersection (where hit happens)")
    logTest(intsct)
    logTest("reflected line")
    rslt = reflectLineOnSphere(lne,sph,intsct)
    logTest(rslt)
    assertTrue(rslt.direction.y == -1,fmt.Sprintf("the reflected line should be %v (and was %v)",
      line{origin:vec3{x:0,y:0,z:0},direction:vec3{x:0,y:-1,z:0}},rslt),t)

    logTestName("TestReflectLineOnsphere (3.)")
    lne = line{ origin: vec3{x:0,y:0,z:0}, direction: vec3{x:0,y:0,z:1}}
    sph = sphere{ center: vec3{x:0,y:0,z:5}, radius: 2.0}
    logTest(fmt.Sprintf("given a line %v",lne))
    logTest(fmt.Sprintf("and a sphere %v",sph))
    intsct_1, intsct_2, _ = lineSphereIntersections(lne,sph)
    intscts = []float64{intsct_1,intsct_2}
    logTest("line sphere intersections are:");logTest(intscts);
    intsct = minSlice(filterAboveZero(intscts))
    logTest("smallest intersection (where hit happens)")
    logTest(intsct)
    logTest("reflected line")
    rslt = reflectLineOnSphere(lne,sph,intsct)
    logTest(rslt)
    assertTrue(rslt.direction.z == -1,fmt.Sprintf("the reflected line should be %v (and was %v)",
      line{origin:vec3{x:0,y:0,z:0},direction:vec3{x:0,y:0,z:-1}},rslt),t)

    logTestName("TestReflectLineOnsphere (4.)")
    lne = line{ origin: vec3{x:0,y:0,z:0}, direction: vec3{x:1,y:1,z:0}}
    lne.direction = toUnitVec(lne.direction)
    sph = sphere{ center: vec3{x:1,y:2,z:0}, radius: 1.0}
    logTest(fmt.Sprintf("given a line %v",lne))
    logTest(fmt.Sprintf("and a sphere %v",sph))
    intsct_1, intsct_2, _ = lineSphereIntersections(lne,sph)
    intscts = []float64{intsct_1,intsct_2}
    logTest("line sphere intersections are:");logTest(intscts);
    intsct = minSlice(filterAboveZero(intscts))
    logTest("smallest intersection (where hit happens)")
    logTest(intsct)
    logTest("reflected line")
    rslt = reflectLineOnSphere(lne,sph,intsct)
    expRslt := line{ origin: vec3{x:0,y:0,z:0}, direction: toUnitVec(vec3{x:1,y:-1,z:0})}
    logTest(rslt)
    assertTrue(vecEquals(rslt.direction,expRslt.direction),fmt.Sprintf("the reflected line should be %v (and was %v)",
      expRslt,rslt),t)
}

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

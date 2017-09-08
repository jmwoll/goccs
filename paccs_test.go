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
	"math"
	"testing"
)

func TestPACCSRotamerSimple(t *testing.T) {
	logTestName("TestPACCSRotamerSimple")
	mol := Loadxyzstring("C 0 0 1")
	ccs := PACCSRotamer(mol, 100000, PAParametersforname("mobcal"))
	logTest("given the single rotamer")
	logTestInline(mol)
	logTest("")
	assertTrue(math.Abs(ccs-23) < 0.3,
		"the single-rotamer-ccs should be 23 +- 0.3", t)
}

func TestPACCSSimple(t *testing.T) {
	logTestName("TestPACCSSimple")
	mol := Loadxyzstring("C 0 0 1")
	ccs := PACCS(mol, 10000, 1000, PAParametersforname("mobcal"))
	logTest("given the molecule")
	logTest(mol)
	assertTrue(math.Abs(ccs-23) < 0.3,
		"the ccs should be 23 +- 0.3", t)
}

func TestPACCSMethane(t *testing.T) {
	logTestName("TestPACCSMethane")
	delta_ccs := 0.3
	mol := Loadxyzfile("xyz/methane.xyz")
	exp_ccs := 27.5 // calculated with mobcal
	ccs := PACCS(mol, 10000, 1000, PAParametersforname("mobcal"))
	logTest("given the molecule methane")
	logTest(mol)
	assertTrue(math.Abs(ccs-exp_ccs) < delta_ccs,
		"the ccs should be 27.5 +- 0.3", t)
}

func TestPACCSEthane(t *testing.T) {
	logTestName("TestPACCSEthane")
	delta_ccs := 0.3
	mol := Loadxyzfile("xyz/ethane.xyz")
	exp_ccs := 35.81 // calculated with mobcal
	ccs := PACCS(mol, 10000, 1000, PAParametersforname("mobcal"))
	logTest("given the molecule ethane")
	logTest(mol)
	assertTrue(math.Abs(ccs-exp_ccs) < delta_ccs,
		"the ccs should be 35.81 +- 0.3", t)
}

func TestPACCSPropane(t *testing.T) {
	logTestName("TestPACCSPropane")
	delta_ccs := 0.3
	mol := Loadxyzfile("xyz/propane.xyz")
	exp_ccs := 42.46 // calculated with mobcal
	ccs := PACCS(mol, 10000, 1000, PAParametersforname("mobcal"))
	logTest("given the molecule propane")
	logTest(mol)
	assertTrue(math.Abs(ccs-exp_ccs) < delta_ccs,
		"the ccs should be 42.46 +- 0.3", t)
}

func TestPACCSButane(t *testing.T) {
	logTestName("TestPACCSButane")
	delta_ccs := 0.3
	mol := Loadxyzfile("xyz/butane.xyz")
	exp_ccs := 50.11 // calculated with mobcal
	ccs := PACCS(mol, 10000, 1000, PAParametersforname("mobcal"))
	logTest("given the molecule butane")
	logTest(mol)
	assertTrue(math.Abs(ccs-exp_ccs) < delta_ccs,
		"the ccs should be 50.11 +- 0.3", t)
}

func TestRotateMolecule(t *testing.T) {
	logTestName("TestRotateMolecule")
	floatDelta := 0.00001
	mol := Loadxyzstring("C 0 0 1\nN 0 0 -1")
	rotmol := RotateMolecule(mol, math.Pi, 0, 0) // 90
	logTest("rotating the molecule")
	logTest(mol)
	logTest("180 degrees around x axis results in the molecule")
	logTest(rotmol)
	assertTrue(rotmol.xs[0]-mol.xs[0] < floatDelta && //
		rotmol.ys[0]-mol.ys[0] < floatDelta && //
		rotmol.xs[1]-mol.xs[1] < floatDelta && //
		rotmol.ys[1]-mol.ys[1] < floatDelta, //
		"the x and y coordinates should not change", t)
	assertTrue(rotmol.zs[0]+mol.zs[0] < floatDelta && //
		rotmol.zs[1]+mol.zs[1] < floatDelta, //
		"the z coordinate should be inverted", t)
}

func TestLoadxyzstring(t *testing.T) {
	logTestName("TestLoadxyzstring")

	var labels []string
	var xs []float64
	var ys []float64
	var zs []float64
	mol := Loadxyzstring("C 0 0 1\n")
	labels = mol.atom_labels
	xs = mol.xs
	ys = mol.ys
	zs = mol.zs
	logTest("Loadxyzstring(C 0 0 1\\n):")
	logTest(labels)
	logTest(xs)
	logTest(ys)
	logTest(zs)
	if labels == nil {
		t.Errorf("Loadxyzstring('C 0 0 1\n') should not be null!")
	}
	if len(labels) != len([]string{"C"}) {
		t.Errorf("len(Loadxyzstring('C 0 0 1\n')) should be 1!")
	}

	mol = Loadxyzstring("C 0 0 1\nN 0 0 2\n")
	labels = mol.atom_labels
	xs = mol.xs
	ys = mol.ys
	zs = mol.zs
	logTest("Loadxyzstring(C 0 0 1\\nN 0 0 2\\n):")
	logTest(labels)
	logTest(xs)
	logTest(ys)
	logTest(zs)
	if labels == nil {
		t.Errorf("Loadxyzstring('C 0 0 1\nN 0 0 2\n') should not be null!")
	}
	if len(labels) != len([]string{"C", "N"}) {
		t.Errorf("len(Loadxyzstring('C 0 0 1\nN 0 0 2\n')) should be 2!")
	}

	mol = Loadxyzstring("3\nA test xyzfile\nC 0 0 1\nN 0 0 2\nBr 11.1 -22.2 -33.3\n")
	labels = mol.atom_labels
	xs = mol.xs
	ys = mol.ys
	zs = mol.zs
	logTest("Loadxyzstring('3\\nA test xyzfile\\nC 0 0 1\\nN 0 0 2\\nBr 11.1 -22.2 -33.3\\n'):")
	logTest(labels)
	logTest(xs)
	logTest(ys)
	logTest(zs)
	if labels == nil {
		t.Errorf("Loadxyzstring('3\\nA test xyzfile\\nC 0 0 1\\nN 0 0 2\\nBr 11.1 -22.2 -33.3\\n') should not be null!")
	}
	if len(labels) != len([]string{"C", "N", "Br"}) {
		t.Errorf("len(Loadxyzstring('3\\nA test xyzfile\\nC 0 0 1\\nN 0 0 2\\nBr 11.1 -22.2 -33.3\\n')) should be 3!")
	}

}

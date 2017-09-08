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
	"flag"
	"fmt"
)

func main() {
	defaultStr := "unset"
	approximationPtr := flag.String("approximation", "PA", "The approximation to use: Must be either 'PA' or 'EHS' (defaults to 'PA')")
	xyzfilePtr := flag.String("xyzfile", defaultStr, "The xyzfile for which to calculate the CCS (alternatively, the xyzstring may be specified).")
	xyzstringPtr := flag.String("xyzstring", defaultStr, "The xyzstring for which to calculate the CCS (alternatively, the xyzfile may be specified.)")
	numrotamersPtr := flag.Int("num_rotamers", 0, "The number of rotamers to consider (defaults to '3000')")
	trialsperrotamerPtr := flag.Int("trials_per_rotamer", 0, "The number of trials per rotamer (defaults to '100000' for PA and '10000' for EHS)")
	parameters := flag.String("parameters", "siu_guo_2010", "The parameters to use. Either one of ['siu_guo_2010','mobcal'] or a JSON, e.g.:\n {'H': 1.23, 'C': 2.34, 'O': 3.45}")
	flag.Parse()
	if *approximationPtr == "PA" {
		fmt.Println("The projection approximation (PA) will be used for CCS calculations.")
	} else {
		if *approximationPtr == "EHS" {
			fmt.Println("The exact hard sphere scattering approximation (EHS) will be used for CCS calculations.")
		} else {
			panic("unknown CCS approximation method must be one of ['PA','EHS']")
		}
	}
	if *xyzfilePtr == defaultStr && *xyzstringPtr == defaultStr {
		panic("Either xyzfile or xyzstring must be specified (at least one must be set).")
	}
	if *xyzfilePtr != defaultStr && *xyzstringPtr != defaultStr {
		panic("Either xyzfile or xyzstring must be specified (not both).")
	}
	if *xyzfilePtr != defaultStr {
		fmt.Println()
		fmt.Print("the contents will be read from the xyzfile:")
		fmt.Println(*xyzfilePtr)
	}
	if *xyzstringPtr != defaultStr {
		fmt.Println()
		fmt.Print("the contents will be read from the string:")
		fmt.Println(*xyzstringPtr)
	}
	if *numrotamersPtr == 0 {
		*numrotamersPtr = 3000
	}
	if *trialsperrotamerPtr == 0 {
		if *approximationPtr == "PA" {
			*trialsperrotamerPtr = 100000
		} else {
			*trialsperrotamerPtr = 10000
		}
	}
	var loadedParams ParameterSet
	if *parameters == "mobcal" || *parameters == "siu_guo_2010" {
		if *approximationPtr == "EHS" {
			loadedParams = EHSParametersforname(*parameters)
		} else {
			loadedParams = PAParametersforname(*parameters)
		}
	} else {
		loadedParams = JSONtoParameterSet(*parameters)
	}

	var mol Molecule
	if *xyzstringPtr != defaultStr {
		mol = Loadxyzstring(*xyzstringPtr)
	} else {
		mol = Loadxyzfile(*xyzfilePtr)
	}
	ccs := 0.0
	if *approximationPtr == "PA" {
		ccs = PACCS(mol, *trialsperrotamerPtr, *numrotamersPtr, loadedParams)
	} else {
		ccs = EHSCCS(mol, *trialsperrotamerPtr, *numrotamersPtr, loadedParams)
	}
	fmt.Println(ccs)

}

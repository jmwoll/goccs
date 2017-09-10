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
	"encoding/json"
)

type ParameterSet map[string]float64

// The parameters to use, i.e. radii of atoms in angstrom.
func PAParametersforname(name string) ParameterSet {
	if name == "mobcal" {
		return map[string]float64{
			"H": 2.2, "C": 2.7, "N": 2.7, "O": 2.7,
		}
	}
	if name == "siu_guo_2010" || name == "" {
		return ParameterSet{"H": 2.01, "C": 2.35, "N": 2.26, "O": 2.26}
	}
	return nil
}

func EHSParametersforname(name string) ParameterSet {
	if name == "mobcal" {
		return map[string]float64{"H": 2.2, "C": 2.7, "N": 2.70, "O": 2.70}
	} else {
		return map[string]float64{"H": 1.5, "C": 2.7, "N": 2.50, "O": 2.50}
	}
	return nil
}

func JSONtoParameterSet(jsonInput string) ParameterSet {
	paramSet := map[string]float64{}
	jsonDat := map[string]interface{}{}
	if err := json.Unmarshal([]byte(jsonInput), &jsonDat); err != nil {
		panic(err)
	}
	for k, v := range jsonDat {
		floatVal, ok := v.(float64)
		if ok {
			paramSet[k] = floatVal
		} else {
			panic("error reading in float from param JSON")
		}
	}
	return paramSet
}

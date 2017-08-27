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


type ParameterSet map[string]float64

// The parameters to use, i.e. radii of atoms in angstrom.
func PAParametersforname(name string) ParameterSet {
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

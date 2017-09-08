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

func TestParametersforname(t *testing.T) {
	rslt := PAParametersforname("mobcal")
	if rslt == nil {
		t.Errorf("parametersforname('mobcal') should not be null!")
	}
	if rslt["H"] != 2.2 || rslt["C"] != 2.7 || rslt["N"] != 2.7 || rslt["O"] != 2.7 {
		t.Errorf("wrong mobcal parameters loaded")
	}
	rslt = PAParametersforname("")
	if rslt == nil {
		t.Errorf("parametersforname('') should not be null!")
	}
	if rslt["H"] != 2.01 || rslt["C"] != 2.35 || rslt["N"] != 2.26 || rslt["O"] != 2.26 {
		t.Errorf("wrong default (siu et al) parameters loaded")
	}
}

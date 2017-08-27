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
      )

func assertTrue (predicate bool, msg interface{}, t *testing.T) {
  if !predicate {
    logFail(msg)
    t.Errorf("")
  } else {
    logPassed(msg)
  }
}

func logPassed (msg interface{}) {
  fmt.Print("\033[92m")
  fmt.Print("\t");fmt.Println(msg);
  fmt.Print("\033[0m")
}

func logTest (msg interface{}) {
  fmt.Print("\t");fmt.Println(msg)
}

func logTestInline (msg interface{}) {
  fmt.Print("\t");fmt.Print(msg)
}

func logFail (msg interface{}) {
  fmt.Println("\033[91m")
  fmt.Print("\t");fmt.Println(msg);
  fmt.Println("\033[0m")
}


func logTestName (msg interface{}) {
  fmt.Println("\033[94m")
  fmt.Println("")
  fmt.Print("-- ");fmt.Println(msg);
  fmt.Println("\033[0m")
}

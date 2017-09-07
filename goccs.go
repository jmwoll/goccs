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

func main(){
  defaultStr := "unset"
  approximationPtr := flag.String("approximation", "PA", "The approximation to use: Must be either 'PA' or 'EHS' (defaults to 'PA')")
  xyzfilePtr := flag.String("xyzfile", defaultStr, "The xyzfile for which to calculate the CCS (alternatively, the xyzstring may be specified).")
  xyzstringPtr := flag.String("xyzstring", defaultStr, "The xyzstring for which to calculate the CCS (alternatively, the xyzfile may be specified.)")
  flag.Parse()
  if *approximationPtr == "PA" {
      fmt.Println("The projection approximation (PA) will be used for CCS calculations.")
  } else {
    if *approximationPtr == "EHS" {
      fmt.Println("The exact hard sphere scattering approximation (EHS) will be used for CCS calculations.")
    }else{panic("unknown CCS approximation method must be one of ['PA','EHS']")}
  }
  if *xyzfilePtr == defaultStr && *xyzstringPtr == defaultStr {
    panic("Either xyzfile or xyzstring must be specified (at least one must be set).")
  }
  if *xyzfilePtr != defaultStr && *xyzstringPtr != defaultStr {
    panic("Either xyzfile or xyzstring must be specified (not both).")
  }
  if *xyzfilePtr != defaultStr {
    fmt.Println()
    fmt.Print("the contents will be read from the xyzfile:"); fmt.Println(*xyzfilePtr);
  }
  if *xyzstringPtr != defaultStr {
    fmt.Println()
    fmt.Print("the contents will be read from the string:"); fmt.Println(*xyzstringPtr);
  }


}

package main

import (
  "testing"
)

func TestParametersforname(t *testing.T) {
  rslt := PAParametersforname("mobcal")
  if rslt == nil{
    t.Errorf("parametersforname('mobcal') should not be null!")
  }
  if rslt["H"] != 2.2 || rslt["C"] != 2.7 || rslt["N"] != 2.7 || rslt["O"] != 2.7 {
    t.Errorf("wrong mobcal parameters loaded")
  }
  rslt = PAParametersforname("")
  if rslt == nil{
    t.Errorf("parametersforname('') should not be null!")
  }
  if rslt["H"] != 2.01 || rslt["C"] != 2.35 || rslt["N"] != 2.26 || rslt["O"] != 2.26 {
    t.Errorf("wrong default (siu et al) parameters loaded")
  }
}

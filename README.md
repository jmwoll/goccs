

## About
goccs is a parallel cross platform CCS prediction software implemented in go.

CCS values can be easily computed in a few steps:

![CCS calculation in simple steps.](https://github.com/jmwoll/goccs/blob/master/doc/animation_usage.gif)

## Overview
Running the command
```
goccs_windows_386.exe -xyzfile example.xyz -approximation PA -parameters mobcal
```
gives out the CCS value in angstrom^2, where approximation is either 'PA' or 'EHS'.
Note that the name of the executable varies depending on the platform used. Custom parameters can be
specified in the JSON format:
```
goccs_windows_386.exe -xyzfile example.xyz -approximation PA -parameters '{"H": 1.23, "C": 2.34, "N": 3.45, "O": 4.56}'
```

Binaries for several platforms (Linux, OSX, Windows) can be found under ```bin/``` .

The number of processes are controlled via the ```-processes``` flag (defaults to 10).
For example, the command
```
goccs_windows_386.exe -xyzfile example.xyz -approximation PA -parameters mobcal -processes 100
```
uses 100 processes in parallel.

Citations for the parameters can be found below:

Mobcal parameters:
```
SM. F. Mesleh, J. M. Hunter, A. A. Shvartsburg, G. C. Schatz, M. F. Jarrold, Structural Information from Ion Mobility Measurements:  Effects of the Long-Range Potential, J. Phys. Chem. 1996, 100, 16082-16086.
J. Phys. Chem. A 1997, 101, 968.
A. A.Shvartsburg, M. F. Jarrold, An exact hard-spheres scattering model for the mobilities of polyatomic ions, Chem. Phys. Lett. 1996, 261, 86-91.
```

Default parameters:
```
C.-K. Siu, Y. Guo, I. S. Saminathan, A. C. Hopkinson, K. M. Siu, Optimization of ion-mobility calculation for conformational analyses, J. Phys. Chem. B, 2010, 114, 1204-1212.
```

## Benchmark
Here a Projection Approximation (PA) benchmark against the established Mobcal software:

![PA Benchmark.](https://github.com/jmwoll/goccs/blob/master/benchmark/benchmark_pa.png)

The PA CCS values exhibit excellent agreement with the literature values. For the
exact hard sphere (EHS) method, we observe slight implementation-dependent
differences to the Mobcal implementation. Generally, values predicted by goccs tend
to lie nearer to the Trajectory Method (TM) values than the Mobcal calculation.

Here an Exact Hard Sphere Scattering (EHS) benchmark against the established Mobcal software:

![PA Benchmark.](https://github.com/jmwoll/goccs/blob/master/benchmark/benchmark_ehs.png)

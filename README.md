
goccs is a parallel cross platform CCS prediction software implemented in go.
CCS values can be easily computed at the command line:
```
goccs -xyzfile example.xyz -approximation PA -parameters mobcal
```
where approximation is either 'PA' or 'EHS'. Custom parameters can be
specified in the JSON format:
```
goccs -xyzfile example.xyz -approximation PA -parameters '{"H": 1.23, "C": 2.34, "N": 3.45, "O": 4.56}'
```

Binaries for several platforms (Linux, OSX, Windows) can be found under ```bin/``` .

The number of processes are controlled via the ```-processes``` flag (defaults to 10).
For example, the command
```
goccs -xyzfile example.xyz -approximation PA -parameters mobcal -processes 100
```
uses 100 processes in parallel.

Citations for the parameters can be found below:

Mobcal parameters:
```
J. Phys. Chem. 1996, 100, 16082-16086;
J. Phys. Chem. A 1997, 101, 968.
Chem. Phys. Lett. 1996, 261, 86-91.
```

Default parameters:
```
J. Phys. Chem. B, Vol. 114, No. 2, 2010.
```

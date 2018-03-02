# Copyright (C) 2017-2018  Jan Wollschläger <janmwoll@gmail.com>
# This file is part of goccs.
#
# goccs is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

import subprocess

from matplotlib import pyplot as plt
from scipy.stats import linregress
import numpy as np


def test_pa_ccs():

	pa_ccs_ref = {
            "methane": 27.499, "ethane": 35.806, "propane": 42.457,
            "butane": 50.114, "pentane": 57.079, "adamantane": 64.799,
            "penguinone": 73.598, "barrelene": 56.733,
            "hirsutene": 88.677, "testosterone": 108.71,
            "paclitaxel": 231.00, "abacavir": 117.27,
            "ciclosporin": 286.18, "codeine": 104.30,
            "talinolol": 165.93, "yangonin": 115.52,
            "vernakalant": 145.82, "octabenzone": 149.15,
            "melatonin": 108.92, "epirubicin": 172.70,
            "alprenolol": 115.17, "trioxifene": 171.34,
            "methylaminoethanol": 53.907,
            "nanokid": 225.19,

	}
	pxs,pys=[],[]
	for name in pa_ccs_ref:
		ref_ccs = pa_ccs_ref[name]
		cmd = "~/go/src/goccs/bin/goccs_linux_386 --xyzfile ~/go/src/goccs/xyz/{}.xyz --parameters mobcal"
		cmd = cmd.format(name)
		ccs = str(subprocess.check_output(cmd, shell=True),'utf-8')
		ccs = ccs.strip().split('\n')[-1]
		ccs = float(ccs)
		#ccs = os.system()
		print('{} (should be {})'.format(ccs,ref_ccs))
		pxs.append(ccs)
		pys.append(ref_ccs)
	slope, intercept, r_value, p_value, std_err = linregress(pxs, pys)
	fit_xs = np.linspace(min(pxs),max(pxs),50)
	fit_ys = [slope*fx+intercept for fx in fit_xs]
	plt.plot(fit_xs,fit_ys,'--',color='black')
	plt.plot(pxs,pys,'ro')
	plt.text(fit_xs[25],fit_ys[25],'R2 = {:.6f}'.format(r_value**2))
	plt.title("PA Benchmark")
	plt.xlabel("CCS (goccs) / A²")
	plt.ylabel("CCS (literature) / A²")
	plt.savefig('benchmark_pa.png')
	plt.show()



if __name__ == '__main__':
	test_pa_ccs()

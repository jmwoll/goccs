
import subprocess


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

	for name in pa_ccs_ref:
		ref_ccs = pa_ccs_ref[name]
		cmd = "~/go/src/goccs/goccs --xyzfile ~/go/src/goccs/xyz/{}.xyz --parameters mobcal"
		cmd = cmd.format(name) 
		ccs = str(subprocess.check_output(cmd, shell=True),'utf-8')
		ccs = ccs.strip().split('\n')[-1]
		ccs = float(ccs)
		#ccs = os.system()
		print('{} (should be {})'.format(ccs,ref_ccs))


if __name__ == '__main__':
	test_pa_ccs()

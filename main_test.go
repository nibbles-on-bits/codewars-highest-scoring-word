package main

import(
	"testing"
)

func TestHighestScoringWord(t *testing.T) {
	arr := [...][2]string{
        {"man i need a taxi up to ubud","taxi"},
        {"what time are we climbing up the volcano","volcano"},
		{"take me to semynak","semynak"},
		{"jhrpuhtfxg xruuf vjlg dyavf xyzbred bxre vids frixc ubdcf zgawkqvegj gzyiypfx vpo fbkkrfsbd qvsohq", "jhrpuhtfxg"},
		{"rkv mbbmcfhcu wvtgqu ecvuwua qdvpqwpg jbtbc gcze mujg htlpzh xovgobf iadzuzrkt czukyyop rryoxd csyhvfpdrw czlwuqf irhy fye fpy gjwijwxa vaz pnekkefqe hlpequl", "csyhvfpdrw"},
		{"smwahanxng bdxxgidh pnouolozh whbhye zslfdicz mlz cljnrnzf oiihhsis aadcht sdlnnm ckztcry dhpvkyrtr", "pnouolozh"},
		{"gugmks bewro rsli rmbo baqw yniydgrt zbferhgyqi nclv wuli gsrc eyzcj", "zbferhgyqi"},
		
	}
    for _,v := range arr {
		got := HighestScoringWord(v[0])
		want := v[1]

		if (got!=want) {
			t.Errorf("\n%v\ngot: %v\nwant: %v", v[0], got, want)
		}
	}
}
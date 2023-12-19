package comparer

import (
	"reflect"
	"sort"
	"testing"
)

func TestComparer_getDiffs(t *testing.T) {
	type args struct {
		m1             map[string]map[string]string
		m2             map[string]map[string]string
		addVersionDiff bool
	}
	type res struct {
		diffPackege1   []Arch
		diffPackegeVer []Arch
	}
	tests := []struct {
		name    string
		args    args
		wantRes res
		wantErr bool
	}{
		{
			name: "Ok1",
			args: args{
				m1: map[string]map[string]string{
					"amd64": {
						"package1": "26.7.9",
						"package2": "0.9.12",
						"package3": "32.1.0",
					},
				},
				m2: map[string]map[string]string{
					"amd64": {
						"package1": "26.6.9",
						"package4": "0.9.12",
						"package5": "32.1.0",
					},
				},
				addVersionDiff: true,
			},
			wantRes: res{
				diffPackege1: []Arch{
					{
						Name:     "amd64",
						Packages: []string{"package2", "package3"},
					},
				},
				diffPackegeVer: []Arch{
					{
						Name:     "amd64",
						Packages: []string{"package1"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Ok2",
			args: args{
				m1: map[string]map[string]string{
					"amd64": {
						"package1": "8.7.9",
						"package2": "0.9.12",
						"package3": "32.1.0",
					},
				},
				m2: map[string]map[string]string{
					"amd64": {
						"package1": "26.6.9",
						"package4": "0.9.12",
						"package5": "32.1.0",
					},
				},
				addVersionDiff: true,
			},
			wantRes: res{
				diffPackege1: []Arch{
					{
						Name:     "amd64",
						Packages: []string{"package2", "package3"},
					},
				},
				diffPackegeVer: make([]Arch, 0),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comp := &Comparer{}
			d1, dv := comp.getDiffs(&tt.args.m1, &tt.args.m2, tt.args.addVersionDiff)
			for _, v := range d1 {
				sort.Strings(v.Packages)
			}
			for _, v := range dv {
				sort.Strings(v.Packages)
			}

			if !reflect.DeepEqual(d1, tt.wantRes.diffPackege1) {
				t.Errorf("getDiffs() return diffsBranch1 = %v, want %v", d1, tt.wantRes.diffPackege1)
			}
			if !reflect.DeepEqual(dv, tt.wantRes.diffPackegeVer) {
				t.Errorf("getDiffs() return diffsVersion = %v, want %v", dv, tt.wantRes.diffPackegeVer)
			}
		})
	}
}

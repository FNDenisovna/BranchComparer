package versioncomparer

import "testing"

func TestVersioncomparer_Compare(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
		wantErr bool
	}{
		{
			name:    "Ok1",
			args:    args{version1: "0:0.9.10-alt3", version2: "0:0.9.10-alt4"},
			wantRes: -1,
			wantErr: false,
		},
		{
			name:    "Ok2",
			args:    args{version1: "0:0.9.1-alt5.qa1", version2: "0:0.9.1-alt5.aa1"},
			wantRes: 1,
			wantErr: false,
		},
		{
			name:    "Ok3",
			args:    args{version1: "1:0.3.0.0.0.1.5e479779-alt1", version2: "1:0.3.4-alt1"},
			wantRes: -1,
			wantErr: false,
		},
		{
			name:    "Ok4",
			args:    args{version1: "0:2.0.0-alt0.3.qa2_10", version2: "0:2.0.0-alt0.3.qa2_10"},
			wantRes: 0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Compare(tt.args.version1, tt.args.version2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compare() return error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if res != tt.wantRes {
				t.Errorf("Compare() return result = %v, want %v", res, tt.wantRes)
			}
		})
	}
}

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
			args:    args{version1: "26.7.9", version2: "26.8.9"},
			wantRes: -1,
			wantErr: false,
		},
		{
			name:    "Ok2",
			args:    args{version1: "0.7.9", version2: "26.8.9"},
			wantRes: -1,
			wantErr: false,
		},
		{
			name:    "Ok3",
			args:    args{version1: "0.8.9", version2: "0.7.9"},
			wantRes: 1,
			wantErr: false,
		},
		{
			name:    "Ok4",
			args:    args{version1: "0.0.1.1", version2: "0.0.1.1"},
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

package branchapi

import (
	"testing"
)

func TestBranchapi_GetPackages(t *testing.T) {
	type args struct {
		branch string
		url    string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
		wantErr   bool
	}{
		{
			name:      "Ok1",
			args:      args{branch: "sisyphus"},
			wantCount: 186260,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := New()
			ps, err := ba.GetPackages(tt.args.branch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPackages() return error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(ps) != tt.wantCount {
				t.Errorf("GetPackages() return count of packages = %v, want %v", len(ps), tt.wantCount)
			}
		})
	}
}

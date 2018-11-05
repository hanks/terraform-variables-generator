package util

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func mockGlobGetFiles(_ string) ([]string, error) {
	files := []string{
		"provider.tf",
		"alb.tf",
		"vpc.tf",
	}
	return files, nil
}

func mockGlobGetNoFiles(_ string) ([]string, error) {
	files := []string{}
	return files, nil
}

func TestGetAllFiles(t *testing.T) {
	type args struct {
		glob func(s string) ([]string, error)
		dir  string
		ext  string
	}
	tests := []struct {
		name      string
		args      args
		wantFiles []string
		wantDir   string
		wantErr   bool
	}{
		{
			name: "Get files from current directory",
			args: args{
				glob: mockGlobGetFiles,
				dir:  "",
				ext:  "*.tf",
			},
			wantFiles: []string{
				"provider.tf",
				"alb.tf",
				"vpc.tf",
			},
			wantDir: "util",
			wantErr: false,
		},
		{
			name: "Get files from specified directory with 0 files",
			args: args{
				glob: mockGlobGetNoFiles,
				dir:  "ad-hoc-dir",
				ext:  "*.tf",
			},
			wantFiles: []string{},
			wantDir:   "ad-hoc-dir",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFiles, gotDir, err := GetAllFiles(tt.args.glob, tt.args.dir, tt.args.ext)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("GetAllFiles() gotFiles = %v, want %v", gotFiles, tt.wantFiles)
			}
			if !strings.Contains(gotDir, tt.wantDir) {
				t.Errorf("GetAllFiles() gotDir = %v, want %v", gotDir, tt.wantDir)
			}
		})
	}
}

func TestFileExists(t *testing.T) {
	dir, _ := os.Getwd()
	curFile := dir + "/" + "file_util.go"

	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Check file existed",
			args: args{
				name: curFile,
			},
			want: true,
		},
		{
			name: "Check file not existed",
			args: args{
				name: "ad-hoc-file",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExists(tt.args.name); got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

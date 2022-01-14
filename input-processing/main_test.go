package main

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func Test_readLargeInput(t *testing.T) {
	f, err := os.OpenFile("test.txt", os.O_RDONLY, 0755)
	if err != nil {
		t.Error(err)
		return
	}

	type args struct {
		rd        io.Reader
		filterStr string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				rd:        f,
				filterStr: "error",
			},
			wantResult: []string{
				"this is test file 1 line with error string",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := readLargeInput(tt.args.rd, tt.args.filterStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("readLargeInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("readLargeInput() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

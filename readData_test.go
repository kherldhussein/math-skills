package main

import (
	"reflect"
	"testing"
)

func TestReadData(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadData(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadData() = %v, want %v", got, tt.want)
			}
		})
	}
}

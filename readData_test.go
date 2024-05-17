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
		name     string
		args     args
		expected []float64
	}{
		{name: "Reading", args: args{path: "data.txt"}, expected: []float64{189, 113, 121, 114, 145, 110}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadData(tt.args.path); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ReadData() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

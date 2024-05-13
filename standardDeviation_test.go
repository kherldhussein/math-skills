package main

import "testing"

func TestStandardDiviation(t *testing.T) {
	type args struct {
		x []float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Known SD", args: args{[]float64{1.0, 2.0, 3.0}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardDiviation(tt.args.x); got != tt.want {
				t.Errorf("StandardDiviation() = %v, want %v", got, tt.want)
			}
		})
	}
}

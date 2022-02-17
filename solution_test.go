package square

import (
	"math"
	"testing"
)

func TestCalcSquare(t *testing.T) {
	type args struct {
		sideLen  float64
		sidesNum anotherInt
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"square", args{sideLen: 10.0, sidesNum: 4}, 100.0},
		{"triangle", args{sideLen: 10.0, sidesNum: 3}, 86.60},
		{"circle", args{sideLen: 10.0, sidesNum: 0}, 314.16},
		{"wrong sides", args{sideLen: 10.0, sidesNum: 5}, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcSquare(tt.args.sideLen, tt.args.sidesNum); math.Round(got*100)/100 != tt.want {
				t.Errorf("CalcSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

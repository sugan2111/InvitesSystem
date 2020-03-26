package services

import (
	"fmt"
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	distance := CalculateDistance(53.339428, -6.257664,52.986375,-6.043701)
	if distance != 41.76872550083617 {
		t.Errorf("Distance was incorrect, got: %f, want: %f.", distance, 41.76872550083617)
	}
}

func TestCalculateDistanceTableDriven(t *testing.T) {
	var tests = []struct {
		a, b,c, d float64
		want float64
	}{
		{53.339428, -6.257664, 51.92893, -10.27699,313.2556337814159},
		{53.339428, -6.257664, 51.8856167,-10.4240951,324.3749120082729},
		{53.339428, -6.257664, 52.3191841,-8.5072391, 188.95936393870804},
		{53.339428, -6.257664, 53.807778, -7.714444,109.3764554298563},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%f,%f,%f,%f", tt.a, tt.b,tt.c,tt.d)
		t.Run(testname, func(t *testing.T) {
			ans := CalculateDistance(tt.a, tt.b, tt.c, tt.d)
			if ans != tt.want {
				t.Errorf("got %f, want %f", ans, tt.want)
			}
		})
	}
}


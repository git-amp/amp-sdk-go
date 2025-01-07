package hexgrid

import (
	"fmt"
	"math"
	"testing"
)

var defaultLayout = Layout{
	Size:        Point{100, 100},
	Origin:      Point{0, 0},
	Orientation: OrientationFlat}

// utility functions
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func TestHexToPixel(t *testing.T) {

	var testCases = []struct {
		hexA     Hex
		expected string
	}{
		{NewHex(0, 0), "0.0;0.0"},
		{NewHex(1, 0), "150.0;86.6"},
		{NewHex(1, -1), "150.0;-86.6"},
		{NewHex(2, -1), "300.0;0.0"},
	}

	for _, tt := range testCases {
		pixel := HexToPixel(defaultLayout, tt.hexA)
		actual := fmt.Sprintf("%.1f;%.1f", pixel.x, pixel.y)
		if actual != tt.expected {
			t.Error("Expected:", tt.expected, "got:", actual)
		}
	}
}

func TestPixelToHex(t *testing.T) {

	var testCases = []struct {
		point    Point
		expected Hex
	}{
		{Point{0, 0}, NewHex(0, 0)},
		{Point{150, 87}, NewHex(1, 0)},
		{Point{300, 10}, NewHex(2, -1)},
	}

	for _, tt := range testCases {
		actual := PixelToHex(defaultLayout, tt.point).Round()
		if actual != tt.expected {
			t.Error("Expected:", tt.expected, "got:", actual)
		}
	}
}

/*
Tests the expected corners of the hexagon, starting at the East vertex and proceeding in CCW order.

	          50     100     50
	         |---|---|---|---|

	 (-50;-86.6) +*******+ (50;-86.6)
	            *         *
	           *           *
	          *    (0,0)    *
	(-100;0) +       +       + (100;0)
	          *             *
	           *           *
	            *         *
	  (-50;86.6) +*******+ (50;86.6)
*/
func TestHexagonCorners(t *testing.T) {
	corners := HexagonCorners(defaultLayout, NewHex(0, 0))
	if len(corners) != 6 {
		t.Error("Invalid length:", len(corners))
	}

	testCase := []struct {
		roundedX float64
		roundedY float64
	}{
		{100, 0},
		{50, -86.6},
		{-50, -86.6},
		{-100, 0},
		{-50, 86.6},
		{50, 86.6},
	}

	for i := 0; i < len(corners); i++ {
		actualX := toFixed(corners[i].x, 1)
		actualY := toFixed(corners[i].y, 1)
		expectedX := testCase[i].roundedX
		expectedY := testCase[i].roundedY
		if actualX != expectedX || actualY != expectedY {
			t.Errorf("Expected: (%.1f,%.1f) got: (%.1f,%.1f)", expectedX, expectedY, actualX, actualY)
		}
	}

}
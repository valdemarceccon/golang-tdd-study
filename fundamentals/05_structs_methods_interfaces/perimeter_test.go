package perimeter_test

import (
	perimeter "github.com/valdemarceccon/golang-tdd-study/fundamentals/05_structs_methods_interfaces"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := perimeter.Rectangle{Width: 10.0, Height: 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		name  string
		shape perimeter.Shape
		want  float64
	}{
		{name: "rectangles", shape: perimeter.Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "circles", shape: perimeter.Circle{Radius: 10}, want: 314.1592653589793},
		{name: "triangles", shape: perimeter.Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %.2f want %.2f", got, tt.want)
			}
		})
	}
}

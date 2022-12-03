package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("expected %.2f got %.2f", got, expected)
	}

}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rec := Rectangle{10.0, 10.0}
		checkArea(t, rec, 100.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestAreaV2(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{{Circle{10}, 314.1592653589793}, {Rectangle{12, 6}, 72.0}, {Triangle{12, 6}, 36.0}}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}

func TestAreaV3(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Triange", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

}

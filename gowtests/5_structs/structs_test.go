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
	t.Run("rectangle", func(t *testing.T) {
		rec := Rectangle{10.0, 10.0}
		got := rec.Area()
		expected := 100.0
		if got != expected {
			t.Errorf("expected %g got %g", got, expected) //%g is used as it's more precise as compared to f
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		expected := 314.1592653589793

		if got != expected {
			t.Errorf("expected %g got %g", got, expected)
		}
	})

}

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
	rec := Rectangle{10.0, 10.0}
	got := Area(rec)
	expected := 100.0

	if got != expected {
		t.Errorf("expected %.2f got %.2f", got, expected)
	}

}

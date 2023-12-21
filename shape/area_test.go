package shape

import "testing"

func TestArea(t *testing.T) {
	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if want != got {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}

	// }

	// t.Run("Rectangle", func(t *testing.T) {
	// 	rectangle := Rectangle{12, 6}
	// 	checkArea(t, rectangle, 72.0)
	// })

	// t.Run("Circle", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle, 314.1592653589793)
	// })

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangke", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}

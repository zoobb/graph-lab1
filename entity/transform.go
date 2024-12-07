package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func translate(v *rl.Vector2, d rl.Vector2) {
	matrix := rl.Matrix{
		M0: 1, M4: 0, M8: 0, M12: d.X,
		M1: 0, M5: 1, M9: 0, M13: d.Y,
		M2: 0, M6: 0, M10: 1, M14: 0,
		M3: 0, M7: 0, M11: 0, M15: 1,
	}
	*v = rl.Vector2Transform(*v, matrix)
}
func rotate(v *rl.Vector2, angle float32, pivot rl.Vector2) {
	cosAngle := float32(math.Cos(float64(angle)))
	sinAngle := float32(math.Sin(float64(angle)))

	translate(v, rl.Vector2{X: -pivot.X, Y: -pivot.Y})
	matrix := rl.Matrix{
		M0: cosAngle, M4: -sinAngle, M8: 0, M12: 0,
		M1: sinAngle, M5: cosAngle, M9: 0, M13: 0,
		M2: 0, M6: 0, M10: 1, M14: 0,
		M3: 0, M7: 0, M11: 0, M15: 1,
	}
	*v = rl.Vector2Transform(*v, matrix)
	translate(v, pivot)
}
func scale(v *rl.Vector2, s rl.Vector2) {
	matrix := rl.Matrix{
		M0: s.X, M4: 0, M8: 0, M12: 0,
		M1: 0, M5: s.Y, M9: 0, M13: 0,
		M2: 0, M6: 0, M10: 1, M14: 0,
		M3: 0, M7: 0, M11: 0, M15: 1,
	}

	*v = rl.Vector2Transform(*v, matrix)
}

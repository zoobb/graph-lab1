package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"math"
	"strconv"
)

type Entity struct {
	Body         entityTexture
	Acceleration float32
	Friction     float32
	Color        color.RGBA
	pos          rl.Vector2
	speed        float32
	velocity     rl.Vector2
	angle        float32
	angularSpeed float32
	angularAcc   float32
	initPosition rl.Vector2
	info         Info
	testPos      rl.Vector2
	testAcc      rl.Vector2
	DeltaTime    float32
}

type entityTexture struct {
	width     float32
	height    float32
	peak      rl.Vector2
	leftWing  rl.Vector2
	rightWing rl.Vector2
	direction rl.Vector2
}

type Info struct {
	FontSize  int32
	FontColor color.RGBA
	Gap       int32
	Pos       rl.Vector2
	temp      map[string]float32
}

func New(w float32, h float32, acceleration float32, angularAcc float32, friction float32, color color.RGBA, initPos rl.Vector2, i Info) *Entity {
	return &Entity{
		Body: entityTexture{
			width:  w,
			height: h,
		},
		Acceleration: acceleration,
		Friction:     1 - float32(math.Pow(float64(friction), 2)),
		Color:        color,
		initPosition: rl.Vector2{
			X: initPos.X,
			Y: initPos.Y,
		},
		pos: rl.Vector2{
			X: initPos.X,
			Y: initPos.Y,
		},
		angularAcc: angularAcc,
		info:       i,
		testPos:    initPos,
	}
}

func (e *Entity) Draw() {
	e.render()
}

func (e *Entity) DrawInfo() {
	e.info.temp = make(map[string]float32)
	names := []string{
		"vel_x",
		"vel_y",
		"vel",
		"ang_deg",
		"ang_spg",
	}

	e.info.temp = map[string]float32{
		names[0]: e.velocity.X,
		names[1]: e.velocity.Y,
		names[2]: float32(e.getVelocity()),
		names[3]: e.angle * rl.Rad2deg,
		names[4]: e.angularSpeed * rl.Rad2deg,
	}

	rl.DrawFPS(int32(e.info.Pos.X), int32(e.info.Pos.Y))

	for index, key := range names {
		value := e.info.temp[key]
		rl.DrawText(key+": "+strconv.FormatFloat(float64(value), 'f', 4, 32), int32(e.info.Pos.X), int32(2+index)*e.info.FontSize+e.info.Gap, e.info.FontSize, e.info.FontColor)
	}
}

func (e *Entity) Control() {
	if rl.IsKeyDown(rl.KeyA) {
		e.angularSpeed -= e.angularAcc * e.DeltaTime
	} else if rl.IsKeyDown(rl.KeyD) {
		e.angularSpeed += e.angularAcc * e.DeltaTime
	} else {
		e.angularSpeed *= e.Friction
	}
	if rl.IsKeyDown(rl.KeyW) {
		e.speed += e.Acceleration * e.DeltaTime
	} else if rl.IsKeyDown(rl.KeyS) {
		e.speed -= e.Acceleration * e.DeltaTime
	} else {
		e.speed *= e.Friction
	}

	if rl.IsKeyDown(rl.KeyC) {
		e.pos = e.initPosition
		e.stop()
	}
	if rl.IsKeyDown(rl.KeyV) {
		e.stop()
	}

	e.moveLinear()
	e.moveAngular()
}
func (e *Entity) render() {
	e.Body.peak = rl.Vector2{
		X: e.pos.X,
		Y: e.pos.Y - e.Body.height/2,
	}
	e.Body.leftWing = rl.Vector2{
		X: e.pos.X - e.Body.width/2,
		Y: e.pos.Y + e.Body.height/2,
	}
	e.Body.rightWing = rl.Vector2{
		X: e.pos.X + e.Body.width/2,
		Y: e.pos.Y + e.Body.height/2,
	}
	e.Body.direction = rl.Vector2{
		X: e.Body.peak.X - e.pos.X,
		Y: e.Body.peak.Y - e.pos.Y,
	}

	rotate(&e.Body.peak, e.angle, e.pos)
	rotate(&e.Body.leftWing, e.angle, e.pos)
	rotate(&e.Body.rightWing, e.angle, e.pos)

	rl.DrawTriangle(e.Body.peak, e.Body.leftWing, e.Body.rightWing, e.Color)
	rl.DrawCircleV(e.pos, 4, rl.Red)
	rl.DrawLineV(e.pos, e.Body.peak, rl.Blue)
}

func (e *Entity) stop() {
	e.speed = 0
	e.angularSpeed = 0
}
func (e *Entity) moveLinear() {
	e.Body.direction = rl.Vector2{
		X: e.Body.peak.X - e.pos.X,
		Y: e.Body.peak.Y - e.pos.Y,
	}
	norm := rl.Vector2Normalize(e.Body.direction)

	e.velocity.X = norm.X * e.speed * e.DeltaTime
	e.velocity.Y = norm.Y * e.speed * e.DeltaTime

	translate(&e.pos, e.velocity)
}

func (e *Entity) moveAngular() {
	e.angle += e.angularSpeed * e.DeltaTime
}

func (e *Entity) getVelocity() float64 {
	return math.Sqrt(math.Pow(float64(e.velocity.X), 2) + math.Pow(float64(e.velocity.Y), 2))
}

package astrophysics

import (
	"Go-Astrophysics/customVector"
	"fmt"
)

type CelestialBody struct {
	position     customVector.Vector2
	velocity     customVector.Vector2
	acceleration customVector.Vector2
	radius       float64
	mass         float64

	collision_resistance_vec customVector.Vector2
	temp_vec                 customVector.Vector2
}

func CreateBody(p customVector.Vector2, m float64, rad float64) *CelestialBody {
	return &CelestialBody{
		position: p,
		radius:   rad,
		mass:     m,
	}
}

func (body *CelestialBody) AddVecToTemp(n_vec customVector.Vector2) customVector.Vector2 {
	body.temp_vec = customVector.VectorAdd(body.temp_vec, n_vec)
	return body.temp_vec
}

func (body *CelestialBody) GetTempVec() customVector.Vector2 {
	return body.temp_vec
}

func (body *CelestialBody) AddVecToColl(n_vec customVector.Vector2) customVector.Vector2 {
	body.collision_resistance_vec = customVector.VectorAdd(body.collision_resistance_vec, n_vec)
	return body.collision_resistance_vec
}

func (body *CelestialBody) GetCollVec() customVector.Vector2 {
	return body.collision_resistance_vec
}

func (body *CelestialBody) MovePosAbs(n_pos customVector.Vector2) customVector.Vector2 {
	body.position = n_pos
	return body.position
}

func (body *CelestialBody) MovePosDelta(delta customVector.Vector2) customVector.Vector2 {
	body.position.X += delta.X
	body.position.Y += delta.Y
	return body.position
}

func (body *CelestialBody) SetVelocity(n_vel customVector.Vector2) {
	body.velocity = n_vel
}

func (body *CelestialBody) GetVelocity() customVector.Vector2 {
	return body.velocity
}

func (body *CelestialBody) SetAcceleration(n_acc customVector.Vector2) {
	body.acceleration = n_acc
}

func (body *CelestialBody) GetAcceleration() customVector.Vector2 {
	return body.acceleration
}

func (body CelestialBody) GetPos() customVector.Vector2 {
	return body.position
}

func (body *CelestialBody) SetMass(m float64) {
	body.mass = m
}

func (body *CelestialBody) GetMass() float64 {
	return body.mass
}

func (body *CelestialBody) SetRadius(r float64) {
	body.radius = r
}

func (body *CelestialBody) GerRadius() float64 {
	return body.radius
}

func (body *CelestialBody) PrintInfo() {
	fmt.Println("V: ", body.velocity, "\na: ", body.acceleration, "\nm: ", body.mass)
}

func (body *CelestialBody) Update() {
	body.velocity = customVector.VectorAdd(
		body.velocity,
		body.acceleration.GetMultiplied(1./60.),
	)
	body.position = customVector.VectorAdd(
		body.position,
		body.velocity,
	)
	body.temp_vec = customVector.Vector2{0, 0}
	body.collision_resistance_vec = customVector.Vector2{0, 0}
}

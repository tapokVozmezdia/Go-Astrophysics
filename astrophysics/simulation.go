package astrophysics

import (
	"Go-Astrophysics/customVector"
	"fmt"
)

type Simulation struct {
	Bodies []CelestialBody
}

func (sim *Simulation) CreateBody(pos customVector.Vector2, m float64) {
	sim.Bodies = append(sim.Bodies, *CreateBody(pos, m, 15))
	sim.Bodies[len(sim.Bodies)-1].SetAcceleration(customVector.Vector2{0, 0})
}

func (sim *Simulation) gravityComputation(ch *bool) {
	for i := 0; i < len(sim.Bodies); i++ {
		for j := i; j < len(sim.Bodies); j++ {
			if i == j {
				continue
			}
			tmp_v_diff := customVector.VectorDiff(sim.Bodies[i].GetPos(), sim.Bodies[j].GetPos())
			tmp_distance := tmp_v_diff.GetLen()

			tmp_v_diff.Normalize()
			acceleration_module := GRAVITATIONAL_CONSTANT * sim.Bodies[j].GetMass() / (tmp_distance * tmp_distance)
			tmp_v_diff.MultLen(acceleration_module)
			sim.Bodies[i].AddVecToTemp(tmp_v_diff)
			tmp_v_diff = tmp_v_diff.GetReversed()
			tmp_v_diff.Normalize()
			acceleration_module = GRAVITATIONAL_CONSTANT * sim.Bodies[i].GetMass() / (tmp_distance * tmp_distance)
			tmp_v_diff.MultLen(acceleration_module)
			sim.Bodies[j].AddVecToTemp(tmp_v_diff)
		}
	}
	*ch = true
}

func (sim *Simulation) collisionComputation(ch *bool) {
	for i := 0; i < len(sim.Bodies); i++ {
		for j := i; j < len(sim.Bodies); j++ {
			if i == j {
				continue
			}
			tmp_v_diff := customVector.VectorDiff(sim.Bodies[i].GetPos(), sim.Bodies[j].GetPos())
			distance := tmp_v_diff.GetLen()
			if distance < sim.Bodies[i].GerRadius()+sim.Bodies[j].GerRadius() {

				// fmt.Println("V1: ", sim.Bodies[i].GetVelocity(), "\nV2: ", sim.Bodies[j].GetVelocity())

				T_Vec := customVector.VectorDiff(sim.Bodies[i].GetVelocity(), sim.Bodies[j].GetVelocity())

				// fmt.Println("T_Vec: ", T_Vec)
				T_Vec.MultLen((sim.Bodies[i].GetMass() + sim.Bodies[j].GetMass()) / 2.)

				tmp_dist := sim.Bodies[i].GerRadius() + sim.Bodies[j].GerRadius() - distance
				tmp_v_diff.Normalize()
				sim.Bodies[j].MovePosDelta(tmp_v_diff.GetMultiplied(tmp_dist / 2))
				tmp_v_diff.Reverse()
				sim.Bodies[i].MovePosDelta(tmp_v_diff.GetMultiplied(tmp_dist / 2))

				fmt.Println("Body 1 (before): \n")
				sim.Bodies[i].PrintInfo()

				sim.Bodies[i].SetVelocity(
					customVector.VectorAdd(
						sim.Bodies[i].GetVelocity(),
						T_Vec.GetMultiplied(1./sim.Bodies[i].GetMass()),
					),
				)
				fmt.Println("Body 1 (after): \n")
				sim.Bodies[i].PrintInfo()

				fmt.Println("Body 2 (before): \n")
				sim.Bodies[j].PrintInfo()

				sim.Bodies[j].SetVelocity(
					customVector.VectorAdd(
						sim.Bodies[j].GetVelocity(),
						T_Vec.GetMultiplied(-1./sim.Bodies[j].GetMass()),
					),
				)

				fmt.Println("Body 2 (after): \n")
				sim.Bodies[j].PrintInfo()
			}
		}
	}
	*ch = true
}

func (sim *Simulation) UpdateAll() {

	ch1 := false
	ch2 := false

	go sim.gravityComputation(&ch1)

	go sim.collisionComputation(&ch2)

	for {
		if ch1 && ch2 {
			break
		}
	}

	for i := 0; i < len(sim.Bodies); i++ {
		sim.Bodies[i].SetAcceleration(sim.Bodies[i].GetTempVec())

		sim.Bodies[i].Update()
	}
}

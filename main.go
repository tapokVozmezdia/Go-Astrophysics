package main

import (
	"Go-Astrophysics/astrophysics"
	"Go-Astrophysics/customVector"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/imdraw"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func DrawBodies(sim *(astrophysics.Simulation), imd *imdraw.IMDraw) {
	for _, body := range sim.Bodies {
		imd.Push(pixel.V(body.GetPos().X, body.GetPos().Y))
		imd.Circle(body.GerRadius(), 1.5)
	}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Physics",
		Bounds: pixel.R(0, 0, 960, 540),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	imd.Color = colornames.Crimson

	simulation := astrophysics.Simulation{}

	paused := false

	for !win.Closed() {
		win.Clear(pixel.RGB(0, 0, 0.1))

		if win.JustPressed(pixelgl.MouseButtonRight) {
			paused = !paused
		}

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			imd.Clear()
			simulation.CreateBody(
				customVector.Vector2(win.MousePosition()),
				5,
			)
		}

		DrawBodies(&simulation, imd)

		imd.Draw(win)
		imd.Clear()
		if !paused {
			simulation.UpdateAll()
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

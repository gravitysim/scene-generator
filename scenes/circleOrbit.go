package scenes

import (
	gravMath "github.com/gravitysim/gravity-commons/math"
	"github.com/gravitysim/gravity-commons/phys"
	"github.com/gravitysim/gravity-commons/scene"
	"math"
)

func CreateCircleOrbitScene() *scene.GravityScene {
	centralBody := phys.NewBody("centralBody", 5e24, 6400e3, gravMath.Point{}, gravMath.Vector{})
	satBody := phys.NewBody("satBody", 1, 1, gravMath.Point{X: 6500e3, Y: 0}, gravMath.Vector{})
	satBody.Velocity.Y = circleVelocity(centralBody.GetMass(), satBody.Position.X)

	gravScene := scene.GravityScene{}
	gravScene.AddBody(&centralBody)
	gravScene.AddBody(&satBody)

	return &gravScene
}

func circleVelocity(mass, dst float64) float64 {
	return math.Sqrt(phys.G * mass / dst)
}

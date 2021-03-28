package scenes

import (
	gravMath "github.com/gravitysim/gravity-commons/math"
	"github.com/gravitysim/gravity-commons/phys"
	"github.com/gravitysim/gravity-commons/scene"
	"math"
)

func CreateSimpleTwoBodiesCollisionScene() *scene.GravityScene {
	sqrt2 := math.Sqrt(2)
	movingBody := phys.NewBody("movingBody", 1e-5, 1, gravMath.Point{X: -10 / sqrt2, Y: -10 / sqrt2}, gravMath.Vector{X: sqrt2, Y: sqrt2})
	staticBody := phys.NewBody("staticBody", 1e-5, 1, gravMath.Point{}, gravMath.Vector{})

	gravScene := scene.GravityScene{}
	gravScene.AddBody(&movingBody)
	gravScene.AddBody(&staticBody)
	gravScene.LinkBodies(0, 1)

	return &gravScene
}

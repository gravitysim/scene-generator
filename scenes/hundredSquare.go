package scenes

import (
	"fmt"
	gravMath "github.com/gravitysim/gravity-commons/math"
	"github.com/gravitysim/gravity-commons/phys"
	"github.com/gravitysim/gravity-commons/scene"
)

func CreateHundredSquareScene() *scene.GravityScene {
	gravScene := scene.GravityScene{}

	const count = 10
	const mass float64 = 1e4
	const radius float64 = 10
	const padding float64 = 0.1
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			x := float64(i) * (2*radius + padding)
			y := float64(j) * (2*radius + padding)
			body := phys.NewBody(fmt.Sprintf("b_%d_%d", i, j), mass, radius, gravMath.Point{X: x, Y: y}, gravMath.Vector{})
			gravScene.AddBody(&body)
		}
	}

	bodies := gravScene.GetBodies()
	for i := 0; i < len(bodies); i++ {
		for j := i + 1; j < len(bodies); j++ {
			b1Num := uint64(gravScene.GetBodyNum(bodies[i]))
			b2Num := uint64(gravScene.GetBodyNum(bodies[j]))
			gravScene.LinkBodies(b1Num, b2Num)
		}
	}

	return &gravScene
}

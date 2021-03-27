package main

import (
	"github.com/gravitysim/gravity-commons/scene/scene_file"
	"github.com/gravitysim/scene-generator/scenes"
)

func main() {
	gravScene := scenes.CreateCircleOrbitScene()

	err := scene_file.SaveScene(gravScene, "build/circleOrbit.scene")
	if err != nil {
		panic(err)
	}
}

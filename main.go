package main

import "github.com/gravitysim/gravity-commons/scene/scene_file"

func main() {
	gravScene := createCircleOrbitScene()

	err := scene_file.SaveScene(gravScene, "scenes/circleOrbit.scene")
	if err != nil {
		panic(err)
	}
}

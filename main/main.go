package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/gravitysim/gravity-commons/scene"
	"github.com/gravitysim/gravity-commons/scene/scene_file"
	"github.com/gravitysim/scene-generator/scenes"
)

type CircleOrbitCmd struct {
}

type SimpleTwoBodiesCollisionCmd struct {
}

type HundredSquareCmd struct {
}

type Args struct {
	CircleOrbit              *CircleOrbitCmd              `arg:"subcommand:circle-orbit"`
	SimpleTwoBodiesCollision *SimpleTwoBodiesCollisionCmd `arg:"subcommand:simple-two-bodies"`
	HundredSquare            *HundredSquareCmd            `arg:"subcommand:hundred-square"`
}

var args Args

func main() {
	arg.MustParse(&args)

	gravScene, fileName := createScene()
	if gravScene != nil {
		err := scene_file.SaveScene(gravScene, fmt.Sprintf("build/%s", fileName))
		if err != nil {
			panic(err)
		}
	}
}

func createScene() (*scene.GravityScene, string) {
	switch {
	case args.CircleOrbit != nil:
		return scenes.CreateCircleOrbitScene(), "circleOrbit.scene"
	case args.SimpleTwoBodiesCollision != nil:
		return scenes.CreateSimpleTwoBodiesCollisionScene(), "simpleTwoBodiesCollision.scene"
	case args.HundredSquare != nil:
		return scenes.CreateHundredSquareScene(), "hundredSquare.scene"
	}
	return nil, ""
}

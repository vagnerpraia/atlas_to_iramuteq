package control

import (
	"github.com/vagnerpraia/atlas_to_iramuteq/model"
	"github.com/vagnerpraia/atlas_to_iramuteq/service/about"
	"github.com/vagnerpraia/atlas_to_iramuteq/service/convertFile"
	"github.com/vagnerpraia/atlas_to_iramuteq/service/help"
	"github.com/vagnerpraia/atlas_to_iramuteq/service/version"
)

var Config model.Config

func Router(service string) bool {
	result := false

	switch service {
	case "--about", "-a":
		result = about.Execute()
	case "--help", "-h":
		result = help.Execute()
	case "--version", "-v":
		result = version.Execute()
	default:
		convertFile.Config = Config
		result = convertFile.Execute()
	}

	return result
}
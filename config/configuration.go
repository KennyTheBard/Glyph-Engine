package config

type Configuration struct {
	StartingStoryID uint `json:"startingStoryID"`
	Port            uint `json:"port"`
}

var GlobalConfiguration Configuration

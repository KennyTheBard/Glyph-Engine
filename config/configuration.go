package config

type Configuration struct {
	StartingStoryID uint   `json:"startingStoryID"`
	DefaultPort     string `json:"defaultPort"`
}

var GlobalConfiguration Configuration

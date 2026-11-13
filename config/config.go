package config

import "github.com/safkhamezadeh/airhockey/geometry"

type PuckConfig struct {
	SpawnPosition geometry.Vector2
	Scale         float64
}

type Config struct {
	ScreenSize geometry.Size
	Puck       PuckConfig
}

var defaultConfig = Config{
	ScreenSize: geometry.Size{Width: 640, Height: 800},
	Puck: PuckConfig{
		SpawnPosition: geometry.Vector2{X: 320, Y: 400},
		Scale:         0.1,
	},
}

func Default() Config {
	return defaultConfig
}

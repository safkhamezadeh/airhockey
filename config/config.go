package config

import "github.com/safkhamezadeh/airhockey/geometry"

type GameObjectConfig struct {
	SpawnPosition geometry.Vector2
	Scale         float64
}

type PuckConfig struct {
	GameObjectConfig
}

type MalletConfig struct {
	GameObjectConfig
}

type Config struct {
	ScreenSize geometry.Size
	Puck       PuckConfig
	Mallet     MalletConfig
	AiMallet   MalletConfig
}

var defaultConfig = Config{
	ScreenSize: geometry.Size{Width: 640, Height: 800},
	Puck: PuckConfig{
		GameObjectConfig: GameObjectConfig{
			SpawnPosition: geometry.Vector2{X: 320, Y: 400},
			Scale:         0.1,
		},
	},
	Mallet: MalletConfig{
		GameObjectConfig: GameObjectConfig{
			SpawnPosition: geometry.Vector2{X: 320, Y: 700},
			Scale:         1.0,
		},
	},
	AiMallet: MalletConfig{
		GameObjectConfig: GameObjectConfig{
			SpawnPosition: geometry.Vector2{X: 320, Y: 150},
			Scale:         1.0,
		},
	},
}

func Default() Config {
	return defaultConfig
}

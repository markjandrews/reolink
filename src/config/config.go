package config

type Camera struct {
	Uid string
}

type Config struct {
	Cameras map[string]Camera
}

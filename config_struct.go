package main

type Config struct {
	// LogLevel string `hcl:"log_level"`
	Bar Bar `hcl:"bar,block"`
}

type Bar struct {
	Name  string       `hcl:"name"`
	Clock *ClockConfig `hcl:"clock,block"`
}

type Image struct {
	Path string `hcl:"path"`
}

type DynamicImage struct {
	Ranges map[int]string `hcl:"ranges"`
}

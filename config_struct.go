package main

type Config struct {
	// LogLevel string `hcl:"log_level"`
	Bar Bar `hcl:"bar,block"`
}

type Bar struct {
	Name        string            `hcl:"name"`
	Clock       *ClockConfig      `hcl:"clock,block"`
	Battery     *BatteryConfig    `hcl:"battery,block"`
	CPU         *CPUConfig        `hcl:"cpu,block"`
	Memory      *MemConfig        `hcl:"memory,block"`
	Network     *NetworkConfig    `hcl:"network,block"`
	PulseAudio  *PulseAudioConfig `hcl:"pulseaudio,block"`
	Temperature *TempConfig       `hcl:"temperature,block"`
}

type Image struct {
	Path string `hcl:"path"`
}

type DynamicImage struct {
	Ranges map[int]string `hcl:"ranges"`
}

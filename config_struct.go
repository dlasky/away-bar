package main

import "dlasky/away-bar/modules"

type Config struct {
	// LogLevel string `hcl:"log_level"`
	Bar    *BarConfig    `hcl:"bar,block"`
	Module *ModuleConfig `hcl:"module,block"`
}

type ModuleConfig struct {
	Name        string                    `hcl:"name"`
	Clock       *modules.ClockConfig      `hcl:"clock,block"`
	Battery     *modules.BatteryConfig    `hcl:"battery,block"`
	CPU         *modules.CPUConfig        `hcl:"cpu,block"`
	Memory      *modules.MemConfig        `hcl:"memory,block"`
	Network     *modules.NetworkConfig    `hcl:"network,block"`
	PulseAudio  *modules.PulseAudioConfig `hcl:"pulseaudio,block"`
	Temperature *modules.TempConfig       `hcl:"temperature,block"`
	Icon        *modules.IconConfig       `hcl:"icon,block"`
	Image       *modules.ImageConfig      `hcl:"image,block"`
}

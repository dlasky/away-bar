package main

type ConfigData struct {
	Bar     Bar       `json:"bar"`
	Modules []Modules `json:"modules"`
}
type Margins struct {
	Left  int `json:"left"`
	Right int `json:"right"`
}
type Bar struct {
	Position string  `json:"position"`
	Margins  Margins `json:"margins"`
	Display  string  `json:"display"`
}
type Modules struct {
	Name             string `json:"name"`
	Type             string `json:"type"`
	Applicationicons bool   `json:"applicationIcons,omitempty"`
	Output           string `json:"output,omitempty"`
	Icon             string `json:"icon,omitempty"`
}

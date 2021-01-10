package main

import "image/color"

var palNames = []string{
	"Commodore 64",
	"Pico-8",
	"Original Game Boy",
}

var palettes = [][]color.Color{
	// C64
	[]color.Color{
		color.RGBA{0, 0, 0, 255},       // black
		color.RGBA{255, 255, 255, 255}, // white
		color.RGBA{136, 0, 0, 255},     // red
		color.RGBA{170, 255, 238, 255}, // cyan
		color.RGBA{204, 68, 204, 255},  // purple
		color.RGBA{0, 204, 85, 255},    // green
		color.RGBA{0, 0, 170, 255},     // blue
		color.RGBA{238, 238, 119, 255}, // yellow
		color.RGBA{221, 136, 85, 255},  // orange
		color.RGBA{102, 68, 0, 255},    // brown
		color.RGBA{255, 119, 119, 255}, // light red
		color.RGBA{51, 51, 51, 255},    // dark grey
		color.RGBA{119, 119, 119, 255}, // grey
		color.RGBA{170, 255, 102, 255}, // light green
		color.RGBA{0, 136, 255, 255},   // light blue
		color.RGBA{187, 187, 187, 255}, // light grey
	},

	// Pico-8
	[]color.Color{
		color.RGBA{0, 0, 0, 255},       // black
		color.RGBA{29, 43, 83, 255},    // dark blue
		color.RGBA{126, 37, 83, 255},   // dark purple
		color.RGBA{0, 135, 81, 255},    // dark green
		color.RGBA{171, 82, 54, 255},   // brown
		color.RGBA{95, 87, 79, 255},    // dark grey
		color.RGBA{194, 195, 199, 255}, // light grey
		color.RGBA{255, 241, 232, 255}, // white
		color.RGBA{255, 0, 77, 255},    // red
		color.RGBA{255, 163, 0, 255},   // orange
		color.RGBA{255, 236, 39, 255},  // yellow
		color.RGBA{0, 228, 54, 255},    // green
		color.RGBA{41, 173, 255, 255},  // blue
		color.RGBA{131, 118, 156, 255}, // lavender
		color.RGBA{255, 119, 168, 255}, // pink
		color.RGBA{255, 204, 170, 255}, // light-peach
	},

	// Original Game Boy (green)
	[]color.Color{
		color.RGBA{15, 56, 15, 255},   // darkest
		color.RGBA{48, 98, 48, 255},   // dark
		color.RGBA{139, 172, 15, 255}, // light
		color.RGBA{155, 188, 15, 255}, // lightest
	},
}

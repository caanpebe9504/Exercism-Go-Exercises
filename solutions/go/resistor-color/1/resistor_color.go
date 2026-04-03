package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	//panic("Please implement the Colors function")

    resistorColors := []string{
        "black",
        "brown",
        "red",
        "orange",
        "yellow",
        "green",
        "blue",
        "violet",
        "grey",
        "white",
    }
    return resistorColors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	//panic("Please implement the ColorCode function")

    resistors := map[string]int{
        "black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
    }
	return resistors[color]
}

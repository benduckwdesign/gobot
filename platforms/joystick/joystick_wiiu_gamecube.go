package joystick

var wiiuGameCubeConfig = joystickConfig{
	Name: "Wii U Adapter GameCube Controller",
	GUID: "3333",
	Axis: []pair{
		{
			Name: "left_x",
			ID:   0,
		},
		{
			Name: "left_y",
			ID:   1,
		},
		{
			Name: "right_x",
			ID:   5,
		},
		{
			Name: "right_y",
			ID:   4,
		},
		{
			Name: "rt",
			ID:   3,
		},
		{
			Name: "lt",
			ID:   2,
		},
	},
	Buttons: []pair{
		{
			Name: "x",
			ID:   2,
		},
		{
			Name: "a",
			ID:   0,
		},
		{
			Name: "b",
			ID:   1,
		},
		{
			Name: "y",
			ID:   3,
		},
		{
			Name: "lb",
			ID:   6,
		},
		{
			Name: "rb",
			ID:   5,
		},
		{
			Name: "start",
			ID:   7,
		},
		{
			Name: "z",
			ID:   4,
		},
		{
			Name: "down",
			ID:   9,
		},
		{
			Name: "left",
			ID:   10,
		},
		{
			Name: "right",
			ID:   11,
		},
		{
			Name: "up",
			ID:   8,
		},
	},
}

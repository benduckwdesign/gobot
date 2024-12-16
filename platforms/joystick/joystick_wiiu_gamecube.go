package joystick

// This config was created for Windows using the following commonly used setup:
// Gamecube USB Adapter Driver v3.2.1 (GCN USB Feeder.exe) -> vJoy
// and Zadig to replace the USB driver with WinUSB

var wiiuGameCubeConfig = joystickConfig{
	// WUP-028
	Name: "Wii U Adapter GameCube Controller",
	// USB ID 057E (Nintendo) 0337 (Wii U USB Adapter Device)
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

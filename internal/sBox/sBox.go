package sBox

var lookup = [4][4]byte{
	{0b00, 0b10, 0b01, 0b11},
	{0b10, 0b00, 0b11, 0b01},
	{0b01, 0b11, 0b00, 0b10},
	{0b11, 0b01, 0b10, 0b00},
}

func sBox(b byte) (byte, error) {
	// para coger los ultimos bits de un byte se aplica una máscara
	col := (b >> 2) & 0b11
	row := b & 0b11
	
	return lookup[row][col], nil
}

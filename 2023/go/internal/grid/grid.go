package grid

// Cartesian
type Coord struct {
	X int
	Y int
}

// Cartesian coordinates of all adjacent Neighbors
// Does not take into consideration edges, maximum
// .....
// .nnn.
// .ncn.
// .nnn.
// .....
func (c Coord) Neighbors() []Coord {
	return []Coord{
		{X: c.X - 1, Y: c.Y - 1},
		{X: c.X - 1, Y: c.Y},
		{X: c.X - 1, Y: c.Y + 1},
		{X: c.X, Y: c.Y - 1},
		{X: c.X, Y: c.Y + 1},
		{X: c.X + 1, Y: c.Y - 1},
		{X: c.X + 1, Y: c.Y},
		{X: c.X + 1, Y: c.Y + 1},
	}
}

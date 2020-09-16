package model

type Board interface {
	CurrentState() [][]*Figure
	Move(figure *Figure, dest Coordinate) (success bool)
	Castle(rook *Rook) (success bool)
}

type Coordinate struct {
	x int
	y int
}

type DefaultBoard struct {
	state           [][]*Figure
	castle_possible bool
}

func (b DefaultBoard) CurrentState() [][]*Figure { return b.state }

func (b DefaultBoard) Move(figure *Figure, dest Coordinate) bool {
	possible := (*figure).movements(b)
	for _, field := range possible {
		if field == dest {
			// TODO move
			break
		}
	}
	return false
}

func (b DefaultBoard) Castle(rook *Rook) bool {
	// TODO
	return false
}

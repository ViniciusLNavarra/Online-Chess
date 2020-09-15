package model

type Board = [8][8]Figure

type Coordinate struct {
	x int
	y int
}

type Figure interface {
	id() uint8
	movements(Board) []Coordinate
}

type King struct{}

func (k King) id() uint8 { return 0 }

func (k King) movements(b Board) []Coordinate {
	// TODO:
}

type Queen struct{}

func (q Queen) id() uint8 { return 1 }

func (q Queen) movements(b Board) []Coordinate {
	//TODO:
}

type Rook struct{}

func (r Rook) id() uint8 { return 2 }

func (r Rook) movements(b Board) []Coordinate {
	//TODO:
}

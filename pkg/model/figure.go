package model

type Figure interface {
	id() uint8
	color() uint8
	movements(Board) []Coordinate
}

type Color struct {
	id uint8
}

func (c Color) color() uint8 { return c.id }

var WHITE = Color{0}
var BLACK = Color{1}

type King struct{ Color }

func (k *King) id() uint8 { return 0 }

func (k *King) movements(b Board) []Coordinate {
	// TODO:
	return []Coordinate{}
}

type Queen struct{ Color }

func (q *Queen) id() uint8 { return 1 }

func (q *Queen) movements(b Board) []Coordinate {
	//TODO:
	return []Coordinate{}
}

type Rook struct{ Color }

func (r *Rook) id() uint8 { return 2 }

func (r *Rook) movements(b Board) []Coordinate {
	//TODO:
	return []Coordinate{}
}

type Bishop struct{ Color }

func (f *Bishop) id() uint8 { return 3 }

func (f *Bishop) movements(b Board) []Coordinate {
	//TODO:
	return []Coordinate{}
}

type Knight struct{ Color }

func (k *Knight) id() uint8 { return 4 }

func (k *Knight) movements(b Board) []Coordinate {
	//TODO:
	return []Coordinate{}
}

type Pawn struct{ Color }

func (p *Pawn) id() uint8 { return 5 }

func (p *Pawn) movements(b Board) []Coordinate {
	//TODO:
	return []Coordinate{}
}

type Invalid struct{}

func (i *Invalid) id() uint8 { return 6 }

func (i *Invalid) movements(b Board) []Coordinate {
	return []Coordinate{}
}

func (i *Invalid) color() uint8 { return 255 }

package simulation

type PlaySession interface {
}

type playSession struct {
	player Player
	table  table
}

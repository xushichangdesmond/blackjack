package simulation

type box struct {
	//sync.Mutex
	player Player
}

func (b box) String() string {
	//b.Lock()
	//defer b.Unlock()

	p := b.player
	if p == nil {
		return "No Player"
	}
	return "Player Name - " + p.Name()

}

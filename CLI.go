package poker

import "io"

type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

func (c *CLI) PlayerPoker() {
	c.playerStore.RecordWin("Chris")
}

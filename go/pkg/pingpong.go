package pkg

import "fmt"

type Match interface {
    AddPoint(p *Player)
}

type Player struct {
    Name string
    Score uint
    Win bool
}

type PingPongMatch struct {

}

const (
    POINT uint = 15
    MAX_POINTS  = 60
)

func (p PingPongMatch) AddPoint(player *Player) {
    player.Score += POINT
    if player.Score >= MAX_POINTS {
        player.Win = true
    }
}

func (p PingPongMatch) ShowResult(player1, player2 Player) string {
    return fmt.Sprintf("%d - %d", player1.Score, player2.Score)
}
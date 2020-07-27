package pkg_test

import (
	"github.com/mercadolibre/fury_shipping-dx-dojo/pkg"
	"github.com/stretchr/testify/assert"
    "testing"
)

func Test_AddPoints_PlayerHas15Points(t *testing.T){
    player := pkg.Player{}
    var match pkg.Match
    match = pkg.PingPongMatch{}

    match.AddPoint(&player)
    assert.Equal(t, uint(15), player.Score)
}

func Test_Match_PlayerWins(t *testing.T){
    player := pkg.Player{}
    var match pkg.Match
    match = pkg.PingPongMatch{}

    match.AddPoint(&player)
    match.AddPoint(&player)
    match.AddPoint(&player)
    match.AddPoint(&player)
    assert.True(t, player.Win)
}

func Test_Match_showResult(t *testing.T){
    player1 := pkg.Player{}
    player2 := pkg.Player{}

    match := pkg.PingPongMatch{}

    match.AddPoint(&player1)
    match.AddPoint(&player1)
    match.AddPoint(&player2)
    match.AddPoint(&player2)
    assert.Equal(t, "30 - 30", match.ShowResult(player1, player2))
}

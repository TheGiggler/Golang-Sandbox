package models

import "sync"

type PlayByPlay struct {
	Text   string
	Index  int
	GameID string
}

type Game struct {
	Description    string
	GameID         int
	LeagueID       int
	HomeTeamID     int
	VisitingTeamID int
}
type GamePlayByPlays []PlayByPlay

func (p PlayByPlay) GetPlayByPlay() string {
	return p.Text
}

func (p *PlayByPlay) IncrementIndex(m *sync.Mutex) {
	m.Lock()
	p.Index++
	m.Unlock()

}

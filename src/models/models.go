package models

import (
		"sync"
		"time"
)

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

type GameFeed struct {
	LastUpdatedOn time.Time `json:"lastUpdatedOn"`
	GameDayDate time.Time "json:gameDayDate"
	Games         []struct {
		Schedule struct {
			ID        int         `json:"id"`
			StartTime time.Time   `json:"startTime"`
			EndedTime interface{} `json:"endedTime"`
			AwayTeam  struct {
				ID           int    `json:"id"`
				Abbreviation string `json:"abbreviation"`
			} `json:"awayTeam"`
			HomeTeam struct {
				ID           int    `json:"id"`
				Abbreviation string `json:"abbreviation"`
			} `json:"homeTeam"`
			Venue struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"venue"`
			VenueAllegiance          string        `json:"venueAllegiance"`
			ScheduleStatus           string        `json:"scheduleStatus"`
			OriginalStartTime        interface{}   `json:"originalStartTime"`
			DelayedOrPostponedReason interface{}   `json:"delayedOrPostponedReason"`
			PlayedStatus             string        `json:"playedStatus"`
			Attendance               interface{}   `json:"attendance"`
			Officials                []interface{} `json:"officials"`
			Broadcasters             []interface{} `json:"broadcasters"`
			Weather                  interface{}   `json:"weather"`
		} `json:"schedule"`
		Score struct {
			CurrentPeriod                 interface{} `json:"currentPeriod"`
			CurrentPeriodSecondsRemaining interface{} `json:"currentPeriodSecondsRemaining"`
			CurrentIntermission           interface{} `json:"currentIntermission"`
			AwayScoreTotal                int         `json:"awayScoreTotal"`
			AwayShotsTotal                int         `json:"awayShotsTotal"`
			HomeScoreTotal                int         `json:"homeScoreTotal"`
			HomeShotsTotal                int         `json:"homeShotsTotal"`
			Periods                       []struct {
				PeriodNumber int `json:"periodNumber"`
				AwayScore    int `json:"awayScore"`
				AwayShots    int `json:"awayShots"`
				HomeScore    int `json:"homeScore"`
				HomeShots    int `json:"homeShots"`
			} `json:"periods"`
		} `json:"score"`
	} `json:"games"`
	References struct {
		TeamReferences []struct {
			ID           int    `json:"id"`
			City         string `json:"city"`
			Name         string `json:"name"`
			Abbreviation string `json:"abbreviation"`
			HomeVenue    struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"homeVenue"`
			TeamColoursHex       []interface{} `json:"teamColoursHex"`
			SocialMediaAccounts  []interface{} `json:"socialMediaAccounts"`
			OfficialLogoImageSrc interface{}   `json:"officialLogoImageSrc"`
		} `json:"teamReferences"`
		VenueReferences []struct {
			ID                    int           `json:"id"`
			Name                  string        `json:"name"`
			City                  string        `json:"city"`
			Country               string        `json:"country"`
			GeoCoordinates        interface{}   `json:"geoCoordinates"`
			CapacitiesByEventType []interface{} `json:"capacitiesByEventType"`
			PlayingSurface        interface{}   `json:"playingSurface"`
			BaseballDimensions    []interface{} `json:"baseballDimensions"`
			HasRoof               interface{}   `json:"hasRoof"`
			HasRetractableRoof    interface{}   `json:"hasRetractableRoof"`
		} `json:"venueReferences"`
	} `json:"references"`
}
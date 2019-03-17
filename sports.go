package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"models"
	"net/http"
	"sync"
	"time"
)

var updateCount int;

var authHeader = "5384639843049-39-"

//test http request
func MakeRequest() {
	resp, err := http.Get("https://www.mlb.com")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func main() {
	fmt.Printf("Welcome to Sports!\n")

	var wg sync.WaitGroup
	var mux = sync.Mutex{}
	var games = LoadGames()
for true{
	for _,game :=range games {

		wg.Add(1)
		go UpdateGame(game.GameID,&wg,&mux)

	}
	wg.Wait()

	time.Sleep(time.Second*5)
}
	// for _, element := range pbps {
	// 	fmt.Printf(element.GetPlayByPlay() + "\n")
	// 	var p = &element
	// 	p.IncrementIndex()
	// 	fmt.Printf("%v\n", element.Index)

	// }

	fmt.Printf("Sports is over!\n")
}

func UpdateGame (GameID int, wg *sync.WaitGroup,m *sync.Mutex){

	fmt.Printf("Updating GameID %v\n", GameID)
	m.Lock();
	updateCount++
	m.Unlock()
	fmt.Printf("updateCount: %v\n", updateCount)
	wg.Done()
}

func GetLivePlayByPlay(gameID int) models.PlayByPlay {

	return models.PlayByPlay{}

}

func LoadGames() []models.Game {
	var games = []models.Game{
		models.Game{GameID:1 },
		models.Game{GameID:2 },
		models.Game{GameID:3 },
		models.Game{GameID:4 },
	}
	

	return games
}

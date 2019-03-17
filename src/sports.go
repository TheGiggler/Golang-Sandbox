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

var updateCount int

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
	updateChan := make(chan string)
	quit := make(chan bool)

	fmt.Printf("Welcome to Sports!\n")

	var wg sync.WaitGroup
	var mux = sync.Mutex{}
	var games = LoadGames()
	

	go OutputGameResult(updateChan,quit)
	//for true {
		for i := 0; i < 10; i++ {


		for _, game := range games {

			wg.Add(1)
			go UpdateGame(game.GameID, &wg, &mux, updateChan)

		}
		wg.Wait()

		time.Sleep(time.Second * 2)
		//quit<-true
	}

	
	time.Sleep(time.Second * 5)
	// for _, element := range pbps {
	// 	fmt.Printf(element.GetPlayByPlay() + "\n")
	// 	var p = &element
	// 	p.IncrementIndex()
	// 	fmt.Printf("%v\n", element.Index)

	// }

	fmt.Printf("Sports is over!\n")
}

func OutputGameResult(ch chan string, quit chan bool) {

	timer := time.NewTimer(time.Second * 6)
	for true {
		select {
		case msg1 := <-ch:
			fmt.Println("received", msg1)
		case <-timer.C:
			fmt.Sprint("Timer expired")
		case<-quit:
		fmt.Sprint("Quitting OutputGameResult")
			return

		//default:
		//	fmt.Println("no activity")
		}
	}

}

func UpdateGame(GameID int, wg *sync.WaitGroup, m *sync.Mutex, ch chan string) {

	fmt.Printf("In Updating GameID %v\n", GameID)
	m.Lock()
	updateCount++
	m.Unlock()
	fmt.Printf("updateCount %v\n", updateCount)
	ch <- fmt.Sprint("Updating GameID ", GameID)
	wg.Done()
}

func GetLivePlayByPlay(gameID int) models.PlayByPlay {

	return models.PlayByPlay{}

}

func LoadGames() []models.Game {
	var games = []models.Game{
		models.Game{GameID: 1},
		models.Game{GameID: 2},
		models.Game{GameID: 3},
		models.Game{GameID: 4},
	}

	return games
}

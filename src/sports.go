package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"models"
	"net/http"
	"strings"
	"sync"
	"time"
	"strconv"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/bson"
)
var updateCount int

var authHeader = "ODA5MWU4OGUtZDQ2Ni00YTdlLTljNTUtZTE2MTZhOk1ZU1BPUlRTRkVFRFM="

//test http request
func RequestPlayByPay(GameID int) {

	id := strconv.Itoa(GameID)
	tmp := "https://api.mysportsfeeds.com/v2.1/pull/nhl/2018-2019/games/{game}/playbyplay.json"

	uri := strings.Replace(tmp, "{game}", id, -1)
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization","Basic " + authHeader)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func GetGamesForToday() []models.Game{



	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))





	//todo: get date parameter from current date
	uri:="https://api.mysportsfeeds.com/v2.1/pull/nhl/2018-2019/date/20190317/games.json"
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization","Basic " + authHeader)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	gameFeed:=&models.GameFeed{}
	newerr := json.Unmarshal([]byte(body), gameFeed)
	if(newerr!=nil) {
		log.Fatal(newerr)
	}

	gameFeed.GameDayDate = time.Now()
	//persist to mongo
	collection := mongoClient.Database("schedule").Collection("games")	
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, gameFeed)
	id := res.InsertedID
	log.Println(id)

	games:=[]models.Game{}
	for _, game := range gameFeed.Games {
		g := models.Game{GameID:game.Schedule.ID}
		games = append(games,g)
    }
	//game:=new(models.Game)

	return games
}

func main() {
	updateChan := make(chan string)
	quit := make(chan bool)

	fmt.Printf("Welcome to Sports!\n")

	var wg sync.WaitGroup
	var mux = sync.Mutex{}
	var games = GetGamesForToday()
	//need to get game list first ... needs to run synchronously?

	go OutputGameResult(updateChan, quit)
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

	//time.Sleep(time.Second * 5)
	// for _, element := range pbps {
	// 	fmt.Printf(element.GetPlayByPlay() + "\n")
	// 	var p = &element
	// 	p.IncrementIndex()
	// 	fmt.Printf("%v\n", element.Index)

	// }
	fmt.Sprint("Sending quit")
	quit <- true
	fmt.Sprint("Quit sent")
	time.Sleep(time.Second * 5)
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
		case <-quit:
			fmt.Println("Quitting OutputGameResult")
			return

		default:
			fmt.Println("no activity")
		}
	}

}

func UpdateGame(GameID int, wg *sync.WaitGroup, m *sync.Mutex, ch chan string) {
	defer wg.Done()
	fmt.Printf("In Updating GameID %v\n", GameID)
	m.Lock()
	updateCount++
	m.Unlock()
	fmt.Printf("updateCount %v\n", updateCount)
	go RequestPlayByPay(GameID)
	ch <- fmt.Sprint("Updating GameID ", GameID)

}

func GetLivePlayByPlay(gameID int) models.PlayByPlay {

	return models.PlayByPlay{}

}

func LoadGames() []models.Game {
	var games = []models.Game{
		models.Game{GameID: 47409},
		models.Game{GameID: 47410},
		models.Game{GameID: 47411},
		models.Game{GameID: 47412},
	}

	return games
}

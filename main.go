package main

import (
	"io"
	"log"
	"os"
	"time"

	nhlApi "nhl-project/nhlApi"
)

func main() {
	// help to benchmark time for request
	now := time.Now()

	rosterFile, err := os.OpenFile("rosters.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file %v", err)
	}

	defer rosterFile.Close()

	writer := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(writer)

	teams, err := nhlApi.GetAllTeams()
	if err != nil {
		log.Fatalf("error getting all teams %v", err)
	}

	for _, team := range teams {
		log.Println("---------------------")
		log.Printf("Team Name: %s", team.Name)
	}

	log.Printf("Time taken: ", time.Now().Sub(now).String())
}

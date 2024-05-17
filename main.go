package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Profile struct {
	TotalSolved int `json:"totalSolved"`
}

func getNumberOfDays(startTime int64) int {
	curTime := time.Now().Unix()

	differenceInSeconds := curTime - startTime
	differenceInDays := differenceInSeconds / (60 * 60 * 24)

	return int(differenceInDays)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	leetCodeHandler := os.Getenv("LEETCODEHANDLER")

	// try not to panic for everything
	resp, err := http.Get("https://leetcode-stats-api.herokuapp.com/" + leetCodeHandler)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var profile Profile
	if err := json.Unmarshal(body, &profile); err != nil {
		panic(err)
	}

	// stated with 200 questions solved at the start date
	startQuestions := 200
	curSolved := profile.TotalSolved
	startDate := time.Date(2024, time.April, 15, 15, 30, 0, 0, time.UTC).Unix()

	expectedToSolve := (getNumberOfDays(startDate) * 3) + startQuestions

	if expectedToSolve > curSolved {
		status := fmt.Sprintf("You are behind by %d questions!!", expectedToSolve-curSolved)
		color.Red(status)
	} else if expectedToSolve < curSolved {
		status := fmt.Sprintf("Keep up the good work")
		color.Green(status)
	} else {
		fmt.Printf("NOt BAd NOT GOod")
	}
}

package main

import (
	"database/sql"
	"fmt"
	"time"
)

func challengesTool(user User, db *sql.DB) {
	fmt.Println("I was running out of options, so instead of an API, I'll give you some challenges")
	fmt.Println("You can do a challenge and add yourself to our leaderboard")
	fmt.Println("1. TryHackMe")
	fmt.Println("2. HackTheBox")
	fmt.Println("3. Other challenges/websites with Rooms?")
	fmt.Println("4. Retrieve Leaderboard")
	var option int
	fmt.Scanln(&option)
	challengeOptions(option, user, db)

}

func challengeOptions(option int, user User, db *sql.DB) {
	if option == 1 {
		fmt.Println("Please give us the Room you did")
		var room string
		fmt.Scanln(&room)
		fmt.Println("Please give us the score you got (https://docs.tryhackme.com/docs/rooms/how-points-work/)")
		var score int
		fmt.Scanln(&score)
		fmt.Println("You will be added using your ID and Username, as well as above options")
		website := "TryHackMe"
		time.Sleep(2 * time.Second)
		leaderboardAdd(user, db, room, score, website)
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 2 {
		fmt.Println("Please give us the Machine you did")
		var machine string
		fmt.Scanln(&machine)
		fmt.Println("How many points do you have?")
		var points int
		fmt.Scanln(&points)
		website := "HackTheBox"
		time.Sleep(2 * time.Second)
		leaderboardAdd(user, db, machine, points, website)
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 3 {
		fmt.Println("What website did you use?")
		var website string
		fmt.Scanln(&website)
		fmt.Println("Do you have a score to add?")
		var score int
		fmt.Scanln(&score)
		fmt.Println("Did you do a room/machine/box, if so does it have a name?")
		var activity string
		fmt.Scanln(&activity)
		leaderboardAdd(user, db, activity, score, website)
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 4 {
		retrieveLeaderboard(user, db)
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else {
		falseOptionFunc(user, db)
	}
}

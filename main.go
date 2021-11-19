package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Rank struct {
	Name string `json:"name"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
}

type UserProfile struct {
	Elo int `json:"elo"`
}

type GetRankResponse struct {
	Rank Rank `json:"rank"`
	Elo  int  `json:"elo"`
}

func getRankFromElo(elo int) Rank {
	var rank Rank

	for i := 0; i < len(ranks); i++ {
		tmpRank := ranks[i]
		if tmpRank.Min > 0 && elo > tmpRank.Min {
			if tmpRank.Max > 0 && elo < tmpRank.Max {
				rank = tmpRank
				break
			}
		}
	}

	return rank
}

func getRank(c *gin.Context) {
	name := c.Param("name")

	resp, err := http.Get(fmt.Sprintf("https://esportal.com/api/user_profile/get?username=%s&rank=1", name))

	if err != nil {
		fmt.Println(err)
	}

	var profile UserProfile
	err = json.NewDecoder(resp.Body).Decode(&profile)

	if err != nil {
		fmt.Print(err)
	}

	rank := GetRankResponse{
		Rank: getRankFromElo(profile.Elo),
		Elo:  profile.Elo,
	}

	json.NewEncoder(c.Writer).Encode(rank)
}

func handleRequests() {
	router := gin.Default()
	router.GET("/rank/:name", getRank)

	router.Run()
}

func main() {
	handleRequests()
}

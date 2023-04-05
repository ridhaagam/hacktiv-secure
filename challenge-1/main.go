package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"secure/challenge-1/domain"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init router
	router := gin.Default()
	
	// Post simulation
	go sumulation()

	// Router port
	err := router.Run(":4000")

	// Handle router
	if err != nil {
		panic("Error When Running")
	}
}

func sumulation(){
	for {
		time.Sleep(15 * time.Second)
		valueWater := rand.Intn(100) + 1
		valueWind := rand.Intn(100) + 1
		postReq := domain.Element{
			Water: valueWater,
			Wind:  valueWind,
		}

		// Hitung status untuk water
		if postReq.Water < 5 {
			postReq.WaterStatus = "aman"
		} else if postReq.Water >= 5 && postReq.Water <= 8 {
			postReq.WaterStatus = "siaga"
		} else {
			postReq.WaterStatus = "bahaya"
		}

		// Hitung status untuk wind
		if postReq.Wind < 6 {
			postReq.WindStatus = "aman"
		} else if postReq.Wind >= 6 && postReq.Wind <= 15 {
			postReq.WindStatus = "siaga"
		} else {
			postReq.WindStatus = "bahaya"
		}

		jsonValue, _ := json.Marshal(postReq)
		resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		// Tampilkan hasil POST di terminal
		var postRes domain.Element
		json.NewDecoder(resp.Body).Decode(&postRes)
		fmt.Printf("{\n   valuewater: %d,\n   valuewind: %d\n}\nstatus water : %s\nstatus wind : %s\n\n", postRes.Water, postRes.Wind, postRes.WaterStatus, postRes.WindStatus)
	}
}
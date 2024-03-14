package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

var PORT = ":8080"

func updateStatus() {
	for {
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		status := Status{Water: water, Wind: wind}
		jsonData, err := json.Marshal(status)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile("status.json", jsonData, 0644)
		if err != nil {
			panic(err)
		}

		time.Sleep(15 * time.Second)
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	jsonData, err := ioutil.ReadFile("status.json")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
	w.Write(jsonData)
}

func main() {
	go updateStatus()

	http.HandleFunc("/status", statusHandler)

	fmt.Println("Server is running on PORT", PORT)
	http.ListenAndServe(PORT, nil)
}

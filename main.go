package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Activity struct {
	TotalKM    float64 `json:"total_km"`
	LastUpdate string  `json:"last_update"`
}

func main() {
	file, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Failed to read file, please ensure data.json exists:", err)
		return
	}

	var data Activity
	json.Unmarshal(file, &data)

	dailyKM := 5.0
	data.TotalKM += dailyKM
	data.LastUpdate = time.Now().Format("2006-01-02 15:04:05")

	updatedData, _ := json.MarshalIndent(data, "", "  ")
	_ = os.WriteFile("data.json", updatedData, 0644)

	fmt.Printf("🚀 Successfully updated! Current total distance: %.2f KM\n", data.TotalKM)
}
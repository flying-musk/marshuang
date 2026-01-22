package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Piano struct {
	PianoHours float64 `json:"piano_hours"`
	LastUpdate string  `json:"last_update"`
}

func main() {
	file, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Failed to read file, please ensure data.json exists:", err)
		return
	}

	var data Piano
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	pianoHoursEnv := os.Getenv("PIANO_HOURS")
	if pianoHoursEnv == "" {
		fmt.Println("No input received for PIANO_HOURS, skipping update.")
	} else {
		hours, err := strconv.ParseFloat(pianoHoursEnv, 64)
		if err != nil {
			fmt.Println("Invalid input for hours:", err)
		} else if hours > 0 {
			data.PianoHours += hours
			fmt.Printf("🎹 Piano Update: +%.1f hrs, Total: %.1f hrs\n", hours, data.PianoHours)
		}
	}

	data.LastUpdate = time.Now().Format("2006-01-02 15:04:05")

	updatedData, _ := json.MarshalIndent(data, "", "  ")
	err = os.WriteFile("data.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Successfully saved data to data.json")
}

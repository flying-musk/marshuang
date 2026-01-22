package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Piano struct {
	PianoHours float64 `json:"piano_hours"`
	LastUpdate string  `json:"last_update"`
}

func main() {
	// 1. Read existing data
	file, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Printf("Error: Failed to read data.json: %v\n", err)
		return
	}

	var data Piano
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// 2. Early Return Logic
	pianoHoursEnv := os.Getenv("PIANO_HOURS")

	if pianoHoursEnv == "" {
		fmt.Println("No input received for PIANO_HOURS, skipping update.")
		return
	}

	hours, err := strconv.ParseFloat(pianoHoursEnv, 64)
	if err != nil || hours <= 0 {
		fmt.Printf("Invalid or zero input (%s), skipping update.\n", pianoHoursEnv)
		return
	}

	// 3. Update Data Object
	data.PianoHours += hours
	data.LastUpdate = time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("🎹 Success: +%.2f hrs | New Total: %.2f hrs\n", hours, data.PianoHours)

	// 4. Save to data.json
	updatedJSON, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile("data.json", updatedJSON, 0644)

	// 5. Update README.md
	updateREADME(data.PianoHours)
}

func updateREADME(totalHours float64) {
	readmePath := "README.md"
	content, err := os.ReadFile(readmePath)
	if err != nil {
		fmt.Printf("Error: Cannot read README: %v\n", err)
		return
	}

	text := string(content)

	// Regex breakdown:
	// ### 🎹 My total piano hours: -> Literal text match
	// \s* -> Matches 0 or more whitespace characters
	// \d+                         -> Matches 1 or more digits (integer part)
	// (\.\d+)?                    -> Optional group: a dot followed by 1 or more digits (decimal part)
	// \s* -> Optional whitespace before unit
	// hrs                         -> Literal unit match
	pattern := `### 🎹 My total piano hours:\s*\d+(\.\d+)?\s*hrs`
	re := regexp.MustCompile(pattern)

	// Guard clause: Check if the target line exists before proceeding
	if !re.MatchString(text) {
		fmt.Println("Warning: Could not find piano hours line in README.md")
		return
	}

	// Format the new total hours to 2 decimal places
	newHoursStr := fmt.Sprintf("%.2f", totalHours)
	newLine := fmt.Sprintf("### 🎹 My total piano hours: %shrs", newHoursStr)

	// Perform the string replacement globally in the text
	newText := re.ReplaceAllString(text, newLine)

	// Write the updated content back to README.md
	err = os.WriteFile(readmePath, []byte(newText), 0644)
	if err == nil {
		fmt.Printf("Success: README.md updated to %s hrs! 🎉\n", newHoursStr)
	}
}

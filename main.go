package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type LogWriter = func()

func getMessage() string {
	messages := []string{
		"Checking user credentials",
		"Sending notification email",
		"Saving entity to database",
		"Database connection lost",
	}

	idx := rand.Intn(len(messages))
	return messages[idx]
}

func getLevel() string {
	levels := []string{
		"DEBUG",
		"DEBUG",
		"DEBUG",
		"DEBUG",
		"INFO",
		"INFO",
		"ERROR",
	}

	idx := rand.Intn(len(levels))
	return levels[idx]
}

func jsonWriter() {
	entry := struct {
		Timestamp time.Time `json:"timestamp"`
		Level     string    `json:"log_level"`
		Message   string    `json:"message""`
	}{
		Timestamp: time.Now(),
		Level:     getLevel(),
		Message:   getMessage(),
	}

	data, _ := json.Marshal(entry)
	fmt.Println(string(data))
}

func plainWriter() {
	ts := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] %s\n", getLevel(), ts, getMessage())
}

func main() {
	mode := os.Getenv("MODE")
	var logWriter LogWriter

	switch mode {
	case "json":
		logWriter = jsonWriter
	case "plain":
		logWriter = plainWriter
	default:
		panic("unsupported MODE")
	}

	for {
		logWriter()
		time.Sleep(time.Second)
	}
}

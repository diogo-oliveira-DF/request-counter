package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	dataFile = "./output/counter.json"
)

var (
	mu       sync.Mutex
	requests []time.Time
)

// Handler Function called when there is interaction with the server (port 8080)
// mu is a mutex that ensures that the requests slice is accessed by only one goroutine at a time
// appends a new timestamp into requests slice
// and then will clean the invalid requests (after the 60 seconds, as specified)
func Handler(w http.ResponseWriter, _ *http.Request) {
	mu.Lock()         // locks the function
	defer mu.Unlock() // at the end of the function, unlocks the function

	now := time.Now()
	requests = append(requests, now)

	// TODO HERE Logic to clean old requests

	count := len(requests)
	_, err := w.Write([]byte(fmt.Sprintf("Number of requests in the last 60 seconds: %d", count)))
	if err != nil {
		log.Printf("failed to send output: %v", err)
	}

	// TODO HERE save new requests data to file
}

// LoadSavedData When the service is booted, it will read a json file containing
// all the records that were registered before.
// If the file does not exist, it will be created
func LoadSavedData() {
	file, err := os.ReadFile(dataFile)
	if os.IsNotExist(err) {
		err = createFile()
		if err != nil {
			log.Printf("error creating file: %v", err)
			return
		}
		file, _ = os.ReadFile(dataFile)
	}
	if err != nil {
		log.Printf("failed to open data file: %v", err)
		return
	}

	err = json.Unmarshal(file, &requests)
	if err != nil {
		log.Printf("Failed to decode requests: %v", err)
	}
}

// createFile Creates a new file in the giving location
// The file is initialized as "[]" because the data
// that will be witten/read inside it will be in form of a slice
func createFile() error {
	err := os.WriteFile(dataFile, []byte("[]"), 0644)
	if err != nil {
		log.Printf("failed to open data file: %v", err)
		return err
	}

	return nil
}

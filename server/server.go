package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// HTTP Handlers

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	response := info()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("ERROR: Failed to encode info response, %s", err)
	}
}

func HandleStart(w http.ResponseWriter, r *http.Request) {
	state := GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode start json, %s", err)
		return
	}

	start(state)

	// Nothing to respond with here
}

func HandleMove(w http.ResponseWriter, r *http.Request) {
	state := GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode move json, %s", err)
		return
	}

	response := move(state)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("ERROR: Failed to encode move response, %s", err)
		return
	}
}

func HandleEnd(w http.ResponseWriter, r *http.Request) {
	state := GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode end json, %s", err)
		return
	}

	end(state)

	// Nothing to respond with here
}

// Middleware

const ServerID = "battlesnake/github/starter-snake-go"

func withServerID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", ServerID)
		next(w, r)
	}
}

func RefreshJavaSource(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("new_src")

	if file == nil {
		log.Print("ERROR: no file new_src provided in form data")
	}

	defer file.Close()

	if err != nil {
		log.Panicf("ERROR: error occurred when fetching uploaded file, %s", err)
		return
	}

	// Copy into source
	f, err := os.OpenFile("./Main.java", os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	io.Copy(f, file)

	cmd := exec.Command("javac", "Main.java")
	_, err = cmd.Output()

	if err != nil {
		log.Panicf("ERROR: new Java source code compiled with error: %s", err)
		return
	}
}

// Start Battlesnake Server

func RunServer() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}

	http.HandleFunc("/", withServerID(HandleIndex))
	http.HandleFunc("/start", withServerID(HandleStart))
	http.HandleFunc("/move", withServerID(HandleMove))
	http.HandleFunc("/end", withServerID(HandleEnd))
	http.HandleFunc("/refresh", withServerID(RefreshJavaSource))

	log.Printf("Running Battlesnake at http://0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

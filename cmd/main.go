package main

import (
	"encoding/json"
	"flag"
	"log"
	mrand "math/rand"
	"net/http"
	"time"
)

const (
	Alphabet    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	AlphabetLen = len(Alphabet)
)

func GenerateRandomString(length int) string {
	seededRand := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = Alphabet[seededRand.Intn(AlphabetLen)]
	}
	return string(b)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handeled request %s:%s %s", r.UserAgent(), r.RemoteAddr, r.RequestURI)
	response := make(map[string]string)
	response["request_at"] = time.Now().String()
	time.Sleep(time.Duration(70+mrand.Intn(80)) * time.Millisecond)
	response["request_processed"] = time.Now().String()
	response["value"] = GenerateRandomString(6)

	respBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(respBytes)
	if err != nil {
		log.Println("Could not write response:", err)
	}
}

func main() {
	var serverAddress string
	flag.StringVar(&serverAddress, "address", "127.0.0.1:43000", "The server address in the format of host:port")
	flag.Parse()

	http.HandleFunc("/", rootHandler)
	log.Printf("Listening on %s", serverAddress)
	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

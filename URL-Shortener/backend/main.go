package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const baseURL = "127.0.0.1:8000/short/"

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/short/", redirect)
	http.ListenAndServe(":8000", nil)
}

func handleURLShortener(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var urlMap map[string]string
	data, err := os.ReadFile("url-map.json")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		err = json.Unmarshal(data, &urlMap)
		if err != nil {
			panic(err)
		} else {
			random := rand.New(rand.NewSource(time.Now().UnixNano()))
			randNumber := random.Intn(1000)
			var dataMap map[string]string
			data, err := io.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			defer r.Body.Close()
			if err := json.Unmarshal(data, &dataMap); err != nil {
				panic(err)
			}

			urlMap[strconv.Itoa(randNumber)] = dataMap["url"]
			writeFile(&urlMap)
			w.Header().Set("Content-Type", "application/json")
			response := map[string]string{"message": "Here is shortened url : " + baseURL + strconv.Itoa(randNumber)}
			json.NewEncoder(w).Encode(response)
		}
	}

}

func writeFile(data *map[string]string) error {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}

	return os.WriteFile("url-map.json", jsonData, 0644)

}
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // or specific origin
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func handler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == http.MethodOptions {
		// Handle preflight
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		handleURLShortener(w, r)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/short/")
	if id == "" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}
	data, err := os.ReadFile("url-map.json")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var urlMap map[string]string
		err = json.Unmarshal(data, &urlMap)
		if err != nil {
			panic(err)
		}
		redirectURL := urlMap[id]
		http.Redirect(w, r, redirectURL, http.StatusFound)
	}

}

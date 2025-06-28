package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/lrstanley/go-ytdlp"
)

// returns a hello world
func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// /isEven?number=1
func isEvenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	inputNumberStr := r.URL.Query().Get("number")
	inputNumber, err := strconv.Atoi(inputNumberStr)
	if err != nil {
		response := map[string]string{"error": "input must be an integer"}
		json.NewEncoder(w).Encode(response)
		return
	}

	isEven := inputNumber%2 == 0
	response := map[string]string{"isEven": strconv.FormatBool(isEven)}
	json.NewEncoder(w).Encode(response)
}

func yMp3Handler(w http.ResponseWriter, r *http.Request) {
	videoID := r.URL.Query().Get("v")
	if videoID == "" {
		http.Error(w, "Missing video ID", http.StatusBadRequest)
		return
	}

	ytdlp.MustInstall(context.TODO(), nil)

	dl := ytdlp.New().
		ExtractAudio().
		AudioFormat("wav").
		AudioQuality("0").
		// save to temp file
		Output("/tmp/%(id)s.%(ext)s")

	w.Header().Set("Content-Type", "audio/mpeg")

	// debug log
	fmt.Println("Downloading: https://www.youtube.com/watch?v=" + videoID)

	_, err := dl.Run(context.TODO(), "https://www.youtube.com/watch?v="+videoID)
	if err != nil {
		http.Error(w, "Download error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, "/tmp/"+videoID+".wav")

	// delete temp file
	err = os.Remove("/tmp/" + videoID + ".wav")
	if err != nil {
		http.Error(w, "Error deleting temp file: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

func main() {
	http.HandleFunc("/isEven", isEvenHandler)
	http.HandleFunc("/y-mp3", yMp3Handler)
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}

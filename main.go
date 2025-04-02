package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fprintf, err := fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	//runFfmpeg()
	os.Create(r.URL.Path[1:] + ".txt")
	//file.WriteTo()
	if err != nil {
		println(fprintf, err)
		return
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "This route is 'POST' only", http.StatusMethodNotAllowed)
	}

	startUp := time.Now()

	file, header, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "Erro ao obter o arquivo", http.StatusBadRequest)
	}

	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", header.Filename)

	endUp := time.Since(startUp)
	fmt.Printf("tempo de upload: %s", endUp)
}

func main() {
	//http.HandleFunc("/", handler)
	//log.Println(http.ListenAndServe(":8080", nil))

	http.HandleFunc("/upload", uploadHandler)
	log.Println(http.ListenAndServe(":8080", nil))

}

func runFfmpeg() {
	cmd := exec.Command("ffmpeg")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("cmd.Run() failed: %v\n\n", err)
	}
}

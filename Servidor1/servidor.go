package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func uploader(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2000)

	file, fileInfo, err := r.FormFile("archivo")

	f, err := os.OpenFile("./files/"+fileInfo.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer f.Close()

	io.Copy(f, file)

	fmt.Fprintf(w, fileInfo.Filename)

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/files", uploader)

	http.ListenAndServe(":8000", nil)
}

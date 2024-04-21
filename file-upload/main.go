package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

type Todo struct {
	Title string
	Done  bool
}

const port = ":9000"

func main() {

	mux := http.NewServeMux()

	fileHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("uploads")))
	client := http.StripPrefix("/client/", http.FileServer(http.Dir("client")))
	mux.HandleFunc("/static/", fileHandler.ServeHTTP)
	mux.HandleFunc("/client/", client.ServeHTTP)
	mux.HandleFunc("/", home)
	mux.HandleFunc("GET /upload", uploadView)
	mux.HandleFunc("POST /upload", upload)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Println(err)
	}
}

const uploadPath = "uploads"
const maxUploadSize = 10 << 20 // 10MB
func uploadView(w http.ResponseWriter, r *http.Request) {
	type pageData struct {
		Title string
	}
	template := template.Must(template.ParseFiles("upload.html"))
	data := pageData{
		Title: "File Upload",
	}
	template.Execute(w, data)
}

func upload(w http.ResponseWriter, r *http.Request) {

	// parse the form data
	// 10 << 20 = 10 * 2^20 = 10 * 1024 * 1024 = 10MB. Max file size
	r.ParseMultipartForm(maxUploadSize)

	// get the file from the form
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a new file in the upload directory
	dst, err := os.Create(filepath.Join(uploadPath, handler.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Successfully Uploaded File")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func home(w http.ResponseWriter, r *http.Request) {
	// read the files in the uploads directory
	type PageData struct {
		Title string
		Files []os.DirEntry
	}
	files, err := os.ReadDir(uploadPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title: "View Uploads",
		Files: files,
	}

	template := template.Must(
		template.ParseFiles("view-uploads.html"),
	)
	template.Execute(w, data)
	// files, err := os.ReadDir(uploadPath)
}

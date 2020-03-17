package app

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

const maxBytes  =10 *1024 *1024

func( receiver *server) handleFilesSave() func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		err := request.ParseMultipartForm(maxBytes)
		if err != nil {
			log.Print(err)
			http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		file, header, err := request.FormFile("data")
		uploadedFiles := ""
		if err != nil {
			log.Print(err)
			http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		defer file.Close()
		contentType := header.Header.Get("Content-Type")
		formFiles := request.MultipartForm
		files := formFiles.File

		for _, file := range files["data"] {
			openFile, err := file.Open(
				)
			if err != nil {
				log.Printf("can't create file: %v", err)
			}

			uploadedFiles, err = receiver.filesSvc.Save(openFile,contentType)
			if err != nil {
				log.Printf("can't save file: %v", err)
			}

		}

		responseWriter.Header().Set("Content-Type", "application/json")
		_, err = responseWriter.Write([]byte(uploadedFiles))
		if err != nil {
			http.Error(
				responseWriter,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)
			return
		}
		http.Redirect(responseWriter, request, "/", http.StatusFound)
	}
}


func(receiver *server) handleGetFile() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, request *http.Request) {
		fileName := strings.TrimPrefix(request.RequestURI, "/")
		log.Print(fileName)

		file, err := ioutil.ReadFile(filepath.Join(receiver.media, fileName))
		if err != nil {
			panic(err)
		}
		_, err = w.Write(file)
		if err != nil {
			log.Print(err)
		}
	}
}
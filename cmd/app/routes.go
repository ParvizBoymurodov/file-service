package app

import (
	"net/http"
)

func (receiver *server ) InitRoutes(address string)  {
	receiver.router.GET("/",receiver.handleGetFile())
	receiver.router.POST("/api/files",receiver.handleFilesSave())
	receiver.router.GET("/media/{id}", http.StripPrefix("/media", http.FileServer(http.Dir(receiver.media))).ServeHTTP)
}
package app

import (
	"errors"
	"fileservice/pkg/services"
	"github.com/ParvizBoymurodov/mux/pkg/mux"
	"net/http"
)

type server struct {
	filesSvc      *services.FilesSvc
	router        *mux.ExactMux
	templatesPath string
	assetsPath    string
	media         string
}

func (receiver *server) ServeHTTP(writer http.ResponseWriter,request *http.Request) {
	receiver.router.ServeHTTP(writer, request)
}

func NewServer(filesSvc *services.FilesSvc, router *mux.ExactMux, templatesPath string, assetsPath string, media string) *server {

	if filesSvc == nil {
		panic(errors.New("burgersSvc can't be nil"))
	}
	if router == nil {
		panic(errors.New("router can't be nil"))
	}
	if templatesPath == "" {
		panic(errors.New("templatesPath can't be empty"))
	}
	if assetsPath == "" {
		panic(errors.New("assetsPath can't be empty"))
	}
	if media == "" {
		panic(errors.New("media can't be empty"))
	}

	return &server{
		filesSvc: filesSvc,
		router:        router,
		templatesPath: templatesPath,
		assetsPath:    assetsPath,
		media:         media,
	}

}



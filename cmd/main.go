package main

import (
	"fileservice/cmd/app"
	"fileservice/pkg/services"
	"flag"
	"github.com/ParvizBoymurodov/mux/pkg/mux"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

var (
	host = flag.String("host", "0.0.0.0", "Server host")
	port = flag.String("port", "9991", "Server port")
)

const envHost = "HOST"
const envPort = "PORT"

func fromFLagOrEnv(flag *string, envName string) (server string, ok bool){
	if *flag != ""{
		return *flag, true
	}
	return os.LookupEnv(envName)
}

func main() {
	flag.Parse()
	hostf, ok := fromFLagOrEnv(host, envHost)
	if !ok {
		hostf = *host
	}
	portf, ok := fromFLagOrEnv(port, envPort)
	if !ok {
		portf = *port
	}
	addr := net.JoinHostPort(hostf, portf)
	start(addr)
}

func start(addr string) {
	router := mux.NewExactMux()
	templatesPath := filepath.Join("web", "templates")
	assetsPath := filepath.Join("web", "assets")
	media := filepath.Join("web", "mediaPath")
	filesSvc :=services.NewFilesSvc(media)
	server :=app.NewServer(
		filesSvc,
		router,
		templatesPath,
		assetsPath,
		media)
	server.InitRoutes(addr)
	panic(http.ListenAndServe(addr, server))
}
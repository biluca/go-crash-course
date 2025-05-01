package main

import (
	"fmt"
	"net/http"

	"go_api/internal/handlers"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	var router *chi.Mux = chi.NewRouter()
	handlers.Handler(router)

	fmt.Println("Starting GO API Service...")
	fmt.Println(`                                                    
  _    _ ______ _      _      ____             _____ _____ 
 | |  | |  ____| |    | |    / __ \      /\   |  __ \_   _|
 | |__| | |__  | |    | |   | |  | |    /  \  | |__) || |  
 |  __  |  __| | |    | |   | |  | |   / /\ \ |  ___/ | |  
 | |  | | |____| |____| |___| |__| |  / ____ \| |    _| |_ 
 |_|  |_|______|______|______\____/  /_/    \_\_|   |_____|                       
	`)

	err := http.ListenAndServe("localhost:8081", router)
	if err != nil {
		log.Error(err)
	}

}

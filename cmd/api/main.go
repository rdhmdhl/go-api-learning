package main

import (
	"fmt"
	"github.com/avukadin/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("starting Go API service...")
	fmt.Println(`
   _____       ___    ___  _____ _____ 
  / ____|     |__ \  / _ \| ____| ____|
 | |  __  ___    ) || | | | |__ | |__  
 | | |_ |/ _ \  / / | | | |___ \|___ \ 
 | |__| | (_) |/ /_ | |_| |___) |___) |
  \_____|\___/|____(_)___/|____/|____/ 
                                       
            🚀 Go API is live!
	`)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}

package main

import (
	"github.com/gorilla/mux"
	"github.com/searchspring/go-basics/clients/render"
	"github.com/searchspring/go-basics/clients/sscore"
	"github.com/searchspring/go-basics/dals/campaigns"
	campaignsHndl "github.com/searchspring/go-basics/handlers/campaigns"
	renderLib "github.com/unrolled/render"
	"net/http"
	"strconv"
)

const port = 8090

func main() {
	// Clients
	renderImpl := render.New(&render.Deps{Render: renderLib.New()})
	sscoreImpl := sscore.New(&sscore.Deps{}, &sscore.Config{
		Address:  "127.0.0.1:3306",
		Username: "root",
		Password: "root",
		DBName:   "ss_core_dev",
	})

	// DALs
	campaignsImpl := campaigns.New(&campaigns.Deps{SSCore: sscoreImpl})

	// Handlers
	campaignsHandler := campaignsHndl.New(&campaignsHndl.Deps{
		Campaigns: campaignsImpl,
		Render:    renderImpl,
	})

	router := mux.NewRouter()
	router.HandleFunc("/campaigns/{id}", campaignsHandler.Exists)

	http.ListenAndServe(":" + strconv.Itoa(port), router)
}

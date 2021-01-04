package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql" // Force the import
)

var renderer *render.Render
var router *mux.Router
var db *sql.DB
const port = 8090

func mustGetDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/?parseTime=true")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE ss_core_dev")
	if err != nil {
		panic(err)
	}

	return db
}

func getCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err:= rows.Scan(&count)
		if err != nil {
			panic(err)
		}
	}
	return count
}

func campaignExists(id string) (bool, error) {
	rows, err := db.Query(`SELECT COUNT(*) FROM MerchandisingCampaigns WHERE id = ?`, id)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	return getCount(rows) > 0, nil
}

func handler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, has := vars["id"]

	if !has {
		renderer.JSON(w, http.StatusBadRequest, map[string]string{
			"error": "No id",
		})
		return
	}

	found, err := campaignExists(id)

	if err != nil {
		renderer.JSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	renderer.JSON(w, http.StatusOK, map[string]string{
		"found": strconv.FormatBool(found),
	})
}

func main() {
	renderer = render.New()
	router = mux.NewRouter()
	db = mustGetDB()

	router.HandleFunc("/campaigns/{id}", handler)

	http.ListenAndServe(":" + strconv.Itoa(port), router)
}

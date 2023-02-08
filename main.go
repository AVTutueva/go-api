package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"tsi.co/go-api/utils"
)

func main() {

	// get data about cute cats
	resp, err := http.Get("https://cataas.com/api/cats?tags=cute")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data []utils.CatStruct
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	//  format data
	utils.PrepareCatStruct(&data)

	// show on the web-site
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	router := chi.NewRouter()

	// Format HTML page
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.HTML(w, r, `
		<html>
			<head>
			<title>Cute Cats</title>
			
			<style>
         		table, th, td {
            	border: 1px solid black;
         		}
      		</style>
			
			</head>
			
			<body>
			<h1>Cute Cats to adopt</h1>
			</body>
			</html>
		`)

		for i, cat := range data {
			render.HTML(w, r, fmt.Sprintf(`
			<hr>
			<h3>Cat %d</h3>
			<img src="https://cataas.com/cat/cute/says/%s">
			<p></p>
			<table>
			<tr>
				<th>Cat owner</th>
				<th>CreatedAt</th>
				<th>UpdatedAt</th>
				<th>Tags</th>
			</tr>
			<tr>
				<th>%s</th>
				<th>%s</th>
				<th>%s</th>
				<th>%s</th>
			</tr>
			</table>

			`, i+1, cat.Id, cat.Owner, cat.CreatedAt, cat.UpdatedAt, strings.Join(cat.Tags, ", ")))
		}
	})

	// log.Printf("Starting server on port %s.", port)
	err_server := http.ListenAndServe(":"+port, router)
	log.Fatal(err_server)
}

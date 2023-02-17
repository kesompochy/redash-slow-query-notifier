package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Data struct {
	QueryName  string `json:"query_name`
	User       string `json:"user`
	MinRuntime string `json:min_runtime`
	QueryID    string `json:id`
}

func main() {
	host := os.Getenv("HOST")
	queryId := os.Getenv("QUERY_ID")
	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf("https://%s/api/queries/%s/results.csv?api_key=%s", host, queryId, apiKey)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	r := csv.NewReader(res.Body)
	rows, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range rows {
		fmt.Println(v)
	}
}

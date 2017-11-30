package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

type Items struct {
	Item []Item `json:"items"`
}

type Item struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

func parseFlag() (string, string, string, string, string) {
	var query *string = flag.String("q", "example word", "Search query")
	var id *string = flag.String("i", "xxxx", "Application id")
	var key *string = flag.String("k", "yyyy", "Api key")
	var name *string = flag.String("n", "example index name", "Index name")
	var url *string = flag.String("u", "http://example.com/", "Blog URL")
	flag.Parse()

	return *query, *id, *key, *name, *url
}

func searchPosts(application_id string, apikey string, index_name string, query string) algoliasearch.QueryRes {
	client := algoliasearch.NewClient(application_id, apikey)
	index := client.InitIndex(index_name)
	res, _ := index.Search(query, nil)

	return res
}

func collectItem(res algoliasearch.QueryRes, url string) []Item {
	var items []Item
	for _, v := range res.Hits {
		title := v["title"].(string)
		status := v["status"].(string)
		slug := v["slug"].(string)
		posturl := url + slug

		items = append(items, Item{
			Title:    title,
			Subtitle: status,
			Arg:      posturl})
	}

	if items == nil {
		items = append(items, Item{
			Title: "No result"})
	}

	return items
}

func marshalItem(items []Item) []byte {
	j, _ := json.Marshal(Items{Item: items})

	return j
}
func showItem(j []byte) {
	os.Stdout.Write(j)
}

func Do() {
	query, application_id, apikey, index_name, url := parseFlag()
	res := searchPosts(application_id, apikey, index_name, query)
	items := collectItem(res, url)
	j := marshalItem(items)

	showItem(j)
}

func main() {
	Do()
}

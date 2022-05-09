package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type BookInfo struct {
	Items []struct {
		VolumeInfo struct {
			Title         string   `json:"title"`
			Subtitle      string   `json:"subtitle"`
			Authors       []string `json:"authors"`
			Publisher     string   `json:"publisher"`
			PublishedDate string   `json:"publishedDate"`
			Description   string   `json:"description"`
			PageCount     int      `json:"pageCount"`
			AverageRating int      `json:"averageRating"`
			RatingsCount  int      `json:"ratingsCount"`
			Language      string   `json:"language"`
			PreviewLink   string   `json:"previewLink"`
		} `json:"volumeInfo"`
	} `json:"items"`
}

func PrettyPrint(i interface{}) string {
	// This function is used to print the JSON in a pretty way

	s, _ := json.MarshalIndent(i, "", "\t")
	string_text := string(s)

	fmt.Println(reflect.TypeOf(string_text))

	return string_text
}

func EnterKey() string {
	// This function is used to enter author, title or relational key

	var key string

	fmt.Print("Enter author, title or relational key: ")
	fmt.Scanf("%s", &key)

	return key
}

func GetAllInfo(key string) BookInfo {
	// This function is used to get all the info of book

	all_article_url := "https://www.googleapis.com/books/v1/volumes?q=" + key

	r, err := http.Get(all_article_url)

	if err != nil {
		fmt.Println(err)
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	var all_info BookInfo

	if err := json.Unmarshal(body, &all_info); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return all_info
}

func main() {
	fmt.Println("Welcome to search book cli!")

	key := EnterKey()
	all_info := GetAllInfo(key)

	fmt.Println(PrettyPrint(all_info))
}

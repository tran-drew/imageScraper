package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	
	resp, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")	
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	
	
	text := string(bs)
	fmt.Printf("%s", text)
	fmt.Print(len(text))
}
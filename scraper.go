package main

import (
	"fmt"
	"net/http"
	"image/png"
	"io/ioutil"
	"io"
/*	"golang.org/x/net/html"*/
	"os"
)

func main() {

	resp, _ := http.Get("https://nationalpositions.com/wp-content/uploads/2016/10/logo-nike.png")
	bytes, _ := ioutil.ReadAll(resp.Body)
	r, _ := 
	imBytes, _ := io.ReadFull(resp.Body, bytes)
	defer resp.Body.Close()

	newF, _ := os.Create("./newF.html")
	defer newF.Close()

	img, _ := png.Decode(io.Reader(imBytes))
	f, err := os.OpenFile("np1.png", os.O_WRONLY|os.O_CREATE) 
	if err != nil {
		panic("Bad png")
	}

	png.Encode(f, bytes)

	n2, _ := newF.WriteString(string(bytes))
	fmt.Printf("Wrote %d bytes\n", n2)

	newF.Sync("http://www.parkavenueorthodonticsnyc.com/images/slide5.jpg")
}
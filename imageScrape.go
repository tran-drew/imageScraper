package main

import (
	"net/http"
	// "io/ioutil"
	"image/png"
	"image/jpeg"
	"os"
	"fmt"
	"strings"
)

func getPNGImage(url []string, name string) {

	resp, err := http.Get(strings.Join(url, "/"))
	if err != nil {
			fmt.Println(err)
	}
	defer resp.Body.Close()
	// bytes, _ := ioutil.ReadAll(resp.Body)

	img, err := png.Decode(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	filename := "./images/" + name
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	png.Encode(f, img)
	return
}

func getJPGImage(url []string, name string) {
	resp, err := http.Get(strings.Join(url, "/"))
	if err != nil {
			fmt.Println(err)
	}
	defer resp.Body.Close()

	img, err := jpeg.Decode(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	filename := "./images/" + name
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	o := jpeg.Options{Quality: 100}
	jpeg.Encode(f, img, *o)
	return
}

/*func parseURL(url string) {
	splitURL := strings.Split(url)
	return splitURL
}
*/

func main() {

	url := strings.Split("http://www.parkavenueorthodonticsnyc.com/images/slide3.jpg", "/")

	switch {
	case strings.Contains(url[len(url)-1], ".png"):
		getPNGImage(url, url[len(url)-1])
	case strings.Contains(url[len(url)-1], ".jpg"):
		getJPGImage(url, url[len(url)-1])
	}
	
}
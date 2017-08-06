package main
import (
	"net/http"
	"bufio"
	"image/png"
	"image/jpeg"
	"os"
	"strings"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getPNGImage(url []string, name string) {

	resp, err := http.Get(strings.Join(url, "/"))
	check(err)
	defer resp.Body.Close()
	// bytes, _ := ioutil.ReadAll(resp.Body)

	img, err := png.Decode(resp.Body)
	check(err)

	filename := "./images/" + name
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
	check(err)
	
	png.Encode(f, img)
	return
}

func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  check(err)
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
  file, err := os.Create(path)
  check(err)
  defer file.Close()

  w := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(w, line)
  }
  return w.Flush()
}
func getJPGImage(url []string, name string) {
	resp, err := http.Get(strings.Join(url, "/"))
	check(err)
	defer resp.Body.Close()

	img, err := jpeg.Decode(resp.Body)
	check(err)

	filename := "./images/" + name
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
	check(err)

	
	jpeg.Encode(f, img, nil)
	return
}

func main() {

	lines, err := readLines("foo.in.txt")
	check(err)
	for i, line := range lines {
		url := strings.Split(line, "/")
		fmt.Println(i, line)
			switch {
				case strings.Contains(url[len(url)-1], ".png"):
					getPNGImage(url, url[len(url)-1])
				case strings.Contains(url[len(url)-1], ".jpg"):
					getJPGImage(url, url[len(url)-1])
				case strings.Contains(url[len(url)-1], ".jpeg"):
					getJPGImage(url, url[len(url)-1])
			}
	}
	
}

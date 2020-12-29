package imgdow

import (
	"bufio"
	"io"
	"net/http"
	"os"
)

// DownloadImage will download a url to a local file.
func DownloadImage(imageURL string) error {
	filepath := "images/" + GetFilename(imageURL)

	// Get the data
	resp, err := http.Get(imageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// GetFilename can get filename from URL
func GetFilename(URL string) string {

	var index int
	for i, v := range URL {
		if string(v) == "/" {
			index = i
		}
	}

	return URL[index+1:]
}

// GetStringLine allows you to get a string from the user
func GetStringLine(prefix string) string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return text[:len(text)-2]
}

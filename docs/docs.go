package docs

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/rvfet/rich-go"
)

func FetchMany(urls []string) string {
	if len(urls) == 0 {
		return "No documentation URLs provided"
	}

	ch := make(chan string, len(urls)) // Buffered channel
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchOneForChannel(url, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var finalRes string
	for result := range ch {
		finalRes += result
	}
	return finalRes
}

func fetchOneForChannel(url string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		ch <- "No content for " + url + " at this time\n\n---\n\n"
		rich.Error("No content fetched for " + url + ". GET request returned an error: " + err.Error() + "❌")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ch <- "No content for " + url + " at this time\n\n---\n\n"
		rich.Error("No content fetched for " + url + ". Fetching it returned status code: " + strconv.Itoa(resp.StatusCode) + "❌")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- "No content for " + url + " at this time\n\n---\n\n"
		rich.Error("No content fetched for "+url+". Reading the response content returned an error: ", err.Error()+"❌")
		return
	}

	rich.Info("Fetched documentation for " + url + "✅")
	ch <- "## Documentation for " + url + "\n\n" + string(body) + "\n\n---\n\n"
}

func WriteFileContent(pathToFile, content string) error {
	dirName := filepath.Dir(pathToFile)
	errDirPath := os.MkdirAll(dirName, 0777)
	if errDirPath != nil {
		rich.Error("Unable to create the parent directory, an error occurred: ", errDirPath.Error()+"❌")
		return errDirPath
	}
	errFlWrt := os.WriteFile(pathToFile, []byte(content), 0777)
	if errFlWrt != nil {
		rich.Error("Unable to write file, an error occurred: ", errFlWrt.Error()+"❌")
	} else {
		rich.Info("File " + pathToFile + " successfully written!✅")
	}
	return errFlWrt
}

func GhToRawUrl(url string) string {
	url = strings.ReplaceAll(url, "github.com", "raw.githubusercontent.com")
	url = strings.ReplaceAll(url, "/blob/", "/")
	url = strings.ReplaceAll(url, "/tree/", "/")
	return url
}

package internal

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

//go:embed data/servers.txt
var dataByte []byte

type result struct {
	URL        string
	Accessible bool
}

var dataList = strings.Split(string(dataByte), "\n")

func GetAccessibleServers(workers uint8, timeout uint8) {
	var wg sync.WaitGroup

	urlChan := make(chan string, workers)
	resultChan := make(chan result, workers)

	client := http.Client{
		Timeout: time.Duration(time.Duration(timeout) * time.Second),
	}

	wg.Add(len(dataList))

	go func() {
		defer close(urlChan)
		for _, data := range dataList {
			urlChan <- data
		}
	}()

	bar := progressbar.NewOptions(len(dataList),
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionShowCount(),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("Scanning servers..."),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	defer bar.Finish()

	for i := uint8(0); i < workers; i++ {
		go func() {
			for url := range urlChan {
				networkRequest(url, resultChan, client)
				bar.Add(1)
				wg.Done()
			}
		}()
	}

	go func() {
		generateList(resultChan)
	}()

	wg.Wait()

	fmt.Println("Finished generating server list")
}

func networkRequest(url string, resultChan chan result, client http.Client) {
	response, _ := client.Get(url)
	if response != nil {
		resultChan <- result{
			URL:        url,
			Accessible: response.StatusCode == http.StatusOK,
		}
	}
	return
}

func generateList(ch chan result) {
	var mutex = &sync.Mutex{}
	mutex.Lock()

	f, err := os.Create("bdix.txt")
	if err != nil {
		log.Fatalln("unable to create file: ", err.Error())
	}

	defer f.Close()

	defer close(ch)

	for res := range ch {
		_, err := f.WriteString(res.URL + "\n")
		if err != nil {
			log.Fatalln("unable to write file: ", err.Error())
		}
		err = f.Sync()
		if err != nil {
			log.Fatalln("unable to sync file: ", err.Error())
		}
	}

	mutex.Unlock()
	return
}

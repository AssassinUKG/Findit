package main

// Notes
// 1. Takes command line arg for what to search for (findit!) eg: alert(0)
// 2. read file from stdin

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/fatih/color"
)

func argsCheck() string {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Enter a value to 'findit'!")
		fmt.Println("Usage: findit -c \"alert(0)\"")
		return ""
	}

	canary := flag.String("c", "--canary", "Canary value to find")
	flag.Parse()

	return *canary
}

func main() {

	searchParam := argsCheck()
	if searchParam == "" {
		return
	}
	//banner
	color.Cyan("%s %s", "ʕ◕ᴥ◕ ʔ", color.Cyan("Findit"))
	fmt.Println("Created by: Richard Jones")
	color.Yellow("Searching for canary value: %s \r\n", color.GreenString(searchParam))

	// read from standard in
	sc := bufio.NewScanner(os.Stdin)
	// make jobs
	jobs := make(chan string)
	var wg sync.WaitGroup
	//fmt.Println("Running for loop…")
	for i := 0; i < 15; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for line := range jobs {
				// line could be domain or any string var etc. Web Request
				resp, err := http.Get(line)
				if err != nil {
					continue
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
				}
				sb := string(body)
				findVal := strings.Contains(sb, searchParam)
				if findVal {

					fmt.Println(color.GreenString("Found Canary in response:"), line)

				}
			}
		}()

	}
	for sc.Scan() {
		line := sc.Text()
		jobs <- line
	}

	close(jobs)
	wg.Wait()

}

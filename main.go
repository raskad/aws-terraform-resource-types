package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	resp, err := http.Get("https://www.terraform.io/docs/providers/aws/index.html")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(body)))
	re := regexp.MustCompile(`.*/docs/providers/aws/r/(.*)\.html.*`)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		if len(matches) > 0 {
			fmt.Println(matches[1])
		}
	}
}

package jobs

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetPullRequests() {
	response, err := http.Get("https://api.github.com/repos/prithvipal/test/pulls?state=all&sort=updated&direction=desc")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}

package jobs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Prithvipal/sailpoint/config"
)

type PullRequest struct {
	URL       string `json:"url"`
	State     string `json:"state"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetPullRequests() []PullRequest {
	cfg := config.GetConfig()
	url := fmt.Sprintf(cfg.Git.URL, cfg.Git.Owner, cfg.Git.Repo)
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var prs []PullRequest
	json.Unmarshal(responseData, &prs)
	return prs
}

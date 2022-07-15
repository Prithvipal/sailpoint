package jobs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Prithvipal/sailpoint/config"
	"github.com/sirupsen/logrus"
)

type PullRequest struct {
	URL       string `json:"url"`
	State     string `json:"state"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetPullRequests() ([]PullRequest, error) {
	cfg := config.GetConfig()
	url := fmt.Sprintf(cfg.Git.URL, cfg.Git.Owner, cfg.Git.Repo)
	response, err := http.Get(url)

	if err != nil {
		logrus.Error("Error while fetching pull requests", err)
		return []PullRequest{}, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Error("Error while reading pull requests response", err)
		return []PullRequest{}, err
	}
	var prs []PullRequest
	err = json.Unmarshal(responseData, &prs)
	if err != nil {
		logrus.Error("unmarshalling error", err)
		return []PullRequest{}, err
	}

	return filterPRs(prs)
}

func filterPRs(prs []PullRequest) ([]PullRequest, error) {
	sevenDayBefore := time.Now().Add(-7 * 24 * time.Hour)
	res := []PullRequest{}
	for _, pr := range prs {
		updatedDate, err := time.Parse(time.RFC3339, pr.UpdatedAt)
		if err != nil {
			logrus.Error("time parsing error", err)
			return []PullRequest{}, err
		}
		if sevenDayBefore.Before(updatedDate) {
			res = append(res, pr)
		}
	}
	return res, nil
}

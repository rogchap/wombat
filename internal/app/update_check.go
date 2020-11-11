package app

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hashicorp/go-version"
)

const latestReleaseURL = "https://api.github.com/repos/rogchap/wombat/releases/latest"

type releaseResponse struct {
	TagName string `json:"tag_name"`
	HTMLURL string `json:"html_url"`
}

var ghClient *http.Client

func init() {
	ghClient = &http.Client{
		Timeout: 5 * time.Second,
	}
}

func checkForUpdate() (*releaseInfo, error) {
	req, err := http.NewRequest("GET", latestReleaseURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	resp, err := ghClient.Do(req)
	if err != nil {
		return nil, err
	}
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := &releaseResponse{}
	if err := json.Unmarshal(raw, r); err != nil {
		return nil, err
	}

	if versionGreaterThanOrEqual(semver, r.TagName) {
		return nil, errors.New("no update available")
	}

	return &releaseInfo{
		OldVersion: semver,
		NewVersion: r.TagName,
		URL:        r.HTMLURL,
	}, nil
}

func versionGreaterThanOrEqual(v, w string) bool {
	vv, ve := version.NewVersion(v)
	vw, we := version.NewVersion(w)

	return ve == nil && we == nil && vv.GreaterThanOrEqual(vw)
}

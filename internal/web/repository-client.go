package web

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
)

type RepositoryRespone struct {
	Items []struct {
		DownloadUrl string `json:"downloadUrl"`
		Maven2      struct {
			Extension  string `json:"extension"`
			ArtifactId string `json:"artifactId"`
			Version    string `json:"version"`
		} `json:"maven2"`
	} `json:"items"`
}

const (
	SEARCH_PATH = "/service/rest/v1/search/assets"
)

// RepositoryClient is a client for interacting with the Repository API.
type RepositoryClient struct {
	// baseUrl is the base URL for the API.
	baseUrl  string
	user     string
	password string
}

// NewRepositoryClient creates a new RepositoryClient.
func NewRepositoryClient(baseURL string, user string, password string) *RepositoryClient {
	return &RepositoryClient{
		baseUrl:  baseURL,
		user:     user,
		password: password,
	}
}

// GetRepository returns the repository.
func (c *RepositoryClient) GetVersions(application string, extension string, version string, snapshots bool) (RepositoryRespone, error) {
	getSearchUrl(c.baseUrl, snapshots, application, extension, version)

	req, err := http.NewRequest("GET", getSearchUrl(c.baseUrl, snapshots, application, extension, version), nil)
	if err != nil {
		return RepositoryRespone{}, err
	}
	req.SetBasicAuth(c.user, c.password)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return RepositoryRespone{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return RepositoryRespone{}, errors.New("failed to get versions: " + response.Status)
	}

	var result RepositoryRespone
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return RepositoryRespone{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return RepositoryRespone{}, err
	}
	return result, nil
}

func getSearchUrl(baseUrl string, snapshots bool, application string, extension string, version string) string {
	repositoryQuery, err := url.Parse(baseUrl + SEARCH_PATH)
	if err != nil {
		log.Fatal(err)
	}
	repository := "release-distributions"
	if snapshots {
		repository = "snapshot-distributions"
	}
	q := repositoryQuery.Query()
	q.Set("sort", "version")
	q.Set("repository", repository)
	q.Set("maven.groupId", "ch.tie")
	q.Set("maven.artifactId", application)
	q.Set("maven.extension", extension)
	q.Set("version", version)

	repositoryQuery.RawQuery = q.Encode()
	return repositoryQuery.String()
}

func (c *RepositoryClient) Login() error {
	response, err := c.GetVersions("asqueue", "jar", "*", false)
	if err != nil {
		return err
	}
	if len(response.Items) == 0 {
		return errors.New("failed to login")
	}
	return nil
}

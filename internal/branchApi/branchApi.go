package branchapi

import (
	"branchcomparer/internal/domain"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type BranchApi struct {
	url string
}

func New(url string) *BranchApi {
	return &BranchApi{
		url: url,
	}
}

type BranchResponse struct {
	Length   int              `json:"length"`
	Packages []domain.Package `json:"packages"`
}

type BranchPackage struct {
	Name      string `json:"name"`
	Epoch     int    `json:"epoch"`
	Version   string `json:"version"`
	Release   string `json:"release"`
	Arch      string `json:"arch"`
	Disttag   string `json:"disttag"`
	Buildtime int    `json:"buildtime"`
	Source    string `json:"source"`
}

func (ba *BranchApi) GetPackages(branch string) ([]domain.Package, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	r, err := client.Get(ba.url + branch)
	if err != nil {
		err = fmt.Errorf("Request to url (%s) is failed. Error: %w", ba.url, err)
		return nil, err
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("Reading of response is failed. Error: %w", err)
		return nil, err
	}

	var resp BranchResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		err = fmt.Errorf("Unmarshal response is failed. Error: %w", err)
		return nil, err
	}

	return resp.Packages, nil
}

/* from swagger:
{
  "request_args": {},
  "length": 0,
  "packages": [
    {
      "name": "string",
      "epoch": 0,
      "version": "string",
      "release": "string",
      "arch": "string",
      "disttag": "string",
      "buildtime": 0,
      "source": "string"
    }
  ]
}
*/

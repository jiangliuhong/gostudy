package module

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type QueryItem struct {
	Id            string    `json:"id"`
	Namespace     string    `json:"namespace"`
	Provider      string    `json:"provider"`
	Name          string    `json:"name"`
	Icon          string    `json:"icon"`
	DownloadCount int       `json:"downloadCount"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Version       string    `json:"version"`
	VcsId         string    `json:"vcsId"`
	RepoId        string    `json:"repoId"`
	RepoAddr      string    `json:"repoAddr"`
	GitTag        string    `json:"gitTag"`
	Category      string    `json:"category"`
	Type          string    `json:"type"`
	Desc          string    `json:"desc"`
}

type QueryResult struct {
	List []QueryItem `json:"list"`
}

type QueryResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  QueryResult `json:"result"`
}

func queryAll() error {
	iacQueryUrl := "http://iacstore-prod.yun.idcos/api/v1/modules?pageSize=10008&page=1&providers=alicloud"
	resp, err := http.Get(iacQueryUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var qr QueryResponse
	err = json.Unmarshal(b, &qr)
	if err != nil {
		return err
	}
	for _, item := range qr.Result.List {
		if item.Namespace == "alicloud" {
			fmt.Println(fmt.Sprintf("deleteModule %s:%s", item.Id, item.Namespace))
			err := deleteModule(item.Id)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func deleteModule(id string) error {

	iacDeleteUrl := "http://iacstore-prod.yun.idcos/api/v1/modules/%s"
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf(iacDeleteUrl, id), nil)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", iacAuthToken))
	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}
	return nil
}

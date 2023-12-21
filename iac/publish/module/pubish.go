package module

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

type Release struct {
	Version string `json:"tag_name"`
}

var filterRepositoryName = map[string]bool{
	"terraform-alicloud-drds":             true,
	"terraform-alicloud-ddoscoo-instance": true,
	"terraform-alicloud-demo":             true,
	"terraform-alicloud-bootstrap":        true,
	"terraform-alicloud-cloud-monitor":    true,
	"terraform-alicloud-vpc-peering":      true,
	"terraform-alicloud-edas":             true,
}

func newProxy() *http.Transport {
	parse, _ := url.Parse("http://127.0.0.1:7890")
	transport := &http.Transport{
		Proxy: http.ProxyURL(parse),
	}
	return transport
}

var proxy = newProxy()

func getLatestRelease(orgName, repoName, authToken string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", orgName, repoName)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))

	client := &http.Client{Transport: proxy}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch latest release for repository '%s/%s'. Status: %s", orgName, repoName, response.Status)
	}

	var release Release
	body, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body))
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &release)
	if err != nil {
		return "", err
	}

	return release.Version, nil
}

var authToken = "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"

func getOrganizationRepositories(orgName string) ([]Repository, error) {
	page := 1
	perPage := 100
	var allRepositories []Repository

	for {
		url := fmt.Sprintf("https://api.github.com/orgs/%s/repos?page=%d&per_page=%d", orgName, page, perPage)
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))
		client := &http.Client{Transport: proxy}
		response, err := client.Do(request)
		//response, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch repositories for organization '%s'. Status: %s", orgName, response.Status)
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var repositories []Repository
		err = json.Unmarshal(body, &repositories)
		if err != nil {
			return nil, err
		}

		//for i := range repositories {
		//	if filterRepositoryName[repositories[i].Name] {
		//		continue
		//	}
		//	version, err := getLatestRelease(orgName, repositories[i].Name, authToken)
		//	if err != nil {
		//		return nil, err
		//	}
		//	repositories[i].Version = version
		//}

		for _, repo := range repositories {
			allRepositories = append(allRepositories, repo)
		}

		// 如果当前页返回的仓库数量少于每页的数量，则说明已经获取完所有仓库，退出循环
		if len(repositories) < perPage {
			break
		}

		page++
	}

	return allRepositories, nil
}

var repoFileName = "repo.json"

func queryAllLatestRelease() {
	orgName := githubOrg
	file, err := ioutil.ReadFile(repoFileName)
	if err != nil {
		panic(err)
	}
	var repositories []*Repository
	err = json.Unmarshal(file, &repositories)
	if err != nil {
		panic(err)
	}
	//version, err := getLatestRelease(orgName, repositories[0].Name, authToken)
	//if err != nil {
	//	fmt.Println("getLatestRelease err:" + err.Error())
	//}
	//fmt.Println(version)

	for _, repo := range repositories {
		version, err := getLatestRelease(orgName, repo.Name, authToken)
		if err != nil {
			fmt.Println("getLatestRelease err:" + err.Error())
		}
		repo.Version = version
	}
	//err = saveRepositoriesToFile(repositories, repoFileName)
	data, err := json.Marshal(repositories)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(repoFileName, data, 0644)
	if err != nil {
		panic(err)
	}
}

func saveRepositoriesToFile(repositories []Repository, filename string) error {
	//data, err := json.MarshalIndent(repositories, "", "  ")
	data, err := json.Marshal(repositories)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Repositories saved to %s\n", filename)
	return nil
}

type Request struct {
	Category  string `json:"category"`
	Desc      string `json:"desc"`
	Namespace string `json:"namespace"`
	Revision  string `json:"revision"`
	Type      string `json:"type"`
	RepoAddr  string `json:"repoAddr"`
	Icon      string `json:"icon"`
}

func requestIacStoreAll() {
	file, err := ioutil.ReadFile(repoFileName)
	if err != nil {
		panic(err)
	}
	var repositories []*Repository
	err = json.Unmarshal(file, &repositories)
	if err != nil {
		panic(err)
	}
	startname := "terraform-alicloud-vpc-ecs-rds-slb"
	isstart := false
	var errReop []string
	for _, repo := range repositories {
		if !isstart {
			if repo.Name == startname {
				isstart = true
			} else {
				continue
			}
		}
		if repo.Version == "" {
			continue
		}
		err = requestIacStore(*repo)
		if err != nil {
			//panic(err)
			errReop = append(errReop, repo.Name)
			fmt.Println(err.Error())
		}

		time.Sleep(5 * 1e9)
	}
	if len(errReop) > 0 {
		fmt.Println(strings.Join(errReop, ";"))
	}
	//repo := repositories[0]
	//err = requestIacStore(*repo)
	//if err != nil {
	//	panic(err)
	//}
}

var iacUrl = "http://iacstore-prod.yun.idcos/api/v1/modules/release"
var iacAuthToken = "eck_8226e4dfcf88418fb8f48c2802246863"
var githubOrg = "alibabacloud-automation"

func requestIacStore(repo Repository) error {

	r := &Request{
		Category:  "alicloud",
		Desc:      repo.Description,
		Namespace: "alicloud",
		Revision:  repo.Version,
		Type:      "common",
		RepoAddr:  "https://github.com/" + githubOrg + "/" + repo.Name + ".git",
		Icon:      "/api/v1/icons?path=icons/default/alicloud.png",
	}
	body, err := json.Marshal(r)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, iacUrl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", iacAuthToken))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		resbody, _ := ioutil.ReadAll(response.Body)
		var m map[string]interface{}
		_ = json.Unmarshal(resbody, &m)
		c := m["code"]
		if c.(float64) != float64(303101) {
			return fmt.Errorf("failed to post release. Status: %s,%s", response.Status, string(resbody))
		}
	}

	fmt.Println("Release posted successfully!")
	return nil
}

func syncRepo() {
	group := githubOrg
	repositories, err := getOrganizationRepositories(group)
	if err != nil {
		panic(err)
	}
	//for _, repo := range repositories {
	//	fmt.Printf("Repository: %s\n", repo.Name)
	//	fmt.Printf("Description: %s\n", repo.Description)
	//	fmt.Printf("Version: %s\n", repo.Version)
	//	fmt.Println()
	//}
	fmt.Println(strconv.Itoa(len(repositories)))
	err = saveRepositoriesToFile(repositories, "repo.json")
	if err != nil {
		fmt.Printf("Failed to save repositories: %s\n", err)
		return
	}
}

func main() {

	//syncRepo()
	//queryAllLatestRelease()
	requestIacStoreAll()
}

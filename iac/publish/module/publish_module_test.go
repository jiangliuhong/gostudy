package module

import (
	"encoding/json"
	"os"
	"testing"
)

type ReqBody struct {
	Category  string `json:"category"`
	Desc      string `json:"desc"`
	Namespace string `json:"namespace"`
	Revision  string `json:"revision"`
	Type      string `json:"type"`
	RepoAddr  string `json:"repoAddr"`
	Icon      string `json:"icon"`
}

func TestGenRepoJson(t *testing.T) {
	syncRepo()
	queryAllLatestRelease()
	t.Logf("success")
}

func TestRequestIacStoreAll(t *testing.T) {
	requestIacStoreAll()
	t.Logf("success")
}

func TestRequestIacStoreDelete(t *testing.T) {
	err := queryAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("success")
}

func TestPublishModule(t *testing.T) {
	file, err := os.ReadFile("repo.json")
	if err != nil {
		t.Fatal(err)
	}
	var modules []Module
	err = json.Unmarshal(file, &modules)
	if err != nil {
		t.Fatal(err)
		return
	}
	//for _, m := range modules {
	//	r := &ReqBody{
	//		Category: m.
	//	}
	//}
}

//repoAddr=`echo $list | jq ".[$index].repoAddr";`
//category=`echo $list | jq ".[$index].category";`
//desc=`echo $list | jq ".[$index].desc";`
//namespace=`echo $list | jq ".[$index].namespace";`
//revision=`echo $list | jq ".[$index].revision";`
//type=`echo $list | jq ".[$index].type";`
//icon=`echo $list | jq ".[$index].icon";`
//# request $repoAddr $category $desc $namespace $revision $type $icon
//raw='{"category":'$category',"desc":'$desc',"namespace":'$namespace',"revision":'$revision',"type":'$type',"repoAddr":'$repoAddr',"icon":'$icon'}'
//echo $raw
//cmd=`curl -s -X POST \
//    --url http://iacstore-prod.yun.idcos/api/v1/modules/release \
//    --header 'Authorization: Bearer eck_8226e4dfcf88418fb8f48c2802246863' \
//    --header 'content-type: application/json' \
//    --data "$raw"`

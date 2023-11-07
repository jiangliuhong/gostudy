package main

import (
	"fmt"
	"github.com/apenella/go-ansible/pkg/options"
	"iac/ansible"
	"math/rand"
	"sync"
)

const base = "/Users/jiangliuhong/Documents/terrform/test"

func randStr(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	var wg sync.WaitGroup
	exe := ansible.ExecRequest{
		Playbooks: []string{"/Users/jiangliuhong/Documents/terrform/test/backup.yml"},
		ConnectionOptions: &options.AnsibleConnectionOptions{
			Connection: "ssh",
			User:       "root",
			PrivateKey: "~/.ssh/aliyun-key",
		},
		Inventory: "cmp-510308-core.yun.idcos,",
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(exec ansible.ExecRequest) {
			defer wg.Done()
			exec.Envs = map[string]interface{}{
				"name": randStr(5),
			}
			_, _, err := ansible.RunPlaybook(&exec)
			if err != nil {
				fmt.Println(err.Error())
			}
		}(exe)
	}
	wg.Wait()
	fmt.Printf("complete")
}

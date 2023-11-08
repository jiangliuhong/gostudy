package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/execute/measure"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
	"github.com/apenella/go-ansible/pkg/stdoutcallback/results"
	"github.com/fatih/color"
	"io"
	"math/rand"
	"time"
)

func randStr(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// https://github.com/apenella/go-ansible/blob/master/examples/ansibleplaybook-with-executor-time-measurament/ansibleplaybook-with-executor-time-measurament.go

func main() {
	var res *results.AnsiblePlaybookJSONResults
	var timeout int
	flag.IntVar(&timeout, "timeout", 32400, "Timeout in seconds")
	flag.Parse()

	buff := new(bytes.Buffer)

	fmt.Printf("Timeout: %d seconds\n", timeout)
	// 配置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	// 配置连接方式， "local" 表示，无论你后面 Inventory 选项怎么配，都是在执行 ansible 的本地执行
	// 要真正连接到 Inventory 配置的机器，注释掉 Connection 选项或者使用 "smart" 或 "ssh" 作为参数值
	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "ssh",
		User:       "root",
		PrivateKey: "~/.ssh/aliyun-key",
	}

	// 资产清单和变量文件，也可以是一个 map 类型来作为变量，就无需引入文件
	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "cmp-510308-core.yun.idcos,",
		//ExtraVars: map[string]interface{}{
		//	"name": randStr(5),
		//},
		ExtraVarsFile: []string{
			"@./test/vars-file.yml",
		},
	}

	durationBuff := new(bytes.Buffer)

	executorTimeMeasurement := measure.NewExecutorTimeMeasurement(
		execute.NewDefaultExecute(
			execute.WithEnvVar("ANSIBLE_FORCE_COLOR", "true"),
			execute.WithTransformers(
				results.Prepend("Go-ansible example"),
				results.LogFormat(results.DefaultLogFormatLayout, results.Now),
			),
			execute.WithShowDuration(),
		),
		measure.WithShowDuration(),
		measure.WithWrite(durationBuff),
	)

	// 构造 ansible
	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         []string{"./test/backup.yml"},
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		Exec:              executorTimeMeasurement,
		// Exec: executorTimeMeasurement,
		StdoutCallback: "json",
	}

	// 执行 playbook
	err := playbook.Run(ctx)
	if err != nil {
		panic(err)
		//fmt.Printf(err.Error())
	}

	color.Cyan("\n\t%s\n", durationBuff.String())

	// 输出结果
	res, err = results.ParseJSONResultsStream(io.Reader(buff))
	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res.Stats)
	fmt.Printf(string(marshal))
}

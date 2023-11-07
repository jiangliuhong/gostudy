package ansible

import (
	"bytes"
	"context"
	"fmt"
	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/execute/measure"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
	"github.com/apenella/go-ansible/pkg/stdoutcallback/results"
	"github.com/fatih/color"
	"io"
	"time"
)

type ExecRequest struct {
	Playbooks         []string                          // 工作文件
	ConnectionOptions *options.AnsibleConnectionOptions // 连接配置
	Inventory         string
	Envs              map[string]interface{}
}

func RunPlaybook(r *ExecRequest) (*results.AnsiblePlaybookJSONResults, string, error) {
	//var timeout int
	//flag.IntVar(&timeout, "timeout", 32400, "Timeout in seconds")
	//flag.Parse()
	timeout := 32400
	jsonBuff := new(bytes.Buffer)
	durationBuff := new(bytes.Buffer)
	fmt.Printf("Timeout: %d seconds\n", timeout)
	// 配置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

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
		Playbooks:         r.Playbooks,
		ConnectionOptions: r.ConnectionOptions,
		Options: &playbook.AnsiblePlaybookOptions{
			Inventory: r.Inventory,
			ExtraVars: r.Envs,
		},
		Exec: executorTimeMeasurement,
		// Exec: executorTimeMeasurement,
		StdoutCallback: "json",
	}

	color.Cyan("\n\t%s\n", durationBuff.String())
	durationByte, err := io.ReadAll(io.Reader(durationBuff))
	if err != nil {
		return nil, "", err
	}
	durationString := string(durationByte)
	err = playbook.Run(ctx)
	if err != nil {
		return nil, durationString, err
	}

	res, err := results.ParseJSONResultsStream(io.Reader(jsonBuff))
	return res, durationString, err
}

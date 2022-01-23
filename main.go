package main

import (
	"fmt"
	"github.com/AgentGuo/kube-scheduler-plugins/pkg/namescore"
	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	command := app.NewSchedulerCommand(
		// app.WithPlugin(mem.Name, mem.New),
		app.WithPlugin(namescore.Name, namescore.New),
	)

	logs.InitLogs()
	defer logs.FlushLogs()
	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

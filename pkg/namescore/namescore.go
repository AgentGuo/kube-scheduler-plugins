package namescore

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const Name = "ScoreByNameLength"

type ScoreByNameLengthPlugin struct {
}

func (s ScoreByNameLengthPlugin) Name() string {
	return Name
}

func (s ScoreByNameLengthPlugin) Score(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) (int64, *framework.Status) {
	var score int64 = 0
	for i := 0; i < len(nodeName); i++ {
		score += int64(nodeName[i])
		score %= 10
	}
	return score, nil
}

func (s ScoreByNameLengthPlugin) ScoreExtensions() framework.ScoreExtensions {
	return nil
}

func New(configuration runtime.Object, f framework.Handle) (framework.Plugin, error) {
	return &ScoreByNameLengthPlugin{}, nil
}

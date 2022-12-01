package pkg

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/apis/extender/v1"
)

const (
	versionPath      = "/version"
	apiPrefix        = "/scheduler"
	bindPath         = apiPrefix + "/bind"
	preemptionPath   = apiPrefix + "/preemption"
	predicatesPrefix = apiPrefix + "/predicates"
	prioritiesPrefix = apiPrefix + "/priorities"
)

var (
	version string // injected via ldflags at build time

	TruePredicate = Predicate{
		Name: "always_true",
		Func: func(pod v1.Pod, node v1.Node) (bool, error) {
			return true, nil
		},
	}

	ZeroPriority = Prioritize{
		Name: "zero_score",
		Func: func(_ v1.Pod, nodes []v1.Node) (*schedulerapi.HostPriorityList, error) {
			var priorityList schedulerapi.HostPriorityList
			priorityList = make([]schedulerapi.HostPriority, len(nodes))
			for i, node := range nodes {
				priorityList[i] = schedulerapi.HostPriority{
					Host:  node.Name,
					Score: 0,
				}
			}
			return &priorityList, nil
		},
	}

	NoBind = Bind{
		Func: func(podName string, podNamespace string, podUID types.UID, node string) error {
			return fmt.Errorf("this extender doesn't support Bind.  Please make 'BindVerb' be empty in your ExtenderConfig")
		},
	}

	EchoPreemption = Preemption{
		Func: func(
			_ v1.Pod,
			_ map[string]*schedulerapi.Victims,
			nodeNameToMetaVictims map[string]*schedulerapi.MetaVictims,
		) map[string]*schedulerapi.MetaVictims {
			return nodeNameToMetaVictims
		},
	}
)

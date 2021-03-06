package api

import (
	"context"
	"github.com/solo-io/valet/pkg/client/aws"
	"github.com/solo-io/valet/pkg/client/helm"
	"github.com/solo-io/valet/pkg/client/kube"
	"github.com/solo-io/valet/pkg/cmd"
	"github.com/solo-io/valet/pkg/render"
)

type Step interface {
	GetDescription(ctx *WorkflowContext, values render.Values) (string, error)
	Run(ctx *WorkflowContext, values render.Values) error
	GetDocs(ctx *WorkflowContext, values render.Values, flags render.Flags) (string, error)
}

type WorkflowContext struct {
	Ctx context.Context
	//Logger
	//SharedState
	Runner       cmd.Runner
	FileStore    render.FileStore
	HelmClient   helm.Client
	KubeClient   kube.Client
	AwsDnsClient aws.DnsClient
}

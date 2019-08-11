package cluster

import (
	"github.com/solo-io/go-utils/cliutils"
	"github.com/solo-io/kube-cluster/cli/internal"
	"github.com/solo-io/kube-cluster/cli/options"
	"github.com/spf13/cobra"
)

func ClusterCmd(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cluster",
		Short:   "interacting with kube clusters",
		RunE: func(cmd *cobra.Command, args []string) error {
			return internal.RootAddError
		},
	}

	cmd.AddCommand(EnsureCmd(opts))
	cliutils.ApplyOptions(cmd, optionsFunc)
	return cmd
}


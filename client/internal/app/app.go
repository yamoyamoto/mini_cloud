package app

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/yamoyamoto/mini_cloud/client/internal/log"
	"os"
)

type App struct{}

func New() *App {
	return &App{}
}

var (
	hostArg       string
	instanceIpArg string
)

func (a *App) Run(ctx context.Context) error {
	rootCmd := cobra.Command{
		Use:   "client-cli",
		Short: "client-cli is a command line interface for mini-cloud",
	}

	createInstanceCmd := &cobra.Command{
		Use:   "create-instance",
		Short: "Create a new instance",
		Run: func(cmd *cobra.Command, args []string) {
			if err := runCreateInstanceCmd(ctx, hostArg, instanceIpArg); err != nil {
				log.FromContext(ctx).Error(err.Error())
				os.Exit(1)
			}
		},
	}
	createInstanceCmd.Flags().StringVar(&hostArg, "host", "h", "firecracker's host")
	createInstanceCmd.Flags().StringVar(&hostArg, "instance-ip", "i", "instance's ip address")
	rootCmd.AddCommand(createInstanceCmd)

	return rootCmd.Execute()
}

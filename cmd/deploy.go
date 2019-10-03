package cmd

import (
	"time"

	"github.com/spf13/cobra"

	"mydevops/pkg"
)

func init() {
	deployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy miaoyun",
		RunE:  runDeploy,
	}
	deployCmd.Flags().StringP("file", "f", "", "Specify the file path")

	RootCmd.AddCommand(deployCmd)
}

func runDeploy(cmd *cobra.Command, args []string) error {
	start := time.Now()
	defer pkg.PrintDone(start)

	path, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}

	deploy, err := pkg.ParseDeployment(path)
	if err != nil {
		return err
	}

	fl := pkg.NewFileLock(path)
	if err := fl.TryLock(1 * time.Hour); err != nil {
		return err
	}
	defer fl.Unlock()

	return deploy.Deploy()
}

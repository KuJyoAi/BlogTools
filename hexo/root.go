package hexo

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use: "hexo",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

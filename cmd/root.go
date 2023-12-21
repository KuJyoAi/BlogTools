package cmd

import (
	"blog-tools/hexo"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "hexo",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(hexo.VersionCmd())
}

package hexo

import (
	"github.com/spf13/cobra"
)

func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Hexo",
		Long:  `All software has versions. This is Hexo's`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Hexo-go beta")
		},
	}
}

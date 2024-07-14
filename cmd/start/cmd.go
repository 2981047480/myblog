package start

import (
	"os"
	"vblog/common/config"

	"github.com/spf13/cobra"
)

var (
	testParam string
)

var Cmd = &cobra.Command{
	Use:   "start",
	Short: "start vblog service",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := os.Getenv("CONFIG_PATH")
		if configPath == "" {
			configPath = config.Filename
		}

		cobra.CheckErr(config.ReadDBConf(configPath).Application.Start())
	},
}

func init() {
	Cmd.Flags().StringVarP(&testParam, "test", "t", "test", "test flag")
}

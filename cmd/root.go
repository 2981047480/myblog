package cmd

import (
	"fmt"
	initCmd "vblog/cmd/init"
	"vblog/cmd/start"
	"vblog/common/config"
	"vblog/ioc"

	"github.com/spf13/cobra"

	_ "vblog/app"
)

var (
	configPath string
)

var RootCmd = &cobra.Command{
	Use:   "vblog",
	Short: "vblog service",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if args[0] == "version" {
				fmt.Println("Current version: v0.0.1")
			}
		} else {
			cmd.Help()
		}
	},
}

func Execute() error {
	// 初始化执行需要的配置
	cobra.OnInitialize(func() {
		cobra.CheckErr(config.LoadFromYaml())
		cobra.CheckErr(ioc.ControllerImpl.Init())
		cobra.CheckErr(ioc.Api.Init())
	})

	return RootCmd.Execute()
}

func init() {
	// --config/-c
	RootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", config.Filename, "The config file of service")

	// Root-->init
	RootCmd.AddCommand(initCmd.Cmd)
	// Root-->start
	RootCmd.AddCommand(start.Cmd)
}

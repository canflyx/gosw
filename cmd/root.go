package cmd

import (
	"fmt"
	"os"

	"github.com/canflyx/gosw/version"
	"github.com/spf13/cobra"
)

var vers bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "demo-api",
	Short: "demo-api 后端API",
	Long:  "demo-api 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
			return nil
		}
		return nil
	},
}

func Execute(defCmd string) {
	var cmdFound bool
	cmd := RootCmd.Commands()
	// 查找注册的命令，并与输入的命令进行对比，判断参数里面已含默认命令
	for _, a := range cmd {
		for _, b := range os.Args[1:] {
			if a.Name() == b {
				cmdFound = true
				break
			}
		}
	}
	if !cmdFound && len(os.Args) < 2 {
		// args := append([]string{defCmd}, os.Args[1:]...)
		RootCmd.SetArgs([]string{defCmd})
	}
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "print demo-api version")
}

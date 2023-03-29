package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/canflyx/gosw/app"
	_ "github.com/canflyx/gosw/apps/all"
	"github.com/canflyx/gosw/conf"
	"github.com/spf13/cobra"
)

var confFile string

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start switches api",
	Long:  "start switches api",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := conf.LoadConfigFromJson(confFile)
		if err != nil {
			fmt.Println(err)
		}
		if err := loadGlobalLogger(); err != nil {
			return err
		}
		// 初始化全局app
		if err := app.InitAllApp(); err != nil {
			return err
		}

		svc := newManager()
		ch := make(chan os.Signal, 1)
		defer close(ch)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
		go svc.WaitStop(ch)
		return svc.Start()
	},
}

func (m *manager) WaitStop(ch <-chan os.Signal) {
	for v := range ch {
		switch v {
		default:
			m.l.Infof("received signal:%s", v)
			m.http.Stop()
		}
	}
}
func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "config.json", "配置文件路径")
	RootCmd.AddCommand(StartCmd)
}

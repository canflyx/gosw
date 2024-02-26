package tools

import (
	"slices"
	"strings"

	"github.com/canflyx/gosw/conf"
)

// 对自定义命令格式化
func FormatCmd(cmd string) []conf.CMD {
	var macCmd []conf.CMD
	if !strings.ContainsAny(cmd, ";,") {
		cmdAndFlag := strings.Split(cmd, ",")
		a := conf.CMD{}
		if len(cmdAndFlag) == 2 {
			a.CMD = cmdAndFlag[0]
			a.CMDFlag = cmdAndFlag[1]

		} else {
			a.CMD = cmd
		}
		macCmd = append(macCmd, a)
		return macCmd
	}

	cmds := strings.Split(cmd, ";")
	for _, cmd := range cmds {
		a := conf.CMD{}
		cmdAndFlag := strings.Split(cmd, ",")
		if len(cmdAndFlag) == 2 {
			a.CMD = cmdAndFlag[0]
			a.CMDFlag = cmdAndFlag[1]

		} else {
			a.CMD = cmd
		}
		macCmd = append(macCmd, a)
	}
	return macCmd
}

// 对命令进行组合
func SplicCmd(cmds conf.TelnetCmd, flag ...int) []conf.CMD {
	var cmdOk []conf.CMD
	// 预处理，比如 user-interface
	if len(cmds.PreCmd) > 0 {
		cmdOk = slices.Concat(cmdOk, cmds.PreCmd)
	}
	for _, val := range flag {
		if val == 1 {
			cmdOk = append(cmdOk, conf.CMD{CMD: cmds.CoreCmd})
		}
		if val == 0 {
			cmdOk = append(cmdOk, conf.CMD{CMD: cmds.AccessCmd})
		}
	}
	// // user 模式下运行
	// if len(cmds.UserCmd) > 0 {
	// 	cmdOk = slices.Concat(cmdOk, cmds.UserCmd)
	// }

	// // en 模式下运行
	// if len(cmds.EnCmd) > 0 {
	// 	cmdOk = append(cmdOk, conf.CMD{CMD: cmds.EnableCmd, CMDFlag: cmds.EnableFlag})
	// 	cmdOk = slices.Concat(cmdOk, cmds.EnCmd)
	// }
	return cmdOk
}

// 对命令进行组合
func SplicShhCmd(cmds conf.TelnetCmd, flag ...int) []string {
	cmdOk := make([]string, 0)
	// 预处理，比如 user-interface
	if len(cmds.PreCmd) > 0 {
		for _, c := range cmds.PreCmd {
			cmdOk = append(cmdOk, c.CMD)
		}
	}
	for _, val := range flag {
		if val == 1 {
			cmdOk = append(cmdOk, cmds.CoreCmd)
		}
		if val == 0 {
			cmdOk = append(cmdOk, cmds.AccessCmd)
		}
	}
	// // user 模式下运行
	// if len(cmds.UserCmd) > 0 {
	// 	for _, c := range cmds.UserCmd {
	// 		cmdOk = append(cmdOk, c.CMD)
	// 	}
	// }

	// // en 模式下运行
	// if len(cmds.EnCmd) > 0 {
	// 	cmdOk = append(cmdOk, cmds.EnableCmd)
	// 	for _, c := range cmds.EnCmd {
	// 		cmdOk = append(cmdOk, c.CMD)
	// 	}
	// }
	if len(cmds.ExitCmd) > 0 {
		for _, c := range cmds.ExitCmd {
			cmdOk = append(cmdOk, c.CMD)
		}
	}
	return cmdOk
}

package impl

import (
	"errors"
	"fmt"
	"log/slog"
	"unsafe"

	"github.com/canflyx/gosw/apps/tools"
	"github.com/canflyx/gosw/conf"
	ssh "github.com/shenbowei/switch-ssh-go"
)

type SshConfig struct {
	log *slog.Logger
}

func (ss *SshConfig) Process(params interface{}) (interface{}, error) {
	if sw, ok := params.(*SwitchesConfig); !ok {
		return errors.New("ssh Config input type error"), nil
	} else {

		ss.log = conf.GetNameLog("ssh")
		cmd := sw.BrandCMD

		// 对自定义进行分别处理,readbytes 或 expect 截断byte
		if sw.Flag == 2 {
			cmds := make([]string, 0)
			for _, c := range cmd.UserCmd {
				cmds = append(cmds, c.CMD)
			}
			result, err := ssh.RunCommands(sw.User, sw.Password, sw.Ip+":22", cmds...)
			if err != nil {
				return nil, err
			}
			return unsafe.Slice(unsafe.StringData(result), len(result)), nil
		}
		cmds := tools.SplicShhCmd(cmd, sw.Flag)
		result, err := ssh.RunCommands(sw.User, sw.Password, sw.Ip+":22", cmds...)
		if err != nil {
			return nil, err
		}
		ss.log.Info("ssh return data len = " + fmt.Sprint(len(result)))
		return unsafe.Slice(unsafe.StringData(result), len(result)), nil
	}
}

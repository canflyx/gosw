package tools

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/canflyx/gosw/conf"
)

func SlogInit() *slog.Logger {
	lc := conf.C().Log
	flog := os.Stdout
	if lc.To == conf.ToFile {
		flog, err := os.OpenFile(filepath.Join(lc.PathDir+"log.txt"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		defer flog.Close()

	}
	if lc.Format == "JSON" {
		return slog.New(slog.NewJSONHandler(flog, nil))
	}
	return slog.New(slog.NewTextHandler(flog, nil))
}

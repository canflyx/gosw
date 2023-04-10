package conf

import (

	// _ "github.com/go-sql-driver/mysql"
	"sync"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func newConfig() *Config {
	return &Config{
		App:        newDefaultAPP(),
		Log:        newDefaultLog(),
		Sqlite:     newDefaultSqlite(),
		TelnetCmds: newDefaultTelnetCmd(),
	}
}

// Config 应用配置
type Config struct {
	App        *App         `json:"app" yaml:"app"`
	Log        *log         `json:"log"  yaml:"log"`
	Sqlite     *Sqlite3     `json:"sql_type" yaml:"sql_type"`
	TelnetCmds []*TelnetCmd `json:"telnet_cmds" yaml:"telnet_cmds"`
}

func (c *Config) TelnetCmd() map[string]TelnetCmd {
	var brandMap = make(map[string]TelnetCmd)
	if len(c.TelnetCmds) < 1 {
		return nil
	}
	for _, u := range c.TelnetCmds {
		brandMap[u.Brand] = *u
	}
	return brandMap
}

type App struct {
	Name          string `json:"name" env:"APP_NAME" yaml:"name"`
	EncryptKey    string `json:"encrypt_key" env:"APP_ENCRYPT_KEY" yaml:"encrypt"`
	HTTP          *http  `json:"http" yaml:"http"`
	TelnetTimeout int64  `json:"telnet_timeout"`
}

func newDefaultAPP() *App {
	return &App{
		Name: "go-switch",
		// EncryptKey: "defualt app encrypt key",
		HTTP:          newDefaultHTTP(),
		TelnetTimeout: 5,
	}
}
func (a *App) GetTime() time.Duration {
	return time.Duration(a.TelnetTimeout) * time.Second
}

type http struct {
	Host      string `json:"host" env:"HTTP_HOST" yaml:"host"`
	Port      string `json:"port" env:"HTTP_PORT" yaml:"port"`
	EnableSSL bool   `json:"enable_ssl" env:"HTTP_ENABLE_SSL" yaml:"enable_ssl"`
	CertFile  string `json:"cert_file" env:"HTTP_CERT_FILE" yaml:"cert_file"`
	KeyFile   string `json:"key_file" env:"HTTP_KEY_FILE" yaml:"key_file"`
}

func (a *http) Addr() string {
	return a.Host + ":" + a.Port
}

func newDefaultHTTP() *http {
	return &http{
		Host: "127.0.0.1",
		Port: "8055",
	}
}

type log struct {
	Level   string    `json:"level" env:"LOG_LEVEL" yaml:"level"`
	PathDir string    `json:"path_dir" env:"LOG_PATH_DIR" yaml:"path_dir"`
	Format  LogFormat `json:"format" env:"LOG_FORMAT" yaml:"format"`
	To      LogTo     `json:"to" env:"LOG_TO" yaml:"to"`
}

func newDefaultLog() *log {
	return &log{
		Level:   "debug",
		PathDir: "logs",
		Format:  "text",
		To:      "stdout",
	}
}

type Sqlite3 struct {
	Path string `json:"path" env:"SQLITE_PATH" yaml:"path"`
}

func newDefaultSqlite() *Sqlite3 {
	return &Sqlite3{
		Path: "./app.db",
	}
}

func (s *Sqlite3) GetDB() *gorm.DB {
	var err error
	var once sync.Once

	// 打开 sqlite 并配置 Logger,不需要则删除
	db, err = gorm.Open(sqlite.Open(s.Path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	once.Do(s.CreateTables)
	if err != nil {
		return nil
	}
	return db
}

func (s *Sqlite3) CreateTables() {
	sql := `CREATE TABLE arp_lists (
		id          INTEGER,
		created_at  DATETIME,
		updated_at  DATETIME,
		deleted_at  DATETIME,
		arp_ip      TEXT     UNIQUE ON CONFLICT REPLACE,
		mac_address TEXT,
		note        TEXT,
		PRIMARY KEY (
			id
		)
	);
	
	
	CREATE TABLE mac_addrs (
		id          INTEGER,
		created_at  DATETIME,
		updated_at  DATETIME,
		deleted_at  DATETIME,
		mac_address TEXT     UNIQUE ON CONFLICT REPLACE,
		port        TEXT,
		switch_ip   TEXT,
		note        TEXT,
		PRIMARY KEY (
			id
		)
	);
	
	CREATE TABLE switches (
		id         INTEGER,
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME,
		ip         TEXT ,
		user       TEXT,
		password   TEXT,
		is_core    INTEGER,
		status     INTEGER,
		brand      STRING,
		note       TEXT,
		PRIMARY KEY (
			id
		)
	);
	CREATE TABLE scan_logs (
		id         INTEGER,
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME,
		switch_id  INTEGER,
		log        TEXT,
		PRIMARY KEY (
			id
		)
	);
	
	`
	if !db.Migrator().HasTable("switches") {
		db.Exec(sql)
	}
}

type TelnetCmd struct {
	Brand        string `json:"brand" yaml:"brand"`
	UserFlag     string `json:"user_flag" yaml:"user_flag"`
	PasswordFlag string `json:"password_flag" yaml:"password_flag"`
	LoginFlag    string `json:"login_flag" yaml:"login_flag"`
	EnableCmd    string `json:"enable_cmd" yaml:"enable_cmd" `
	EnableFlag   string `json:"enable_flag" yaml:"enable_flag"`
	Cmds         []CMD  `json:"cmds" yaml:"cmds"`
	ReadCmd      []CMD  `json:"read_cmd" yaml:"read_cmd"`
	AccessCmd    string `json:"access_cmd" yaml:"access_cmd"`
	CoreCmd      string `json:"core_cmd" yaml:"core_cmd"`
	ReadFlag     string `json:"read_flag" yaml:"read_flag"`
	ExitCmds     []CMD  `json:"exit_cmds"  yaml:"exit_cmds"`
}
type CMD struct {
	CMD     string `json:"cmd" yaml:"cmd"`
	CMDFlag string `json:"cmd_flag" yaml:"cmd_flag"`
}

func newDefaultTelnetCmd() []*TelnetCmd {
	var cmds []*TelnetCmd
	defaultCmd := &TelnetCmd{
		Brand:        "default",
		UserFlag:     "name:",
		PasswordFlag: "ssword:",
		LoginFlag:    ">",
		EnableCmd:    "sys",
		EnableFlag:   "]",
		Cmds:         []CMD{{"user-interface vty 0 4", "]"}, {"screen-length 0", "]"}},
		AccessCmd:    "dis mac-address",
		CoreCmd:      "dis arp",
		ReadFlag:     "]",
		ExitCmds:     []CMD{{"screen-length 50", "]"}},
	}
	cmds = append(cmds, defaultCmd)
	return cmds
}

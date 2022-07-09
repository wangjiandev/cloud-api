package conf

// 全局的config对象
var config *Config

func C() *Config {
	return config
}

func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		MySQL: NewDefaultMySQL(),
		Log:   NewDefaultLog(),
	}
}

type Config struct {
	App   *App   `toml:"app"`
	MySQL *MySQL `toml:"mysql"`
	Log   *Log   `toml:"log"`
}

func NewDefaultApp() *App {
	return &App{
		Name: "demo",
		Host: "0.0.0.0",
		Port: "8080",
	}
}

type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
	// Key  string `toml:"key" env:"APP_KEY"`
	// EnableSSL bool   `toml:"enable_ssl" env:"APP_ENABLE_SSL"`
	// CertFile  string `toml:"cert_file" env:"APP_CERT_FILE"`
	// KeyFile   string `toml:"key_file" env:"APP_KEY_FILE"`
}

func NewDefaultMySQL() *MySQL {
	return &MySQL{
		Host:        "127.0.0.1",
		Port:        "3306",
		Username:    "root",
		Password:    "123456",
		Database:    "cloud-station",
		MaxOpenConn: 10,
		MaxIdleConn: 5,
		MaxLifeTime: 200,
		MaxIdleTime: 100,
	}
}

type MySQL struct {
	Host        string `toml:"host" env:"MYSQL_HOST"`
	Port        string `toml:"port" env:"MYSQL_PORT"`
	Username    string `toml:"username" env:"MYSQL_USERNAME"`
	Password    string `toml:"password" env:"MYSQL_PASSWORD"`
	Database    string `toml:"database" env:"MYSQL_DATABASE"`
	MaxOpenConn int64  `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int64  `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int64  `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int64  `toml:"max_idle_time" env:"MYSQL_MAX_IDLE_TIME"`
}

func NewDefaultLog() *Log {
	return &Log{
		Level:  "debug",
		Format: TextFormat,
		To:     ToStdout,
	}
}

type Log struct {
	Level   string    `toml:"level" env:"LOGGER_LEVEL"`
	Format  LogFormat `toml:"format" env:"LOGGER_FORMAT"`
	To      LogTo     `toml:"to" env:"LOGGER_TO"`
	PathDir string    `toml:"path_dir" env:"LOGGER_PATH_DIR"`
}

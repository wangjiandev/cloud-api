package conf

type Config struct {
	App   *app   `toml:"app"`
	MySQL *MySQL `toml:"mysql"`
	Log   *Log   `toml:"log"`
}

type app struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
	Key  string `toml:"key" env:"APP_KEY"`
	// EnableSSL bool   `toml:"enable_ssl" env:"APP_ENABLE_SSL"`
	// CertFile  string `toml:"cert_file" env:"APP_CERT_FILE"`
	// KeyFile   string `toml:"key_file" env:"APP_KEY_FILE"`
}

type MySQL struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	Username string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
}

type Log struct {
	Level  string    `toml:"level" env:"LOGGER_LEVEL"`
	Dir    string    `toml:"dir" env:"LOGGER_DIR"`
	Format LogFormat `toml:"format" env:"LOGGER_FORMAT"`
	To     LogTo     `toml:"to" env:"LOGGER_TO"`
}

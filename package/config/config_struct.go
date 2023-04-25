package config

// Config Config檔
type Config struct {
	Server Server   `toml:"server"`
	DB     Database `toml:"db"`
}

// Server 服務器設定檔
type Server struct {
	Port string `toml:"port"`
}

// Database 資料庫
type Database struct {
	Account         string `toml:"account"`
	Password        string `toml:"password"`
	Host            string `toml:"host"`
	DBName          string `toml:"db_name"`
	Port            string `toml:"port"`
	ConnMaxLifetime int    `toml:"connMaxLifetime"`
}

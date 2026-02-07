package config

type Config struct {
	AppConfig  AppConfig      `json:"app_config"`
	Postgres   PostgresConfig `json:"postgres"`
	HashConfig HashConfig     `json:"hash"`
	JWTConfig  JWTConfig      `json:"jwt"`
}

type AppConfig struct {
	AppHost string `json:"app_host"`
	AppPort string `json:"app_port"`
}

type PostgresConfig struct {
	URI             string `json:"uri"`
	Database        string `json:"database"`
	ConnectTimeout  int    `json:"connect_timeout"`
	MaxIdleTime     int    `json:"max_idle_time"`
	MaxConnLifetime int    `json:"max_conn_lifetime"`
	MaxPoolSize     uint64 `json:"max_pool_size"`
	MinPoolSize     uint64 `json:"min_pool_size"`
}

type HashConfig struct {
	Salt      string `json:"salt"`
	MinLength int    `json:"min_length"`
}

type JWTConfig struct {
	Secret       string `json:"secret"`
	ExpiryHours  int    `json:"expiry_hours"`
}

package config

type Config struct {
	Addr string
	Env  string
	DB   DBConfig
}

type DBConfig struct {
	ConnStr      string
	MaxOpenConns int    //
	MaxIdleConns int    //
	MaxIdleTime  string //
}

package bootstrap

type Database struct {
	Main  Main
	Slave Slave
}

type Main struct {
	Host       string
	Port       string
	DBName     string `mapstructure:"db_name"`
	Username   string
	Password   string
	MaxCon     int `mapstructure:"max_con"`
	MaxIdleCon int `mapstructure:"max_idle_con"`
}

type Slave struct {
	Host       string
	Port       string
	DBName     string `mapstructure:"db_name"`
	Username   string
	Password   string
	MaxCon     int `mapstructure:"max_con"`
	MaxIdleCon int `mapstructure:"max_idle_con"`
}

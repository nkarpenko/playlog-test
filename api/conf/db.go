package conf

// DBConfig - mysql
type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"pwd"`
	Database string `mapstructure:"db"`
}

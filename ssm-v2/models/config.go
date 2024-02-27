package models

type Server struct {
	HostName string `mapstructure:"hostname"`
	IP       string `mapstructure:"ip"`
	User     string `mapstructure:"user"`
	KeyAuth  bool   `mapstructure:"keyAuth"`
	Alias    string `mapstructure:"alias"`
}

type Environment struct {
	Name    string   `mapstructure:"name"`
	Servers []Server `mapstructure:"servers"`
}

type Group struct {
	Name        string        `mapstructure:"name"`
	User        string        `mapstructure:"user"`
	Environment []Environment `mapstructure:"environment"`
}

type Config struct {
	Groups []Group `mapstructure:"groups"`
}

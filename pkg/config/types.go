package config

type Config struct {
	PackageName string    `yaml:"packageName"`
	Dockerfile  bool      `yaml:"dockerfile"`
	GitIgnore   bool      `yaml:"gitignore"`
	EnvFile     bool      `yaml:"envFile"`
	Container   Container `yaml:"container"`
}

type Container struct {
	DB     DB     `yaml:"db"`
	Server Server `yaml:"server"`
	JWT    JWT    `yaml:"jwt"`
}

type DB struct {
	Connection string `yaml:"connection"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	User       string `yaml:"user"`
	Name       string `yaml:"name"`
	Password   string `yaml:"password"`
}

type Server struct {
	URL  string `yaml:"url"`
	Port string `yaml:"port"`
}

type JWT struct {
	Secret string `yaml:"secret"`
	Expire int    `yaml:"expire"`
}

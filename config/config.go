package config

type Config struct {
	Server   ServerConf
	Database DatabaseConf
	Authz    AuthzConf
	Frontend FrontendConf
}

type ServerConf struct {
	Port int
}

type DatabaseConf struct {
	Host       string
	Port       int
	DBName     string
	DBUser     string
	DBPassword string
	Timezone   string
}

type AuthzConf struct {
	Endpoint string
}

type FrontendConf struct {
	Endpoint string
}

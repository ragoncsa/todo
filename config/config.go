package config

type Config struct {
	Server   ServerConf
	Database DatabaseConf
	Authz    AuthzConf
	Authn    AuthnConf
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
	Disable  bool
}

type AuthnConf struct {
	ClientId    string
	NotEnforced bool
}

type FrontendConf struct {
	Endpoints []string
}

package configs

// ServerConfig config.
type ServerConfigStruct struct {
	Port               string
	BodyLimit          string
	AllowCorsOrigin    string
	AllowCorsMethods   string
	DisableMaintenance string
	ProductName        string
	ModuleName         string
}

var ServerConfig ServerConfigStruct

package config

type Config struct {
	Path             string
	Host             string
	Port             int
	SupplierFilePath string
}

func DefaultConfig() Config {
	return Config{
		Path:             "/api",
		Host:             "0.0.0.0",
		Port:             8081,
		SupplierFilePath: "/suppliers.json",
	}
}

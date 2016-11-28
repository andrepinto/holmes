package config

type Service struct {
	Name string `json:"name"`
	Version string `json:"version"`
	Description string `json:"description"`
	Consul string `json:"consul"`
}

type Api  struct {
	Port int `json:"port"`
}


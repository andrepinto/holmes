package discovery

type Service struct {
	Name string `json:"name"`
	Version string `json:"version"`
	Description string `json:"description"`
}

type Api  struct {
	Port int `json:"port"`
}

type Dependency struct {
	Name string `json:"name"`
	Version string `json:"version"`
}

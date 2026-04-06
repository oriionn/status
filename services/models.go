package services

type Service struct {
 	Name string `toml:"name"`
    URL  string `toml:"url"`
    ShowURL bool `toml:"show_url"`
	Up uint
	Total uint
	Status bool
	Latency int
}

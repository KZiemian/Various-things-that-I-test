type Config struct {}

func WithReticulatedSplines(c *Config) {}

type Terrain struct {
	config Config
}

func NewTerrain(options ...func(*Config)) *Terrain {
	var t Terrain

	for _, option := range options {
		option(&t.config)
	}

	return &t
}

func WithCities(n int) func(*Config) {}

func main() {
	// t := NewTerrain(WithReticulatedSplines)
	t := NewTerrain(WithCities(9))
}

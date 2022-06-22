type Option interface {
	Apply(*Config)
}

func NewTerrain(options ...Option) *Terrain {
	var config Config

	for _, option := range options {
		option.Apply(&t.config)
	}
}

type splines struct {}

func (s *splines) Apply(c *Config) {}

func WithReticulatedSplines() Option {
	return new(splines)
}

type cities struct {
	cities int
}

func (c *cities) Apply(c *Config) {}

func WithCities(n int) Option {
	return &cities{
		cities: n
	}
}

func main() {
	t := NewTerrain(
		WithReticulatedSplines(),
		WithCities(9),
	)
}

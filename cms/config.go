package cms

type Config struct {
	ConfigDir     string
	ContentDir    string
	StaticDir     string
	WebAddr       string
}

func (c *Config) fillDefaultValues() {
	c.ConfigDir = stringOr(c.ConfigDir, "./config")
	c.ContentDir = stringOr(c.ContentDir, "./content")
	c.StaticDir = stringOr(c.StaticDir, "./static")
	c.WebAddr = stringOr(c.WebAddr, ":8080")
}

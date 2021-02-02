package cms

import "path"

type Config struct {
	DataDir    string
	WebAddr    string
	AdminPanel bool
}

func (c Config) ConfigDir() string {
	return path.Join(c.DataDir, "config")
}

func (c Config) StaticDir() string {
	return path.Join(c.DataDir, "static")
}

func (c Config) ContentDir() string {
	return path.Join(c.DataDir, "content")
}

func (c *Config) fillDefaultValues() {
	c.DataDir = stringOr(c.DataDir, "./data")
	c.WebAddr = stringOr(c.WebAddr, ":8080")
}

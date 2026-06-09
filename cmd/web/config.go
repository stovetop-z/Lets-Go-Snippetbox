package main

import (
	"flag"
)

type Config struct {
	addr      string
	staticDir string
}

func (c *Config) ConfigureFlags() {
	flag.StringVar(&c.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&c.staticDir, "static-dir", "./ui/static/", "Path to static assets")
	flag.Parse()
}

package url

import (
	"errors"
	"strings"
)

type URL struct {
	Scheme string
	Host string
	Path string
}

func Parse(rawUrl string) (*URL, error) {
	i := strings.Index(rawUrl, "://")
	if i < 0 {
		return nil, errors.New("Missing scheme")
	}
	scheme, rest := rawUrl[:i], rawUrl[i+3:]
	host, path := rest, ""
	if i := strings.Index(rest, "/"); i >= 0 {
		host, path = rest[:i], rest[i+1:]
	}
	return &URL{scheme, host, path}, nil
}
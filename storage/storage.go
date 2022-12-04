package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"

)

type Storage interface {
	Save(*Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

var ErrNoSavedPage = errors.New("no saved pages")

type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {

	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", fmt.Errorf("writeString %w", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", fmt.Errorf("writeString %w", err)
	}

	return fmt.Sprint(h.Sum(nil)), nil
}

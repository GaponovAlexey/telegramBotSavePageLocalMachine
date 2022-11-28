package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"tg/sitesess.ca/storage"
	"time"
)

type Storage struct {
	basePath string
}

const (
	defaultPerm = 0774
)

var ErrNoSavePage = errors.New("no saved page")

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	fPath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return fmt.Errorf("mkdirall %w", err)
	}

	fName, err := fileName(page)
	if err != nil {
		return fmt.Errorf("filename %w", err)
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return fmt.Errorf("os.Create %w", err)
	}
	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil
}

func (s Storage) PickRandom(userName string) (page *storage.Page, err error) {
	path := filepath.Join(s.basePath, userName)

	files, err := os.ReadDir()
	if err != nil {
		return nil, fmt.Errorf("ReadDei %w", err)
	}

	if len(files) == 0 {
		return nil, ErrNoSavePage
	}
	// 0
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))

	file := files[n]
	return s.decodePage(filepath.Join(path, file.Name()))
}

func (s Storage) Remove(p *storage.Page) error {
	fileName, err := fileName(p)
	if err != nil {
		return fmt.Errorf("fileName Remove %w", err)
	}
	
	path := filepath.Join(s.basePath, p.UserName, fileName)

	if err := os.Remove(path): err != nil {
		return fmt.Errorf("can't ramove file", err)
	}
	return nil
}

func(s Storage) isExist(p *storage.Page) (bool, error) {
	fileName, err := fileName(p)
	if err != nil {
		return false, fmt.Errorf("fileName Remove %w", err)
	}
	path := filepath.Join(s.basePath, p.UserName, fileName)

	switch _,err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err!= nil:
		msg := fmt.Sprintf("can't check if file %s exists", path)
		return false, e.Wrap(msg,err) 
	
		return true, nil
	}
	
}

func (s Storage) decodePage(filePath string) (*storage.Page, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Open filepath %w", err)
	}
	defer func() { _ = f.Close() }()
	var p storage.Page
	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, fmt.Errorf("newDecoder decoderPage %w", err)
	}
	return &p, nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}

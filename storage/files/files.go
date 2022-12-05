package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"tg/sitesess.ca/lib/e"
	"tg/sitesess.ca/storage"

)

type Storage struct {
	basePath string
}

const defaultPerm = 0774

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

// SAVE
func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIfErr("can't save", err) }()

	fPath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}

	fName, err := fileName(page)
	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil
}

// PickRandom
func (s Storage) PickRandom(userName string) (page *storage.Page, err error) {
	defer func() { err = e.WrapIfErr("can't save PickRandom", err) }()
	path := filepath.Join(s.basePath, userName)

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, storage.ErrNoSavedPage
	}
	// 0
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))

	file := files[n]

	return s.decodePage(filepath.Join(path, file.Name()))
}

// Remove
func (s Storage) Remove(p *storage.Page) (err error) {
	defer func() { err = e.WrapIfErr("can't Remove", err) }()
	fileName, err := fileName(p)
	if err != nil {
		return err
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)

	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}

func (s Storage) isExist(p *storage.Page) (b bool, err error) {
	defer func() { err = e.WrapIfErr("can't isExist", err) }()
	fileName, err := fileName(p)
	if err != nil {
		return false, err
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)

	//switch
	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		msg := fmt.Sprintf("can't check if file %s exists", path)
		return false, e.Wrap(msg, err)
	}

	return true, nil
}

func (s Storage) decodePage(filePath string) (sp *storage.Page, err error) {
	defer func() { err = e.WrapIfErr("can't decodePage", err) }()
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	var p storage.Page
	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}

package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

func NewFileStore(storeID, filePath string) *FileStore {
	return &FileStore{
		storeID: storeID,
		filepath: filePath,
	}
}

type FileStore struct {
	storeID string
	filepath string
	mutex sync.Mutex
	keyValues map[string]string
}

func (fs *FileStore) loadKeyValues() error {
	f, err := os.Open(fs.filepath)
	if err != nil {
		fs.keyValues = make(map[string]string)
		return nil
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	err = dec.Decode(&fs.keyValues)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FileStore) saveKeyValues() error {
	var sb strings.Builder
	enc := json.NewEncoder(&sb)
	err := enc.Encode(fs.keyValues)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(fs.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, sb.String())
	if err != nil {
		return err
	}

	return nil
}

func (fs *FileStore) Get(key string) (string, error) {
	err := fs.loadKeyValues()
	if err != nil {
		return "", err
	}
	value, ok := fs.keyValues[key]
	if !ok {
		return "", errors.New(fmt.Sprintf("No value for the key: %s", key))
	}
	return value, nil
}

func (fs *FileStore) Set(key, value string) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()
	err := fs.loadKeyValues()
	if err != nil {
		return err
	}
	fs.keyValues[key] = value
	err = fs.saveKeyValues()
	if err != nil {
		return err
	}
	return nil
}

func (fs *FileStore) Delete(key string) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()
	err := fs.loadKeyValues()
	if err != nil {
		return err
	}
	delete(fs.keyValues, fs.keyValues[key])
	return nil
}
package command

import (
	"errors"
	"os"
	"path/filepath"
	"time"
)

func getTodaysDir() (string, error) {
	// Currently kyouRoot is $HOME/kyou
	home, exists := os.LookupEnv("HOME")
	if !exists {
		return "", errors.New("environment variable $HOME must be defined")
	}

	const layout = "2006-01-02"
	date := time.Now().Format(layout)

	return filepath.Join(home, "kyou", date), nil
}

package filesys

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

var defaultDir = ".bh"

// FilePath returns the default location to read and write files.
func FilePath(filename string) string {
	home := os.Getenv("HOME")
	return path.Join(home, defaultDir, filename)
}

// DirectorySetup is called to ensure that the path to the default
// directory with all the Command Line Tool data exists. If it does not
// exist then this function will create the directory.
func DirectorySetup() {
	bhDirPath := FilePath("")
	if _, err := os.Stat(bhDirPath); os.IsNotExist(err) {
		err := os.Mkdir(bhDirPath, 7666)
		if err != nil {
			log.Fatalf("Unable to create the default CLI data directory: %v", err)
		}
	}
}

func CreateFile(filename string) (*os.File, error) {
	return os.Create(FilePath(filename))
}

// WriteFile writes the file data to the OS/FS.
func WriteFile(filename string, xb []byte) {
	filePath := FilePath(filename)
	err := os.WriteFile(filePath, xb, 0666)
	if err != nil {
		log.Fatalf("Unable to write data to file: %s: %v", filePath, err)
	}
}

// ReadFile reads a file from the OS/FS.
func ReadFile(filename string) []byte {
	filePath := FilePath(filename)
	xb, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Unable to read file: %s: %v", filePath, err)
	}
	return xb
}

func LoadInterface(filename string, v DataStructure) {
	filePath := FilePath(filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		v.Save()
	}
	data := ReadFile(filename)
	err := json.Unmarshal(data, &v)
	if err != nil {
		log.Fatalf("Unable to unmarshal data: %v", err)
	}
}

package services

import (
	"errors"
	"fmt"
	jsonwriter "github.com/ParvizBoymurodov/rest/pkg"
	"github.com/google/uuid"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FilesSvc struct {
	media string
}

func NewFilesSvc(media string) *FilesSvc {
	if media == "" {
		panic(errors.New("media path can't be nil"))
	}
	return &FilesSvc{media: media}
}

func (receiver *FilesSvc) Save(sources io.Reader, contentType string) (name string, err error) {
	var path string
	extension := (strings.Split(contentType, "/"))[1]
	if len(extension) == 0 {
		return "", errors.New("invalid extension")
	}
	uuidV4 := uuid.New().String()
	name = fmt.Sprintf("%s.%s", uuidV4, extension)
	path = filepath.Join(receiver.media, name)
	log.Print(name)
	dst, err := os.Create(path)
	if err != nil {
		log.Print("can't close file")
	}
	defer func() {
		if dst.Close() != nil {
			log.Print("can't close file")
		}
	}()
	_, err = io.Copy(dst, sources)
	if err != nil {
		log.Printf("ca't save file: %v", sources)
	}

	upload, err := jsonwriter.JsonFileUpload(name)
	return upload, nil
}

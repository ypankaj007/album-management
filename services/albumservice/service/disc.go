package service

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"sync"

	"album-management/services/albumservice/config"

	"github.com/google/uuid"
	"github.com/micro/go-micro/util/log"
)

var once sync.Once

// Disk create/delete folder and upload/delete image
type Disk struct {
	RootPath  string
	Permision os.FileMode
	//mutex     sync.RWMutex
}

var (
	diskInstance *Disk
)

// NewDisk is singleton disk object
func NewDisk() *Disk {

	once.Do(func() { // <-- atomic, does not allow repeating

		diskInstance = &Disk{RootPath: config.GetENV("ROOT_PATH"), Permision: 0755} // <-- thread safe

	})

	return diskInstance
}

// CreateFolder create folder in disk
func (disk *Disk) CreateFolder(path string) error {
	path = disk.RootPath + path
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, disk.Permision)
		return err
	}
	if os.IsExist(err) {
		return nil
	}
	return err
}

// Remove folder or image
func (disk *Disk) Remove(path string) error {
	path = disk.RootPath + path
	return os.RemoveAll(path)
}

// Save image to disk
func (disk *Disk) Save(path string, imageData bytes.Buffer) (string, error) {
	path = disk.RootPath + path
	imageID, err := uuid.NewRandom()
	if err != nil {
		log.Error(err.Error())
		return "", errors.New("Cannot generate imageID")
	}

	imagePath := fmt.Sprintf("%s/%s", path, imageID)

	file, err := os.Create(imagePath)
	if err != nil {
		log.Error(err.Error())
		return "", errors.New("Cannot create image file")
	}

	_, err = imageData.WriteTo(file)
	if err != nil {
		log.Error(err.Error())
		return "", errors.New("Cannot write image to file")
	}
	return imageID.String(), nil
}

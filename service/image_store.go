package service

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
)

type ImageStore interface {
	Save(laptopID string, imageType string, imageData bytes.Buffer) (string, error)
}

type DiskImagesStore struct {
	mutex sync.RWMutex
	imageFolder string
	images map[string]*ImageInfo
}

type ImageInfo struct {
	LaptopID string
	Type string
	Path string
}

func NewDiskImagesStore(imageFolder string) *DiskImagesStore {
	return &DiskImagesStore{
		imageFolder: imageFolder,
		images:      make(map[string]*ImageInfo),
	}
}

func (s *DiskImagesStore) Save(laptopID string, imageType string, imageData bytes.Buffer) (string, error) {
	imageID, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("cannot generate image id: %w", err)
	}

	imagePath := fmt.Sprintf("%s/%s%s", s.imageFolder, imageID, imageType)
	file, err := os.Create(imagePath)
	if err != nil {
		return "", fmt.Errorf("cannot generate image file: %w", err)
	}

	_, err = imageData.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("cannot write image to file: %w", err)
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.images[imageID.String()] = &ImageInfo{
		LaptopID: laptopID,
		Type:     imageType,
		Path:     imagePath,
	}
	return imageID.String(), nil
}


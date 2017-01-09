package storage

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/storage"
	gcpgcs "cloud.google.com/go/storage"
)

type gcs struct {
	client *gcpgcs.Client
	model  Model
}

const (
	defaultGCSRegion = "us-east1"
	keyPath          = "/tmp/key.json"
)

func NewGCS(m Model) Storage {

	key := m.ServiceAccountKey
	err := ioutil.WriteFile(keyPath, []byte(key), 0644)
	if err != nil {
		log.Fatalf("Failed to create google service account key file: %v", err)
	}

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", keyPath)

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &gcs{
		client: client,
		model:  m,
	}
}

func (g *gcs) Delete(filename string) error {
	return nil
}

func (g *gcs) Download(filename string, writer io.Writer) (Version, error) {
	return Version{}, nil
}

func (g *gcs) Upload(filename string, reader io.Reader) (Version, error) {
	return Version{}, nil
}

func (g *gcs) LatestVersion(filename string) (Version, error) {
	return Version{}, nil
}

func (g *gcs) Version(filename string) (Version, error) {
	return Version{}, nil
}

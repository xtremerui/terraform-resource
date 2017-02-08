package storage

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"

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
	ctx              = context.Background()
)

func NewGCS(m Model) Storage {

	key := m.ServiceAccountKey
	err := ioutil.WriteFile(keyPath, []byte(key), 0644)
	if err != nil {
		log.Fatalf("Failed to create google service account key file: %v", err)
	}

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", keyPath)

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

func (g *gcs) Download(filename string, destination io.Writer) (Version, error) {

	key := path.Join(g.model.BucketPath, filename)

	objHandle := client.Bucket(g.model.Bucket).Object(key)
	r, err := objHandle.NewReader(ctx)
	util.FatalIfError(err)
	data, err := ioutil.ReadAll(r)
	util.FatalIfError(err)
	err = r.Close()
	util.FatalIfError(err)
	attrs, err := objHandle.Attrs(ctx)
	util.FatalIfError(err)
	sha1Hex := attrs.Metadata["SHA1"]
	fmt.Printf("Read '%s' of size %d, sha1: %s\n", path, len(data), sha1Hex)

	version := Version{
		LastModified: data,
		StateFile:    filename,
	}

	return Version, nil
}

func (g *gcs) Upload(filename string, content io.Reader) (Version, error) {
	return Version{}, nil
}

func (g *gcs) LatestVersion(filename string) (Version, error) {
	return Version{}, nil
}

func (g *gcs) Version(filename string) (Version, error) {
	return Version{}, nil
}

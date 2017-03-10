package storage

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"cloud.google.com/go/storage"
)

type gcs struct {
	client *storage.Client
	model  Model
}

const (
	defaultGCSRegion = "us-east1"
	keyPath          = "/tmp/key.json"
)

var ctx = context.Background()

func NewGCS(m Model) Storage {

	// key := m.ServiceAccountKey
	// err := ioutil.WriteFile(keyPath, []byte(key), 0644)
	// if err != nil {
	// 	log.Fatalf("Failed to create google service account key file: %v", err)
	// }

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

	objHandle := g.client.Bucket(g.model.Bucket).Object(key)
	r, err := objHandle.NewReader(ctx)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	err = r.Close()
	if err != nil {
		log.Fatal(err)
	}
	attrs, err := objHandle.Attrs(ctx)
	if err != nil {
		log.Fatal(err)
	}
	sha1Hex := attrs.Metadata["SHA1"]
	fmt.Printf("Read '%s' of size %d, sha1: %s\n", keyPath, len(data), sha1Hex)
	if _, err := io.Copy(destination, r); err != nil {
		// TODO: Handle error.
	}

	objAttrs, err := objHandle.Attrs(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	fmt.Printf("object %s has size %d and is last updated at %s\n",
		objAttrs.Name, objAttrs.Size, objAttrs.Updated)

	version := Version{
		LastModified: time.Now(),
		StateFile:    filename,
	}

	return version, nil
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

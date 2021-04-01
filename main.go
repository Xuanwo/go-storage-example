package main

import (
	"bytes"
	"log"

	"github.com/aos-dev/go-service-fs/v2"
	qingstor "github.com/aos-dev/go-service-qingstor/v2"
	"github.com/aos-dev/go-storage/v3/pairs"
	"github.com/aos-dev/go-storage/v3/pkg/credential"
	"github.com/aos-dev/go-storage/v3/types"
)

func initFS() (store types.Storager) {
	// Init a service.
	store, err := fs.NewStorager(pairs.WithWorkDir("/tmp"))
	if err != nil {
		log.Fatalf("service init failed: %v", err)
	}

	return store
}

func initQingStor() (store types.Storager) {
	store, err := qingstor.NewStorager(
		pairs.WithWorkDir("/tmp"),
		pairs.WithCredential(credential.NewHmac("access_key", "secret_key").String()),
	)
	if err != nil {
		log.Fatalf("service init failed: %v", err)
	}

	return store
}

func main() {
	store := initFS()
	// store := initQingStor()
	content := []byte("Hello, world!")
	length := int64(len(content))
	r := bytes.NewReader(content)

	_, err := store.Write("hello", r, length)
	if err != nil {
		log.Fatalf("write failed: %v", err)
	}

	var buf bytes.Buffer

	_, err = store.Read("hello", &buf)
	if err != nil {
		log.Fatalf("storager read: %v", err)
	}

	log.Printf("%s", buf.String())
}

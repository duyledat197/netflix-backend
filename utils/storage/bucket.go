package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
)

type Bucket struct {
	Client     *storage.Client
	BucketName string
	Bucket     *storage.BucketHandle
}

func NewBucket(bucketName string) (*Bucket, error) {
	ctx := context.Background()
	// b, err := os.ReadFile(path.Join("./google_cloud_key.json"))
	// if err != nil {
	// 	return nil, err
	// }
	// log.Println(string(b))
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	bucket := client.Bucket(bucketName)

	return &Bucket{
		Client:     client,
		BucketName: bucketName,
		Bucket:     bucket,
	}, nil
}

func (b *Bucket) UploadImage(ctx context.Context, file multipart.File, fileName string) (string, error) {
	t := time.Now()
	obj := b.Bucket.Object(fmt.Sprintf("images/%v-%v", t.Format(time.RFC3339), fileName))
	wc := obj.NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		log.Println("io.Copy(wc, file)", err)
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		log.Println("wc.Close()", err)
		return "", fmt.Errorf("Writer.Close: %v", err)
	}
	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		log.Println(err)
		return "", err
	}
	attr := wc.Attrs()

	return objectURL(attr), nil
}

// variable global
const (
	Endpoint = "https://storage.googleapis.com/"
)

// parse object url
func objectURL(objAttrs *storage.ObjectAttrs) string {
	return fmt.Sprintf("%v%s/%s", Endpoint, objAttrs.Bucket, objAttrs.Name)
}

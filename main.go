package main

import (
	"fmt"
	"github.com/minio/minio-go"
	"log"
)

func main() {
	endpoint := ""
	accessKeyID := ""
	secretAccessKey := ""

	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		log.Fatalln(err)
	}

	buckets, err := minioClient.ListBuckets()
	if err != nil {
		log.Fatalln(err)
	}
	for _, bucket := range buckets {
		fmt.Println(bucket.Name)
	}

	obj, err := minioClient.GetObject("garage", "image.jpg", minio.GetObjectOptions{})
	log.Println(obj)
}

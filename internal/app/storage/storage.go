package storage

import (
	"context"
	"fmt"
	"github.com/Point-AI/backend/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func ConnectToStorage(cfg *config.Config) *minio.Client {
	minioClient, err := minio.New("play.min.io", &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinIo.AccessKey, cfg.MinIo.SecretKey, ""),
		Secure: true,
	})

	if err != nil {
		panic(fmt.Errorf("failed to create MinIO client: %w", err))
	}

	//Comment it out once the bucket is created once
	if err = minioClient.MakeBucket(context.Background(), cfg.MinIo.BucketName, minio.MakeBucketOptions{}); err != nil {
		panic(fmt.Errorf("failed to create bucket: %w", err))
	}

	return minioClient
}

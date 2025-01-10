package cloudstorage

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"hangout.com/core/storage-service/config"
	"hangout.com/core/storage-service/files"
	"hangout.com/core/storage-service/logger"
)

func Connect(ctx context.Context, cfg *config.Config, log logger.Log) (*minio.Client, error) {
	log.Info("connecting to Minio/s3")
	useSsl := false
	minioClient, err := minio.New(cfg.Minio.BaseUrl, &minio.Options{Creds: credentials.NewStaticV4(cfg.Minio.AccessKey, cfg.Minio.SecretKey, ""), Secure: useSsl})
	if err != nil {
		log.Error("could not connect to Minio/s3", "url", cfg.Minio.BaseUrl)
	}
	log.Info("Checking if buckets exist")
	_, err = minioClient.BucketExists(ctx, cfg.Minio.UploadBucket)
	if err != nil {
		log.Debug("Upload bucket exists skipping creation")
	} else {
		log.Info("Upload bucket does not exist, Creating a new one")
		err = minioClient.MakeBucket(ctx, cfg.Minio.UploadBucket, minio.MakeBucketOptions{})
		if err != nil {
			log.Error("Error in creating new upload bucket")
		} else {
			log.Info("Upload bucket successfully created")
		}
	}
	_, err = minioClient.BucketExists(ctx, cfg.Minio.StorageBucket)
	if err != nil {
		log.Debug("Storage bucket exists skipping creation")
	} else {
		log.Info("Storage bucket does not exist, Creating a new one")
		err = minioClient.MakeBucket(ctx, cfg.Minio.StorageBucket, minio.MakeBucketOptions{})
		if err != nil {
			log.Error("Error in creating new storage bucket")
		} else {
			log.Info("Storage bucket successfully created")
		}
	}
	return minioClient, err
}
func Download(minioClient *minio.Client, file *files.File, cfg *config.Config, log logger.Log) string {
	return "hello"
}

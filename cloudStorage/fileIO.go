package cloudstorage

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"hangout.com/core/storage-service/config"
	"hangout.com/core/storage-service/files"
	"hangout.com/core/storage-service/logger"
)

func Connect(workerId int, ctx context.Context, cfg *config.Config, log logger.Log) (*minio.Client, error) {
	log.Info("connecting to Minio/s3", "worker-id", workerId)
	useSsl := false
	minioClient, err := minio.New(cfg.Minio.BaseUrl, &minio.Options{Creds: credentials.NewStaticV4(cfg.Minio.AccessKey, cfg.Minio.SecretKey, ""), Secure: useSsl})
	if err != nil {
		log.Error("could not connect to Minio/s3", "url", cfg.Minio.BaseUrl, "error", err, "worker-id", workerId)
	}
	log.Info("Checking if buckets exist", "worker-id", workerId)
	_, err = minioClient.BucketExists(ctx, cfg.Minio.UploadBucket)
	if err != nil {
		log.Info("Upload bucket does not exist, Creating a new one", "worker-id", workerId)
		err = minioClient.MakeBucket(ctx, cfg.Minio.UploadBucket, minio.MakeBucketOptions{})
		if err != nil {
			log.Error("Error in creating new upload bucket", "worker-id", workerId)
		} else {
			log.Info("Upload bucket successfully created", "worker-id", workerId)
		}
	} else {
		log.Debug("Upload bucket exists skipping creation", "worker-id", workerId)
	}
	_, err = minioClient.BucketExists(ctx, cfg.Minio.StorageBucket)
	if err != nil {
		log.Info("Storage bucket does not exist, Creating a new one", "worker-id", workerId)
		err = minioClient.MakeBucket(ctx, cfg.Minio.StorageBucket, minio.MakeBucketOptions{})
		if err != nil {
			log.Error("Error in creating new storage bucket", "worker-id", workerId)
		} else {
			log.Info("Storage bucket successfully created", "worker-id", workerId)
		}
	} else {
		log.Debug("Storage bucket exists skipping creation", "worker-id", workerId)
	}
	return minioClient, err
}
func Download(workerId int, ctx context.Context, minioClient *minio.Client, file *files.File, cfg *config.Config, log logger.Log) {
	log.Info("Downloading file", "file", file.Filename, "worker-id", workerId)
	err := minioClient.FGetObject(ctx, cfg.Minio.UploadBucket, file.Filename, "/tmp/"+file.Filename, minio.GetObjectOptions{})
	if err != nil {
		log.Error("Error occured while downloading file", "file", file.Filename, "worker-id", workerId)
	}
}

package storage

import (
	"context"
	"mime"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	appconfig "github.com/Benson-14/file-upload-service/internal/config"
)

type S3Client struct {
	client *s3.Client
	bucket string
}

func NewS3Client(cfg *appconfig.Config) (*S3Client, error) {
	awsCfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(cfg.S3.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.S3.AccessKeyID,
			cfg.S3.SecretAccessKey,
			"",
		)),
	)
	if err != nil {
		return nil, err
	}

	return &S3Client{
		client: s3.NewFromConfig(awsCfg),
		bucket: cfg.S3.Bucket,
	}, nil
}

// Upload streams the file directly to S3 without buffering it on disk.
func (s *S3Client) Upload(ctx context.Context, key string, body interface{ Read([]byte) (int, error) }, size int64) error {
	contentType := mime.TypeByExtension(filepath.Ext(key))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(s.bucket),
		Key:           aws.String(key),
		Body:          body,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(contentType),
	})
	return err
}

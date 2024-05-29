package helper

import (
	cfg "ConnetHub_job/pkg/config"
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Helper struct {
	cfg cfg.Config
}

func NewHelper(cfg cfg.Config) *Helper {
	return &Helper{cfg: cfg}
}

func (h *Helper) AddImageToAwsS3(file []byte) (string, error) {

	fileUID := uuid.New()
	fileName := fileUID.String()

	config, err := cfg.LoadConfig()
	if err != nil {
		return "", err
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			config.AWSAccesskeyID,
			config.AWSSecretaccesskey,
			"",
		),
	})
	if err != nil {
		return "", err
	}
	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.BucketName),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(file),
	})

	expo := "connectHub"
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", expo, fileName)
	return url, nil
}

func FromProtoTimestamp(ts *timestamppb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	return ts.AsTime()
}

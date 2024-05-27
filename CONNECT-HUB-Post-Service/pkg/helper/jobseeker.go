package helper

import (
	cfg "ConnetHub_post/pkg/config"
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
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

	// fmt.Println("pppppppp", config.DBHost)

	// fmt.Println("print1", config.AWSRegion)
	// fmt.Println("print2", config.AWSAccesskeyID)
	// fmt.Println("print3", config.AWSSecretaccesskey)
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
		Key:    aws.String("test1"),
		Body:   bytes.NewReader(file),
	})

	expo := "connectHub"
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", expo, fileName)
	return url, nil
}

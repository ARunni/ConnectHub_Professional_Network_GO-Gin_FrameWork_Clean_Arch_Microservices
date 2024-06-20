package helper

import (
	logging "github.com/ARunni/ConnetHub_post/Logging"
	cfg "github.com/ARunni/ConnetHub_post/pkg/config"
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Helper struct {
	cfg     cfg.Config
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewHelper(cfg cfg.Config) *Helper {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Post.log")
	return &Helper{
		cfg:     cfg,
		Logger:  logger,
		LogFile: logFile,
	}
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

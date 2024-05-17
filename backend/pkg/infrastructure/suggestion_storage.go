package infrastructure

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io/ioutil"
)

func (s *S3Storage) SaveFile(file []byte, filename string) (string, error) {
	if s.bucketName == "" {
		return "", errors.New("Bucket name is empty")
	}

	_, err := s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &s.bucketName,
		Key:    &filename,
		Body:   bytes.NewReader(file),
	})
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://%s.storage.yandexcloud.net/%s", s.bucketName, filename)
	return url, nil
}

func (s *S3Storage) GetFile(filename string) ([]byte, error) {
	resp, err := s.client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &s.bucketName,
		Key:    &filename,
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fileContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return fileContent, nil
}

func (s *S3Storage) ListFiles() ([]string, error) {
	result, err := s.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &s.bucketName,
	})
	if err != nil {
		return nil, err
	}

	var files []string
	for _, item := range result.Contents {
		files = append(files, *item.Key)
	}
	return files, nil
}

func (s *S3Storage) DeleteFile(filename string) error {
	_, err := s.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &s.bucketName,
		Key:    &filename,
	})
	if err != nil {
		return err
	}
	return nil
}

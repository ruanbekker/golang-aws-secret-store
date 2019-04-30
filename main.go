package main

import (
	"os"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func encryptSecret(secretValue string) (string, error) {

	var kmsRegion string
	var kmsKey string

	kmsRegion = os.Getenv("KMS_REGION")
	kmsKey = os.Getenv("KMS_KEY")

	s := session.Must(session.NewSession())
	svc := kms.New(s, &aws.Config{
		Region: aws.String(kmsRegion),
	})

	params := &kms.EncryptInput{
		KeyId:     aws.String(kmsKey),
		Plaintext: []byte(secretValue),
	}

	resp, err := svc.Encrypt(params)
	if err != nil {
		return "", err
	}

	cipher := string(resp.CiphertextBlob[:])
	return cipher, nil

}

func decryptSecret(encryptedCipher string) (string, error) {

	var kmsRegion string
	kmsRegion = os.Getenv("KMS_REGION")

	s := session.Must(session.NewSession())
	svc := kms.New(s, &aws.Config{
		Region: aws.String(kmsRegion),
	})

	blob := []byte(encryptedCipher)

	params := &kms.DecryptInput{
		CiphertextBlob: blob,
	}

	resp, err := svc.Decrypt(params)

	if err != nil {
		fmt.Println("Error decrypting:", err)
		os.Exit(1)
	}

	blobString := string(resp.Plaintext)
	return blobString, nil
}

func putSecret(region string, secretName string, secretValue string) string {

	var secretBucketName string

	secretBucketName = os.Getenv("S3_BUCKET")

	s := session.Must(session.NewSession())
	svc := s3.New(s, &aws.Config{
		Region: aws.String(region),
	})

	params := &s3.PutObjectInput{
		Bucket:               aws.String(secretBucketName),
		Key:                  aws.String(secretName),
		Body:                 bytes.NewReader([]byte(secretValue)),
		ServerSideEncryption: aws.String("AES256"),
		//ServerSideEncryption: aws.String("aws:kms"),
		//SSEKMSKeyId:          aws.String(kmsKey),
	}

	resp, err := svc.PutObject(params)

	if err != nil {
		fmt.Println(err.Error())
	}
	if resp == nil {
		fmt.Println(err.Error())
	}

	putResponse := "OK"
	return putResponse
}

func getSecret(region string, secretName string) (string, error) {

	var secretBucketName string

	secretBucketName = os.Getenv("S3_BUCKET")

	s := session.Must(session.NewSession())
	svc := s3.New(s, &aws.Config{
		Region: aws.String(region),
	})

	params := &s3.GetObjectInput{
		Bucket: aws.String(secretBucketName),
		Key:    aws.String(secretName),
	}

	resp, err := svc.GetObject(params)

	if err != nil {
		fmt.Println(err.Error())
	}

	var reader io.ReadCloser
	reader = resp.Body
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to Read Key")
		os.Exit(1)
	}

	return string(data[:]), nil
}

func main() {
	fmt.Println("Hello, World")
}

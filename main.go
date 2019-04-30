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

func main() {
	fmt.Println("Hello, World")
}

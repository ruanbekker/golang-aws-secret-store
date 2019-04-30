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

func main() {
	fmt.Println("Hello, World")
}

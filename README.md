# golang-aws-secret-store
Secrets Management on AWS with S3 and KMS using Golang

## About

Learning golang at the moment and wanted a easy and cheap way to store encrypted application secrets on S3.

## What does it do

The cli tool allows you to: 

- encrypt a string to a S3 path styled key eg. `secrets/production/db.domain.com/username` 
- gets encrypted with your KMS key
- stores the data on S3 with Server Side Encryption (SSE). 
- when the key is downloaded from S3 directly, the data will be in a encrypted form making it unusable.
- when using the tool's get method, it will decrypt using the kms key and return the secret to stdout.
- authentication: iam roles/users

As it can be run from a binary, it makes it easy to read application secrets.

## Examples

Build a binary:

```bash
$ go build -o secretstore main.go
```

Store your database username as a secret to `secrets/production/db.domain.com/username`

```bash
$ ./secretstore -put -secretName=secrets/production/db.domain.com/username -secretValue=rds_admin
```

Read the database username from the secret:

```bash
$ ./secretstore -get -secretName=secrets/production/db.domain.com/username
rds_admin
```

Read the S3 key directly using the cli:

```bash
$ aws --profile test s3 cp s3://<your_s3_bucket>/secrets/production/db.domain.com/username ./username
$ cat ./username
0x3W[encrypted_data]9231Pw2
...
```



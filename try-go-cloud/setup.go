package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/go-cloud/blob"
	"github.com/google/go-cloud/blob/gcsblob"
	"github.com/google/go-cloud/blob/s3blob"
	"github.com/google/go-cloud/gcp"
)

// setupBucket creates a connection to a particular cloud provider's blob storage.
func setupBucket(ctx context.Context, cloud, bucket string) (*blob.Bucket, error) {
	switch cloud {
	case "aws":
		return setupAWS(ctx, bucket)
	case "gcp":
		return setupGCP(ctx, bucket)
	default:
		return nil, fmt.Errorf("invalid cloud provider: %s", cloud)
	}
}

// setupGCP creates a connection to Google Cloud Storage (GCS).
func setupGCP(ctx context.Context, bucket string) (*blob.Bucket, error) {
	// DefaultCredentials assumes a user has logged in with gcloud.
	// See here for more information:
	// https://cloud.google.com/docs/authentication/getting-started
	creds, err := gcp.DefaultCredentials(ctx)
	if err != nil {
		return nil, err
	}
	c, err := gcp.NewHTTPClient(gcp.DefaultTransport(), gcp.CredentialsTokenSource(creds))
	if err != nil {
		return nil, err
	}
	return gcsblob.OpenBucket(ctx, bucket, c)
}

// setupAWS creates a connection to Simple Cloud Storage Service (S3).
func setupAWS(ctx context.Context, bucket string) (*blob.Bucket, error) {
	c := &aws.Config{
		// Either hard-code the region or use AWS_REGION.
		Region: aws.String("us-east-2"),
		// credentials.NewEnvCredentials assumes two environment variables are
		// present:
		// 1. AWS_ACCESS_KEY_ID, and
		// 2. AWS_SECRET_ACCESS_KEY.
		Credentials: credentials.NewEnvCredentials(),
	}
	s := session.Must(session.NewSession(c))
	return s3blob.OpenBucket(ctx, s, bucket)
}

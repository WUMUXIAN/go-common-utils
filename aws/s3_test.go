package aws

import (
	"fmt"
	"testing"
	"time"
)

var (
	bucketName string
)

func TestMain(m *testing.M) {
	fmt.Println("Create the bucket first.")
	bucketName = "wumuxian-test"
	err := GetS3Service("ap-southeast-1").CreateBucket("wumuxian-test", "ap-southeast-1")
	if err != nil {
		fmt.Println("Failed to create the bucket.")
		// os.Exit(1)
	}
	m.Run()
	fmt.Println("Delete the bucket.")
	err = GetS3Service("ap-southeast-1").DeleteBucket("wumuxian-test")
	if err != nil {
		fmt.Println("Failed to delete the bucket.")
		// os.Exit(1)
	}
}

func TestS3Upload(t *testing.T) {
	content := make([]byte, 0)
	for i := 0; i < 1*1024*1024; i++ {
		content = append(content, 'b')
	}

	tm := time.Now().UTC().Unix()
	err := GetS3Service("ap-southeast-1").UploadToS3Concurrently(content, bucketName, "/test", true)
	delta := time.Now().UTC().Unix() - tm
	fmt.Printf("Upload %d MB of data Concurrently takes %d seconds\n", 50, delta)

	tm = time.Now().UTC().Unix()
	err = GetS3Service("ap-southeast-1").UploadToS3(content, bucketName, "/test", true)
	delta = time.Now().UTC().Unix() - tm
	fmt.Printf("Upload %d MB of data Normally takes %d seconds\n", 50, delta)

	if err != nil {
		t.Fatal(err)
	}
}

func TestS3Read(t *testing.T) {
	tm := time.Now().UTC().Unix()
	_, err := GetS3Service("ap-southeast-1").ReadFromS3Concurrently(bucketName, "/test")
	delta := time.Now().UTC().Unix() - tm
	fmt.Printf("Download %d MB of data Concurrently takes %d seconds\n", 50, delta)

	if err != nil {
		t.Fatal(err)
	}

	tm = time.Now().UTC().Unix()
	_, err = GetS3Service("ap-southeast-1").ReadFromS3(bucketName, "/test")
	delta = time.Now().UTC().Unix() - tm
	fmt.Printf("Download %d MB of data Normally takes %d seconds\n", 50, delta)

	if err != nil {
		t.Fatal(err)
	}
}

func TestS3List(t *testing.T) {
	content := make([]byte, 0)
	for i := 0; i < 1*1024; i++ {
		content = append(content, 'b')
	}

	tm := time.Now().UTC().Unix()
	err := GetS3Service("ap-southeast-1").UploadToS3Concurrently(content, bucketName, "/test")
	delta := time.Now().UTC().Unix() - tm
	fmt.Printf("Upload %d MB of data Concurrently takes %d seconds\n", 1, delta)

	if err != nil {
		t.Fatal(err)
	}

	objects, err := GetS3Service("ap-southeast-1").ListAllS3(bucketName, "test")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Objects:", objects)
}

func TestS3Remove(t *testing.T) {
	content := make([]byte, 0)
	for i := 0; i < 1*1024; i++ {
		content = append(content, 'b')
	}

	tm := time.Now().UTC().Unix()
	err := GetS3Service("ap-southeast-1").UploadToS3Concurrently(content, bucketName, "/test")
	delta := time.Now().UTC().Unix() - tm
	fmt.Printf("Upload %d MB of data Concurrently takes %d seconds\n", 1, delta)

	if err != nil {
		t.Fatal(err)
	}

	err = GetS3Service("ap-southeast-1").RemoveFromS3(bucketName, "/test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestS3RemoveAll(t *testing.T) {
	var err error
	content := make([]byte, 0)
	for i := 0; i < 1024; i++ {
		content = append(content, 'b')
	}
	err = GetS3Service("ap-southeast-1").UploadToS3(content, bucketName, "/test", true)
	if err != nil {
		t.Fatal(err)
	}
	err = GetS3Service("ap-southeast-1").RemoveAllFromS3(bucketName, "test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestS3Copy(t *testing.T) {
	var err error
	content := make([]byte, 0)
	for i := 0; i < 1024; i++ {
		content = append(content, 'b')
	}
	err = GetS3Service("ap-southeast-1").UploadToS3(content, bucketName, "/test", true)
	if err != nil {
		t.Fatal(err)
	}

	err = GetS3Service("ap-southeast-1").CopyWithInS3(bucketName, "test", bucketName, "/test1", false)
	if err != nil {
		t.Fatal(err)
	}

	GetS3Service("ap-southeast-1").ListAllS3(bucketName, "test")

	err = GetS3Service("ap-southeast-1").CopyWithInS3(bucketName, "/test", bucketName, "test1", true)
	if err != nil {
		t.Fatal(err)
	}

	GetS3Service("ap-southeast-1").ListAllS3(bucketName, "/test")

	err = GetS3Service("ap-southeast-1").RemoveFromS3(bucketName, "/test1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestS3PreSignedURL(t *testing.T) {
	_, err := GetS3Service("ap-southeast-1").GetPreSignedURL(bucketName, "/test", 15*time.Minute)
	if err != nil {
		t.Fatal(err)
	}
}

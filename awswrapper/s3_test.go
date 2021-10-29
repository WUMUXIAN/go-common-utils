package awswrapper

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
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
		return
	}

	// Try to create again, this time it will fail.
	err = GetS3Service("ap-southeast-1").CreateBucket("wumuxian-test", "ap-southeast-1")
	if err == nil {
		fmt.Println("You Should Not Be Allowed to Create Bucket With the Same Name.")
		return
	}

	m.Run()

	fmt.Println("Delete the bucket.")
	err = GetS3Service("ap-southeast-1").DeleteBucket(bucketName)
	if err != nil {
		fmt.Println("Failed to delete the bucket.")
		return
	}

	// Delete again.
	err = GetS3Service("ap-southeast-1").DeleteBucket(bucketName)
	if err == nil {
		fmt.Println("You Should Not Be Allowed to Delete Bucket That Does Not Exist.")
		return
	}
}

func TestS3Upload(t *testing.T) {
	content := make([]byte, 0)
	for i := 0; i < 1*1024*1024; i++ {
		content = append(content, 'b')
	}

	tm := time.Now().UTC().Unix()
	err := GetS3Service("ap-southeast-1").UploadToS3Concurrently(content, bucketName, "/test", true)
	Convey("Upload To S3 Concurrently", t, func() {
		So(err, ShouldBeNil)
		delta := time.Now().UTC().Unix() - tm
		fmt.Printf("Upload %d MB of data Concurrently takes %d seconds\n", 50, delta)
	})

	tm = time.Now().UTC().Unix()
	err = GetS3Service("ap-southeast-1").UploadToS3(content, bucketName, "/test", true)
	Convey("Upload To S3 Normally", t, func() {
		So(err, ShouldBeNil)
		delta := time.Now().UTC().Unix() - tm
		fmt.Printf("Upload %d MB of data Normally takes %d seconds\n", 50, delta)
	})

	err = GetS3Service("ap-southeast-1").UploadToS3(content, "badBucketName", "/test", true)
	Convey("Upload To S3 Normally", t, func() {
		So(err, ShouldNotBeNil)
	})
}

func TestS3Read(t *testing.T) {
	tm := time.Now().UTC().Unix()
	_, err := GetS3Service("ap-southeast-1").ReadFromS3Concurrently(bucketName, "/test")
	Convey("Download From S3 Concurrently", t, func() {
		So(err, ShouldBeNil)
		delta := time.Now().UTC().Unix() - tm
		fmt.Printf("Download %d MB of data Concurrently takes %d seconds\n", 50, delta)
	})

	tm = time.Now().UTC().Unix()
	_, err = GetS3Service("ap-southeast-1").ReadFromS3(bucketName, "/test")
	Convey("Download From S3 Normally", t, func() {
		So(err, ShouldBeNil)
		delta := time.Now().UTC().Unix() - tm
		fmt.Printf("Download %d MB of data Normally takes %d seconds\n", 50, delta)
	})

	_, err = GetS3Service("ap-southeast-1").ReadFromS3("badBucketName", "/test")
	Convey("Download To S3 Normally", t, func() {
		So(err, ShouldNotBeNil)
	})

	headObject, err := GetS3Service("ap-southeast-1").ReadHeadObject(bucketName, "/test")
	Convey("Get Metadata For Existing Object", t, func() {
		So(err, ShouldBeNil)
		So(headObject, ShouldNotBeNil)
		So(headObject.ContentLength, ShouldNotBeNil)
		So(*headObject.ContentLength, ShouldEqual, 1*1024*1024)
	})

	existed, err := GetS3Service("ap-southeast-1").Exists(bucketName, "/test")
	Convey("Check Object Existence For Existing Object", t, func() {
		So(err, ShouldBeNil)
		So(existed, ShouldBeTrue)
	})

	existed, err = GetS3Service("ap-southeast-1").Exists(bucketName, "/test_asdfdsa")
	Convey("Check Object Existence For Non-Existing Object", t, func() {
		So(err, ShouldNotBeNil)
		So(existed, ShouldBeFalse)
	})
}

func TestS3List(t *testing.T) {
	objects, err := GetS3Service("ap-southeast-1").ListAllS3(bucketName, "test")
	Convey("List All Objects With Prefix", t, func() {
		So(err, ShouldBeNil)
		So(objects, ShouldResemble, []string{"test"})
	})

	objects, err = GetS3Service("us-east-1").ListAllS3("tectus-dreamlab-dev", "")
	Convey("List All Objects With Prefix", t, func() {
		So(err, ShouldBeNil)
		So(len(objects), ShouldBeGreaterThan, 1000)
	})
}

func TestS3PreSignedURL(t *testing.T) {
	_, err := GetS3Service("ap-southeast-1").GetPreSignedURL(bucketName, "/test", 15*time.Minute)
	Convey("S3 PreSigned URL", t, func() {
		So(err, ShouldBeNil)
	})
}

func TestS3Copy(t *testing.T) {
	err := GetS3Service("ap-southeast-1").CopyWithInS3(bucketName, "test", bucketName, "/test1", false)
	Convey("Copy Object Within S3 Bucket", t, func() {
		So(err, ShouldBeNil)
	})

	objects, err := GetS3Service("ap-southeast-1").ListAllS3(bucketName, "test")
	Convey("List All Objects With Prefix", t, func() {
		So(err, ShouldBeNil)
		So(objects, ShouldResemble, []string{"test", "test1"})
	})

	err = GetS3Service("ap-southeast-1").CopyWithInS3(bucketName, "/test1", bucketName, "test2", true)
	Convey("Copy Object Within S3 Bucket", t, func() {
		So(err, ShouldBeNil)
	})

	objects, err = GetS3Service("ap-southeast-1").ListAllS3(bucketName, "test")
	Convey("List All Objects With Prefix", t, func() {
		So(err, ShouldBeNil)
		So(objects, ShouldResemble, []string{"test", "test2"})
	})

	err = GetS3Service("ap-southeast-1").CopyWithInS3(bucketName, "test", bucketName, "/test1", false)
	Convey("Copy Object Within S3 Bucket", t, func() {
		So(err, ShouldBeNil)
	})

	objects, err = GetS3Service("ap-southeast-1").ListAllS3(bucketName, "test")
	Convey("List All Objects With Prefix", t, func() {
		So(err, ShouldBeNil)
		So(objects, ShouldResemble, []string{"test", "test1", "test2"})
	})

	err = GetS3Service("ap-southeast-1").CopyWithInS3(bucketName, "test3", bucketName, "/test1", false)
	Convey("Copy Object Within S3 Bucket", t, func() {
		So(err, ShouldNotBeNil)
	})
}

func TestS3Remove(t *testing.T) {
	err := GetS3Service("ap-southeast-1").RemoveFromS3(bucketName, "/test")
	Convey("Copy Object Within S3 Bucket", t, func() {
		So(err, ShouldBeNil)
	})
	objects, err := GetS3Service("ap-southeast-1").ListAllS3(bucketName, "test")
	Convey("List All Objects With Prefix", t, func() {
		So(err, ShouldBeNil)
		So(objects, ShouldResemble, []string{"test1", "test2"})
	})
	err = GetS3Service("ap-southeast-1").RemoveFromS3("non-exited-bucket", "/test3")
	Convey("Copy Object Within S3 Bucket", t, func() {
		So(err, ShouldNotBeNil)
	})
}

func TestS3RemoveAll(t *testing.T) {
	err := GetS3Service("ap-southeast-1").RemoveAllFromS3(bucketName, "test")
	Convey("Copy Object Within S3 Bucket", t, func() {
		So(err, ShouldBeNil)
	})
	objects, err := GetS3Service("ap-southeast-1").ListAllS3(bucketName, "test")
	Convey("List All Objects With Prefix", t, func() {
		So(err, ShouldBeNil)
		So(objects, ShouldBeEmpty)
	})
}

func TestConcurrentAccessService(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go func() {
			s3svc := GetS3Service("ap-southeast-1")
			if s3svc == nil {
				panic("")
			}
		}()
		go func() {
			s3svc := GetS3Service("us-east-1")
			if s3svc == nil {
				panic("")
			}
		}()
	}
}

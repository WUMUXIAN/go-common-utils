package aws

// ConcurrentUploader is used to upload multiple files to S3 concurrently.
type ConcurrentUploader struct {
	uploadStatus  map[string]int // status 0 = pending, 1 = completed, -1 = failed
	uploadContent map[string][]byte

	region     string
	bucketName string

	uploadChan     chan map[string]error
	UploadComplete chan bool
}

// NewConcurrentUploader creates a new uploader
func NewConcurrentUploader(bucketName, region string) (uploader *ConcurrentUploader) {
	uploader = new(ConcurrentUploader)

	uploader.uploadStatus = make(map[string]int)
	uploader.uploadContent = make(map[string][]byte)
	uploader.bucketName = bucketName
	uploader.region = region

	uploader.uploadChan = make(chan map[string]error)
	uploader.UploadComplete = make(chan bool)

	return
}

// Upload marks the content with the upload path, it will not execute the actual uploading.
func (uploader *ConcurrentUploader) Upload(content []byte, path string) {
	uploader.uploadContent[path] = content
	uploader.uploadStatus[path] = 0
}

// Execute executes the uploading process.
func (uploader *ConcurrentUploader) Execute() {
	for path, content := range uploader.uploadContent {
		go func(path string, content []byte) {
			err := GetS3Service(uploader.region).UploadToS3Concurrently(content, uploader.bucketName, path)
			uploader.uploadChan <- map[string]error{
				path: err,
			}
		}(path, content)
	}

	uploader.checkProgress()
}

func (uploader *ConcurrentUploader) checkProgress() {
	go func() {
		for {
			for path, err := range <-uploader.uploadChan {
				if err == nil {
					uploader.uploadStatus[path] = 1
				} else {
					uploader.uploadStatus[path] = -1
				}
			}

			completed := true
			err := false

			for _, status := range uploader.uploadStatus {
				if status == 0 {
					completed = false
					break
				} else if status == -1 {
					err = true
					break
				}
			}

			if completed {
				uploader.UploadComplete <- !err
				break
			}
		}
	}()
}

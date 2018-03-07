package aws

// ConcurrentDownloader is used to download multiple files to S3 concurrently.
type ConcurrentDownloader struct {
	DownloadContent map[string][]byte

	bucketName string
	region     string

	downloadChan     chan map[string][]byte
	DownloadComplete chan bool
}

// NewConcurrentDownloader creates a new downloader
func NewConcurrentDownloader(bucketName, region string) (downloader *ConcurrentDownloader) {
	downloader = new(ConcurrentDownloader)

	downloader.DownloadContent = make(map[string][]byte)
	downloader.bucketName = bucketName
	downloader.region = region

	downloader.downloadChan = make(chan map[string][]byte)
	downloader.DownloadComplete = make(chan bool)

	return
}

// Download marks to be downloaded path, it will not execute the actual downloading.
func (downloader *ConcurrentDownloader) Download(path string) {
	downloader.DownloadContent[path] = nil
}

// Execute executes the downloading process.
func (downloader *ConcurrentDownloader) Execute() {
	for path := range downloader.DownloadContent {
		go func(path string) {
			content, err := GetS3Service(downloader.region).ReadFromS3Concurrently(downloader.bucketName, path)
			if err == nil {
				downloader.downloadChan <- map[string][]byte{
					path: content,
				}
			} else {
				downloader.downloadChan <- map[string][]byte{
					path: []byte("error"),
				}
			}

		}(path)
	}

	downloader.checkProgress()
}

func (downloader *ConcurrentDownloader) checkProgress() {
	go func() {
		for {
			for path, content := range <-downloader.downloadChan {
				downloader.DownloadContent[path] = content
			}

			completed := true
			err := false

			for _, content := range downloader.DownloadContent {
				if content == nil {
					completed = false
					break
				} else if string(content) == "error" {
					err = true
					break
				}
			}

			if completed {
				downloader.DownloadComplete <- !err
				break
			}
		}
	}()
}

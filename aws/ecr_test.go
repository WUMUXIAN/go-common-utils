package aws

import "testing"

func TestListImages(t *testing.T) {
	images := GetECRService("us-east-1").ListImages("vinspection")
	for _, image := range images {
		t.Log(image)
	}
}

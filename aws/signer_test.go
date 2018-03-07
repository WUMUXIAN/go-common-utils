package aws

import (
	"crypto/x509"
	"fmt"
	"testing"

	"github.com/WUMUXIAN/go-common-utils/codec"
)

func TestUrlSigner(t *testing.T) {

	keyID := "APKAJ7HSYGBOXOVCZBQB"
	privateKeyString := `MIIEpAIBAAKCAQEAzHcEg5rjG3p8FzDzCbzmkMalOJJiBATufaBRZcf2kyzhyCdMWDubWbC+D53NCxw8yM0InC0vda4uCQ4y+MCjDArYJ7eX50dTXgCNDl6ueUpGsqOGZ66bHZUnE8nZ1qWf/usAJPL9wVXggYM09MxvXw9ozOiTgW/vfwaIImI7PoB/HOE7pO1LJbBGBEyH4C9vtRJsmb1IVieQYk2VOitu/WPGQ4MjAVU6VL4p/5bBGPT+ntop7trz3kwhW+QM+IFEAP44//v+U8b83DzzK/h1Sb/p5iIRqpSoUeIARaVjMvc2Tl897UGLDFx7eSD5gfnlYQ/HyWOC6EaJRonpFhpTLQIDAQABAoIBADHKHcBKfFlZp1QiaFyLsn240c0H4HRoGwdefdPSMNpACK8r2qx1v2vE3VOCMscs1iRzwU/+tNrsUmuEGKd4iXvKPD7Xt1briIKQkcjZB5Wjn7mqlnUzELTQeFaaRcC+TkrOQRe+UEq/Nc9z+vZNviUg5H1ZeWoArwp4tbfhwdmDE3vPBMrV/hl2aIFKbc7EaVAkv5xPXICFvvu++qElNMBw+D5r9wLre7RKRL+NFDGtGrbvp1efPZ6p8axHswzdEd9x//05gD7hn1gjm/JrhOB1M9HuZs8Er8bp5S5yVKOxIF6v7yFWCdLS8T4cd7fcMXLTjQTTwAneZ7IfVqiihWECgYEA6mM1gKFGy/j5Fwx+cEy95rIEPcli7cY2aCuVG0zmHAef/t8+rX4ghF94l9qizrAA1zH4SOezIRNEao6l6z/whf/JalRTVozmyTp7tWsvqb58/AjMPPXZRhsQ6WEt6VWP2dvinQxhWrTQ9g7Zs7PN3XWdIrCTdbdUn71b6K7fBeUCgYEA31F5zQv6ql22hrvUEuejxk/oSBwZeHLfHmedUAqg77xDCbs9K19UfoYPKKky1QbPOJrNV4kD1t56nHFOHpFgA3hxHWmoaD95lUcjryGRpBlqnELnbJORKq+s+rgCMrjhxRH+pP8kgV4oXrkgBdgecG1dYHupZQW5HiBNad+6w6kCgYEAu9DBl5AkLeAUoW6GhrBH32s4UNZl6ohRIooB0j19implv5LeI6GUpt3lwTEWEq8gDVBiVvErLc7FnOkvdOHod0eu+wAVQ55mdErjxEzehZM5ja/zEMojz7RyicAwTPAd9AHphdTc0hVf+DuQIRpsVRAg2SJLFyHPsqzG0B4IYi0CgYBcdMIFt4jvEd9oxsjVjtuKVjjn6eJNsNlZIDLMGapptrrWg5Oeqlg0DdKm1e46rhgK6mRLcmmJgxCmRm6+Txe+OBY0xDK5/lWbDRnj/vTqSK+PxE9F745xaswl/RrD3zFxwrJ3oz585Pu3w9NTOBfaGh1HvcrzDTyIEX2bcMpFCQKBgQDFVxhsCl7cELI1WFhrHwt2TulpclF4+T+fMlIusCLCwlBkoyXOondkuUFSLQPkwtDe6GuhejznfaM9V1xSFmbDYKsCm4gfV30N+OYTBpP7PDvGxtUAOsT8ZjE3cFJq/EL/XvxaJn3Juh8uLuNUblcMRu1MXQHeIOAgqLpESlV3NQ==`

	privateKeyBytes, _ := codec.Base64ToBytes(privateKeyString)
	privateKey, _ := x509.ParsePKCS1PrivateKey(privateKeyBytes)

	url, err := GetURLSigner(keyID, privateKey).SignURL(
		"d2cjabszqfh4q8.cloudfront.net",
		"/company-29e9035a-b47d-7e47-18c1-d7c50cf143e1/project-a1a8d8c6-6748-3031-bc77-dfeac13255ee/attachments/cd26c0dd-2fb7-b1d8-294a-e0fd8a792350",
		60, "testVideo.mov")

	if err == nil {
		fmt.Println("Url: ", url)
	} else {
		t.Fatal(err)
	}

	url, err = GetURLSigner(keyID, privateKey).SignURL(
		"dk70nib18pl2w.cloudfront.net",
		"/17000000/17-00000-0-V1_1_0_APP.lod",
		60)

	if err == nil {
		fmt.Println("Url: ", url)
	} else {
		t.Fatal(err)
	}
}

package util

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"

	"Demo/pkg/setting"
)

func UploadImageToTencentCloudCos(file *os.File, mode string) (img_url string, err error) {
	// Open the file to be uploaded
	// file, err := os.Open(fileName)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return "", err
	// }
	// defer file.Close()

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	fileSize := fileInfo.Size()

	// Read the file contents into a byte slice
	data := make([]byte, fileSize)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	// Rename the file with a timestamp
	timestamp := time.Now().Unix()
	tmp := ".jpg"
	newFileName := fmt.Sprintf("%d_%s%s", timestamp, mode, tmp)

	// Create a new COS client
	bucketURL, err := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", setting.BucketName, setting.Region))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	// Create a new COS client with the appropriate authorization
	client := cos.NewClient(
		&cos.BaseURL{BucketURL: bucketURL},
		&http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  setting.SecretID,
				SecretKey: setting.SecretKey,
			},
		},
	)

	// Upload the file to COS
	reader := bytes.NewReader(data)
	_, err = client.Object.Put(context.Background(), setting.ObjectKey+newFileName, reader, nil)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	// Get the URL of the uploaded file
	url := fmt.Sprintf("https://%s.cos.%s.myqcloud.com/%s", setting.BucketName, setting.Region, setting.ObjectKey+newFileName)
	return url, err
}

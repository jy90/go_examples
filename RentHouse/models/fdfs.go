package models

import (
	"fmt"
	"github.com/weilaihui/fdfs_client"
)

func FDFSUploadByFilename(fileName string) (groupName string, fileId string, err error) {
	fdfsClient, err := fdfs_client.NewFdfsClient("/etc/fdfs/client.conf")
	if err != nil {
		fmt.Printf("New FdfsClient error %s", err.Error())
		return "", "", err
	}

	uploadResponse, err := fdfsClient.UploadByFilename(fileName)
	if err != nil {
		fmt.Printf("UploadByfilename error %s", err.Error())
		return "", "", err
	}
	fmt.Println("groupName ==", uploadResponse.GroupName)
	fmt.Println("fileId ==", uploadResponse.RemoteFileId)
	return uploadResponse.GroupName, uploadResponse.RemoteFileId, nil
}

func FDFSUploadByBuffer(buffer []byte, suffix string) (groupName string, fileId string, err error) {
	fdfsClient, err := fdfs_client.NewFdfsClient("/etc/fdfs/client.conf")
	if err != nil {
		fmt.Printf("New FdfsClient error %s", err.Error())
		return "", "", err
	}

	uploadResponse, err := fdfsClient.UploadAppenderByBuffer(buffer, suffix)
	if err != nil {
		fmt.Println("uploadByFilename error ", err)
		return "", "", err
	}
	fmt.Println("groupName ==", uploadResponse.GroupName)
	fmt.Println("fileId ==", uploadResponse.RemoteFileId)
	return uploadResponse.GroupName, uploadResponse.RemoteFileId, nil
}

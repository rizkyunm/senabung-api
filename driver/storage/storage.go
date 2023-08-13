package storage

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"os"
)

var (
	accountName, accessKey, serviceURL, containerName string
	client                                            *azblob.Client
)

func newStorage() *azblob.Client {
	accountName = os.Getenv("STORAGE_ACCOUNT_NAME")
	accessKey = os.Getenv("STORAGE_ACCESS_KEY")
	serviceURL = fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)
	containerName = os.Getenv("STORAGE_CONTAINER_NAME")

	credential, err := azblob.NewSharedKeyCredential(accountName, accessKey)
	if err != nil {
		panic(err.Error())
	}

	client, err = azblob.NewClientWithSharedKeyCredential(serviceURL, credential, nil)
	if err != nil {
		panic(err.Error())
	}

	return client
}

func GetStorageClient() *azblob.Client {
	if client != nil {
		return client
	}

	return newStorage()
}

func GetContainerName() string {
	return containerName
}

func GetAccountName() string {
	return accountName
}

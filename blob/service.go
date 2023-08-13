package blob

import (
	"context"
	"github.com/rizkyunm/senabung-api/driver/storage"
	"io"
	"mime/multipart"
)

func UploadObject(file *multipart.FileHeader, path string, c context.Context) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	fileBytes, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	blobClient := storage.GetStorageClient()
	_, err = blobClient.UploadBuffer(c, storage.GetContainerName(), path, fileBytes, nil)
	if err != nil {
		return "", err
	}

	filePath := "https://" + storage.GetAccountName() + ".blob.core.windows.net/" + storage.GetContainerName() + "/" + path
	return filePath, nil
}

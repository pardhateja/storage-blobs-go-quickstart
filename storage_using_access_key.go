package main

import (
	"context"
	"fmt"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

func main2() {
	// Define your Azure Blob Storage connection string or use your credentials.
	accountName := "your-storage-account-name"
	accountKey := "your-storage-account-key" // or use environment variables
	containerName := "your-container-name"
	folderName := "your-folder-name"

	// Create a pipeline to handle the request/response.
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	pipeline := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	// Create a container URL.
	containerURL := azblob.NewContainerURL(
		azblob.NewServiceURL(containerName, pipeline),
		containerName,
	)

	// Create a virtual directory URL for the specified folder.
	folderURL := containerURL.NewBlobURL(folderName + "/")

	// List blobs in the folder.
	listBlob, err := folderURL.ListBlobHierarchySegment(
		context.TODO(),
		azblob.Marker{},
		azblob.ListBlobsSegmentOptions{
			Delimiter: "",
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, blobInfo := range listBlob.Segment.BlobItems {
		fmt.Printf("Blob Name: %s\n", blobInfo.Name)
	}
}

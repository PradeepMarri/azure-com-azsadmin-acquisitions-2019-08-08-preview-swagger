package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Acquisition represents the Acquisition schema from the OpenAPI specification
type Acquisition struct {
	Acquisitionid string `json:"acquisitionid,omitempty"` // The ID of page BLOB acquisition.
	Storageaccount string `json:"storageaccount,omitempty"` // The storage account that holds the page BLOB.
	Filepath string `json:"filePath,omitempty"` // The file path of the page BLOB file on storage cluster.
	Maximumblobsize int64 `json:"maximumblobsize,omitempty"` // The maximum size of the page BLOB.
	Susbcriptionid string `json:"susbcriptionid,omitempty"` // ID of the subscription associated with the page BLOB.
	Blob string `json:"blob,omitempty"` // The name of the page BLOB.
	Filepathunc string `json:"filePathUnc,omitempty"` // The file path unc of the page BLOB file on storage cluster.
	Status string `json:"status,omitempty"` // The status of page BLOB acquisition.
	Container string `json:"container,omitempty"` // The container associated with the page BLOB.
}

// AcquisitionList represents the AcquisitionList schema from the OpenAPI specification
type AcquisitionList struct {
	Value []Acquisition `json:"value,omitempty"` // List of acquisitions.
}

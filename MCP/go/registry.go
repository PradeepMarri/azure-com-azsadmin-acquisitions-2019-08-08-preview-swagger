package main

import (
	"github.com/storagemanagementclient/mcp-server/config"
	"github.com/storagemanagementclient/mcp-server/models"
	tools_acquisitions "github.com/storagemanagementclient/mcp-server/tools/acquisitions"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_acquisitions.CreateAcquisitions_listTool(cfg),
	}
}

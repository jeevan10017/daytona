// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package workspace

import (
	"errors"
)

type Workspace struct {
	Id       string     `json:"id"`
	Name     string     `json:"name"`
	Projects []*Project `json:"projects"`
	Target   string     `json:"target"`
	ApiKey   string     `json:"-"`
} // @name Workspace

type WorkspaceInfo struct {
	Name             string         `json:"name"`
	Projects         []*ProjectInfo `json:"projects"`
	ProviderMetadata string         `json:"providerMetadata,omitempty"`
} // @name WorkspaceInfo

func (w *Workspace) GetProject(projectName string) (*Project, error) {
	for _, project := range w.Projects {
		if project.Name == projectName {
			return project, nil
		}
	}
	return nil, errors.New("project not found")
}

type WorkspaceEnvVarParams struct {
	ApiUrl        string
	ApiKey        string
	ServerUrl     string
	ServerVersion string
}

func GetWorkspaceEnvVars(workspace *Workspace, params WorkspaceEnvVarParams) map[string]string {
	envVars := map[string]string{
		"DAYTONA_WS_ID":          workspace.Id,
		"DAYTONA_SERVER_API_KEY": params.ApiKey,
		"DAYTONA_SERVER_VERSION": params.ServerVersion,
		"DAYTONA_SERVER_URL":     params.ServerUrl,
		"DAYTONA_SERVER_API_URL": params.ApiUrl,
		// $HOME will be replaced at runtime
		"DAYTONA_AGENT_LOG_FILE_PATH": "$HOME/.daytona-agent.log",
	}

	return envVars
}

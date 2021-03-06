// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package workspacesiface provides an interface for the Amazon WorkSpaces.
package workspacesiface

import (
	"github.com/aws/aws-sdk-go/service/workspaces"
)

// WorkSpacesAPI is the interface type for workspaces.WorkSpaces.
type WorkSpacesAPI interface {
	CreateWorkspaces(*workspaces.CreateWorkspacesInput) (*workspaces.CreateWorkspacesOutput, error)

	DescribeWorkspaceBundles(*workspaces.DescribeWorkspaceBundlesInput) (*workspaces.DescribeWorkspaceBundlesOutput, error)

	DescribeWorkspaceDirectories(*workspaces.DescribeWorkspaceDirectoriesInput) (*workspaces.DescribeWorkspaceDirectoriesOutput, error)

	DescribeWorkspaces(*workspaces.DescribeWorkspacesInput) (*workspaces.DescribeWorkspacesOutput, error)

	RebootWorkspaces(*workspaces.RebootWorkspacesInput) (*workspaces.RebootWorkspacesOutput, error)

	RebuildWorkspaces(*workspaces.RebuildWorkspacesInput) (*workspaces.RebuildWorkspacesOutput, error)

	TerminateWorkspaces(*workspaces.TerminateWorkspacesInput) (*workspaces.TerminateWorkspacesOutput, error)
}

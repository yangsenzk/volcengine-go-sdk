// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package emriface provides an interface to enable mocking the EMR service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package emr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// EMRAPI provides an interface to enable mocking the
// emr.EMR service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // EMR.
//    func myFunc(svc EMRAPI) bool {
//        // Make svc.CheckUserCredentials request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := emr.New(sess)
//
//        myFunc(svc)
//    }
//
type EMRAPI interface {
	CheckUserCredentialsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CheckUserCredentialsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CheckUserCredentialsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CheckUserCredentials(*CheckUserCredentialsInput) (*CheckUserCredentialsOutput, error)
	CheckUserCredentialsWithContext(volcengine.Context, *CheckUserCredentialsInput, ...request.Option) (*CheckUserCredentialsOutput, error)
	CheckUserCredentialsRequest(*CheckUserCredentialsInput) (*request.Request, *CheckUserCredentialsOutput)

	CreateClusterCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateClusterCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateClusterCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateCluster(*CreateClusterInput) (*CreateClusterOutput, error)
	CreateClusterWithContext(volcengine.Context, *CreateClusterInput, ...request.Option) (*CreateClusterOutput, error)
	CreateClusterRequest(*CreateClusterInput) (*request.Request, *CreateClusterOutput)

	CreateClusterUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateClusterUserCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateClusterUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateClusterUser(*CreateClusterUserInput) (*CreateClusterUserOutput, error)
	CreateClusterUserWithContext(volcengine.Context, *CreateClusterUserInput, ...request.Option) (*CreateClusterUserOutput, error)
	CreateClusterUserRequest(*CreateClusterUserInput) (*request.Request, *CreateClusterUserOutput)

	CreateClusterUserGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateClusterUserGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateClusterUserGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateClusterUserGroup(*CreateClusterUserGroupInput) (*CreateClusterUserGroupOutput, error)
	CreateClusterUserGroupWithContext(volcengine.Context, *CreateClusterUserGroupInput, ...request.Option) (*CreateClusterUserGroupOutput, error)
	CreateClusterUserGroupRequest(*CreateClusterUserGroupInput) (*request.Request, *CreateClusterUserGroupOutput)

	CreateNodeGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateNodeGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateNodeGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateNodeGroup(*CreateNodeGroupInput) (*CreateNodeGroupOutput, error)
	CreateNodeGroupWithContext(volcengine.Context, *CreateNodeGroupInput, ...request.Option) (*CreateNodeGroupOutput, error)
	CreateNodeGroupRequest(*CreateNodeGroupInput) (*request.Request, *CreateNodeGroupOutput)

	DeleteClusterUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteClusterUserCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteClusterUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteClusterUser(*DeleteClusterUserInput) (*DeleteClusterUserOutput, error)
	DeleteClusterUserWithContext(volcengine.Context, *DeleteClusterUserInput, ...request.Option) (*DeleteClusterUserOutput, error)
	DeleteClusterUserRequest(*DeleteClusterUserInput) (*request.Request, *DeleteClusterUserOutput)

	DeleteClusterUserGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteClusterUserGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteClusterUserGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteClusterUserGroup(*DeleteClusterUserGroupInput) (*DeleteClusterUserGroupOutput, error)
	DeleteClusterUserGroupWithContext(volcengine.Context, *DeleteClusterUserGroupInput, ...request.Option) (*DeleteClusterUserGroupOutput, error)
	DeleteClusterUserGroupRequest(*DeleteClusterUserGroupInput) (*request.Request, *DeleteClusterUserGroupOutput)

	DeleteNodeGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteNodeGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteNodeGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteNodeGroup(*DeleteNodeGroupInput) (*DeleteNodeGroupOutput, error)
	DeleteNodeGroupWithContext(volcengine.Context, *DeleteNodeGroupInput, ...request.Option) (*DeleteNodeGroupOutput, error)
	DeleteNodeGroupRequest(*DeleteNodeGroupInput) (*request.Request, *DeleteNodeGroupOutput)

	GetApplicationConfigFileCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetApplicationConfigFileCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetApplicationConfigFileCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetApplicationConfigFile(*GetApplicationConfigFileInput) (*GetApplicationConfigFileOutput, error)
	GetApplicationConfigFileWithContext(volcengine.Context, *GetApplicationConfigFileInput, ...request.Option) (*GetApplicationConfigFileOutput, error)
	GetApplicationConfigFileRequest(*GetApplicationConfigFileInput) (*request.Request, *GetApplicationConfigFileOutput)

	GetClusterCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetClusterCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetClusterCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetCluster(*GetClusterInput) (*GetClusterOutput, error)
	GetClusterWithContext(volcengine.Context, *GetClusterInput, ...request.Option) (*GetClusterOutput, error)
	GetClusterRequest(*GetClusterInput) (*request.Request, *GetClusterOutput)

	ListApplicationConfigFilesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListApplicationConfigFilesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListApplicationConfigFilesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListApplicationConfigFiles(*ListApplicationConfigFilesInput) (*ListApplicationConfigFilesOutput, error)
	ListApplicationConfigFilesWithContext(volcengine.Context, *ListApplicationConfigFilesInput, ...request.Option) (*ListApplicationConfigFilesOutput, error)
	ListApplicationConfigFilesRequest(*ListApplicationConfigFilesInput) (*request.Request, *ListApplicationConfigFilesOutput)

	ListApplicationConfigHistoriesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListApplicationConfigHistoriesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListApplicationConfigHistoriesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListApplicationConfigHistories(*ListApplicationConfigHistoriesInput) (*ListApplicationConfigHistoriesOutput, error)
	ListApplicationConfigHistoriesWithContext(volcengine.Context, *ListApplicationConfigHistoriesInput, ...request.Option) (*ListApplicationConfigHistoriesOutput, error)
	ListApplicationConfigHistoriesRequest(*ListApplicationConfigHistoriesInput) (*request.Request, *ListApplicationConfigHistoriesOutput)

	ListApplicationConfigsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListApplicationConfigsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListApplicationConfigsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListApplicationConfigs(*ListApplicationConfigsInput) (*ListApplicationConfigsOutput, error)
	ListApplicationConfigsWithContext(volcengine.Context, *ListApplicationConfigsInput, ...request.Option) (*ListApplicationConfigsOutput, error)
	ListApplicationConfigsRequest(*ListApplicationConfigsInput) (*request.Request, *ListApplicationConfigsOutput)

	ListApplicationsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListApplicationsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListApplicationsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListApplications(*ListApplicationsInput) (*ListApplicationsOutput, error)
	ListApplicationsWithContext(volcengine.Context, *ListApplicationsInput, ...request.Option) (*ListApplicationsOutput, error)
	ListApplicationsRequest(*ListApplicationsInput) (*request.Request, *ListApplicationsOutput)

	ListClusterUserGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListClusterUserGroupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListClusterUserGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListClusterUserGroups(*ListClusterUserGroupsInput) (*ListClusterUserGroupsOutput, error)
	ListClusterUserGroupsWithContext(volcengine.Context, *ListClusterUserGroupsInput, ...request.Option) (*ListClusterUserGroupsOutput, error)
	ListClusterUserGroupsRequest(*ListClusterUserGroupsInput) (*request.Request, *ListClusterUserGroupsOutput)

	ListClusterUsersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListClusterUsersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListClusterUsersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListClusterUsers(*ListClusterUsersInput) (*ListClusterUsersOutput, error)
	ListClusterUsersWithContext(volcengine.Context, *ListClusterUsersInput, ...request.Option) (*ListClusterUsersOutput, error)
	ListClusterUsersRequest(*ListClusterUsersInput) (*request.Request, *ListClusterUsersOutput)

	ListClustersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListClustersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListClustersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListClusters(*ListClustersInput) (*ListClustersOutput, error)
	ListClustersWithContext(volcengine.Context, *ListClustersInput, ...request.Option) (*ListClustersOutput, error)
	ListClustersRequest(*ListClustersInput) (*request.Request, *ListClustersOutput)

	ListComponentInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListComponentInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListComponentInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListComponentInstances(*ListComponentInstancesInput) (*ListComponentInstancesOutput, error)
	ListComponentInstancesWithContext(volcengine.Context, *ListComponentInstancesInput, ...request.Option) (*ListComponentInstancesOutput, error)
	ListComponentInstancesRequest(*ListComponentInstancesInput) (*request.Request, *ListComponentInstancesOutput)

	ListComponentsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListComponentsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListComponentsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListComponents(*ListComponentsInput) (*ListComponentsOutput, error)
	ListComponentsWithContext(volcengine.Context, *ListComponentsInput, ...request.Option) (*ListComponentsOutput, error)
	ListComponentsRequest(*ListComponentsInput) (*request.Request, *ListComponentsOutput)

	ListNodeGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListNodeGroupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListNodeGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListNodeGroups(*ListNodeGroupsInput) (*ListNodeGroupsOutput, error)
	ListNodeGroupsWithContext(volcengine.Context, *ListNodeGroupsInput, ...request.Option) (*ListNodeGroupsOutput, error)
	ListNodeGroupsRequest(*ListNodeGroupsInput) (*request.Request, *ListNodeGroupsOutput)

	ListNodesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListNodesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListNodesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListNodes(*ListNodesInput) (*ListNodesOutput, error)
	ListNodesWithContext(volcengine.Context, *ListNodesInput, ...request.Option) (*ListNodesOutput, error)
	ListNodesRequest(*ListNodesInput) (*request.Request, *ListNodesOutput)

	ListOperationsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListOperationsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListOperationsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListOperations(*ListOperationsInput) (*ListOperationsOutput, error)
	ListOperationsWithContext(volcengine.Context, *ListOperationsInput, ...request.Option) (*ListOperationsOutput, error)
	ListOperationsRequest(*ListOperationsInput) (*request.Request, *ListOperationsOutput)

	ReleaseClusterCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ReleaseClusterCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ReleaseClusterCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ReleaseCluster(*ReleaseClusterInput) (*ReleaseClusterOutput, error)
	ReleaseClusterWithContext(volcengine.Context, *ReleaseClusterInput, ...request.Option) (*ReleaseClusterOutput, error)
	ReleaseClusterRequest(*ReleaseClusterInput) (*request.Request, *ReleaseClusterOutput)

	RunApplicationActionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RunApplicationActionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RunApplicationActionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RunApplicationAction(*RunApplicationActionInput) (*RunApplicationActionOutput, error)
	RunApplicationActionWithContext(volcengine.Context, *RunApplicationActionInput, ...request.Option) (*RunApplicationActionOutput, error)
	RunApplicationActionRequest(*RunApplicationActionInput) (*request.Request, *RunApplicationActionOutput)

	ScaleInNodeGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ScaleInNodeGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ScaleInNodeGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ScaleInNodeGroup(*ScaleInNodeGroupInput) (*ScaleInNodeGroupOutput, error)
	ScaleInNodeGroupWithContext(volcengine.Context, *ScaleInNodeGroupInput, ...request.Option) (*ScaleInNodeGroupOutput, error)
	ScaleInNodeGroupRequest(*ScaleInNodeGroupInput) (*request.Request, *ScaleInNodeGroupOutput)

	ScaleOutNodeGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ScaleOutNodeGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ScaleOutNodeGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ScaleOutNodeGroup(*ScaleOutNodeGroupInput) (*ScaleOutNodeGroupOutput, error)
	ScaleOutNodeGroupWithContext(volcengine.Context, *ScaleOutNodeGroupInput, ...request.Option) (*ScaleOutNodeGroupOutput, error)
	ScaleOutNodeGroupRequest(*ScaleOutNodeGroupInput) (*request.Request, *ScaleOutNodeGroupOutput)

	ScaleUpNodeGroupDiskCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ScaleUpNodeGroupDiskCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ScaleUpNodeGroupDiskCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ScaleUpNodeGroupDisk(*ScaleUpNodeGroupDiskInput) (*ScaleUpNodeGroupDiskOutput, error)
	ScaleUpNodeGroupDiskWithContext(volcengine.Context, *ScaleUpNodeGroupDiskInput, ...request.Option) (*ScaleUpNodeGroupDiskOutput, error)
	ScaleUpNodeGroupDiskRequest(*ScaleUpNodeGroupDiskInput) (*request.Request, *ScaleUpNodeGroupDiskOutput)

	UpdateApplicationConfigCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateApplicationConfigCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateApplicationConfigCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateApplicationConfig(*UpdateApplicationConfigInput) (*UpdateApplicationConfigOutput, error)
	UpdateApplicationConfigWithContext(volcengine.Context, *UpdateApplicationConfigInput, ...request.Option) (*UpdateApplicationConfigOutput, error)
	UpdateApplicationConfigRequest(*UpdateApplicationConfigInput) (*request.Request, *UpdateApplicationConfigOutput)

	UpdateClusterAttributeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateClusterAttributeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateClusterAttributeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateClusterAttribute(*UpdateClusterAttributeInput) (*UpdateClusterAttributeOutput, error)
	UpdateClusterAttributeWithContext(volcengine.Context, *UpdateClusterAttributeInput, ...request.Option) (*UpdateClusterAttributeOutput, error)
	UpdateClusterAttributeRequest(*UpdateClusterAttributeInput) (*request.Request, *UpdateClusterAttributeOutput)

	UpdateClusterUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateClusterUserCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateClusterUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateClusterUser(*UpdateClusterUserInput) (*UpdateClusterUserOutput, error)
	UpdateClusterUserWithContext(volcengine.Context, *UpdateClusterUserInput, ...request.Option) (*UpdateClusterUserOutput, error)
	UpdateClusterUserRequest(*UpdateClusterUserInput) (*request.Request, *UpdateClusterUserOutput)

	UpdateClusterUserGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateClusterUserGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateClusterUserGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateClusterUserGroup(*UpdateClusterUserGroupInput) (*UpdateClusterUserGroupOutput, error)
	UpdateClusterUserGroupWithContext(volcengine.Context, *UpdateClusterUserGroupInput, ...request.Option) (*UpdateClusterUserGroupOutput, error)
	UpdateClusterUserGroupRequest(*UpdateClusterUserGroupInput) (*request.Request, *UpdateClusterUserGroupOutput)

	UpdateClusterUserPasswordCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateClusterUserPasswordCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateClusterUserPasswordCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateClusterUserPassword(*UpdateClusterUserPasswordInput) (*UpdateClusterUserPasswordOutput, error)
	UpdateClusterUserPasswordWithContext(volcengine.Context, *UpdateClusterUserPasswordInput, ...request.Option) (*UpdateClusterUserPasswordOutput, error)
	UpdateClusterUserPasswordRequest(*UpdateClusterUserPasswordInput) (*request.Request, *UpdateClusterUserPasswordOutput)

	UpdateNodeGroupAttributeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateNodeGroupAttributeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateNodeGroupAttributeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateNodeGroupAttribute(*UpdateNodeGroupAttributeInput) (*UpdateNodeGroupAttributeOutput, error)
	UpdateNodeGroupAttributeWithContext(volcengine.Context, *UpdateNodeGroupAttributeInput, ...request.Option) (*UpdateNodeGroupAttributeOutput, error)
	UpdateNodeGroupAttributeRequest(*UpdateNodeGroupAttributeInput) (*request.Request, *UpdateNodeGroupAttributeOutput)

	UpdateNodeGroupChargeTypeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateNodeGroupChargeTypeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateNodeGroupChargeTypeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateNodeGroupChargeType(*UpdateNodeGroupChargeTypeInput) (*UpdateNodeGroupChargeTypeOutput, error)
	UpdateNodeGroupChargeTypeWithContext(volcengine.Context, *UpdateNodeGroupChargeTypeInput, ...request.Option) (*UpdateNodeGroupChargeTypeOutput, error)
	UpdateNodeGroupChargeTypeRequest(*UpdateNodeGroupChargeTypeInput) (*request.Request, *UpdateNodeGroupChargeTypeOutput)

	UpdateNodeGroupEcsSpecCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateNodeGroupEcsSpecCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateNodeGroupEcsSpecCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateNodeGroupEcsSpec(*UpdateNodeGroupEcsSpecInput) (*UpdateNodeGroupEcsSpecOutput, error)
	UpdateNodeGroupEcsSpecWithContext(volcengine.Context, *UpdateNodeGroupEcsSpecInput, ...request.Option) (*UpdateNodeGroupEcsSpecOutput, error)
	UpdateNodeGroupEcsSpecRequest(*UpdateNodeGroupEcsSpecInput) (*request.Request, *UpdateNodeGroupEcsSpecOutput)
}

var _ EMRAPI = (*EMR)(nil)

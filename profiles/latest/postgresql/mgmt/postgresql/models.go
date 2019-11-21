// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package postgresql

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CreateMode = original.CreateMode

const (
	CreateModeDefault                   CreateMode = original.CreateModeDefault
	CreateModeGeoRestore                CreateMode = original.CreateModeGeoRestore
	CreateModePointInTimeRestore        CreateMode = original.CreateModePointInTimeRestore
	CreateModeReplica                   CreateMode = original.CreateModeReplica
	CreateModeServerPropertiesForCreate CreateMode = original.CreateModeServerPropertiesForCreate
)

type GeoRedundantBackup = original.GeoRedundantBackup

const (
	Disabled GeoRedundantBackup = original.Disabled
	Enabled  GeoRedundantBackup = original.Enabled
)

type IdentityType = original.IdentityType

const (
	None           IdentityType = original.None
	SystemAssigned IdentityType = original.SystemAssigned
)

type OperationOrigin = original.OperationOrigin

const (
	NotSpecified OperationOrigin = original.NotSpecified
	System       OperationOrigin = original.System
	User         OperationOrigin = original.User
)

type ServerSecurityAlertPolicyState = original.ServerSecurityAlertPolicyState

const (
	ServerSecurityAlertPolicyStateDisabled ServerSecurityAlertPolicyState = original.ServerSecurityAlertPolicyStateDisabled
	ServerSecurityAlertPolicyStateEnabled  ServerSecurityAlertPolicyState = original.ServerSecurityAlertPolicyStateEnabled
)

type ServerState = original.ServerState

const (
	ServerStateDisabled ServerState = original.ServerStateDisabled
	ServerStateDropping ServerState = original.ServerStateDropping
	ServerStateReady    ServerState = original.ServerStateReady
)

type ServerVersion = original.ServerVersion

const (
	NineFullStopFive    ServerVersion = original.NineFullStopFive
	NineFullStopSix     ServerVersion = original.NineFullStopSix
	OneOne              ServerVersion = original.OneOne
	OneZero             ServerVersion = original.OneZero
	OneZeroFullStopTwo  ServerVersion = original.OneZeroFullStopTwo
	OneZeroFullStopZero ServerVersion = original.OneZeroFullStopZero
)

type SkuTier = original.SkuTier

const (
	Basic           SkuTier = original.Basic
	GeneralPurpose  SkuTier = original.GeneralPurpose
	MemoryOptimized SkuTier = original.MemoryOptimized
)

type SslEnforcementEnum = original.SslEnforcementEnum

const (
	SslEnforcementEnumDisabled SslEnforcementEnum = original.SslEnforcementEnumDisabled
	SslEnforcementEnumEnabled  SslEnforcementEnum = original.SslEnforcementEnumEnabled
)

type StorageAutogrow = original.StorageAutogrow

const (
	StorageAutogrowDisabled StorageAutogrow = original.StorageAutogrowDisabled
	StorageAutogrowEnabled  StorageAutogrow = original.StorageAutogrowEnabled
)

type VirtualNetworkRuleState = original.VirtualNetworkRuleState

const (
	Deleting     VirtualNetworkRuleState = original.Deleting
	Initializing VirtualNetworkRuleState = original.Initializing
	InProgress   VirtualNetworkRuleState = original.InProgress
	Ready        VirtualNetworkRuleState = original.Ready
	Unknown      VirtualNetworkRuleState = original.Unknown
)

type BaseClient = original.BaseClient
type BasicServerPropertiesForCreate = original.BasicServerPropertiesForCreate
type CheckNameAvailabilityClient = original.CheckNameAvailabilityClient
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type Configuration = original.Configuration
type ConfigurationListResult = original.ConfigurationListResult
type ConfigurationProperties = original.ConfigurationProperties
type ConfigurationsClient = original.ConfigurationsClient
type ConfigurationsCreateOrUpdateFuture = original.ConfigurationsCreateOrUpdateFuture
type Database = original.Database
type DatabaseListResult = original.DatabaseListResult
type DatabaseProperties = original.DatabaseProperties
type DatabasesClient = original.DatabasesClient
type DatabasesCreateOrUpdateFuture = original.DatabasesCreateOrUpdateFuture
type DatabasesDeleteFuture = original.DatabasesDeleteFuture
type FirewallRule = original.FirewallRule
type FirewallRuleListResult = original.FirewallRuleListResult
type FirewallRuleProperties = original.FirewallRuleProperties
type FirewallRulesClient = original.FirewallRulesClient
type FirewallRulesCreateOrUpdateFuture = original.FirewallRulesCreateOrUpdateFuture
type FirewallRulesDeleteFuture = original.FirewallRulesDeleteFuture
type LocationBasedPerformanceTierClient = original.LocationBasedPerformanceTierClient
type LogFile = original.LogFile
type LogFileListResult = original.LogFileListResult
type LogFileProperties = original.LogFileProperties
type LogFilesClient = original.LogFilesClient
type NameAvailability = original.NameAvailability
type NameAvailabilityRequest = original.NameAvailabilityRequest
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type PerformanceTierListResult = original.PerformanceTierListResult
type PerformanceTierProperties = original.PerformanceTierProperties
type PerformanceTierServiceLevelObjectives = original.PerformanceTierServiceLevelObjectives
type ProxyResource = original.ProxyResource
type ReplicasClient = original.ReplicasClient
type ResourceIdentity = original.ResourceIdentity
type SecurityAlertPolicyProperties = original.SecurityAlertPolicyProperties
type Server = original.Server
type ServerForCreate = original.ServerForCreate
type ServerListResult = original.ServerListResult
type ServerProperties = original.ServerProperties
type ServerPropertiesForCreate = original.ServerPropertiesForCreate
type ServerPropertiesForDefaultCreate = original.ServerPropertiesForDefaultCreate
type ServerPropertiesForGeoRestore = original.ServerPropertiesForGeoRestore
type ServerPropertiesForReplica = original.ServerPropertiesForReplica
type ServerPropertiesForRestore = original.ServerPropertiesForRestore
type ServerSecurityAlertPoliciesClient = original.ServerSecurityAlertPoliciesClient
type ServerSecurityAlertPoliciesCreateOrUpdateFuture = original.ServerSecurityAlertPoliciesCreateOrUpdateFuture
type ServerSecurityAlertPolicy = original.ServerSecurityAlertPolicy
type ServerUpdateParameters = original.ServerUpdateParameters
type ServerUpdateParametersProperties = original.ServerUpdateParametersProperties
type ServersClient = original.ServersClient
type ServersCreateFuture = original.ServersCreateFuture
type ServersDeleteFuture = original.ServersDeleteFuture
type ServersRestartFuture = original.ServersRestartFuture
type ServersUpdateFuture = original.ServersUpdateFuture
type Sku = original.Sku
type StorageProfile = original.StorageProfile
type TrackedResource = original.TrackedResource
type VirtualNetworkRule = original.VirtualNetworkRule
type VirtualNetworkRuleListResult = original.VirtualNetworkRuleListResult
type VirtualNetworkRuleListResultIterator = original.VirtualNetworkRuleListResultIterator
type VirtualNetworkRuleListResultPage = original.VirtualNetworkRuleListResultPage
type VirtualNetworkRuleProperties = original.VirtualNetworkRuleProperties
type VirtualNetworkRulesClient = original.VirtualNetworkRulesClient
type VirtualNetworkRulesCreateOrUpdateFuture = original.VirtualNetworkRulesCreateOrUpdateFuture
type VirtualNetworkRulesDeleteFuture = original.VirtualNetworkRulesDeleteFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewCheckNameAvailabilityClient(subscriptionID string) CheckNameAvailabilityClient {
	return original.NewCheckNameAvailabilityClient(subscriptionID)
}
func NewCheckNameAvailabilityClientWithBaseURI(baseURI string, subscriptionID string) CheckNameAvailabilityClient {
	return original.NewCheckNameAvailabilityClientWithBaseURI(baseURI, subscriptionID)
}
func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClient(subscriptionID)
}
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabasesClient(subscriptionID string) DatabasesClient {
	return original.NewDatabasesClient(subscriptionID)
}
func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string) DatabasesClient {
	return original.NewDatabasesClientWithBaseURI(baseURI, subscriptionID)
}
func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClient(subscriptionID)
}
func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationBasedPerformanceTierClient(subscriptionID string) LocationBasedPerformanceTierClient {
	return original.NewLocationBasedPerformanceTierClient(subscriptionID)
}
func NewLocationBasedPerformanceTierClientWithBaseURI(baseURI string, subscriptionID string) LocationBasedPerformanceTierClient {
	return original.NewLocationBasedPerformanceTierClientWithBaseURI(baseURI, subscriptionID)
}
func NewLogFilesClient(subscriptionID string) LogFilesClient {
	return original.NewLogFilesClient(subscriptionID)
}
func NewLogFilesClientWithBaseURI(baseURI string, subscriptionID string) LogFilesClient {
	return original.NewLogFilesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewReplicasClient(subscriptionID string) ReplicasClient {
	return original.NewReplicasClient(subscriptionID)
}
func NewReplicasClientWithBaseURI(baseURI string, subscriptionID string) ReplicasClient {
	return original.NewReplicasClientWithBaseURI(baseURI, subscriptionID)
}
func NewServerSecurityAlertPoliciesClient(subscriptionID string) ServerSecurityAlertPoliciesClient {
	return original.NewServerSecurityAlertPoliciesClient(subscriptionID)
}
func NewServerSecurityAlertPoliciesClientWithBaseURI(baseURI string, subscriptionID string) ServerSecurityAlertPoliciesClient {
	return original.NewServerSecurityAlertPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewServersClient(subscriptionID string) ServersClient {
	return original.NewServersClient(subscriptionID)
}
func NewServersClientWithBaseURI(baseURI string, subscriptionID string) ServersClient {
	return original.NewServersClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualNetworkRuleListResultIterator(page VirtualNetworkRuleListResultPage) VirtualNetworkRuleListResultIterator {
	return original.NewVirtualNetworkRuleListResultIterator(page)
}
func NewVirtualNetworkRuleListResultPage(getNextPage func(context.Context, VirtualNetworkRuleListResult) (VirtualNetworkRuleListResult, error)) VirtualNetworkRuleListResultPage {
	return original.NewVirtualNetworkRuleListResultPage(getNextPage)
}
func NewVirtualNetworkRulesClient(subscriptionID string) VirtualNetworkRulesClient {
	return original.NewVirtualNetworkRulesClient(subscriptionID)
}
func NewVirtualNetworkRulesClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkRulesClient {
	return original.NewVirtualNetworkRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreateModeValues() []CreateMode {
	return original.PossibleCreateModeValues()
}
func PossibleGeoRedundantBackupValues() []GeoRedundantBackup {
	return original.PossibleGeoRedundantBackupValues()
}
func PossibleIdentityTypeValues() []IdentityType {
	return original.PossibleIdentityTypeValues()
}
func PossibleOperationOriginValues() []OperationOrigin {
	return original.PossibleOperationOriginValues()
}
func PossibleServerSecurityAlertPolicyStateValues() []ServerSecurityAlertPolicyState {
	return original.PossibleServerSecurityAlertPolicyStateValues()
}
func PossibleServerStateValues() []ServerState {
	return original.PossibleServerStateValues()
}
func PossibleServerVersionValues() []ServerVersion {
	return original.PossibleServerVersionValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleSslEnforcementEnumValues() []SslEnforcementEnum {
	return original.PossibleSslEnforcementEnumValues()
}
func PossibleStorageAutogrowValues() []StorageAutogrow {
	return original.PossibleStorageAutogrowValues()
}
func PossibleVirtualNetworkRuleStateValues() []VirtualNetworkRuleState {
	return original.PossibleVirtualNetworkRuleStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}

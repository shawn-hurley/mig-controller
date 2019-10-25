package containerregistryapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01-preview/containerregistry"
)

// RegistriesClientAPI contains the set of methods on the RegistriesClient type.
type RegistriesClientAPI interface {
	CheckNameAvailability(ctx context.Context, registryNameCheckRequest containerregistry.RegistryNameCheckRequest) (result containerregistry.RegistryNameStatus, err error)
	Create(ctx context.Context, resourceGroupName string, registryName string, registry containerregistry.Registry) (result containerregistry.RegistriesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.RegistriesDeleteFuture, err error)
	GenerateCredentials(ctx context.Context, resourceGroupName string, registryName string, generateCredentialsParameters containerregistry.GenerateCredentialsParameters) (result containerregistry.RegistriesGenerateCredentialsFuture, err error)
	Get(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.Registry, err error)
	ImportImage(ctx context.Context, resourceGroupName string, registryName string, parameters containerregistry.ImportImageParameters) (result containerregistry.RegistriesImportImageFuture, err error)
	List(ctx context.Context) (result containerregistry.RegistryListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerregistry.RegistryListResultPage, err error)
	ListCredentials(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.RegistryListCredentialsResult, err error)
	ListPolicies(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.RegistryPolicies, err error)
	ListUsages(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.RegistryUsageListResult, err error)
	RegenerateCredential(ctx context.Context, resourceGroupName string, registryName string, regenerateCredentialParameters containerregistry.RegenerateCredentialParameters) (result containerregistry.RegistryListCredentialsResult, err error)
	Update(ctx context.Context, resourceGroupName string, registryName string, registryUpdateParameters containerregistry.RegistryUpdateParameters) (result containerregistry.RegistriesUpdateFuture, err error)
	UpdatePolicies(ctx context.Context, resourceGroupName string, registryName string, registryPoliciesUpdateParameters containerregistry.RegistryPolicies) (result containerregistry.RegistriesUpdatePoliciesFuture, err error)
}

var _ RegistriesClientAPI = (*containerregistry.RegistriesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result containerregistry.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*containerregistry.OperationsClient)(nil)

// ReplicationsClientAPI contains the set of methods on the ReplicationsClient type.
type ReplicationsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replication containerregistry.Replication) (result containerregistry.ReplicationsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, registryName string, replicationName string) (result containerregistry.ReplicationsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, registryName string, replicationName string) (result containerregistry.Replication, err error)
	List(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.ReplicationListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replicationUpdateParameters containerregistry.ReplicationUpdateParameters) (result containerregistry.ReplicationsUpdateFuture, err error)
}

var _ ReplicationsClientAPI = (*containerregistry.ReplicationsClient)(nil)

// WebhooksClientAPI contains the set of methods on the WebhooksClient type.
type WebhooksClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, registryName string, webhookName string, webhookCreateParameters containerregistry.WebhookCreateParameters) (result containerregistry.WebhooksCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result containerregistry.WebhooksDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result containerregistry.Webhook, err error)
	GetCallbackConfig(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result containerregistry.CallbackConfig, err error)
	List(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.WebhookListResultPage, err error)
	ListEvents(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result containerregistry.EventListResultPage, err error)
	Ping(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result containerregistry.EventInfo, err error)
	Update(ctx context.Context, resourceGroupName string, registryName string, webhookName string, webhookUpdateParameters containerregistry.WebhookUpdateParameters) (result containerregistry.WebhooksUpdateFuture, err error)
}

var _ WebhooksClientAPI = (*containerregistry.WebhooksClient)(nil)

// ScopeMapsClientAPI contains the set of methods on the ScopeMapsClient type.
type ScopeMapsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, registryName string, scopeMapName string, scopeMapCreateParameters containerregistry.ScopeMap) (result containerregistry.ScopeMapsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, registryName string, scopeMapName string) (result containerregistry.ScopeMapsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, registryName string, scopeMapName string) (result containerregistry.ScopeMap, err error)
	List(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.ScopeMapListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, registryName string, scopeMapName string, scopeMapUpdateParameters containerregistry.ScopeMapUpdateParameters) (result containerregistry.ScopeMapsUpdateFuture, err error)
}

var _ ScopeMapsClientAPI = (*containerregistry.ScopeMapsClient)(nil)

// TokensClientAPI contains the set of methods on the TokensClient type.
type TokensClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, registryName string, tokenName string, tokenCreateParameters containerregistry.Token) (result containerregistry.TokensCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, registryName string, tokenName string) (result containerregistry.TokensDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, registryName string, tokenName string) (result containerregistry.Token, err error)
	List(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.TokenListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, registryName string, tokenName string, tokenUpdateParameters containerregistry.TokenUpdateParameters) (result containerregistry.TokensUpdateFuture, err error)
}

var _ TokensClientAPI = (*containerregistry.TokensClient)(nil)
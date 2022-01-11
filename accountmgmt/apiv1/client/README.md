# Go API client for accountmgmtclient

Manage user subscriptions and clusters

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./accountmgmtclient"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.openshift.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AppServicesApi* | [**ApiAccountsMgmtV1AccessTokenPost**](docs/AppServicesApi.md#apiaccountsmgmtv1accesstokenpost) | **Post** /api/accounts_mgmt/v1/access_token | Return access token generated from registries in docker format
*AppServicesApi* | [**ApiAccountsMgmtV1CurrentAccountGet**](docs/AppServicesApi.md#apiaccountsmgmtv1currentaccountget) | **Get** /api/accounts_mgmt/v1/current_account | Get the authenticated account
*AppServicesApi* | [**ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet**](docs/AppServicesApi.md#apiaccountsmgmtv1organizationsorgidquotacostget) | **Get** /api/accounts_mgmt/v1/organizations/{orgId}/quota_cost | Returns a summary of quota cost
*AppServicesApi* | [**ApiAuthorizationsV1SelfTermsReviewPost**](docs/AppServicesApi.md#apiauthorizationsv1selftermsreviewpost) | **Post** /api/authorizations/v1/self_terms_review | Review your status of Terms
*DefaultApi* | [**ApiAccountsMgmtV1AccountsGet**](docs/DefaultApi.md#apiaccountsmgmtv1accountsget) | **Get** /api/accounts_mgmt/v1/accounts | Returns a list of accounts
*DefaultApi* | [**ApiAccountsMgmtV1AccountsIdGet**](docs/DefaultApi.md#apiaccountsmgmtv1accountsidget) | **Get** /api/accounts_mgmt/v1/accounts/{id} | Get an account by id
*DefaultApi* | [**ApiAccountsMgmtV1AccountsIdLabelsGet**](docs/DefaultApi.md#apiaccountsmgmtv1accountsidlabelsget) | **Get** /api/accounts_mgmt/v1/accounts/{id}/labels | Returns a list of labels
*DefaultApi* | [**ApiAccountsMgmtV1AccountsIdLabelsKeyDelete**](docs/DefaultApi.md#apiaccountsmgmtv1accountsidlabelskeydelete) | **Delete** /api/accounts_mgmt/v1/accounts/{id}/labels/{key} | Delete a label
*DefaultApi* | [**ApiAccountsMgmtV1AccountsIdLabelsKeyGet**](docs/DefaultApi.md#apiaccountsmgmtv1accountsidlabelskeyget) | **Get** /api/accounts_mgmt/v1/accounts/{id}/labels/{key} | Get subscription labels by label key
*DefaultApi* | [**ApiAccountsMgmtV1AccountsIdLabelsKeyPatch**](docs/DefaultApi.md#apiaccountsmgmtv1accountsidlabelskeypatch) | **Patch** /api/accounts_mgmt/v1/accounts/{id}/labels/{key} | Create a new label or update an existing label
*DefaultApi* | [**ApiAccountsMgmtV1AccountsIdLabelsPost**](docs/DefaultApi.md#apiaccountsmgmtv1accountsidlabelspost) | **Post** /api/accounts_mgmt/v1/accounts/{id}/labels | Create a new label or update an existing label
*DefaultApi* | [**ApiAccountsMgmtV1AccountsIdPatch**](docs/DefaultApi.md#apiaccountsmgmtv1accountsidpatch) | **Patch** /api/accounts_mgmt/v1/accounts/{id} | Update an account
*DefaultApi* | [**ApiAccountsMgmtV1AccountsPost**](docs/DefaultApi.md#apiaccountsmgmtv1accountspost) | **Post** /api/accounts_mgmt/v1/accounts | Create a new account
*DefaultApi* | [**ApiAccountsMgmtV1ClusterAuthorizationsPost**](docs/DefaultApi.md#apiaccountsmgmtv1clusterauthorizationspost) | **Post** /api/accounts_mgmt/v1/cluster_authorizations | Authorizes new cluster creation against an exsiting RH Subscription.
*DefaultApi* | [**ApiAccountsMgmtV1ClusterRegistrationsPost**](docs/DefaultApi.md#apiaccountsmgmtv1clusterregistrationspost) | **Post** /api/accounts_mgmt/v1/cluster_registrations | Finds or creates a cluster registration with a registy credential token and cluster ID
*DefaultApi* | [**ApiAccountsMgmtV1ErrorsGet**](docs/DefaultApi.md#apiaccountsmgmtv1errorsget) | **Get** /api/accounts_mgmt/v1/errors | Returns a list of errors
*DefaultApi* | [**ApiAccountsMgmtV1ErrorsIdGet**](docs/DefaultApi.md#apiaccountsmgmtv1errorsidget) | **Get** /api/accounts_mgmt/v1/errors/{id} | Get an error by id
*DefaultApi* | [**ApiAccountsMgmtV1FeatureTogglesIdQueryPost**](docs/DefaultApi.md#apiaccountsmgmtv1featuretogglesidquerypost) | **Post** /api/accounts_mgmt/v1/feature_toggles/{id}/query | Query a feature toggle by id
*DefaultApi* | [**ApiAccountsMgmtV1LabelsGet**](docs/DefaultApi.md#apiaccountsmgmtv1labelsget) | **Get** /api/accounts_mgmt/v1/labels | Returns a list of labels
*DefaultApi* | [**ApiAccountsMgmtV1MetricsGet**](docs/DefaultApi.md#apiaccountsmgmtv1metricsget) | **Get** /api/accounts_mgmt/v1/metrics | Returns a list of metrics
*DefaultApi* | [**ApiAccountsMgmtV1NotifyPost**](docs/DefaultApi.md#apiaccountsmgmtv1notifypost) | **Post** /api/accounts_mgmt/v1/notify | Notify the owner of cluster/subscription
*DefaultApi* | [**ApiAccountsMgmtV1OrganizationsGet**](docs/DefaultApi.md#apiaccountsmgmtv1organizationsget) | **Get** /api/accounts_mgmt/v1/organizations | Returns a list of organizations
*DefaultApi* | [**ApiAccountsMgmtV1OrganizationsIdGet**](docs/DefaultApi.md#apiaccountsmgmtv1organizationsidget) | **Get** /api/accounts_mgmt/v1/organizations/{id} | Get an organization by id
*DefaultApi* | [**ApiAccountsMgmtV1OrganizationsIdLabelsGet**](docs/DefaultApi.md#apiaccountsmgmtv1organizationsidlabelsget) | **Get** /api/accounts_mgmt/v1/organizations/{id}/labels | Returns a list of labels
*DefaultApi* | [**ApiAccountsMgmtV1OrganizationsIdLabelsKeyDelete**](docs/DefaultApi.md#apiaccountsmgmtv1organizationsidlabelskeydelete) | **Delete** /api/accounts_mgmt/v1/organizations/{id}/labels/{key} | Delete a label
*DefaultApi* | [**ApiAccountsMgmtV1OrganizationsIdLabelsKeyGet**](docs/DefaultApi.md#apiaccountsmgmtv1organizationsidlabelskeyget) | **Get** /api/accounts_mgmt/v1/organizations/{id}/labels/{key} | Get subscription labels by label key
*DefaultApi* | [**ApiAccountsMgmtV1OrganizationsIdLabelsKeyPatch**](docs/DefaultApi.md#apiaccountsmgmtv1organizationsidlabelskeypatch) | **Patch** /api/accounts_mgmt/v1/organizations/{id}/labels/{key} | Create a new label or update an existing label
*DefaultApi* | [**ApiAccountsMgmtV1OrganizationsIdLabelsPost**](docs/DefaultApi.md#apiaccountsmgmtv1organizationsidlabelspost) | **Post** /api/accounts_mgmt/v1/organizations/{id}/labels | Create a new label or update an existing label
*DefaultApi* | [**ApiAccountsMgmtV1OrganizationsIdPatch**](docs/DefaultApi.md#apiaccountsmgmtv1organizationsidpatch) | **Patch** /api/accounts_mgmt/v1/organizations/{id} | Update an organization
*DefaultApi* | [**ApiAccountsMgmtV1OrganizationsIdSummaryDashboardGet**](docs/DefaultApi.md#apiaccountsmgmtv1organizationsidsummarydashboardget) | **Get** /api/accounts_mgmt/v1/organizations/{id}/summary_dashboard | Returns a summary of organizations clusters based on metrics
*DefaultApi* | [**ApiAccountsMgmtV1OrganizationsOrgIdQuotaSummaryGet**](docs/DefaultApi.md#apiaccountsmgmtv1organizationsorgidquotasummaryget) | **Get** /api/accounts_mgmt/v1/organizations/{orgId}/quota_summary | Returns a summary of resource quota
*DefaultApi* | [**ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaGet**](docs/DefaultApi.md#apiaccountsmgmtv1organizationsorgidresourcequotaget) | **Get** /api/accounts_mgmt/v1/organizations/{orgId}/resource_quota | Returns a list of resource quota objects
*DefaultApi* | [**ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaPost**](docs/DefaultApi.md#apiaccountsmgmtv1organizationsorgidresourcequotapost) | **Post** /api/accounts_mgmt/v1/organizations/{orgId}/resource_quota | Create a new resource quota
*DefaultApi* | [**ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdDelete**](docs/DefaultApi.md#apiaccountsmgmtv1organizationsorgidresourcequotaquotaiddelete) | **Delete** /api/accounts_mgmt/v1/organizations/{orgId}/resource_quota/{quotaId} | Delete a resource quota
*DefaultApi* | [**ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdGet**](docs/DefaultApi.md#apiaccountsmgmtv1organizationsorgidresourcequotaquotaidget) | **Get** /api/accounts_mgmt/v1/organizations/{orgId}/resource_quota/{quotaId} | Get a resource quota by id
*DefaultApi* | [**ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdPatch**](docs/DefaultApi.md#apiaccountsmgmtv1organizationsorgidresourcequotaquotaidpatch) | **Patch** /api/accounts_mgmt/v1/organizations/{orgId}/resource_quota/{quotaId} | Update a resource quota
*DefaultApi* | [**ApiAccountsMgmtV1OrganizationsPost**](docs/DefaultApi.md#apiaccountsmgmtv1organizationspost) | **Post** /api/accounts_mgmt/v1/organizations | Create a new organization
*DefaultApi* | [**ApiAccountsMgmtV1PlansGet**](docs/DefaultApi.md#apiaccountsmgmtv1plansget) | **Get** /api/accounts_mgmt/v1/plans | Get all plans
*DefaultApi* | [**ApiAccountsMgmtV1PlansIdGet**](docs/DefaultApi.md#apiaccountsmgmtv1plansidget) | **Get** /api/accounts_mgmt/v1/plans/{id} | Get a plan by id
*DefaultApi* | [**ApiAccountsMgmtV1PullSecretsExternalResourceIdDelete**](docs/DefaultApi.md#apiaccountsmgmtv1pullsecretsexternalresourceiddelete) | **Delete** /api/accounts_mgmt/v1/pull_secrets/{externalResourceId} | Delete a pull secret
*DefaultApi* | [**ApiAccountsMgmtV1PullSecretsPost**](docs/DefaultApi.md#apiaccountsmgmtv1pullsecretspost) | **Post** /api/accounts_mgmt/v1/pull_secrets | Return access token generated from registries in docker format
*DefaultApi* | [**ApiAccountsMgmtV1RegistriesGet**](docs/DefaultApi.md#apiaccountsmgmtv1registriesget) | **Get** /api/accounts_mgmt/v1/registries | Returns a list of registries
*DefaultApi* | [**ApiAccountsMgmtV1RegistriesIdGet**](docs/DefaultApi.md#apiaccountsmgmtv1registriesidget) | **Get** /api/accounts_mgmt/v1/registries/{id} | Get an registry by id
*DefaultApi* | [**ApiAccountsMgmtV1RegistryCredentialsGet**](docs/DefaultApi.md#apiaccountsmgmtv1registrycredentialsget) | **Get** /api/accounts_mgmt/v1/registry_credentials | 
*DefaultApi* | [**ApiAccountsMgmtV1RegistryCredentialsIdDelete**](docs/DefaultApi.md#apiaccountsmgmtv1registrycredentialsiddelete) | **Delete** /api/accounts_mgmt/v1/registry_credentials/{id} | Delete a registry credential by id
*DefaultApi* | [**ApiAccountsMgmtV1RegistryCredentialsIdGet**](docs/DefaultApi.md#apiaccountsmgmtv1registrycredentialsidget) | **Get** /api/accounts_mgmt/v1/registry_credentials/{id} | Get a registry credentials by id
*DefaultApi* | [**ApiAccountsMgmtV1RegistryCredentialsIdPatch**](docs/DefaultApi.md#apiaccountsmgmtv1registrycredentialsidpatch) | **Patch** /api/accounts_mgmt/v1/registry_credentials/{id} | Update a registry credential
*DefaultApi* | [**ApiAccountsMgmtV1RegistryCredentialsPost**](docs/DefaultApi.md#apiaccountsmgmtv1registrycredentialspost) | **Post** /api/accounts_mgmt/v1/registry_credentials | Request the creation of a registry credential
*DefaultApi* | [**ApiAccountsMgmtV1ReservedResourcesGet**](docs/DefaultApi.md#apiaccountsmgmtv1reservedresourcesget) | **Get** /api/accounts_mgmt/v1/reserved_resources | Returns a list of reserved resources
*DefaultApi* | [**ApiAccountsMgmtV1ResourceQuotaGet**](docs/DefaultApi.md#apiaccountsmgmtv1resourcequotaget) | **Get** /api/accounts_mgmt/v1/resource_quota | Returns a list of resource quota objects
*DefaultApi* | [**ApiAccountsMgmtV1RoleBindingsGet**](docs/DefaultApi.md#apiaccountsmgmtv1rolebindingsget) | **Get** /api/accounts_mgmt/v1/role_bindings | Returns a list of role bindings
*DefaultApi* | [**ApiAccountsMgmtV1RoleBindingsIdDelete**](docs/DefaultApi.md#apiaccountsmgmtv1rolebindingsiddelete) | **Delete** /api/accounts_mgmt/v1/role_bindings/{id} | Delete a role binding
*DefaultApi* | [**ApiAccountsMgmtV1RoleBindingsIdGet**](docs/DefaultApi.md#apiaccountsmgmtv1rolebindingsidget) | **Get** /api/accounts_mgmt/v1/role_bindings/{id} | Get a role binding
*DefaultApi* | [**ApiAccountsMgmtV1RoleBindingsIdPatch**](docs/DefaultApi.md#apiaccountsmgmtv1rolebindingsidpatch) | **Patch** /api/accounts_mgmt/v1/role_bindings/{id} | Update a role binding
*DefaultApi* | [**ApiAccountsMgmtV1RoleBindingsPost**](docs/DefaultApi.md#apiaccountsmgmtv1rolebindingspost) | **Post** /api/accounts_mgmt/v1/role_bindings | Create a new role binding
*DefaultApi* | [**ApiAccountsMgmtV1RolesGet**](docs/DefaultApi.md#apiaccountsmgmtv1rolesget) | **Get** /api/accounts_mgmt/v1/roles | Returns a list of roles
*DefaultApi* | [**ApiAccountsMgmtV1RolesIdGet**](docs/DefaultApi.md#apiaccountsmgmtv1rolesidget) | **Get** /api/accounts_mgmt/v1/roles/{id} | Get a role by id
*DefaultApi* | [**ApiAccountsMgmtV1SkuRulesGet**](docs/DefaultApi.md#apiaccountsmgmtv1skurulesget) | **Get** /api/accounts_mgmt/v1/sku_rules | Returns a list of UHC product SKU Rules
*DefaultApi* | [**ApiAccountsMgmtV1SkuRulesIdGet**](docs/DefaultApi.md#apiaccountsmgmtv1skurulesidget) | **Get** /api/accounts_mgmt/v1/sku_rules/{id} | Get a sku rules by id
*DefaultApi* | [**ApiAccountsMgmtV1SkusGet**](docs/DefaultApi.md#apiaccountsmgmtv1skusget) | **Get** /api/accounts_mgmt/v1/skus | Returns a list of UHC product SKUs
*DefaultApi* | [**ApiAccountsMgmtV1SkusIdGet**](docs/DefaultApi.md#apiaccountsmgmtv1skusidget) | **Get** /api/accounts_mgmt/v1/skus/{id} | Get a sku by id
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsGet**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionsget) | **Get** /api/accounts_mgmt/v1/subscriptions | Returns a list of subscriptions
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsIdDelete**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionsiddelete) | **Delete** /api/accounts_mgmt/v1/subscriptions/{id} | Deletes a subscription by id
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsIdGet**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionsidget) | **Get** /api/accounts_mgmt/v1/subscriptions/{id} | Get a subscription by id
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsIdLabelsGet**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionsidlabelsget) | **Get** /api/accounts_mgmt/v1/subscriptions/{id}/labels | Returns a list of labels
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsIdLabelsKeyDelete**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionsidlabelskeydelete) | **Delete** /api/accounts_mgmt/v1/subscriptions/{id}/labels/{key} | Delete a label
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsIdLabelsKeyGet**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionsidlabelskeyget) | **Get** /api/accounts_mgmt/v1/subscriptions/{id}/labels/{key} | Get subscription labels by label key
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsIdLabelsKeyPatch**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionsidlabelskeypatch) | **Patch** /api/accounts_mgmt/v1/subscriptions/{id}/labels/{key} | Create a new label or update an existing label
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsIdLabelsPost**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionsidlabelspost) | **Post** /api/accounts_mgmt/v1/subscriptions/{id}/labels | Create a new label or update an existing label
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsIdNotifyPost**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionsidnotifypost) | **Post** /api/accounts_mgmt/v1/subscriptions/{id}/notify | Notify the owner of a subscription
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsIdPatch**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionsidpatch) | **Patch** /api/accounts_mgmt/v1/subscriptions/{id} | Update a subscription
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsIdReservedResourcesGet**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionsidreservedresourcesget) | **Get** /api/accounts_mgmt/v1/subscriptions/{id}/reserved_resources | Returns a list of reserved resources
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsIdSupportCasesGet**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionsidsupportcasesget) | **Get** /api/accounts_mgmt/v1/subscriptions/{id}/support_cases | Returns a list of open support creates opened against the external cluster id of this subscrption
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsPost**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionspost) | **Post** /api/accounts_mgmt/v1/subscriptions | Create a new subscription
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsAccountIdDelete**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionssubidnotificationcontactsaccountiddelete) | **Delete** /api/accounts_mgmt/v1/subscriptions/{subId}/notification_contacts/{accountId} | Deletes a notification contact by subscription and account id
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsGet**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionssubidnotificationcontactsget) | **Get** /api/accounts_mgmt/v1/subscriptions/{subId}/notification_contacts | Returns a list of notification contacts for the given subscription
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsPost**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionssubidnotificationcontactspost) | **Post** /api/accounts_mgmt/v1/subscriptions/{subId}/notification_contacts | Add an account as a notification contact to this subscription
*DefaultApi* | [**ApiAccountsMgmtV1SubscriptionsSubIdReservedResourcesReservedResourceIdGet**](docs/DefaultApi.md#apiaccountsmgmtv1subscriptionssubidreservedresourcesreservedresourceidget) | **Get** /api/accounts_mgmt/v1/subscriptions/{subId}/reserved_resources/{reservedResourceId} | Get reserved resources by id
*DefaultApi* | [**ApiAccountsMgmtV1SupportCasesCaseIdDelete**](docs/DefaultApi.md#apiaccountsmgmtv1supportcasescaseiddelete) | **Delete** /api/accounts_mgmt/v1/support_cases/{caseId} | Delete a support case
*DefaultApi* | [**ApiAccountsMgmtV1SupportCasesPost**](docs/DefaultApi.md#apiaccountsmgmtv1supportcasespost) | **Post** /api/accounts_mgmt/v1/support_cases | create a support case for the subscription
*DefaultApi* | [**ApiAccountsMgmtV1TokenAuthorizationPost**](docs/DefaultApi.md#apiaccountsmgmtv1tokenauthorizationpost) | **Post** /api/accounts_mgmt/v1/token_authorization | Finds the account owner of the provided token
*DefaultApi* | [**ApiAuthorizationsV1AccessReviewPost**](docs/DefaultApi.md#apiauthorizationsv1accessreviewpost) | **Post** /api/authorizations/v1/access_review | Review an account&#39;s access to perform an action on a particular resource or resource type
*DefaultApi* | [**ApiAuthorizationsV1CapabilityReviewPost**](docs/DefaultApi.md#apiauthorizationsv1capabilityreviewpost) | **Post** /api/authorizations/v1/capability_review | Review an account&#39;s capabilities
*DefaultApi* | [**ApiAuthorizationsV1ExportControlReviewPost**](docs/DefaultApi.md#apiauthorizationsv1exportcontrolreviewpost) | **Post** /api/authorizations/v1/export_control_review | Determine whether a user is restricted from downloading Red Hat software based on export control compliance. 
*DefaultApi* | [**ApiAuthorizationsV1FeatureReviewPost**](docs/DefaultApi.md#apiauthorizationsv1featurereviewpost) | **Post** /api/authorizations/v1/feature_review | Review feature to perform an action on it such as toggle a feature on/off
*DefaultApi* | [**ApiAuthorizationsV1ResourceReviewPost**](docs/DefaultApi.md#apiauthorizationsv1resourcereviewpost) | **Post** /api/authorizations/v1/resource_review | Obtain resource ids for resources an account may perform the specified action upon. Resource ids returned as [\&quot;*\&quot;] is shorthand for all ids.
*DefaultApi* | [**ApiAuthorizationsV1SelfAccessReviewPost**](docs/DefaultApi.md#apiauthorizationsv1selfaccessreviewpost) | **Post** /api/authorizations/v1/self_access_review | Review your ability to perform an action on a particular resource or resource type
*DefaultApi* | [**ApiAuthorizationsV1SelfFeatureReviewPost**](docs/DefaultApi.md#apiauthorizationsv1selffeaturereviewpost) | **Post** /api/authorizations/v1/self_feature_review | Review your ability to toggle a feature
*DefaultApi* | [**ApiAuthorizationsV1SelfResourceReviewPost**](docs/DefaultApi.md#apiauthorizationsv1selfresourcereviewpost) | **Post** /api/authorizations/v1/self_resource_review | Obtain resource ids for resources you may perform the specified action upon. Resource ids returned as [\&quot;*\&quot;] is shorthand for all ids.
*DefaultApi* | [**ApiAuthorizationsV1TermsReviewPost**](docs/DefaultApi.md#apiauthorizationsv1termsreviewpost) | **Post** /api/authorizations/v1/terms_review | Review an account&#39;s status of Terms


## Documentation For Models

 - [AccessReview](docs/AccessReview.md)
 - [AccessReviewResponse](docs/AccessReviewResponse.md)
 - [AccessTokenCfg](docs/AccessTokenCfg.md)
 - [Account](docs/Account.md)
 - [AccountAllOf](docs/AccountAllOf.md)
 - [AccountList](docs/AccountList.md)
 - [AccountListAllOf](docs/AccountListAllOf.md)
 - [AccountPatchRequest](docs/AccountPatchRequest.md)
 - [AccountReference](docs/AccountReference.md)
 - [AccountReferenceAllOf](docs/AccountReferenceAllOf.md)
 - [Capability](docs/Capability.md)
 - [CapabilityAllOf](docs/CapabilityAllOf.md)
 - [CapabilityReview](docs/CapabilityReview.md)
 - [CapabilityReviewRequest](docs/CapabilityReviewRequest.md)
 - [ClusterAuthorizationRequest](docs/ClusterAuthorizationRequest.md)
 - [ClusterAuthorizationResponse](docs/ClusterAuthorizationResponse.md)
 - [ClusterMetricsNodes](docs/ClusterMetricsNodes.md)
 - [ClusterRegistrationRequest](docs/ClusterRegistrationRequest.md)
 - [ClusterRegistrationResponse](docs/ClusterRegistrationResponse.md)
 - [ClusterResource](docs/ClusterResource.md)
 - [ClusterResourceTotal](docs/ClusterResourceTotal.md)
 - [ClusterUpgrade](docs/ClusterUpgrade.md)
 - [EphemeralResourceQuota](docs/EphemeralResourceQuota.md)
 - [Error](docs/Error.md)
 - [ErrorAllOf](docs/ErrorAllOf.md)
 - [ErrorList](docs/ErrorList.md)
 - [ErrorListAllOf](docs/ErrorListAllOf.md)
 - [ExcessResource](docs/ExcessResource.md)
 - [ExcessResourceAllOf](docs/ExcessResourceAllOf.md)
 - [ExportControlReview](docs/ExportControlReview.md)
 - [ExportControlReviewRequest](docs/ExportControlReviewRequest.md)
 - [FeatureReview](docs/FeatureReview.md)
 - [FeatureReviewResponse](docs/FeatureReviewResponse.md)
 - [FeatureToggle](docs/FeatureToggle.md)
 - [FeatureToggleAllOf](docs/FeatureToggleAllOf.md)
 - [FeatureToggleQueryRequest](docs/FeatureToggleQueryRequest.md)
 - [FeatureToggleQueryRequestAllOf](docs/FeatureToggleQueryRequestAllOf.md)
 - [Label](docs/Label.md)
 - [LabelAllOf](docs/LabelAllOf.md)
 - [LabelList](docs/LabelList.md)
 - [LabelListAllOf](docs/LabelListAllOf.md)
 - [List](docs/List.md)
 - [Metric](docs/Metric.md)
 - [MetricAllOf](docs/MetricAllOf.md)
 - [MetricsList](docs/MetricsList.md)
 - [MetricsListAllOf](docs/MetricsListAllOf.md)
 - [NotificationContactCreateRequest](docs/NotificationContactCreateRequest.md)
 - [NotificationRequest](docs/NotificationRequest.md)
 - [ObjectReference](docs/ObjectReference.md)
 - [OneMetric](docs/OneMetric.md)
 - [Organization](docs/Organization.md)
 - [OrganizationAllOf](docs/OrganizationAllOf.md)
 - [OrganizationList](docs/OrganizationList.md)
 - [OrganizationListAllOf](docs/OrganizationListAllOf.md)
 - [OrganizationPatchRequest](docs/OrganizationPatchRequest.md)
 - [Permission](docs/Permission.md)
 - [PermissionAllOf](docs/PermissionAllOf.md)
 - [PermissionList](docs/PermissionList.md)
 - [PermissionListAllOf](docs/PermissionListAllOf.md)
 - [Plan](docs/Plan.md)
 - [PlanAllOf](docs/PlanAllOf.md)
 - [PlanList](docs/PlanList.md)
 - [PlanListAllOf](docs/PlanListAllOf.md)
 - [PullSecretRequest](docs/PullSecretRequest.md)
 - [QuotaCost](docs/QuotaCost.md)
 - [QuotaCostAllOf](docs/QuotaCostAllOf.md)
 - [QuotaCostList](docs/QuotaCostList.md)
 - [QuotaCostListAllOf](docs/QuotaCostListAllOf.md)
 - [QuotaSummary](docs/QuotaSummary.md)
 - [QuotaSummaryAllOf](docs/QuotaSummaryAllOf.md)
 - [QuotaSummaryList](docs/QuotaSummaryList.md)
 - [QuotaSummaryListAllOf](docs/QuotaSummaryListAllOf.md)
 - [Registry](docs/Registry.md)
 - [RegistryAllOf](docs/RegistryAllOf.md)
 - [RegistryCreateRequest](docs/RegistryCreateRequest.md)
 - [RegistryCredential](docs/RegistryCredential.md)
 - [RegistryCredentialAllOf](docs/RegistryCredentialAllOf.md)
 - [RegistryCredentialList](docs/RegistryCredentialList.md)
 - [RegistryCredentialListAllOf](docs/RegistryCredentialListAllOf.md)
 - [RegistryCredentialPatchRequest](docs/RegistryCredentialPatchRequest.md)
 - [RegistryList](docs/RegistryList.md)
 - [RegistryListAllOf](docs/RegistryListAllOf.md)
 - [RegistryRequest](docs/RegistryRequest.md)
 - [RelatedResource](docs/RelatedResource.md)
 - [RelatedResourceAllOf](docs/RelatedResourceAllOf.md)
 - [ReservedResource](docs/ReservedResource.md)
 - [ReservedResourceAllOf](docs/ReservedResourceAllOf.md)
 - [ReservedResourceList](docs/ReservedResourceList.md)
 - [ReservedResourceListAllOf](docs/ReservedResourceListAllOf.md)
 - [ResourceQuota](docs/ResourceQuota.md)
 - [ResourceQuotaAllOf](docs/ResourceQuotaAllOf.md)
 - [ResourceQuotaList](docs/ResourceQuotaList.md)
 - [ResourceQuotaListAllOf](docs/ResourceQuotaListAllOf.md)
 - [ResourceQuotaRequest](docs/ResourceQuotaRequest.md)
 - [ResourceReview](docs/ResourceReview.md)
 - [ResourceReviewRequest](docs/ResourceReviewRequest.md)
 - [Role](docs/Role.md)
 - [RoleAllOf](docs/RoleAllOf.md)
 - [RoleBinding](docs/RoleBinding.md)
 - [RoleBindingAllOf](docs/RoleBindingAllOf.md)
 - [RoleBindingCreateRequest](docs/RoleBindingCreateRequest.md)
 - [RoleBindingList](docs/RoleBindingList.md)
 - [RoleBindingListAllOf](docs/RoleBindingListAllOf.md)
 - [RoleBindingRequest](docs/RoleBindingRequest.md)
 - [RoleList](docs/RoleList.md)
 - [RoleListAllOf](docs/RoleListAllOf.md)
 - [SKU](docs/SKU.md)
 - [SKUAllOf](docs/SKUAllOf.md)
 - [SelfAccessReview](docs/SelfAccessReview.md)
 - [SelfFeatureReview](docs/SelfFeatureReview.md)
 - [SelfResourceReview](docs/SelfResourceReview.md)
 - [SelfResourceReviewRequest](docs/SelfResourceReviewRequest.md)
 - [SelfTermsReview](docs/SelfTermsReview.md)
 - [SkuList](docs/SkuList.md)
 - [SkuListAllOf](docs/SkuListAllOf.md)
 - [SkuRules](docs/SkuRules.md)
 - [SkuRulesAllOf](docs/SkuRulesAllOf.md)
 - [SkuRulesList](docs/SkuRulesList.md)
 - [SkuRulesListAllOf](docs/SkuRulesListAllOf.md)
 - [Subscription](docs/Subscription.md)
 - [SubscriptionAllOf](docs/SubscriptionAllOf.md)
 - [SubscriptionCreateRequest](docs/SubscriptionCreateRequest.md)
 - [SubscriptionList](docs/SubscriptionList.md)
 - [SubscriptionListAllOf](docs/SubscriptionListAllOf.md)
 - [SubscriptionPatchRequest](docs/SubscriptionPatchRequest.md)
 - [Summary](docs/Summary.md)
 - [SummaryAllOf](docs/SummaryAllOf.md)
 - [SummaryMetrics](docs/SummaryMetrics.md)
 - [SummaryMetricsAllOf](docs/SummaryMetricsAllOf.md)
 - [SummaryVector](docs/SummaryVector.md)
 - [SummaryVectorAllOf](docs/SummaryVectorAllOf.md)
 - [SupportCasesCreatedResponse](docs/SupportCasesCreatedResponse.md)
 - [SupportCasesRequest](docs/SupportCasesRequest.md)
 - [TemplateParameter](docs/TemplateParameter.md)
 - [TermsReview](docs/TermsReview.md)
 - [TermsReviewResponse](docs/TermsReviewResponse.md)
 - [TokenAuthorizationRequest](docs/TokenAuthorizationRequest.md)
 - [TokenAuthorizationResponse](docs/TokenAuthorizationResponse.md)


## Documentation For Authorization



### Bearer

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARERTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



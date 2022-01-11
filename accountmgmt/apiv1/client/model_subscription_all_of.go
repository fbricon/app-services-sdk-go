/*
 * Account Management Service API
 *
 * Manage user subscriptions and clusters
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmtclient

import (
	"encoding/json"
	"time"
)

// SubscriptionAllOf struct for SubscriptionAllOf
type SubscriptionAllOf struct {
	// If set, the date the subscription expires based on the billing model
	BillingExpirationDate *time.Time `json:"billing_expiration_date,omitempty"`
	Capabilities *[]Capability `json:"capabilities,omitempty"`
	CloudAccountId *string `json:"cloud_account_id,omitempty"`
	CloudProviderId *string `json:"cloud_provider_id,omitempty"`
	ClusterBillingModel *string `json:"cluster_billing_model,omitempty"`
	ClusterId *string `json:"cluster_id,omitempty"`
	ConsoleUrl *string `json:"console_url,omitempty"`
	ConsumerUuid *string `json:"consumer_uuid,omitempty"`
	CpuTotal *int32 `json:"cpu_total,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Creator *AccountReference `json:"creator,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	ExternalClusterId *string `json:"external_cluster_id,omitempty"`
	Labels *[]Label `json:"labels,omitempty"`
	// Last time this subscription were reconciled about cluster usage
	LastReconcileDate *time.Time `json:"last_reconcile_date,omitempty"`
	// Last time status was set to Released for this cluster/subscription in Unix time
	LastReleasedAt *time.Time `json:"last_released_at,omitempty"`
	// Last telemetry authorization request for this cluster/subscription in Unix time
	LastTelemetryDate *time.Time `json:"last_telemetry_date,omitempty"`
	Managed bool `json:"managed"`
	Metrics *[]OneMetric `json:"metrics,omitempty"`
	NotificationContacts *[]Account `json:"notification_contacts,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	Plan *Plan `json:"plan,omitempty"`
	ProductBundle *string `json:"product_bundle,omitempty"`
	Provenance *string `json:"provenance,omitempty"`
	RegionId *string `json:"region_id,omitempty"`
	Released *bool `json:"released,omitempty"`
	ServiceLevel *string `json:"service_level,omitempty"`
	SocketTotal *int32 `json:"socket_total,omitempty"`
	Status *string `json:"status,omitempty"`
	SupportLevel *string `json:"support_level,omitempty"`
	SystemUnits *string `json:"system_units,omitempty"`
	// If the subscription is a trial, date the trial ends
	TrialEndDate *time.Time `json:"trial_end_date,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Usage *string `json:"usage,omitempty"`
}

// NewSubscriptionAllOf instantiates a new SubscriptionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionAllOf(managed bool) *SubscriptionAllOf {
	this := SubscriptionAllOf{}
	this.Managed = managed
	return &this
}

// NewSubscriptionAllOfWithDefaults instantiates a new SubscriptionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionAllOfWithDefaults() *SubscriptionAllOf {
	this := SubscriptionAllOf{}
	return &this
}

// GetBillingExpirationDate returns the BillingExpirationDate field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetBillingExpirationDate() time.Time {
	if o == nil || o.BillingExpirationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.BillingExpirationDate
}

// GetBillingExpirationDateOk returns a tuple with the BillingExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetBillingExpirationDateOk() (*time.Time, bool) {
	if o == nil || o.BillingExpirationDate == nil {
		return nil, false
	}
	return o.BillingExpirationDate, true
}

// HasBillingExpirationDate returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasBillingExpirationDate() bool {
	if o != nil && o.BillingExpirationDate != nil {
		return true
	}

	return false
}

// SetBillingExpirationDate gets a reference to the given time.Time and assigns it to the BillingExpirationDate field.
func (o *SubscriptionAllOf) SetBillingExpirationDate(v time.Time) {
	o.BillingExpirationDate = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetCapabilities() []Capability {
	if o == nil || o.Capabilities == nil {
		var ret []Capability
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetCapabilitiesOk() (*[]Capability, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasCapabilities() bool {
	if o != nil && o.Capabilities != nil {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []Capability and assigns it to the Capabilities field.
func (o *SubscriptionAllOf) SetCapabilities(v []Capability) {
	o.Capabilities = &v
}

// GetCloudAccountId returns the CloudAccountId field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetCloudAccountId() string {
	if o == nil || o.CloudAccountId == nil {
		var ret string
		return ret
	}
	return *o.CloudAccountId
}

// GetCloudAccountIdOk returns a tuple with the CloudAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetCloudAccountIdOk() (*string, bool) {
	if o == nil || o.CloudAccountId == nil {
		return nil, false
	}
	return o.CloudAccountId, true
}

// HasCloudAccountId returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasCloudAccountId() bool {
	if o != nil && o.CloudAccountId != nil {
		return true
	}

	return false
}

// SetCloudAccountId gets a reference to the given string and assigns it to the CloudAccountId field.
func (o *SubscriptionAllOf) SetCloudAccountId(v string) {
	o.CloudAccountId = &v
}

// GetCloudProviderId returns the CloudProviderId field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetCloudProviderId() string {
	if o == nil || o.CloudProviderId == nil {
		var ret string
		return ret
	}
	return *o.CloudProviderId
}

// GetCloudProviderIdOk returns a tuple with the CloudProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetCloudProviderIdOk() (*string, bool) {
	if o == nil || o.CloudProviderId == nil {
		return nil, false
	}
	return o.CloudProviderId, true
}

// HasCloudProviderId returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasCloudProviderId() bool {
	if o != nil && o.CloudProviderId != nil {
		return true
	}

	return false
}

// SetCloudProviderId gets a reference to the given string and assigns it to the CloudProviderId field.
func (o *SubscriptionAllOf) SetCloudProviderId(v string) {
	o.CloudProviderId = &v
}

// GetClusterBillingModel returns the ClusterBillingModel field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetClusterBillingModel() string {
	if o == nil || o.ClusterBillingModel == nil {
		var ret string
		return ret
	}
	return *o.ClusterBillingModel
}

// GetClusterBillingModelOk returns a tuple with the ClusterBillingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetClusterBillingModelOk() (*string, bool) {
	if o == nil || o.ClusterBillingModel == nil {
		return nil, false
	}
	return o.ClusterBillingModel, true
}

// HasClusterBillingModel returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasClusterBillingModel() bool {
	if o != nil && o.ClusterBillingModel != nil {
		return true
	}

	return false
}

// SetClusterBillingModel gets a reference to the given string and assigns it to the ClusterBillingModel field.
func (o *SubscriptionAllOf) SetClusterBillingModel(v string) {
	o.ClusterBillingModel = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *SubscriptionAllOf) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetConsoleUrl returns the ConsoleUrl field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetConsoleUrl() string {
	if o == nil || o.ConsoleUrl == nil {
		var ret string
		return ret
	}
	return *o.ConsoleUrl
}

// GetConsoleUrlOk returns a tuple with the ConsoleUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetConsoleUrlOk() (*string, bool) {
	if o == nil || o.ConsoleUrl == nil {
		return nil, false
	}
	return o.ConsoleUrl, true
}

// HasConsoleUrl returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasConsoleUrl() bool {
	if o != nil && o.ConsoleUrl != nil {
		return true
	}

	return false
}

// SetConsoleUrl gets a reference to the given string and assigns it to the ConsoleUrl field.
func (o *SubscriptionAllOf) SetConsoleUrl(v string) {
	o.ConsoleUrl = &v
}

// GetConsumerUuid returns the ConsumerUuid field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetConsumerUuid() string {
	if o == nil || o.ConsumerUuid == nil {
		var ret string
		return ret
	}
	return *o.ConsumerUuid
}

// GetConsumerUuidOk returns a tuple with the ConsumerUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetConsumerUuidOk() (*string, bool) {
	if o == nil || o.ConsumerUuid == nil {
		return nil, false
	}
	return o.ConsumerUuid, true
}

// HasConsumerUuid returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasConsumerUuid() bool {
	if o != nil && o.ConsumerUuid != nil {
		return true
	}

	return false
}

// SetConsumerUuid gets a reference to the given string and assigns it to the ConsumerUuid field.
func (o *SubscriptionAllOf) SetConsumerUuid(v string) {
	o.ConsumerUuid = &v
}

// GetCpuTotal returns the CpuTotal field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetCpuTotal() int32 {
	if o == nil || o.CpuTotal == nil {
		var ret int32
		return ret
	}
	return *o.CpuTotal
}

// GetCpuTotalOk returns a tuple with the CpuTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetCpuTotalOk() (*int32, bool) {
	if o == nil || o.CpuTotal == nil {
		return nil, false
	}
	return o.CpuTotal, true
}

// HasCpuTotal returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasCpuTotal() bool {
	if o != nil && o.CpuTotal != nil {
		return true
	}

	return false
}

// SetCpuTotal gets a reference to the given int32 and assigns it to the CpuTotal field.
func (o *SubscriptionAllOf) SetCpuTotal(v int32) {
	o.CpuTotal = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SubscriptionAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetCreator() AccountReference {
	if o == nil || o.Creator == nil {
		var ret AccountReference
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetCreatorOk() (*AccountReference, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasCreator() bool {
	if o != nil && o.Creator != nil {
		return true
	}

	return false
}

// SetCreator gets a reference to the given AccountReference and assigns it to the Creator field.
func (o *SubscriptionAllOf) SetCreator(v AccountReference) {
	o.Creator = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SubscriptionAllOf) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExternalClusterId returns the ExternalClusterId field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetExternalClusterId() string {
	if o == nil || o.ExternalClusterId == nil {
		var ret string
		return ret
	}
	return *o.ExternalClusterId
}

// GetExternalClusterIdOk returns a tuple with the ExternalClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetExternalClusterIdOk() (*string, bool) {
	if o == nil || o.ExternalClusterId == nil {
		return nil, false
	}
	return o.ExternalClusterId, true
}

// HasExternalClusterId returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasExternalClusterId() bool {
	if o != nil && o.ExternalClusterId != nil {
		return true
	}

	return false
}

// SetExternalClusterId gets a reference to the given string and assigns it to the ExternalClusterId field.
func (o *SubscriptionAllOf) SetExternalClusterId(v string) {
	o.ExternalClusterId = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetLabels() []Label {
	if o == nil || o.Labels == nil {
		var ret []Label
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetLabelsOk() (*[]Label, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []Label and assigns it to the Labels field.
func (o *SubscriptionAllOf) SetLabels(v []Label) {
	o.Labels = &v
}

// GetLastReconcileDate returns the LastReconcileDate field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetLastReconcileDate() time.Time {
	if o == nil || o.LastReconcileDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastReconcileDate
}

// GetLastReconcileDateOk returns a tuple with the LastReconcileDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetLastReconcileDateOk() (*time.Time, bool) {
	if o == nil || o.LastReconcileDate == nil {
		return nil, false
	}
	return o.LastReconcileDate, true
}

// HasLastReconcileDate returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasLastReconcileDate() bool {
	if o != nil && o.LastReconcileDate != nil {
		return true
	}

	return false
}

// SetLastReconcileDate gets a reference to the given time.Time and assigns it to the LastReconcileDate field.
func (o *SubscriptionAllOf) SetLastReconcileDate(v time.Time) {
	o.LastReconcileDate = &v
}

// GetLastReleasedAt returns the LastReleasedAt field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetLastReleasedAt() time.Time {
	if o == nil || o.LastReleasedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastReleasedAt
}

// GetLastReleasedAtOk returns a tuple with the LastReleasedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetLastReleasedAtOk() (*time.Time, bool) {
	if o == nil || o.LastReleasedAt == nil {
		return nil, false
	}
	return o.LastReleasedAt, true
}

// HasLastReleasedAt returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasLastReleasedAt() bool {
	if o != nil && o.LastReleasedAt != nil {
		return true
	}

	return false
}

// SetLastReleasedAt gets a reference to the given time.Time and assigns it to the LastReleasedAt field.
func (o *SubscriptionAllOf) SetLastReleasedAt(v time.Time) {
	o.LastReleasedAt = &v
}

// GetLastTelemetryDate returns the LastTelemetryDate field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetLastTelemetryDate() time.Time {
	if o == nil || o.LastTelemetryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastTelemetryDate
}

// GetLastTelemetryDateOk returns a tuple with the LastTelemetryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetLastTelemetryDateOk() (*time.Time, bool) {
	if o == nil || o.LastTelemetryDate == nil {
		return nil, false
	}
	return o.LastTelemetryDate, true
}

// HasLastTelemetryDate returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasLastTelemetryDate() bool {
	if o != nil && o.LastTelemetryDate != nil {
		return true
	}

	return false
}

// SetLastTelemetryDate gets a reference to the given time.Time and assigns it to the LastTelemetryDate field.
func (o *SubscriptionAllOf) SetLastTelemetryDate(v time.Time) {
	o.LastTelemetryDate = &v
}

// GetManaged returns the Managed field value
func (o *SubscriptionAllOf) GetManaged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Managed
}

// GetManagedOk returns a tuple with the Managed field value
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetManagedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Managed, true
}

// SetManaged sets field value
func (o *SubscriptionAllOf) SetManaged(v bool) {
	o.Managed = v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetMetrics() []OneMetric {
	if o == nil || o.Metrics == nil {
		var ret []OneMetric
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetMetricsOk() (*[]OneMetric, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given []OneMetric and assigns it to the Metrics field.
func (o *SubscriptionAllOf) SetMetrics(v []OneMetric) {
	o.Metrics = &v
}

// GetNotificationContacts returns the NotificationContacts field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetNotificationContacts() []Account {
	if o == nil || o.NotificationContacts == nil {
		var ret []Account
		return ret
	}
	return *o.NotificationContacts
}

// GetNotificationContactsOk returns a tuple with the NotificationContacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetNotificationContactsOk() (*[]Account, bool) {
	if o == nil || o.NotificationContacts == nil {
		return nil, false
	}
	return o.NotificationContacts, true
}

// HasNotificationContacts returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasNotificationContacts() bool {
	if o != nil && o.NotificationContacts != nil {
		return true
	}

	return false
}

// SetNotificationContacts gets a reference to the given []Account and assigns it to the NotificationContacts field.
func (o *SubscriptionAllOf) SetNotificationContacts(v []Account) {
	o.NotificationContacts = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *SubscriptionAllOf) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetPlan() Plan {
	if o == nil || o.Plan == nil {
		var ret Plan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetPlanOk() (*Plan, bool) {
	if o == nil || o.Plan == nil {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasPlan() bool {
	if o != nil && o.Plan != nil {
		return true
	}

	return false
}

// SetPlan gets a reference to the given Plan and assigns it to the Plan field.
func (o *SubscriptionAllOf) SetPlan(v Plan) {
	o.Plan = &v
}

// GetProductBundle returns the ProductBundle field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetProductBundle() string {
	if o == nil || o.ProductBundle == nil {
		var ret string
		return ret
	}
	return *o.ProductBundle
}

// GetProductBundleOk returns a tuple with the ProductBundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetProductBundleOk() (*string, bool) {
	if o == nil || o.ProductBundle == nil {
		return nil, false
	}
	return o.ProductBundle, true
}

// HasProductBundle returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasProductBundle() bool {
	if o != nil && o.ProductBundle != nil {
		return true
	}

	return false
}

// SetProductBundle gets a reference to the given string and assigns it to the ProductBundle field.
func (o *SubscriptionAllOf) SetProductBundle(v string) {
	o.ProductBundle = &v
}

// GetProvenance returns the Provenance field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetProvenance() string {
	if o == nil || o.Provenance == nil {
		var ret string
		return ret
	}
	return *o.Provenance
}

// GetProvenanceOk returns a tuple with the Provenance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetProvenanceOk() (*string, bool) {
	if o == nil || o.Provenance == nil {
		return nil, false
	}
	return o.Provenance, true
}

// HasProvenance returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasProvenance() bool {
	if o != nil && o.Provenance != nil {
		return true
	}

	return false
}

// SetProvenance gets a reference to the given string and assigns it to the Provenance field.
func (o *SubscriptionAllOf) SetProvenance(v string) {
	o.Provenance = &v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetRegionId() string {
	if o == nil || o.RegionId == nil {
		var ret string
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetRegionIdOk() (*string, bool) {
	if o == nil || o.RegionId == nil {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasRegionId() bool {
	if o != nil && o.RegionId != nil {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given string and assigns it to the RegionId field.
func (o *SubscriptionAllOf) SetRegionId(v string) {
	o.RegionId = &v
}

// GetReleased returns the Released field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetReleased() bool {
	if o == nil || o.Released == nil {
		var ret bool
		return ret
	}
	return *o.Released
}

// GetReleasedOk returns a tuple with the Released field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetReleasedOk() (*bool, bool) {
	if o == nil || o.Released == nil {
		return nil, false
	}
	return o.Released, true
}

// HasReleased returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasReleased() bool {
	if o != nil && o.Released != nil {
		return true
	}

	return false
}

// SetReleased gets a reference to the given bool and assigns it to the Released field.
func (o *SubscriptionAllOf) SetReleased(v bool) {
	o.Released = &v
}

// GetServiceLevel returns the ServiceLevel field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetServiceLevel() string {
	if o == nil || o.ServiceLevel == nil {
		var ret string
		return ret
	}
	return *o.ServiceLevel
}

// GetServiceLevelOk returns a tuple with the ServiceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetServiceLevelOk() (*string, bool) {
	if o == nil || o.ServiceLevel == nil {
		return nil, false
	}
	return o.ServiceLevel, true
}

// HasServiceLevel returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasServiceLevel() bool {
	if o != nil && o.ServiceLevel != nil {
		return true
	}

	return false
}

// SetServiceLevel gets a reference to the given string and assigns it to the ServiceLevel field.
func (o *SubscriptionAllOf) SetServiceLevel(v string) {
	o.ServiceLevel = &v
}

// GetSocketTotal returns the SocketTotal field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetSocketTotal() int32 {
	if o == nil || o.SocketTotal == nil {
		var ret int32
		return ret
	}
	return *o.SocketTotal
}

// GetSocketTotalOk returns a tuple with the SocketTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetSocketTotalOk() (*int32, bool) {
	if o == nil || o.SocketTotal == nil {
		return nil, false
	}
	return o.SocketTotal, true
}

// HasSocketTotal returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasSocketTotal() bool {
	if o != nil && o.SocketTotal != nil {
		return true
	}

	return false
}

// SetSocketTotal gets a reference to the given int32 and assigns it to the SocketTotal field.
func (o *SubscriptionAllOf) SetSocketTotal(v int32) {
	o.SocketTotal = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SubscriptionAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetSupportLevel returns the SupportLevel field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetSupportLevel() string {
	if o == nil || o.SupportLevel == nil {
		var ret string
		return ret
	}
	return *o.SupportLevel
}

// GetSupportLevelOk returns a tuple with the SupportLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetSupportLevelOk() (*string, bool) {
	if o == nil || o.SupportLevel == nil {
		return nil, false
	}
	return o.SupportLevel, true
}

// HasSupportLevel returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasSupportLevel() bool {
	if o != nil && o.SupportLevel != nil {
		return true
	}

	return false
}

// SetSupportLevel gets a reference to the given string and assigns it to the SupportLevel field.
func (o *SubscriptionAllOf) SetSupportLevel(v string) {
	o.SupportLevel = &v
}

// GetSystemUnits returns the SystemUnits field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetSystemUnits() string {
	if o == nil || o.SystemUnits == nil {
		var ret string
		return ret
	}
	return *o.SystemUnits
}

// GetSystemUnitsOk returns a tuple with the SystemUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetSystemUnitsOk() (*string, bool) {
	if o == nil || o.SystemUnits == nil {
		return nil, false
	}
	return o.SystemUnits, true
}

// HasSystemUnits returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasSystemUnits() bool {
	if o != nil && o.SystemUnits != nil {
		return true
	}

	return false
}

// SetSystemUnits gets a reference to the given string and assigns it to the SystemUnits field.
func (o *SubscriptionAllOf) SetSystemUnits(v string) {
	o.SystemUnits = &v
}

// GetTrialEndDate returns the TrialEndDate field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetTrialEndDate() time.Time {
	if o == nil || o.TrialEndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.TrialEndDate
}

// GetTrialEndDateOk returns a tuple with the TrialEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetTrialEndDateOk() (*time.Time, bool) {
	if o == nil || o.TrialEndDate == nil {
		return nil, false
	}
	return o.TrialEndDate, true
}

// HasTrialEndDate returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasTrialEndDate() bool {
	if o != nil && o.TrialEndDate != nil {
		return true
	}

	return false
}

// SetTrialEndDate gets a reference to the given time.Time and assigns it to the TrialEndDate field.
func (o *SubscriptionAllOf) SetTrialEndDate(v time.Time) {
	o.TrialEndDate = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SubscriptionAllOf) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetUsage() string {
	if o == nil || o.Usage == nil {
		var ret string
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetUsageOk() (*string, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given string and assigns it to the Usage field.
func (o *SubscriptionAllOf) SetUsage(v string) {
	o.Usage = &v
}

func (o SubscriptionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BillingExpirationDate != nil {
		toSerialize["billing_expiration_date"] = o.BillingExpirationDate
	}
	if o.Capabilities != nil {
		toSerialize["capabilities"] = o.Capabilities
	}
	if o.CloudAccountId != nil {
		toSerialize["cloud_account_id"] = o.CloudAccountId
	}
	if o.CloudProviderId != nil {
		toSerialize["cloud_provider_id"] = o.CloudProviderId
	}
	if o.ClusterBillingModel != nil {
		toSerialize["cluster_billing_model"] = o.ClusterBillingModel
	}
	if o.ClusterId != nil {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if o.ConsoleUrl != nil {
		toSerialize["console_url"] = o.ConsoleUrl
	}
	if o.ConsumerUuid != nil {
		toSerialize["consumer_uuid"] = o.ConsumerUuid
	}
	if o.CpuTotal != nil {
		toSerialize["cpu_total"] = o.CpuTotal
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.ExternalClusterId != nil {
		toSerialize["external_cluster_id"] = o.ExternalClusterId
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.LastReconcileDate != nil {
		toSerialize["last_reconcile_date"] = o.LastReconcileDate
	}
	if o.LastReleasedAt != nil {
		toSerialize["last_released_at"] = o.LastReleasedAt
	}
	if o.LastTelemetryDate != nil {
		toSerialize["last_telemetry_date"] = o.LastTelemetryDate
	}
	if true {
		toSerialize["managed"] = o.Managed
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.NotificationContacts != nil {
		toSerialize["notification_contacts"] = o.NotificationContacts
	}
	if o.OrganizationId != nil {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if o.Plan != nil {
		toSerialize["plan"] = o.Plan
	}
	if o.ProductBundle != nil {
		toSerialize["product_bundle"] = o.ProductBundle
	}
	if o.Provenance != nil {
		toSerialize["provenance"] = o.Provenance
	}
	if o.RegionId != nil {
		toSerialize["region_id"] = o.RegionId
	}
	if o.Released != nil {
		toSerialize["released"] = o.Released
	}
	if o.ServiceLevel != nil {
		toSerialize["service_level"] = o.ServiceLevel
	}
	if o.SocketTotal != nil {
		toSerialize["socket_total"] = o.SocketTotal
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SupportLevel != nil {
		toSerialize["support_level"] = o.SupportLevel
	}
	if o.SystemUnits != nil {
		toSerialize["system_units"] = o.SystemUnits
	}
	if o.TrialEndDate != nil {
		toSerialize["trial_end_date"] = o.TrialEndDate
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionAllOf struct {
	value *SubscriptionAllOf
	isSet bool
}

func (v NullableSubscriptionAllOf) Get() *SubscriptionAllOf {
	return v.value
}

func (v *NullableSubscriptionAllOf) Set(val *SubscriptionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionAllOf(val *SubscriptionAllOf) *NullableSubscriptionAllOf {
	return &NullableSubscriptionAllOf{value: val, isSet: true}
}

func (v NullableSubscriptionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


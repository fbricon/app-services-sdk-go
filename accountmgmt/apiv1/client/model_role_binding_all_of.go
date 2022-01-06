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

// RoleBindingAllOf struct for RoleBindingAllOf
type RoleBindingAllOf struct {
	Account *ObjectReference `json:"account,omitempty"`
	ConfigManaged *bool `json:"config_managed,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Organization *ObjectReference `json:"organization,omitempty"`
	Role *ObjectReference `json:"role,omitempty"`
	Subscription *ObjectReference `json:"subscription,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewRoleBindingAllOf instantiates a new RoleBindingAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleBindingAllOf() *RoleBindingAllOf {
	this := RoleBindingAllOf{}
	return &this
}

// NewRoleBindingAllOfWithDefaults instantiates a new RoleBindingAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleBindingAllOfWithDefaults() *RoleBindingAllOf {
	this := RoleBindingAllOf{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *RoleBindingAllOf) GetAccount() ObjectReference {
	if o == nil || o.Account == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBindingAllOf) GetAccountOk() (*ObjectReference, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *RoleBindingAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ObjectReference and assigns it to the Account field.
func (o *RoleBindingAllOf) SetAccount(v ObjectReference) {
	o.Account = &v
}

// GetConfigManaged returns the ConfigManaged field value if set, zero value otherwise.
func (o *RoleBindingAllOf) GetConfigManaged() bool {
	if o == nil || o.ConfigManaged == nil {
		var ret bool
		return ret
	}
	return *o.ConfigManaged
}

// GetConfigManagedOk returns a tuple with the ConfigManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBindingAllOf) GetConfigManagedOk() (*bool, bool) {
	if o == nil || o.ConfigManaged == nil {
		return nil, false
	}
	return o.ConfigManaged, true
}

// HasConfigManaged returns a boolean if a field has been set.
func (o *RoleBindingAllOf) HasConfigManaged() bool {
	if o != nil && o.ConfigManaged != nil {
		return true
	}

	return false
}

// SetConfigManaged gets a reference to the given bool and assigns it to the ConfigManaged field.
func (o *RoleBindingAllOf) SetConfigManaged(v bool) {
	o.ConfigManaged = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RoleBindingAllOf) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBindingAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RoleBindingAllOf) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RoleBindingAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *RoleBindingAllOf) GetOrganization() ObjectReference {
	if o == nil || o.Organization == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBindingAllOf) GetOrganizationOk() (*ObjectReference, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *RoleBindingAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given ObjectReference and assigns it to the Organization field.
func (o *RoleBindingAllOf) SetOrganization(v ObjectReference) {
	o.Organization = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *RoleBindingAllOf) GetRole() ObjectReference {
	if o == nil || o.Role == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBindingAllOf) GetRoleOk() (*ObjectReference, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *RoleBindingAllOf) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given ObjectReference and assigns it to the Role field.
func (o *RoleBindingAllOf) SetRole(v ObjectReference) {
	o.Role = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *RoleBindingAllOf) GetSubscription() ObjectReference {
	if o == nil || o.Subscription == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBindingAllOf) GetSubscriptionOk() (*ObjectReference, bool) {
	if o == nil || o.Subscription == nil {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *RoleBindingAllOf) HasSubscription() bool {
	if o != nil && o.Subscription != nil {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given ObjectReference and assigns it to the Subscription field.
func (o *RoleBindingAllOf) SetSubscription(v ObjectReference) {
	o.Subscription = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RoleBindingAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBindingAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoleBindingAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RoleBindingAllOf) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RoleBindingAllOf) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBindingAllOf) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RoleBindingAllOf) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RoleBindingAllOf) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o RoleBindingAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.ConfigManaged != nil {
		toSerialize["config_managed"] = o.ConfigManaged
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Subscription != nil {
		toSerialize["subscription"] = o.Subscription
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableRoleBindingAllOf struct {
	value *RoleBindingAllOf
	isSet bool
}

func (v NullableRoleBindingAllOf) Get() *RoleBindingAllOf {
	return v.value
}

func (v *NullableRoleBindingAllOf) Set(val *RoleBindingAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleBindingAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleBindingAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleBindingAllOf(val *RoleBindingAllOf) *NullableRoleBindingAllOf {
	return &NullableRoleBindingAllOf{value: val, isSet: true}
}

func (v NullableRoleBindingAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleBindingAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



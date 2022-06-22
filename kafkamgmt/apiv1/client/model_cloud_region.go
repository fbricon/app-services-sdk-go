/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.10.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// CloudRegion Description of a region of a cloud provider.
type CloudRegion struct {
	// Indicates the type of this object. Will be 'CloudRegion'.
	Kind *string `json:"kind,omitempty"`
	// Unique identifier of the object.
	Id *string `json:"id,omitempty"`
	// Name of the region for display purposes, for example `N. Virginia`.
	DisplayName *string `json:"display_name,omitempty"`
	// Whether the region is enabled for deploying an OSD cluster.
	Enabled bool `json:"enabled"`
	// Indicates whether there is capacity left per instance type
	Capacity []RegionCapacityListItem `json:"capacity"`
}

// NewCloudRegion instantiates a new CloudRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRegion(enabled bool, capacity []RegionCapacityListItem) *CloudRegion {
	this := CloudRegion{}
	this.Enabled = enabled
	this.Capacity = capacity
	return &this
}

// NewCloudRegionWithDefaults instantiates a new CloudRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRegionWithDefaults() *CloudRegion {
	this := CloudRegion{}
	var enabled bool = false
	this.Enabled = enabled
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CloudRegion) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegion) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CloudRegion) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *CloudRegion) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CloudRegion) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegion) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudRegion) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudRegion) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CloudRegion) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegion) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CloudRegion) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CloudRegion) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value
func (o *CloudRegion) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CloudRegion) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CloudRegion) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCapacity returns the Capacity field value
func (o *CloudRegion) GetCapacity() []RegionCapacityListItem {
	if o == nil {
		var ret []RegionCapacityListItem
		return ret
	}

	return o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value
// and a boolean to check if the value has been set.
func (o *CloudRegion) GetCapacityOk() (*[]RegionCapacityListItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Capacity, true
}

// SetCapacity sets field value
func (o *CloudRegion) SetCapacity(v []RegionCapacityListItem) {
	o.Capacity = v
}

func (o CloudRegion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["capacity"] = o.Capacity
	}
	return json.Marshal(toSerialize)
}

type NullableCloudRegion struct {
	value *CloudRegion
	isSet bool
}

func (v NullableCloudRegion) Get() *CloudRegion {
	return v.value
}

func (v *NullableCloudRegion) Set(val *CloudRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRegion(val *CloudRegion) *NullableCloudRegion {
	return &NullableCloudRegion{value: val, isSet: true}
}

func (v NullableCloudRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



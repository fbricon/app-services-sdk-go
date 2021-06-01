/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 1.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmt

import (
	"encoding/json"
	"time"
)

// ServiceAccountListItem struct for ServiceAccountListItem
type ServiceAccountListItem struct {
	// server generated unique id of the service account
	Id *string `json:"id,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Href *string `json:"href,omitempty"`
	// client id of the service account
	ClientID *string `json:"clientID,omitempty"`
	// name of the service account
	Name *string `json:"name,omitempty"`
	// owner of the service account
	Owner *string `json:"owner,omitempty"`
	// service account creation timestamp
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// description of the service account
	Description *string `json:"description,omitempty"`
}

// NewServiceAccountListItem instantiates a new ServiceAccountListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountListItem() *ServiceAccountListItem {
	this := ServiceAccountListItem{}
	return &this
}

// NewServiceAccountListItemWithDefaults instantiates a new ServiceAccountListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountListItemWithDefaults() *ServiceAccountListItem {
	this := ServiceAccountListItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceAccountListItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceAccountListItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceAccountListItem) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ServiceAccountListItem) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItem) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ServiceAccountListItem) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ServiceAccountListItem) SetKind(v string) {
	o.Kind = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ServiceAccountListItem) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItem) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ServiceAccountListItem) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ServiceAccountListItem) SetHref(v string) {
	o.Href = &v
}

// GetClientID returns the ClientID field value if set, zero value otherwise.
func (o *ServiceAccountListItem) GetClientID() string {
	if o == nil || o.ClientID == nil {
		var ret string
		return ret
	}
	return *o.ClientID
}

// GetClientIDOk returns a tuple with the ClientID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItem) GetClientIDOk() (*string, bool) {
	if o == nil || o.ClientID == nil {
		return nil, false
	}
	return o.ClientID, true
}

// HasClientID returns a boolean if a field has been set.
func (o *ServiceAccountListItem) HasClientID() bool {
	if o != nil && o.ClientID != nil {
		return true
	}

	return false
}

// SetClientID gets a reference to the given string and assigns it to the ClientID field.
func (o *ServiceAccountListItem) SetClientID(v string) {
	o.ClientID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceAccountListItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceAccountListItem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceAccountListItem) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ServiceAccountListItem) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItem) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ServiceAccountListItem) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ServiceAccountListItem) SetOwner(v string) {
	o.Owner = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServiceAccountListItem) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceAccountListItem) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ServiceAccountListItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceAccountListItem) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItem) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAccountListItem) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceAccountListItem) SetDescription(v string) {
	o.Description = &v
}

func (o ServiceAccountListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.ClientID != nil {
		toSerialize["clientID"] = o.ClientID
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableServiceAccountListItem struct {
	value *ServiceAccountListItem
	isSet bool
}

func (v NullableServiceAccountListItem) Get() *ServiceAccountListItem {
	return v.value
}

func (v *NullableServiceAccountListItem) Set(val *ServiceAccountListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountListItem(val *ServiceAccountListItem) *NullableServiceAccountListItem {
	return &NullableServiceAccountListItem{value: val, isSet: true}
}

func (v NullableServiceAccountListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



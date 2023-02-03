/*
 * Apicurio Registry API [v2]
 *
 * Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 
 *
 * API version: 2.4.x
 * Contact: apicurio@lists.jboss.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryinstanceclient

import (
	"encoding/json"
)

// GroupMetaData struct for GroupMetaData
type GroupMetaData struct {
	// An ID of a single artifact group.
	Id string `json:"id"`
	Description string `json:"description"`
	CreatedBy string `json:"createdBy"`
	CreatedOn string `json:"createdOn"`
	ModifiedBy string `json:"modifiedBy"`
	ModifiedOn string `json:"modifiedOn"`
	// User-defined name-value pairs. Name and value must be strings.
	Properties map[string]string `json:"properties"`
}

// NewGroupMetaData instantiates a new GroupMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMetaData(id string, description string, createdBy string, createdOn string, modifiedBy string, modifiedOn string, properties map[string]string) *GroupMetaData {
	this := GroupMetaData{}
	this.Id = id
	this.Description = description
	this.CreatedBy = createdBy
	this.CreatedOn = createdOn
	this.ModifiedBy = modifiedBy
	this.ModifiedOn = modifiedOn
	this.Properties = properties
	return &this
}

// NewGroupMetaDataWithDefaults instantiates a new GroupMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMetaDataWithDefaults() *GroupMetaData {
	this := GroupMetaData{}
	return &this
}

// GetId returns the Id field value
func (o *GroupMetaData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupMetaData) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupMetaData) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value
func (o *GroupMetaData) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GroupMetaData) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GroupMetaData) SetDescription(v string) {
	o.Description = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *GroupMetaData) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *GroupMetaData) GetCreatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *GroupMetaData) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *GroupMetaData) GetCreatedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *GroupMetaData) GetCreatedOnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *GroupMetaData) SetCreatedOn(v string) {
	o.CreatedOn = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *GroupMetaData) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *GroupMetaData) GetModifiedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *GroupMetaData) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetModifiedOn returns the ModifiedOn field value
func (o *GroupMetaData) GetModifiedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value
// and a boolean to check if the value has been set.
func (o *GroupMetaData) GetModifiedOnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedOn, true
}

// SetModifiedOn sets field value
func (o *GroupMetaData) SetModifiedOn(v string) {
	o.ModifiedOn = v
}

// GetProperties returns the Properties field value
func (o *GroupMetaData) GetProperties() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *GroupMetaData) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *GroupMetaData) SetProperties(v map[string]string) {
	o.Properties = v
}

func (o GroupMetaData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if true {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if true {
		toSerialize["modifiedOn"] = o.ModifiedOn
	}
	if true {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableGroupMetaData struct {
	value *GroupMetaData
	isSet bool
}

func (v NullableGroupMetaData) Get() *GroupMetaData {
	return v.value
}

func (v *NullableGroupMetaData) Set(val *GroupMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMetaData(val *GroupMetaData) *NullableGroupMetaData {
	return &NullableGroupMetaData{value: val, isSet: true}
}

func (v NullableGroupMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.10.1-SNAPSHOT
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// Topic Kafka Topic (A feed where records are stored and published)
type Topic struct {

	// The name of the topic.
	Name *string `json:"name,omitempty"`

	IsInternal *bool `json:"isInternal,omitempty"`

	// Partitions for this topic.
	Partitions *[]Partition `json:"partitions,omitempty"`

	// Topic configuration entry.
	Config *[]ConfigEntry `json:"config,omitempty"`

	// Unique identifier for the object. Not supported for all object kinds.
	Id *string `json:"id,omitempty"`

	Kind string `json:"kind"`

	// Link path to request the object. Not supported for all object kinds.
	Href *string `json:"href,omitempty"`

}

// NewTopic instantiates a new Topic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopic(kind string) *Topic {
	this := Topic{}
	this.Kind = kind
	return &this
}

// NewTopicWithDefaults instantiates a new Topic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopicWithDefaults() *Topic {
	this := Topic{}








	return &this
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *Topic) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Topic) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Topic) SetName(v string) {
	o.Name = &v
}


// GetIsInternal returns the IsInternal field value if set, zero value otherwise.
func (o *Topic) GetIsInternal() bool {
	if o == nil || o.IsInternal == nil {
		var ret bool
		return ret
	}
	return *o.IsInternal
}

// GetIsInternalOk returns a tuple with the IsInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic) GetIsInternalOk() (*bool, bool) {
	if o == nil || o.IsInternal == nil {
		return nil, false
	}
	return o.IsInternal, true
}

// HasIsInternal returns a boolean if a field has been set.
func (o *Topic) HasIsInternal() bool {
	if o != nil && o.IsInternal != nil {
		return true
	}

	return false
}

// SetIsInternal gets a reference to the given bool and assigns it to the IsInternal field.
func (o *Topic) SetIsInternal(v bool) {
	o.IsInternal = &v
}


// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *Topic) GetPartitions() []Partition {
	if o == nil || o.Partitions == nil {
		var ret []Partition
		return ret
	}
	return *o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic) GetPartitionsOk() (*[]Partition, bool) {
	if o == nil || o.Partitions == nil {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *Topic) HasPartitions() bool {
	if o != nil && o.Partitions != nil {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []Partition and assigns it to the Partitions field.
func (o *Topic) SetPartitions(v []Partition) {
	o.Partitions = &v
}


// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Topic) GetConfig() []ConfigEntry {
	if o == nil || o.Config == nil {
		var ret []ConfigEntry
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic) GetConfigOk() (*[]ConfigEntry, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Topic) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given []ConfigEntry and assigns it to the Config field.
func (o *Topic) SetConfig(v []ConfigEntry) {
	o.Config = &v
}


// GetId returns the Id field value if set, zero value otherwise.
func (o *Topic) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Topic) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Topic) SetId(v string) {
	o.Id = &v
}


// GetKind returns the Kind field value
func (o *Topic) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *Topic) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *Topic) SetKind(v string) {
	o.Kind = v
}


// GetHref returns the Href field value if set, zero value otherwise.
func (o *Topic) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Topic) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Topic) SetHref(v string) {
	o.Href = &v
}


func (o Topic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
    
	if o.IsInternal != nil {
		toSerialize["isInternal"] = o.IsInternal
	}
    
	if o.Partitions != nil {
		toSerialize["partitions"] = o.Partitions
	}
    
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
    
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
    
	if true {
		toSerialize["kind"] = o.Kind
	}
    
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
    
	return json.Marshal(toSerialize)
}

type NullableTopic struct {
	value *Topic
	isSet bool
}

func (v NullableTopic) Get() *Topic {
	return v.value
}

func (v *NullableTopic) Set(val *Topic) {
	v.value = val
	v.isSet = true
}

func (v NullableTopic) IsSet() bool {
	return v.isSet
}

func (v *NullableTopic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopic(val *Topic) *NullableTopic {
	return &NullableTopic{value: val, isSet: true}
}

func (v NullableTopic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


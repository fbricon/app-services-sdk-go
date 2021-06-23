/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// TopicsList A list of topics.
type TopicsList struct {

	// The page
	Page *int32 `json:"page,omitempty"`

	// number of entries per page
	Size *int32 `json:"size,omitempty"`

	// Deprecated offset of the topic list
	Offset *int32 `json:"offset,omitempty"`

	// Deprecated maximum of returned topics
	Limit *int32 `json:"limit,omitempty"`

	// Total number of topics
	Total *int32 `json:"total,omitempty"`

	// List of topics
	Items *[]Topic `json:"items,omitempty"`

}

// NewTopicsList instantiates a new TopicsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopicsList() *TopicsList {
	this := TopicsList{}
	return &this
}

// NewTopicsListWithDefaults instantiates a new TopicsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopicsListWithDefaults() *TopicsList {
	this := TopicsList{}







	return &this
}


// GetPage returns the Page field value if set, zero value otherwise.
func (o *TopicsList) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicsList) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *TopicsList) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *TopicsList) SetPage(v int32) {
	o.Page = &v
}


// GetSize returns the Size field value if set, zero value otherwise.
func (o *TopicsList) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicsList) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *TopicsList) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *TopicsList) SetSize(v int32) {
	o.Size = &v
}


// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *TopicsList) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicsList) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *TopicsList) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *TopicsList) SetOffset(v int32) {
	o.Offset = &v
}


// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *TopicsList) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicsList) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *TopicsList) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *TopicsList) SetLimit(v int32) {
	o.Limit = &v
}


// GetTotal returns the Total field value if set, zero value otherwise.
func (o *TopicsList) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicsList) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *TopicsList) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *TopicsList) SetTotal(v int32) {
	o.Total = &v
}


// GetItems returns the Items field value if set, zero value otherwise.
func (o *TopicsList) GetItems() []Topic {
	if o == nil || o.Items == nil {
		var ret []Topic
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicsList) GetItemsOk() (*[]Topic, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *TopicsList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Topic and assigns it to the Items field.
func (o *TopicsList) SetItems(v []Topic) {
	o.Items = &v
}


func (o TopicsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
    
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
    
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
    
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
    
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
    
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
    
	return json.Marshal(toSerialize)
}

type NullableTopicsList struct {
	value *TopicsList
	isSet bool
}

func (v NullableTopicsList) Get() *TopicsList {
	return v.value
}

func (v *NullableTopicsList) Set(val *TopicsList) {
	v.value = val
	v.isSet = true
}

func (v NullableTopicsList) IsSet() bool {
	return v.isSet
}

func (v *NullableTopicsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopicsList(val *TopicsList) *NullableTopicsList {
	return &NullableTopicsList{value: val, isSet: true}
}

func (v NullableTopicsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopicsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


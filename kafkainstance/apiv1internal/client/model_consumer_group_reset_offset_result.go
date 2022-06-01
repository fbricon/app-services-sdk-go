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

// ConsumerGroupResetOffsetResult struct for ConsumerGroupResetOffsetResult
type ConsumerGroupResetOffsetResult struct {

	Items *[]ConsumerGroupResetOffsetResultItem `json:"items,omitempty"`

	Kind *string `json:"kind,omitempty"`

	// Total number of entries in the full result set
	Total int32 `json:"total"`

	// Number of entries per page (returned for fetch requests)
	Size *int32 `json:"size,omitempty"`

	// Current page number (returned for fetch requests)
	Page *int32 `json:"page,omitempty"`

}

// NewConsumerGroupResetOffsetResult instantiates a new ConsumerGroupResetOffsetResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerGroupResetOffsetResult(items []map[string]interface{}, total int32) *ConsumerGroupResetOffsetResult {
	this := ConsumerGroupResetOffsetResult{}
	this.Items = items
	this.Total = total
	return &this
}

// NewConsumerGroupResetOffsetResultWithDefaults instantiates a new ConsumerGroupResetOffsetResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerGroupResetOffsetResultWithDefaults() *ConsumerGroupResetOffsetResult {
	this := ConsumerGroupResetOffsetResult{}






	return &this
}


// GetItems returns the Items field value if set, zero value otherwise.
func (o *ConsumerGroupResetOffsetResult) GetItems() []ConsumerGroupResetOffsetResultItem {
	if o == nil || o.Items == nil {
		var ret []ConsumerGroupResetOffsetResultItem
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroupResetOffsetResult) GetItemsOk() (*[]ConsumerGroupResetOffsetResultItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ConsumerGroupResetOffsetResult) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ConsumerGroupResetOffsetResultItem and assigns it to the Items field.
func (o *ConsumerGroupResetOffsetResult) SetItems(v []ConsumerGroupResetOffsetResultItem) {
	o.Items = &v
}


// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ConsumerGroupResetOffsetResult) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroupResetOffsetResult) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ConsumerGroupResetOffsetResult) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ConsumerGroupResetOffsetResult) SetKind(v string) {
	o.Kind = &v
}


// GetTotal returns the Total field value
func (o *ConsumerGroupResetOffsetResult) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupResetOffsetResult) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ConsumerGroupResetOffsetResult) SetTotal(v int32) {
	o.Total = v
}


// GetSize returns the Size field value if set, zero value otherwise.
func (o *ConsumerGroupResetOffsetResult) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroupResetOffsetResult) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ConsumerGroupResetOffsetResult) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *ConsumerGroupResetOffsetResult) SetSize(v int32) {
	o.Size = &v
}


// GetPage returns the Page field value if set, zero value otherwise.
func (o *ConsumerGroupResetOffsetResult) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroupResetOffsetResult) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ConsumerGroupResetOffsetResult) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ConsumerGroupResetOffsetResult) SetPage(v int32) {
	o.Page = &v
}


func (o ConsumerGroupResetOffsetResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
    
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
    
	if true {
		toSerialize["total"] = o.Total
	}
    
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
    
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
    
	return json.Marshal(toSerialize)
}

type NullableConsumerGroupResetOffsetResult struct {
	value *ConsumerGroupResetOffsetResult
	isSet bool
}

func (v NullableConsumerGroupResetOffsetResult) Get() *ConsumerGroupResetOffsetResult {
	return v.value
}

func (v *NullableConsumerGroupResetOffsetResult) Set(val *ConsumerGroupResetOffsetResult) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerGroupResetOffsetResult) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerGroupResetOffsetResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerGroupResetOffsetResult(val *ConsumerGroupResetOffsetResult) *NullableConsumerGroupResetOffsetResult {
	return &NullableConsumerGroupResetOffsetResult{value: val, isSet: true}
}

func (v NullableConsumerGroupResetOffsetResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerGroupResetOffsetResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


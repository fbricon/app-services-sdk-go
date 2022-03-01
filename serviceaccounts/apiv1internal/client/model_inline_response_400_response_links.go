/*
 * Service Accounts API Documentation
 *
 * This is the API documentation for Service Accounts
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceaccountsclient

import (
	"encoding/json"
)

// InlineResponse400ResponseLinks struct for InlineResponse400ResponseLinks
type InlineResponse400ResponseLinks struct {

	UriBuilder *map[string]interface{} `json:"uriBuilder,omitempty"`

	Params *map[string]string `json:"params,omitempty"`

	Title *string `json:"title,omitempty"`

	Uri *string `json:"uri,omitempty"`

	Rel *string `json:"rel,omitempty"`

	Rels *[]string `json:"rels,omitempty"`

	Type *string `json:"type,omitempty"`

}

// NewInlineResponse400ResponseLinks instantiates a new InlineResponse400ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse400ResponseLinks() *InlineResponse400ResponseLinks {
	this := InlineResponse400ResponseLinks{}
	return &this
}

// NewInlineResponse400ResponseLinksWithDefaults instantiates a new InlineResponse400ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse400ResponseLinksWithDefaults() *InlineResponse400ResponseLinks {
	this := InlineResponse400ResponseLinks{}








	return &this
}


// GetUriBuilder returns the UriBuilder field value if set, zero value otherwise.
func (o *InlineResponse400ResponseLinks) GetUriBuilder() map[string]interface{} {
	if o == nil || o.UriBuilder == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.UriBuilder
}

// GetUriBuilderOk returns a tuple with the UriBuilder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400ResponseLinks) GetUriBuilderOk() (*map[string]interface{}, bool) {
	if o == nil || o.UriBuilder == nil {
		return nil, false
	}
	return o.UriBuilder, true
}

// HasUriBuilder returns a boolean if a field has been set.
func (o *InlineResponse400ResponseLinks) HasUriBuilder() bool {
	if o != nil && o.UriBuilder != nil {
		return true
	}

	return false
}

// SetUriBuilder gets a reference to the given map[string]interface{} and assigns it to the UriBuilder field.
func (o *InlineResponse400ResponseLinks) SetUriBuilder(v map[string]interface{}) {
	o.UriBuilder = &v
}


// GetParams returns the Params field value if set, zero value otherwise.
func (o *InlineResponse400ResponseLinks) GetParams() map[string]string {
	if o == nil || o.Params == nil {
		var ret map[string]string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400ResponseLinks) GetParamsOk() (*map[string]string, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *InlineResponse400ResponseLinks) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]string and assigns it to the Params field.
func (o *InlineResponse400ResponseLinks) SetParams(v map[string]string) {
	o.Params = &v
}


// GetTitle returns the Title field value if set, zero value otherwise.
func (o *InlineResponse400ResponseLinks) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400ResponseLinks) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *InlineResponse400ResponseLinks) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *InlineResponse400ResponseLinks) SetTitle(v string) {
	o.Title = &v
}


// GetUri returns the Uri field value if set, zero value otherwise.
func (o *InlineResponse400ResponseLinks) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400ResponseLinks) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *InlineResponse400ResponseLinks) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *InlineResponse400ResponseLinks) SetUri(v string) {
	o.Uri = &v
}


// GetRel returns the Rel field value if set, zero value otherwise.
func (o *InlineResponse400ResponseLinks) GetRel() string {
	if o == nil || o.Rel == nil {
		var ret string
		return ret
	}
	return *o.Rel
}

// GetRelOk returns a tuple with the Rel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400ResponseLinks) GetRelOk() (*string, bool) {
	if o == nil || o.Rel == nil {
		return nil, false
	}
	return o.Rel, true
}

// HasRel returns a boolean if a field has been set.
func (o *InlineResponse400ResponseLinks) HasRel() bool {
	if o != nil && o.Rel != nil {
		return true
	}

	return false
}

// SetRel gets a reference to the given string and assigns it to the Rel field.
func (o *InlineResponse400ResponseLinks) SetRel(v string) {
	o.Rel = &v
}


// GetRels returns the Rels field value if set, zero value otherwise.
func (o *InlineResponse400ResponseLinks) GetRels() []string {
	if o == nil || o.Rels == nil {
		var ret []string
		return ret
	}
	return *o.Rels
}

// GetRelsOk returns a tuple with the Rels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400ResponseLinks) GetRelsOk() (*[]string, bool) {
	if o == nil || o.Rels == nil {
		return nil, false
	}
	return o.Rels, true
}

// HasRels returns a boolean if a field has been set.
func (o *InlineResponse400ResponseLinks) HasRels() bool {
	if o != nil && o.Rels != nil {
		return true
	}

	return false
}

// SetRels gets a reference to the given []string and assigns it to the Rels field.
func (o *InlineResponse400ResponseLinks) SetRels(v []string) {
	o.Rels = &v
}


// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse400ResponseLinks) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400ResponseLinks) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse400ResponseLinks) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse400ResponseLinks) SetType(v string) {
	o.Type = &v
}


func (o InlineResponse400ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.UriBuilder != nil {
		toSerialize["uriBuilder"] = o.UriBuilder
	}
    
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
    
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
    
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
    
	if o.Rel != nil {
		toSerialize["rel"] = o.Rel
	}
    
	if o.Rels != nil {
		toSerialize["rels"] = o.Rels
	}
    
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
    
	return json.Marshal(toSerialize)
}

type NullableInlineResponse400ResponseLinks struct {
	value *InlineResponse400ResponseLinks
	isSet bool
}

func (v NullableInlineResponse400ResponseLinks) Get() *InlineResponse400ResponseLinks {
	return v.value
}

func (v *NullableInlineResponse400ResponseLinks) Set(val *InlineResponse400ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse400ResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse400ResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse400ResponseLinks(val *InlineResponse400ResponseLinks) *NullableInlineResponse400ResponseLinks {
	return &NullableInlineResponse400ResponseLinks{value: val, isSet: true}
}

func (v NullableInlineResponse400ResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse400ResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


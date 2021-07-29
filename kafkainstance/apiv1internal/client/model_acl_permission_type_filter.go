/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.3.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
	"fmt"
)

// AclPermissionTypeFilter - struct for AclPermissionTypeFilter
type AclPermissionTypeFilter struct {
	AclFilterAny *AclFilterAny
	AclPermissionType *AclPermissionType
}

// AclFilterAnyAsAclPermissionTypeFilter is a convenience function that returns AclFilterAny wrapped in AclPermissionTypeFilter
func AclFilterAnyAsAclPermissionTypeFilter(v *AclFilterAny) AclPermissionTypeFilter {
	return AclPermissionTypeFilter{ AclFilterAny: v}
}

// AclPermissionTypeAsAclPermissionTypeFilter is a convenience function that returns AclPermissionType wrapped in AclPermissionTypeFilter
func AclPermissionTypeAsAclPermissionTypeFilter(v *AclPermissionType) AclPermissionTypeFilter {
	return AclPermissionTypeFilter{ AclPermissionType: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AclPermissionTypeFilter) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AclFilterAny
	err = json.Unmarshal(data, &dst.AclFilterAny)
	if err == nil {
		jsonAclFilterAny, _ := json.Marshal(dst.AclFilterAny)
		if string(jsonAclFilterAny) == "{}" { // empty struct
			dst.AclFilterAny = nil
		} else {
			match++
		}
	} else {
		dst.AclFilterAny = nil
	}

	// try to unmarshal data into AclPermissionType
	err = json.Unmarshal(data, &dst.AclPermissionType)
	if err == nil {
		jsonAclPermissionType, _ := json.Marshal(dst.AclPermissionType)
		if string(jsonAclPermissionType) == "{}" { // empty struct
			dst.AclPermissionType = nil
		} else {
			match++
		}
	} else {
		dst.AclPermissionType = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AclFilterAny = nil
		dst.AclPermissionType = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(AclPermissionTypeFilter)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(AclPermissionTypeFilter)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AclPermissionTypeFilter) MarshalJSON() ([]byte, error) {
	if src.AclFilterAny != nil {
		return json.Marshal(&src.AclFilterAny)
	}

	if src.AclPermissionType != nil {
		return json.Marshal(&src.AclPermissionType)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AclPermissionTypeFilter) GetActualInstance() (interface{}) {
	if obj.AclFilterAny != nil {
		return obj.AclFilterAny
	}

	if obj.AclPermissionType != nil {
		return obj.AclPermissionType
	}

	// all schemas are nil
	return nil
}

type NullableAclPermissionTypeFilter struct {
	value *AclPermissionTypeFilter
	isSet bool
}

func (v NullableAclPermissionTypeFilter) Get() *AclPermissionTypeFilter {
	return v.value
}

func (v *NullableAclPermissionTypeFilter) Set(val *AclPermissionTypeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAclPermissionTypeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAclPermissionTypeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAclPermissionTypeFilter(val *AclPermissionTypeFilter) *NullableAclPermissionTypeFilter {
	return &NullableAclPermissionTypeFilter{value: val, isSet: true}
}

func (v NullableAclPermissionTypeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAclPermissionTypeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


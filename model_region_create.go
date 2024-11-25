/*
 * IONOS Cloud - Object Storage Management API
 *
 * Object Storage Management API is a RESTful API that manages the object storage service configuration for IONOS Cloud.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// RegionCreate struct for RegionCreate
type RegionCreate struct {
	// Metadata
	Metadata   *map[string]interface{} `json:"metadata,omitempty"`
	Properties *Region                 `json:"properties"`
}

// NewRegionCreate instantiates a new RegionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionCreate(properties Region) *RegionCreate {
	this := RegionCreate{}

	this.Properties = &properties

	return &this
}

// NewRegionCreateWithDefaults instantiates a new RegionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionCreateWithDefaults() *RegionCreate {
	this := RegionCreate{}
	return &this
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RegionCreate) GetMetadata() *map[string]interface{} {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegionCreate) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *RegionCreate) SetMetadata(v map[string]interface{}) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *RegionCreate) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for Region will be returned
func (o *RegionCreate) GetProperties() *Region {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegionCreate) GetPropertiesOk() (*Region, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *RegionCreate) SetProperties(v Region) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *RegionCreate) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o RegionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullableRegionCreate struct {
	value *RegionCreate
	isSet bool
}

func (v NullableRegionCreate) Get() *RegionCreate {
	return v.value
}

func (v *NullableRegionCreate) Set(val *RegionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionCreate(val *RegionCreate) *NullableRegionCreate {
	return &NullableRegionCreate{value: val, isSet: true}
}

func (v NullableRegionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

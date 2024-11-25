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

// RegionRead struct for RegionRead
type RegionRead struct {
	// The Region of the Region.
	Id *string `json:"id"`
	// The type of the resource.
	Type *string `json:"type"`
	// The URL of the Region.
	Href       *string                 `json:"href"`
	Metadata   *map[string]interface{} `json:"metadata"`
	Properties *Region                 `json:"properties"`
}

// NewRegionRead instantiates a new RegionRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionRead(id string, type_ string, href string, metadata map[string]interface{}, properties Region) *RegionRead {
	this := RegionRead{}

	this.Id = &id
	this.Type = &type_
	this.Href = &href
	this.Metadata = &metadata
	this.Properties = &properties

	return &this
}

// NewRegionReadWithDefaults instantiates a new RegionRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionReadWithDefaults() *RegionRead {
	this := RegionRead{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RegionRead) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegionRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *RegionRead) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *RegionRead) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RegionRead) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegionRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *RegionRead) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *RegionRead) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RegionRead) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegionRead) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *RegionRead) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *RegionRead) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RegionRead) GetMetadata() *map[string]interface{} {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegionRead) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *RegionRead) SetMetadata(v map[string]interface{}) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *RegionRead) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for Region will be returned
func (o *RegionRead) GetProperties() *Region {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegionRead) GetPropertiesOk() (*Region, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *RegionRead) SetProperties(v Region) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *RegionRead) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o RegionRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullableRegionRead struct {
	value *RegionRead
	isSet bool
}

func (v NullableRegionRead) Get() *RegionRead {
	return v.value
}

func (v *NullableRegionRead) Set(val *RegionRead) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionRead) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionRead(val *RegionRead) *NullableRegionRead {
	return &NullableRegionRead{value: val, isSet: true}
}

func (v NullableRegionRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
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

// StorageClassReadListAllOf struct for StorageClassReadListAllOf
type StorageClassReadListAllOf struct {
	// ID of the list of StorageClass resources.
	Id *string `json:"id"`
	// The type of the resource.
	Type *string `json:"type"`
	// The URL of the list of StorageClass resources.
	Href *string `json:"href"`
	// The list of StorageClass resources.
	Items *[]StorageClassRead `json:"items,omitempty"`
}

// NewStorageClassReadListAllOf instantiates a new StorageClassReadListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageClassReadListAllOf(id string, type_ string, href string) *StorageClassReadListAllOf {
	this := StorageClassReadListAllOf{}

	this.Id = &id
	this.Type = &type_
	this.Href = &href

	return &this
}

// NewStorageClassReadListAllOfWithDefaults instantiates a new StorageClassReadListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageClassReadListAllOfWithDefaults() *StorageClassReadListAllOf {
	this := StorageClassReadListAllOf{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorageClassReadListAllOf) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageClassReadListAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *StorageClassReadListAllOf) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *StorageClassReadListAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorageClassReadListAllOf) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageClassReadListAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *StorageClassReadListAllOf) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *StorageClassReadListAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorageClassReadListAllOf) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageClassReadListAllOf) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *StorageClassReadListAllOf) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *StorageClassReadListAllOf) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []StorageClassRead will be returned
func (o *StorageClassReadListAllOf) GetItems() *[]StorageClassRead {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageClassReadListAllOf) GetItemsOk() (*[]StorageClassRead, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *StorageClassReadListAllOf) SetItems(v []StorageClassRead) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *StorageClassReadListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

func (o StorageClassReadListAllOf) MarshalJSON() ([]byte, error) {
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

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	return json.Marshal(toSerialize)
}

type NullableStorageClassReadListAllOf struct {
	value *StorageClassReadListAllOf
	isSet bool
}

func (v NullableStorageClassReadListAllOf) Get() *StorageClassReadListAllOf {
	return v.value
}

func (v *NullableStorageClassReadListAllOf) Set(val *StorageClassReadListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageClassReadListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageClassReadListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageClassReadListAllOf(val *StorageClassReadListAllOf) *NullableStorageClassReadListAllOf {
	return &NullableStorageClassReadListAllOf{value: val, isSet: true}
}

func (v NullableStorageClassReadListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageClassReadListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

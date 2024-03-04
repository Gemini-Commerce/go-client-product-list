/*
Collection Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package productlist

import (
	"encoding/json"
)

// checks if the ProductlistDeleteProductListAssociationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistDeleteProductListAssociationRequest{}

// ProductlistDeleteProductListAssociationRequest struct for ProductlistDeleteProductListAssociationRequest
type ProductlistDeleteProductListAssociationRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewProductlistDeleteProductListAssociationRequest instantiates a new ProductlistDeleteProductListAssociationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistDeleteProductListAssociationRequest() *ProductlistDeleteProductListAssociationRequest {
	this := ProductlistDeleteProductListAssociationRequest{}
	return &this
}

// NewProductlistDeleteProductListAssociationRequestWithDefaults instantiates a new ProductlistDeleteProductListAssociationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistDeleteProductListAssociationRequestWithDefaults() *ProductlistDeleteProductListAssociationRequest {
	this := ProductlistDeleteProductListAssociationRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductlistDeleteProductListAssociationRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistDeleteProductListAssociationRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductlistDeleteProductListAssociationRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductlistDeleteProductListAssociationRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductlistDeleteProductListAssociationRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistDeleteProductListAssociationRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductlistDeleteProductListAssociationRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductlistDeleteProductListAssociationRequest) SetId(v string) {
	o.Id = &v
}

func (o ProductlistDeleteProductListAssociationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistDeleteProductListAssociationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableProductlistDeleteProductListAssociationRequest struct {
	value *ProductlistDeleteProductListAssociationRequest
	isSet bool
}

func (v NullableProductlistDeleteProductListAssociationRequest) Get() *ProductlistDeleteProductListAssociationRequest {
	return v.value
}

func (v *NullableProductlistDeleteProductListAssociationRequest) Set(val *ProductlistDeleteProductListAssociationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistDeleteProductListAssociationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistDeleteProductListAssociationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistDeleteProductListAssociationRequest(val *ProductlistDeleteProductListAssociationRequest) *NullableProductlistDeleteProductListAssociationRequest {
	return &NullableProductlistDeleteProductListAssociationRequest{value: val, isSet: true}
}

func (v NullableProductlistDeleteProductListAssociationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistDeleteProductListAssociationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



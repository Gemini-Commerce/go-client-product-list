/*
Collection Service

API for managing collection

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product-list

import (
	"encoding/json"
)

// checks if the ProductlistListProductListsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistListProductListsRequest{}

// ProductlistListProductListsRequest struct for ProductlistListProductListsRequest
type ProductlistListProductListsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	PageSize *int64 `json:"pageSize,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
}

// NewProductlistListProductListsRequest instantiates a new ProductlistListProductListsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistListProductListsRequest() *ProductlistListProductListsRequest {
	this := ProductlistListProductListsRequest{}
	return &this
}

// NewProductlistListProductListsRequestWithDefaults instantiates a new ProductlistListProductListsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistListProductListsRequestWithDefaults() *ProductlistListProductListsRequest {
	this := ProductlistListProductListsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductlistListProductListsRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistListProductListsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductlistListProductListsRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductlistListProductListsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ProductlistListProductListsRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistListProductListsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ProductlistListProductListsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *ProductlistListProductListsRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *ProductlistListProductListsRequest) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistListProductListsRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *ProductlistListProductListsRequest) HasPageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *ProductlistListProductListsRequest) SetPageToken(v string) {
	o.PageToken = &v
}

func (o ProductlistListProductListsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistListProductListsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	return toSerialize, nil
}

type NullableProductlistListProductListsRequest struct {
	value *ProductlistListProductListsRequest
	isSet bool
}

func (v NullableProductlistListProductListsRequest) Get() *ProductlistListProductListsRequest {
	return v.value
}

func (v *NullableProductlistListProductListsRequest) Set(val *ProductlistListProductListsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistListProductListsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistListProductListsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistListProductListsRequest(val *ProductlistListProductListsRequest) *NullableProductlistListProductListsRequest {
	return &NullableProductlistListProductListsRequest{value: val, isSet: true}
}

func (v NullableProductlistListProductListsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistListProductListsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


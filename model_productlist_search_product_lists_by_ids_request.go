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

// checks if the ProductlistSearchProductListsByIdsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistSearchProductListsByIdsRequest{}

// ProductlistSearchProductListsByIdsRequest struct for ProductlistSearchProductListsByIdsRequest
type ProductlistSearchProductListsByIdsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Ids []string `json:"ids,omitempty"`
	PageSize *int64 `json:"pageSize,omitempty"`
	PageNumber *int64 `json:"pageNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistSearchProductListsByIdsRequest ProductlistSearchProductListsByIdsRequest

// NewProductlistSearchProductListsByIdsRequest instantiates a new ProductlistSearchProductListsByIdsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistSearchProductListsByIdsRequest() *ProductlistSearchProductListsByIdsRequest {
	this := ProductlistSearchProductListsByIdsRequest{}
	return &this
}

// NewProductlistSearchProductListsByIdsRequestWithDefaults instantiates a new ProductlistSearchProductListsByIdsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistSearchProductListsByIdsRequestWithDefaults() *ProductlistSearchProductListsByIdsRequest {
	this := ProductlistSearchProductListsByIdsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductlistSearchProductListsByIdsRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSearchProductListsByIdsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductlistSearchProductListsByIdsRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductlistSearchProductListsByIdsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ProductlistSearchProductListsByIdsRequest) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSearchProductListsByIdsRequest) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ProductlistSearchProductListsByIdsRequest) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *ProductlistSearchProductListsByIdsRequest) SetIds(v []string) {
	o.Ids = v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ProductlistSearchProductListsByIdsRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSearchProductListsByIdsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ProductlistSearchProductListsByIdsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *ProductlistSearchProductListsByIdsRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *ProductlistSearchProductListsByIdsRequest) GetPageNumber() int64 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int64
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSearchProductListsByIdsRequest) GetPageNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *ProductlistSearchProductListsByIdsRequest) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int64 and assigns it to the PageNumber field.
func (o *ProductlistSearchProductListsByIdsRequest) SetPageNumber(v int64) {
	o.PageNumber = &v
}

func (o ProductlistSearchProductListsByIdsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistSearchProductListsByIdsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.PageNumber) {
		toSerialize["pageNumber"] = o.PageNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductlistSearchProductListsByIdsRequest) UnmarshalJSON(data []byte) (err error) {
	varProductlistSearchProductListsByIdsRequest := _ProductlistSearchProductListsByIdsRequest{}

	err = json.Unmarshal(data, &varProductlistSearchProductListsByIdsRequest)

	if err != nil {
		return err
	}

	*o = ProductlistSearchProductListsByIdsRequest(varProductlistSearchProductListsByIdsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "ids")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "pageNumber")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistSearchProductListsByIdsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ProductlistSearchProductListsByIdsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableProductlistSearchProductListsByIdsRequest struct {
	value *ProductlistSearchProductListsByIdsRequest
	isSet bool
}

func (v NullableProductlistSearchProductListsByIdsRequest) Get() *ProductlistSearchProductListsByIdsRequest {
	return v.value
}

func (v *NullableProductlistSearchProductListsByIdsRequest) Set(val *ProductlistSearchProductListsByIdsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistSearchProductListsByIdsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistSearchProductListsByIdsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistSearchProductListsByIdsRequest(val *ProductlistSearchProductListsByIdsRequest) *NullableProductlistSearchProductListsByIdsRequest {
	return &NullableProductlistSearchProductListsByIdsRequest{value: val, isSet: true}
}

func (v NullableProductlistSearchProductListsByIdsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistSearchProductListsByIdsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



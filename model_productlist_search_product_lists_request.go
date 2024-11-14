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

// checks if the ProductlistSearchProductListsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistSearchProductListsRequest{}

// ProductlistSearchProductListsRequest struct for ProductlistSearchProductListsRequest
type ProductlistSearchProductListsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Query *SearchProductListsRequestQuery `json:"query,omitempty"`
	PageSize *int64 `json:"pageSize,omitempty"`
	PageNumber *int64 `json:"pageNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistSearchProductListsRequest ProductlistSearchProductListsRequest

// NewProductlistSearchProductListsRequest instantiates a new ProductlistSearchProductListsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistSearchProductListsRequest() *ProductlistSearchProductListsRequest {
	this := ProductlistSearchProductListsRequest{}
	return &this
}

// NewProductlistSearchProductListsRequestWithDefaults instantiates a new ProductlistSearchProductListsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistSearchProductListsRequestWithDefaults() *ProductlistSearchProductListsRequest {
	this := ProductlistSearchProductListsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductlistSearchProductListsRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSearchProductListsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// &#39;Has&#39;TenantId returns a boolean if a field has been set.
func (o *ProductlistSearchProductListsRequest) &#39;Has&#39;TenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductlistSearchProductListsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ProductlistSearchProductListsRequest) GetQuery() SearchProductListsRequestQuery {
	if o == nil || IsNil(o.Query) {
		var ret SearchProductListsRequestQuery
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSearchProductListsRequest) GetQueryOk() (*SearchProductListsRequestQuery, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// &#39;Has&#39;Query returns a boolean if a field has been set.
func (o *ProductlistSearchProductListsRequest) &#39;Has&#39;Query() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given SearchProductListsRequestQuery and assigns it to the Query field.
func (o *ProductlistSearchProductListsRequest) SetQuery(v SearchProductListsRequestQuery) {
	o.Query = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ProductlistSearchProductListsRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSearchProductListsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// &#39;Has&#39;PageSize returns a boolean if a field has been set.
func (o *ProductlistSearchProductListsRequest) &#39;Has&#39;PageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *ProductlistSearchProductListsRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *ProductlistSearchProductListsRequest) GetPageNumber() int64 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int64
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSearchProductListsRequest) GetPageNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// &#39;Has&#39;PageNumber returns a boolean if a field has been set.
func (o *ProductlistSearchProductListsRequest) &#39;Has&#39;PageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int64 and assigns it to the PageNumber field.
func (o *ProductlistSearchProductListsRequest) SetPageNumber(v int64) {
	o.PageNumber = &v
}

func (o ProductlistSearchProductListsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistSearchProductListsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
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

func (o *ProductlistSearchProductListsRequest) UnmarshalJSON(data []byte) (err error) {
	varProductlistSearchProductListsRequest := _ProductlistSearchProductListsRequest{}

	err = json.Unmarshal(data, &varProductlistSearchProductListsRequest)

	if err != nil {
		return err
	}

	*o = ProductlistSearchProductListsRequest(varProductlistSearchProductListsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "query")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "pageNumber")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistSearchProductListsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ProductlistSearchProductListsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableProductlistSearchProductListsRequest struct {
	value *ProductlistSearchProductListsRequest
	isSet bool
}

func (v NullableProductlistSearchProductListsRequest) Get() *ProductlistSearchProductListsRequest {
	return v.value
}

func (v *NullableProductlistSearchProductListsRequest) Set(val *ProductlistSearchProductListsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistSearchProductListsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistSearchProductListsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistSearchProductListsRequest(val *ProductlistSearchProductListsRequest) *NullableProductlistSearchProductListsRequest {
	return &NullableProductlistSearchProductListsRequest{value: val, isSet: true}
}

func (v NullableProductlistSearchProductListsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistSearchProductListsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



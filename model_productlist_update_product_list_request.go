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

// checks if the ProductlistUpdateProductListRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistUpdateProductListRequest{}

// ProductlistUpdateProductListRequest struct for ProductlistUpdateProductListRequest
type ProductlistUpdateProductListRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Id *string `json:"id,omitempty"`
	UrlKey *ProductlistLocalizedText `json:"urlKey,omitempty"`
	Attributes *map[string]ProtobufAny `json:"attributes,omitempty"`
	AttributesMask *string `json:"attributesMask,omitempty"`
}

// NewProductlistUpdateProductListRequest instantiates a new ProductlistUpdateProductListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistUpdateProductListRequest() *ProductlistUpdateProductListRequest {
	this := ProductlistUpdateProductListRequest{}
	return &this
}

// NewProductlistUpdateProductListRequestWithDefaults instantiates a new ProductlistUpdateProductListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistUpdateProductListRequestWithDefaults() *ProductlistUpdateProductListRequest {
	this := ProductlistUpdateProductListRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductlistUpdateProductListRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistUpdateProductListRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductlistUpdateProductListRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductlistUpdateProductListRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductlistUpdateProductListRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistUpdateProductListRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductlistUpdateProductListRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductlistUpdateProductListRequest) SetId(v string) {
	o.Id = &v
}

// GetUrlKey returns the UrlKey field value if set, zero value otherwise.
func (o *ProductlistUpdateProductListRequest) GetUrlKey() ProductlistLocalizedText {
	if o == nil || IsNil(o.UrlKey) {
		var ret ProductlistLocalizedText
		return ret
	}
	return *o.UrlKey
}

// GetUrlKeyOk returns a tuple with the UrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistUpdateProductListRequest) GetUrlKeyOk() (*ProductlistLocalizedText, bool) {
	if o == nil || IsNil(o.UrlKey) {
		return nil, false
	}
	return o.UrlKey, true
}

// HasUrlKey returns a boolean if a field has been set.
func (o *ProductlistUpdateProductListRequest) HasUrlKey() bool {
	if o != nil && !IsNil(o.UrlKey) {
		return true
	}

	return false
}

// SetUrlKey gets a reference to the given ProductlistLocalizedText and assigns it to the UrlKey field.
func (o *ProductlistUpdateProductListRequest) SetUrlKey(v ProductlistLocalizedText) {
	o.UrlKey = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ProductlistUpdateProductListRequest) GetAttributes() map[string]ProtobufAny {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistUpdateProductListRequest) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ProductlistUpdateProductListRequest) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *ProductlistUpdateProductListRequest) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

// GetAttributesMask returns the AttributesMask field value if set, zero value otherwise.
func (o *ProductlistUpdateProductListRequest) GetAttributesMask() string {
	if o == nil || IsNil(o.AttributesMask) {
		var ret string
		return ret
	}
	return *o.AttributesMask
}

// GetAttributesMaskOk returns a tuple with the AttributesMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistUpdateProductListRequest) GetAttributesMaskOk() (*string, bool) {
	if o == nil || IsNil(o.AttributesMask) {
		return nil, false
	}
	return o.AttributesMask, true
}

// HasAttributesMask returns a boolean if a field has been set.
func (o *ProductlistUpdateProductListRequest) HasAttributesMask() bool {
	if o != nil && !IsNil(o.AttributesMask) {
		return true
	}

	return false
}

// SetAttributesMask gets a reference to the given string and assigns it to the AttributesMask field.
func (o *ProductlistUpdateProductListRequest) SetAttributesMask(v string) {
	o.AttributesMask = &v
}

func (o ProductlistUpdateProductListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistUpdateProductListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UrlKey) {
		toSerialize["urlKey"] = o.UrlKey
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.AttributesMask) {
		toSerialize["attributesMask"] = o.AttributesMask
	}
	return toSerialize, nil
}

type NullableProductlistUpdateProductListRequest struct {
	value *ProductlistUpdateProductListRequest
	isSet bool
}

func (v NullableProductlistUpdateProductListRequest) Get() *ProductlistUpdateProductListRequest {
	return v.value
}

func (v *NullableProductlistUpdateProductListRequest) Set(val *ProductlistUpdateProductListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistUpdateProductListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistUpdateProductListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistUpdateProductListRequest(val *ProductlistUpdateProductListRequest) *NullableProductlistUpdateProductListRequest {
	return &NullableProductlistUpdateProductListRequest{value: val, isSet: true}
}

func (v NullableProductlistUpdateProductListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistUpdateProductListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



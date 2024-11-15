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

// checks if the ProductlistCreateProductListRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistCreateProductListRequest{}

// ProductlistCreateProductListRequest struct for ProductlistCreateProductListRequest
type ProductlistCreateProductListRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Code *string `json:"code,omitempty"`
	UrlKey *ProductlistLocalizedText `json:"urlKey,omitempty"`
	EntityType *string `json:"entityType,omitempty"`
	EntityCode *string `json:"entityCode,omitempty"`
	Attributes *map[string]ProtobufAny `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistCreateProductListRequest ProductlistCreateProductListRequest

// NewProductlistCreateProductListRequest instantiates a new ProductlistCreateProductListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistCreateProductListRequest() *ProductlistCreateProductListRequest {
	this := ProductlistCreateProductListRequest{}
	return &this
}

// NewProductlistCreateProductListRequestWithDefaults instantiates a new ProductlistCreateProductListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistCreateProductListRequestWithDefaults() *ProductlistCreateProductListRequest {
	this := ProductlistCreateProductListRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductlistCreateProductListRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistCreateProductListRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductlistCreateProductListRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductlistCreateProductListRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProductlistCreateProductListRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistCreateProductListRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProductlistCreateProductListRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProductlistCreateProductListRequest) SetCode(v string) {
	o.Code = &v
}

// GetUrlKey returns the UrlKey field value if set, zero value otherwise.
func (o *ProductlistCreateProductListRequest) GetUrlKey() ProductlistLocalizedText {
	if o == nil || IsNil(o.UrlKey) {
		var ret ProductlistLocalizedText
		return ret
	}
	return *o.UrlKey
}

// GetUrlKeyOk returns a tuple with the UrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistCreateProductListRequest) GetUrlKeyOk() (*ProductlistLocalizedText, bool) {
	if o == nil || IsNil(o.UrlKey) {
		return nil, false
	}
	return o.UrlKey, true
}

// HasUrlKey returns a boolean if a field has been set.
func (o *ProductlistCreateProductListRequest) HasUrlKey() bool {
	if o != nil && !IsNil(o.UrlKey) {
		return true
	}

	return false
}

// SetUrlKey gets a reference to the given ProductlistLocalizedText and assigns it to the UrlKey field.
func (o *ProductlistCreateProductListRequest) SetUrlKey(v ProductlistLocalizedText) {
	o.UrlKey = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *ProductlistCreateProductListRequest) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistCreateProductListRequest) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *ProductlistCreateProductListRequest) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *ProductlistCreateProductListRequest) SetEntityType(v string) {
	o.EntityType = &v
}

// GetEntityCode returns the EntityCode field value if set, zero value otherwise.
func (o *ProductlistCreateProductListRequest) GetEntityCode() string {
	if o == nil || IsNil(o.EntityCode) {
		var ret string
		return ret
	}
	return *o.EntityCode
}

// GetEntityCodeOk returns a tuple with the EntityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistCreateProductListRequest) GetEntityCodeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityCode) {
		return nil, false
	}
	return o.EntityCode, true
}

// HasEntityCode returns a boolean if a field has been set.
func (o *ProductlistCreateProductListRequest) HasEntityCode() bool {
	if o != nil && !IsNil(o.EntityCode) {
		return true
	}

	return false
}

// SetEntityCode gets a reference to the given string and assigns it to the EntityCode field.
func (o *ProductlistCreateProductListRequest) SetEntityCode(v string) {
	o.EntityCode = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ProductlistCreateProductListRequest) GetAttributes() map[string]ProtobufAny {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistCreateProductListRequest) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ProductlistCreateProductListRequest) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *ProductlistCreateProductListRequest) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

func (o ProductlistCreateProductListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistCreateProductListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.UrlKey) {
		toSerialize["urlKey"] = o.UrlKey
	}
	if !IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !IsNil(o.EntityCode) {
		toSerialize["entityCode"] = o.EntityCode
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductlistCreateProductListRequest) UnmarshalJSON(data []byte) (err error) {
	varProductlistCreateProductListRequest := _ProductlistCreateProductListRequest{}

	err = json.Unmarshal(data, &varProductlistCreateProductListRequest)

	if err != nil {
		return err
	}

	*o = ProductlistCreateProductListRequest(varProductlistCreateProductListRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "code")
		delete(additionalProperties, "urlKey")
		delete(additionalProperties, "entityType")
		delete(additionalProperties, "entityCode")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistCreateProductListRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ProductlistCreateProductListRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableProductlistCreateProductListRequest struct {
	value *ProductlistCreateProductListRequest
	isSet bool
}

func (v NullableProductlistCreateProductListRequest) Get() *ProductlistCreateProductListRequest {
	return v.value
}

func (v *NullableProductlistCreateProductListRequest) Set(val *ProductlistCreateProductListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistCreateProductListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistCreateProductListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistCreateProductListRequest(val *ProductlistCreateProductListRequest) *NullableProductlistCreateProductListRequest {
	return &NullableProductlistCreateProductListRequest{value: val, isSet: true}
}

func (v NullableProductlistCreateProductListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistCreateProductListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



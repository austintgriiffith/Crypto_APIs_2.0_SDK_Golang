/*
CryptoAPIs

Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.

API version: 2.0.0
Contact: developers@cryptoapis.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
	"fmt"
)

// ListConfirmedTransactionsByAddressAndTimeRangeRIBS - struct for ListConfirmedTransactionsByAddressAndTimeRangeRIBS
type ListConfirmedTransactionsByAddressAndTimeRangeRIBS struct {
	ListConfirmedTransactionsByAddressAndTimeRangeRIBSB *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB
	ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC
	ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC
	ListConfirmedTransactionsByAddressAndTimeRangeRIBSD *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD
	ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2
	ListConfirmedTransactionsByAddressAndTimeRangeRIBSE *ListConfirmedTransactionsByAddressAndTimeRangeRIBSE
	ListConfirmedTransactionsByAddressAndTimeRangeRIBSEC *ListConfirmedTransactionsByAddressAndTimeRangeRIBSEC
	ListConfirmedTransactionsByAddressAndTimeRangeRIBSL *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL
	ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ
}

// ListConfirmedTransactionsByAddressAndTimeRangeRIBSBAsListConfirmedTransactionsByAddressAndTimeRangeRIBS is a convenience function that returns ListConfirmedTransactionsByAddressAndTimeRangeRIBSB wrapped in ListConfirmedTransactionsByAddressAndTimeRangeRIBS
func ListConfirmedTransactionsByAddressAndTimeRangeRIBSBAsListConfirmedTransactionsByAddressAndTimeRangeRIBS(v *ListConfirmedTransactionsByAddressAndTimeRangeRIBSB) ListConfirmedTransactionsByAddressAndTimeRangeRIBS {
	return ListConfirmedTransactionsByAddressAndTimeRangeRIBS{
		ListConfirmedTransactionsByAddressAndTimeRangeRIBSB: v,
	}
}

// ListConfirmedTransactionsByAddressAndTimeRangeRIBSBCAsListConfirmedTransactionsByAddressAndTimeRangeRIBS is a convenience function that returns ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC wrapped in ListConfirmedTransactionsByAddressAndTimeRangeRIBS
func ListConfirmedTransactionsByAddressAndTimeRangeRIBSBCAsListConfirmedTransactionsByAddressAndTimeRangeRIBS(v *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) ListConfirmedTransactionsByAddressAndTimeRangeRIBS {
	return ListConfirmedTransactionsByAddressAndTimeRangeRIBS{
		ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC: v,
	}
}

// ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSCAsListConfirmedTransactionsByAddressAndTimeRangeRIBS is a convenience function that returns ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC wrapped in ListConfirmedTransactionsByAddressAndTimeRangeRIBS
func ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSCAsListConfirmedTransactionsByAddressAndTimeRangeRIBS(v *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC) ListConfirmedTransactionsByAddressAndTimeRangeRIBS {
	return ListConfirmedTransactionsByAddressAndTimeRangeRIBS{
		ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC: v,
	}
}

// ListConfirmedTransactionsByAddressAndTimeRangeRIBSDAsListConfirmedTransactionsByAddressAndTimeRangeRIBS is a convenience function that returns ListConfirmedTransactionsByAddressAndTimeRangeRIBSD wrapped in ListConfirmedTransactionsByAddressAndTimeRangeRIBS
func ListConfirmedTransactionsByAddressAndTimeRangeRIBSDAsListConfirmedTransactionsByAddressAndTimeRangeRIBS(v *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD) ListConfirmedTransactionsByAddressAndTimeRangeRIBS {
	return ListConfirmedTransactionsByAddressAndTimeRangeRIBS{
		ListConfirmedTransactionsByAddressAndTimeRangeRIBSD: v,
	}
}

// ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2AsListConfirmedTransactionsByAddressAndTimeRangeRIBS is a convenience function that returns ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 wrapped in ListConfirmedTransactionsByAddressAndTimeRangeRIBS
func ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2AsListConfirmedTransactionsByAddressAndTimeRangeRIBS(v *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) ListConfirmedTransactionsByAddressAndTimeRangeRIBS {
	return ListConfirmedTransactionsByAddressAndTimeRangeRIBS{
		ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2: v,
	}
}

// ListConfirmedTransactionsByAddressAndTimeRangeRIBSEAsListConfirmedTransactionsByAddressAndTimeRangeRIBS is a convenience function that returns ListConfirmedTransactionsByAddressAndTimeRangeRIBSE wrapped in ListConfirmedTransactionsByAddressAndTimeRangeRIBS
func ListConfirmedTransactionsByAddressAndTimeRangeRIBSEAsListConfirmedTransactionsByAddressAndTimeRangeRIBS(v *ListConfirmedTransactionsByAddressAndTimeRangeRIBSE) ListConfirmedTransactionsByAddressAndTimeRangeRIBS {
	return ListConfirmedTransactionsByAddressAndTimeRangeRIBS{
		ListConfirmedTransactionsByAddressAndTimeRangeRIBSE: v,
	}
}

// ListConfirmedTransactionsByAddressAndTimeRangeRIBSECAsListConfirmedTransactionsByAddressAndTimeRangeRIBS is a convenience function that returns ListConfirmedTransactionsByAddressAndTimeRangeRIBSEC wrapped in ListConfirmedTransactionsByAddressAndTimeRangeRIBS
func ListConfirmedTransactionsByAddressAndTimeRangeRIBSECAsListConfirmedTransactionsByAddressAndTimeRangeRIBS(v *ListConfirmedTransactionsByAddressAndTimeRangeRIBSEC) ListConfirmedTransactionsByAddressAndTimeRangeRIBS {
	return ListConfirmedTransactionsByAddressAndTimeRangeRIBS{
		ListConfirmedTransactionsByAddressAndTimeRangeRIBSEC: v,
	}
}

// ListConfirmedTransactionsByAddressAndTimeRangeRIBSLAsListConfirmedTransactionsByAddressAndTimeRangeRIBS is a convenience function that returns ListConfirmedTransactionsByAddressAndTimeRangeRIBSL wrapped in ListConfirmedTransactionsByAddressAndTimeRangeRIBS
func ListConfirmedTransactionsByAddressAndTimeRangeRIBSLAsListConfirmedTransactionsByAddressAndTimeRangeRIBS(v *ListConfirmedTransactionsByAddressAndTimeRangeRIBSL) ListConfirmedTransactionsByAddressAndTimeRangeRIBS {
	return ListConfirmedTransactionsByAddressAndTimeRangeRIBS{
		ListConfirmedTransactionsByAddressAndTimeRangeRIBSL: v,
	}
}

// ListConfirmedTransactionsByAddressAndTimeRangeRIBSZAsListConfirmedTransactionsByAddressAndTimeRangeRIBS is a convenience function that returns ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ wrapped in ListConfirmedTransactionsByAddressAndTimeRangeRIBS
func ListConfirmedTransactionsByAddressAndTimeRangeRIBSZAsListConfirmedTransactionsByAddressAndTimeRangeRIBS(v *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) ListConfirmedTransactionsByAddressAndTimeRangeRIBS {
	return ListConfirmedTransactionsByAddressAndTimeRangeRIBS{
		ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListConfirmedTransactionsByAddressAndTimeRangeRIBS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListConfirmedTransactionsByAddressAndTimeRangeRIBSB
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSB)
	if err == nil {
		jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSB, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSB)
		if string(jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSB) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSB = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSB = nil
	}

	// try to unmarshal data into ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC)
	if err == nil {
		jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSBC, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC)
		if string(jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC = nil
	}

	// try to unmarshal data into ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC)
	if err == nil {
		jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC)
		if string(jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC = nil
	}

	// try to unmarshal data into ListConfirmedTransactionsByAddressAndTimeRangeRIBSD
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD)
	if err == nil {
		jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSD, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD)
		if string(jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSD) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD = nil
	}

	// try to unmarshal data into ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2)
	if err == nil {
		jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSD2, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2)
		if string(jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 = nil
	}

	// try to unmarshal data into ListConfirmedTransactionsByAddressAndTimeRangeRIBSE
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSE)
	if err == nil {
		jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSE, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSE)
		if string(jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSE) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSE = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSE = nil
	}

	// try to unmarshal data into ListConfirmedTransactionsByAddressAndTimeRangeRIBSEC
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSEC)
	if err == nil {
		jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSEC, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSEC)
		if string(jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSEC) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSEC = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSEC = nil
	}

	// try to unmarshal data into ListConfirmedTransactionsByAddressAndTimeRangeRIBSL
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSL)
	if err == nil {
		jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSL, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSL)
		if string(jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSL) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSL = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSL = nil
	}

	// try to unmarshal data into ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ
	err = newStrictDecoder(data).Decode(&dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ)
	if err == nil {
		jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSZ, _ := json.Marshal(dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ)
		if string(jsonListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) == "{}" { // empty struct
			dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ = nil
		} else {
			match++
		}
	} else {
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSB = nil
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC = nil
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC = nil
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD = nil
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 = nil
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSE = nil
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSEC = nil
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSL = nil
		dst.ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ListConfirmedTransactionsByAddressAndTimeRangeRIBS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ListConfirmedTransactionsByAddressAndTimeRangeRIBS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListConfirmedTransactionsByAddressAndTimeRangeRIBS) MarshalJSON() ([]byte, error) {
	if src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSB != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSB)
	}

	if src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC)
	}

	if src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC)
	}

	if src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD)
	}

	if src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2)
	}

	if src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSE != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSE)
	}

	if src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSEC != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSEC)
	}

	if src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSL != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSL)
	}

	if src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ != nil {
		return json.Marshal(&src.ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListConfirmedTransactionsByAddressAndTimeRangeRIBS) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSB != nil {
		return obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSB
	}

	if obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC != nil {
		return obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC
	}

	if obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC != nil {
		return obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSBSC
	}

	if obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD != nil {
		return obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD
	}

	if obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 != nil {
		return obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2
	}

	if obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSE != nil {
		return obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSE
	}

	if obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSEC != nil {
		return obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSEC
	}

	if obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSL != nil {
		return obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSL
	}

	if obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ != nil {
		return obj.ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ
	}

	// all schemas are nil
	return nil
}

type NullableListConfirmedTransactionsByAddressAndTimeRangeRIBS struct {
	value *ListConfirmedTransactionsByAddressAndTimeRangeRIBS
	isSet bool
}

func (v NullableListConfirmedTransactionsByAddressAndTimeRangeRIBS) Get() *ListConfirmedTransactionsByAddressAndTimeRangeRIBS {
	return v.value
}

func (v *NullableListConfirmedTransactionsByAddressAndTimeRangeRIBS) Set(val *ListConfirmedTransactionsByAddressAndTimeRangeRIBS) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTransactionsByAddressAndTimeRangeRIBS) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTransactionsByAddressAndTimeRangeRIBS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTransactionsByAddressAndTimeRangeRIBS(val *ListConfirmedTransactionsByAddressAndTimeRangeRIBS) *NullableListConfirmedTransactionsByAddressAndTimeRangeRIBS {
	return &NullableListConfirmedTransactionsByAddressAndTimeRangeRIBS{value: val, isSet: true}
}

func (v NullableListConfirmedTransactionsByAddressAndTimeRangeRIBS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTransactionsByAddressAndTimeRangeRIBS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


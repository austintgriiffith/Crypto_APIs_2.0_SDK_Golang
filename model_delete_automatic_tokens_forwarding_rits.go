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

// DeleteAutomaticTokensForwardingRITS - struct for DeleteAutomaticTokensForwardingRITS
type DeleteAutomaticTokensForwardingRITS struct {
	DeleteAutomaticTokensForwardingRITSBOT *DeleteAutomaticTokensForwardingRITSBOT
	DeleteAutomaticTokensForwardingRITSET *DeleteAutomaticTokensForwardingRITSET
}

// DeleteAutomaticTokensForwardingRITSBOTAsDeleteAutomaticTokensForwardingRITS is a convenience function that returns DeleteAutomaticTokensForwardingRITSBOT wrapped in DeleteAutomaticTokensForwardingRITS
func DeleteAutomaticTokensForwardingRITSBOTAsDeleteAutomaticTokensForwardingRITS(v *DeleteAutomaticTokensForwardingRITSBOT) DeleteAutomaticTokensForwardingRITS {
	return DeleteAutomaticTokensForwardingRITS{ DeleteAutomaticTokensForwardingRITSBOT: v}
}

// DeleteAutomaticTokensForwardingRITSETAsDeleteAutomaticTokensForwardingRITS is a convenience function that returns DeleteAutomaticTokensForwardingRITSET wrapped in DeleteAutomaticTokensForwardingRITS
func DeleteAutomaticTokensForwardingRITSETAsDeleteAutomaticTokensForwardingRITS(v *DeleteAutomaticTokensForwardingRITSET) DeleteAutomaticTokensForwardingRITS {
	return DeleteAutomaticTokensForwardingRITS{ DeleteAutomaticTokensForwardingRITSET: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DeleteAutomaticTokensForwardingRITS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DeleteAutomaticTokensForwardingRITSBOT
	err = json.Unmarshal(data, &dst.DeleteAutomaticTokensForwardingRITSBOT)
	if err == nil {
		jsonDeleteAutomaticTokensForwardingRITSBOT, _ := json.Marshal(dst.DeleteAutomaticTokensForwardingRITSBOT)
		if string(jsonDeleteAutomaticTokensForwardingRITSBOT) == "{}" { // empty struct
			dst.DeleteAutomaticTokensForwardingRITSBOT = nil
		} else {
			match++
		}
	} else {
		dst.DeleteAutomaticTokensForwardingRITSBOT = nil
	}

	// try to unmarshal data into DeleteAutomaticTokensForwardingRITSET
	err = json.Unmarshal(data, &dst.DeleteAutomaticTokensForwardingRITSET)
	if err == nil {
		jsonDeleteAutomaticTokensForwardingRITSET, _ := json.Marshal(dst.DeleteAutomaticTokensForwardingRITSET)
		if string(jsonDeleteAutomaticTokensForwardingRITSET) == "{}" { // empty struct
			dst.DeleteAutomaticTokensForwardingRITSET = nil
		} else {
			match++
		}
	} else {
		dst.DeleteAutomaticTokensForwardingRITSET = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DeleteAutomaticTokensForwardingRITSBOT = nil
		dst.DeleteAutomaticTokensForwardingRITSET = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(DeleteAutomaticTokensForwardingRITS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(DeleteAutomaticTokensForwardingRITS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DeleteAutomaticTokensForwardingRITS) MarshalJSON() ([]byte, error) {
	if src.DeleteAutomaticTokensForwardingRITSBOT != nil {
		return json.Marshal(&src.DeleteAutomaticTokensForwardingRITSBOT)
	}

	if src.DeleteAutomaticTokensForwardingRITSET != nil {
		return json.Marshal(&src.DeleteAutomaticTokensForwardingRITSET)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DeleteAutomaticTokensForwardingRITS) GetActualInstance() (interface{}) {
	if obj.DeleteAutomaticTokensForwardingRITSBOT != nil {
		return obj.DeleteAutomaticTokensForwardingRITSBOT
	}

	if obj.DeleteAutomaticTokensForwardingRITSET != nil {
		return obj.DeleteAutomaticTokensForwardingRITSET
	}

	// all schemas are nil
	return nil
}

type NullableDeleteAutomaticTokensForwardingRITS struct {
	value *DeleteAutomaticTokensForwardingRITS
	isSet bool
}

func (v NullableDeleteAutomaticTokensForwardingRITS) Get() *DeleteAutomaticTokensForwardingRITS {
	return v.value
}

func (v *NullableDeleteAutomaticTokensForwardingRITS) Set(val *DeleteAutomaticTokensForwardingRITS) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAutomaticTokensForwardingRITS) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAutomaticTokensForwardingRITS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAutomaticTokensForwardingRITS(val *DeleteAutomaticTokensForwardingRITS) *NullableDeleteAutomaticTokensForwardingRITS {
	return &NullableDeleteAutomaticTokensForwardingRITS{value: val, isSet: true}
}

func (v NullableDeleteAutomaticTokensForwardingRITS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAutomaticTokensForwardingRITS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



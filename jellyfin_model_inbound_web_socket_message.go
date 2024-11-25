/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// JellyfinInboundWebSocketMessage - Represents the list of possible inbound websocket types
type JellyfinInboundWebSocketMessage struct {
	JellyfinActivityLogEntryStartMessage *JellyfinActivityLogEntryStartMessage
	JellyfinActivityLogEntryStopMessage *JellyfinActivityLogEntryStopMessage
	JellyfinInboundKeepAliveMessage *JellyfinInboundKeepAliveMessage
	JellyfinScheduledTasksInfoStartMessage *JellyfinScheduledTasksInfoStartMessage
	JellyfinScheduledTasksInfoStopMessage *JellyfinScheduledTasksInfoStopMessage
	JellyfinSessionsStartMessage *JellyfinSessionsStartMessage
	JellyfinSessionsStopMessage *JellyfinSessionsStopMessage
}

// JellyfinActivityLogEntryStartMessageAsJellyfinInboundWebSocketMessage is a convenience function that returns JellyfinActivityLogEntryStartMessage wrapped in JellyfinInboundWebSocketMessage
func JellyfinActivityLogEntryStartMessageAsJellyfinInboundWebSocketMessage(v *JellyfinActivityLogEntryStartMessage) JellyfinInboundWebSocketMessage {
	return JellyfinInboundWebSocketMessage{
		JellyfinActivityLogEntryStartMessage: v,
	}
}

// JellyfinActivityLogEntryStopMessageAsJellyfinInboundWebSocketMessage is a convenience function that returns JellyfinActivityLogEntryStopMessage wrapped in JellyfinInboundWebSocketMessage
func JellyfinActivityLogEntryStopMessageAsJellyfinInboundWebSocketMessage(v *JellyfinActivityLogEntryStopMessage) JellyfinInboundWebSocketMessage {
	return JellyfinInboundWebSocketMessage{
		JellyfinActivityLogEntryStopMessage: v,
	}
}

// JellyfinInboundKeepAliveMessageAsJellyfinInboundWebSocketMessage is a convenience function that returns JellyfinInboundKeepAliveMessage wrapped in JellyfinInboundWebSocketMessage
func JellyfinInboundKeepAliveMessageAsJellyfinInboundWebSocketMessage(v *JellyfinInboundKeepAliveMessage) JellyfinInboundWebSocketMessage {
	return JellyfinInboundWebSocketMessage{
		JellyfinInboundKeepAliveMessage: v,
	}
}

// JellyfinScheduledTasksInfoStartMessageAsJellyfinInboundWebSocketMessage is a convenience function that returns JellyfinScheduledTasksInfoStartMessage wrapped in JellyfinInboundWebSocketMessage
func JellyfinScheduledTasksInfoStartMessageAsJellyfinInboundWebSocketMessage(v *JellyfinScheduledTasksInfoStartMessage) JellyfinInboundWebSocketMessage {
	return JellyfinInboundWebSocketMessage{
		JellyfinScheduledTasksInfoStartMessage: v,
	}
}

// JellyfinScheduledTasksInfoStopMessageAsJellyfinInboundWebSocketMessage is a convenience function that returns JellyfinScheduledTasksInfoStopMessage wrapped in JellyfinInboundWebSocketMessage
func JellyfinScheduledTasksInfoStopMessageAsJellyfinInboundWebSocketMessage(v *JellyfinScheduledTasksInfoStopMessage) JellyfinInboundWebSocketMessage {
	return JellyfinInboundWebSocketMessage{
		JellyfinScheduledTasksInfoStopMessage: v,
	}
}

// JellyfinSessionsStartMessageAsJellyfinInboundWebSocketMessage is a convenience function that returns JellyfinSessionsStartMessage wrapped in JellyfinInboundWebSocketMessage
func JellyfinSessionsStartMessageAsJellyfinInboundWebSocketMessage(v *JellyfinSessionsStartMessage) JellyfinInboundWebSocketMessage {
	return JellyfinInboundWebSocketMessage{
		JellyfinSessionsStartMessage: v,
	}
}

// JellyfinSessionsStopMessageAsJellyfinInboundWebSocketMessage is a convenience function that returns JellyfinSessionsStopMessage wrapped in JellyfinInboundWebSocketMessage
func JellyfinSessionsStopMessageAsJellyfinInboundWebSocketMessage(v *JellyfinSessionsStopMessage) JellyfinInboundWebSocketMessage {
	return JellyfinInboundWebSocketMessage{
		JellyfinSessionsStopMessage: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *JellyfinInboundWebSocketMessage) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into JellyfinActivityLogEntryStartMessage
	err = newStrictDecoder(data).Decode(&dst.JellyfinActivityLogEntryStartMessage)
	if err == nil {
		jsonJellyfinActivityLogEntryStartMessage, _ := json.Marshal(dst.JellyfinActivityLogEntryStartMessage)
		if string(jsonJellyfinActivityLogEntryStartMessage) == "{}" { // empty struct
			dst.JellyfinActivityLogEntryStartMessage = nil
		} else {
			if err = validator.Validate(dst.JellyfinActivityLogEntryStartMessage); err != nil {
				dst.JellyfinActivityLogEntryStartMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.JellyfinActivityLogEntryStartMessage = nil
	}

	// try to unmarshal data into JellyfinActivityLogEntryStopMessage
	err = newStrictDecoder(data).Decode(&dst.JellyfinActivityLogEntryStopMessage)
	if err == nil {
		jsonJellyfinActivityLogEntryStopMessage, _ := json.Marshal(dst.JellyfinActivityLogEntryStopMessage)
		if string(jsonJellyfinActivityLogEntryStopMessage) == "{}" { // empty struct
			dst.JellyfinActivityLogEntryStopMessage = nil
		} else {
			if err = validator.Validate(dst.JellyfinActivityLogEntryStopMessage); err != nil {
				dst.JellyfinActivityLogEntryStopMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.JellyfinActivityLogEntryStopMessage = nil
	}

	// try to unmarshal data into JellyfinInboundKeepAliveMessage
	err = newStrictDecoder(data).Decode(&dst.JellyfinInboundKeepAliveMessage)
	if err == nil {
		jsonJellyfinInboundKeepAliveMessage, _ := json.Marshal(dst.JellyfinInboundKeepAliveMessage)
		if string(jsonJellyfinInboundKeepAliveMessage) == "{}" { // empty struct
			dst.JellyfinInboundKeepAliveMessage = nil
		} else {
			if err = validator.Validate(dst.JellyfinInboundKeepAliveMessage); err != nil {
				dst.JellyfinInboundKeepAliveMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.JellyfinInboundKeepAliveMessage = nil
	}

	// try to unmarshal data into JellyfinScheduledTasksInfoStartMessage
	err = newStrictDecoder(data).Decode(&dst.JellyfinScheduledTasksInfoStartMessage)
	if err == nil {
		jsonJellyfinScheduledTasksInfoStartMessage, _ := json.Marshal(dst.JellyfinScheduledTasksInfoStartMessage)
		if string(jsonJellyfinScheduledTasksInfoStartMessage) == "{}" { // empty struct
			dst.JellyfinScheduledTasksInfoStartMessage = nil
		} else {
			if err = validator.Validate(dst.JellyfinScheduledTasksInfoStartMessage); err != nil {
				dst.JellyfinScheduledTasksInfoStartMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.JellyfinScheduledTasksInfoStartMessage = nil
	}

	// try to unmarshal data into JellyfinScheduledTasksInfoStopMessage
	err = newStrictDecoder(data).Decode(&dst.JellyfinScheduledTasksInfoStopMessage)
	if err == nil {
		jsonJellyfinScheduledTasksInfoStopMessage, _ := json.Marshal(dst.JellyfinScheduledTasksInfoStopMessage)
		if string(jsonJellyfinScheduledTasksInfoStopMessage) == "{}" { // empty struct
			dst.JellyfinScheduledTasksInfoStopMessage = nil
		} else {
			if err = validator.Validate(dst.JellyfinScheduledTasksInfoStopMessage); err != nil {
				dst.JellyfinScheduledTasksInfoStopMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.JellyfinScheduledTasksInfoStopMessage = nil
	}

	// try to unmarshal data into JellyfinSessionsStartMessage
	err = newStrictDecoder(data).Decode(&dst.JellyfinSessionsStartMessage)
	if err == nil {
		jsonJellyfinSessionsStartMessage, _ := json.Marshal(dst.JellyfinSessionsStartMessage)
		if string(jsonJellyfinSessionsStartMessage) == "{}" { // empty struct
			dst.JellyfinSessionsStartMessage = nil
		} else {
			if err = validator.Validate(dst.JellyfinSessionsStartMessage); err != nil {
				dst.JellyfinSessionsStartMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.JellyfinSessionsStartMessage = nil
	}

	// try to unmarshal data into JellyfinSessionsStopMessage
	err = newStrictDecoder(data).Decode(&dst.JellyfinSessionsStopMessage)
	if err == nil {
		jsonJellyfinSessionsStopMessage, _ := json.Marshal(dst.JellyfinSessionsStopMessage)
		if string(jsonJellyfinSessionsStopMessage) == "{}" { // empty struct
			dst.JellyfinSessionsStopMessage = nil
		} else {
			if err = validator.Validate(dst.JellyfinSessionsStopMessage); err != nil {
				dst.JellyfinSessionsStopMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.JellyfinSessionsStopMessage = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.JellyfinActivityLogEntryStartMessage = nil
		dst.JellyfinActivityLogEntryStopMessage = nil
		dst.JellyfinInboundKeepAliveMessage = nil
		dst.JellyfinScheduledTasksInfoStartMessage = nil
		dst.JellyfinScheduledTasksInfoStopMessage = nil
		dst.JellyfinSessionsStartMessage = nil
		dst.JellyfinSessionsStopMessage = nil

		return fmt.Errorf("data matches more than one schema in oneOf(JellyfinInboundWebSocketMessage)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(JellyfinInboundWebSocketMessage)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src JellyfinInboundWebSocketMessage) MarshalJSON() ([]byte, error) {
	if src.JellyfinActivityLogEntryStartMessage != nil {
		return json.Marshal(&src.JellyfinActivityLogEntryStartMessage)
	}

	if src.JellyfinActivityLogEntryStopMessage != nil {
		return json.Marshal(&src.JellyfinActivityLogEntryStopMessage)
	}

	if src.JellyfinInboundKeepAliveMessage != nil {
		return json.Marshal(&src.JellyfinInboundKeepAliveMessage)
	}

	if src.JellyfinScheduledTasksInfoStartMessage != nil {
		return json.Marshal(&src.JellyfinScheduledTasksInfoStartMessage)
	}

	if src.JellyfinScheduledTasksInfoStopMessage != nil {
		return json.Marshal(&src.JellyfinScheduledTasksInfoStopMessage)
	}

	if src.JellyfinSessionsStartMessage != nil {
		return json.Marshal(&src.JellyfinSessionsStartMessage)
	}

	if src.JellyfinSessionsStopMessage != nil {
		return json.Marshal(&src.JellyfinSessionsStopMessage)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *JellyfinInboundWebSocketMessage) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.JellyfinActivityLogEntryStartMessage != nil {
		return obj.JellyfinActivityLogEntryStartMessage
	}

	if obj.JellyfinActivityLogEntryStopMessage != nil {
		return obj.JellyfinActivityLogEntryStopMessage
	}

	if obj.JellyfinInboundKeepAliveMessage != nil {
		return obj.JellyfinInboundKeepAliveMessage
	}

	if obj.JellyfinScheduledTasksInfoStartMessage != nil {
		return obj.JellyfinScheduledTasksInfoStartMessage
	}

	if obj.JellyfinScheduledTasksInfoStopMessage != nil {
		return obj.JellyfinScheduledTasksInfoStopMessage
	}

	if obj.JellyfinSessionsStartMessage != nil {
		return obj.JellyfinSessionsStartMessage
	}

	if obj.JellyfinSessionsStopMessage != nil {
		return obj.JellyfinSessionsStopMessage
	}

	// all schemas are nil
	return nil
}

type NullableJellyfinInboundWebSocketMessage struct {
	value *JellyfinInboundWebSocketMessage
	isSet bool
}

func (v NullableJellyfinInboundWebSocketMessage) Get() *JellyfinInboundWebSocketMessage {
	return v.value
}

func (v *NullableJellyfinInboundWebSocketMessage) Set(val *JellyfinInboundWebSocketMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinInboundWebSocketMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinInboundWebSocketMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinInboundWebSocketMessage(val *JellyfinInboundWebSocketMessage) *NullableJellyfinInboundWebSocketMessage {
	return &NullableJellyfinInboundWebSocketMessage{value: val, isSet: true}
}

func (v NullableJellyfinInboundWebSocketMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinInboundWebSocketMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



package protogeo

import (
	"encoding/json"
	"fmt"
)

// EncodeIntID ...
func EncodeIntID(id interface{}) (*Data_Feature_IntId, error) {
	switch t := id.(type) {
	case int:
		return encodeIntID(int64(t)), nil
	case int8:
		return encodeIntID(int64(t)), nil
	case int16:
		return encodeIntID(int64(t)), nil
	case int32:
		return encodeIntID(int64(t)), nil
	case int64:
		return encodeIntID(t), nil
	case uint8:
		return encodeIntID(int64(t)), nil
	case uint16:
		return encodeIntID(int64(t)), nil
	case uint32:
		return encodeIntID(int64(t)), nil
	case uint64:
		return encodeIntID(int64(t)), nil
	default:
		return nil, fmt.Errorf("Value type is not an int")
	}
}

// EncodeID ...
func EncodeID(id interface{}) (*Data_Feature_Id, error) {
	switch t := id.(type) {
	case string:
		return encodeID(t), nil
	case *string:
		return encodeID(*t), nil
	default:
		val, err := json.Marshal(id)
		if err != nil {
			return nil, err
		}
		return encodeID(string(val)), nil
	}
}

func encodeIntID(id int64) *Data_Feature_IntId {
	return &Data_Feature_IntId{
		IntId: id,
	}
}

func encodeID(id string) *Data_Feature_Id {
	return &Data_Feature_Id{
		Id: id,
	}
}

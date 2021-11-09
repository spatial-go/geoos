package utils

import (
	"encoding/json"
	"fmt"

	"github.com/spatial-go/geoos/geobuf/proto"
)

func EncodeIntId(id interface{}) (*proto.Data_Feature_IntId, error) {
	switch t := id.(type) {
	case int:
		return encodeIntId(int64(t)), nil
	case int8:
		return encodeIntId(int64(t)), nil
	case int16:
		return encodeIntId(int64(t)), nil
	case int32:
		return encodeIntId(int64(t)), nil
	case int64:
		return encodeIntId(t), nil
	case uint8:
		return encodeIntId(int64(t)), nil
	case uint16:
		return encodeIntId(int64(t)), nil
	case uint32:
		return encodeIntId(int64(t)), nil
	case uint64:
		return encodeIntId(int64(t)), nil
	default:
		return nil, fmt.Errorf("Value type is not an int")
	}
}

func EncodeId(id interface{}) (*proto.Data_Feature_Id, error) {
	switch t := id.(type) {
	case string:
		return encodeId(t), nil
	case *string:
		return encodeId(*t), nil
	default:
		val, err := json.Marshal(id)
		if err != nil {
			return nil, err
		}
		return encodeId(string(val)), nil
	}
}

func encodeIntId(id int64) *proto.Data_Feature_IntId {
	return &proto.Data_Feature_IntId{
		IntId: id,
	}
}

func encodeId(id string) *proto.Data_Feature_Id {
	return &proto.Data_Feature_Id{
		Id: id,
	}
}

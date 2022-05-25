package protogeo

import (
	"encoding/json"
	"reflect"
)

// EncodeValue ...
func EncodeValue(val interface{}) (*Data_Value, error) {
	v := reflect.ValueOf(val)
	return encodeValue(v, val)
}

func encodeValue(v reflect.Value, val interface{}) (*Data_Value, error) {
	switch v.Kind() {
	case reflect.Bool:
		return encodeBool(v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intval := v.Int()
		if intval < 0 {
			return encodeInt(uint64(v.Int()*-1), false)
		}
		return encodeInt(uint64(v.Int()), true)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return encodeInt(uint64(v.Uint()), true)
	case reflect.Float32, reflect.Float64:
		return encodeDouble(v.Float())
	case reflect.String:
		return encodeString(v.String())
	case reflect.Ptr:
		return encodeValue(v.Elem(), val)
	default:
		return encodeJSON(v.Interface())
	}
}

func encodeInt(val uint64, positive bool) (*Data_Value, error) {
	if positive {
		return &Data_Value{
			ValueType: &Data_Value_PosIntValue{
				PosIntValue: val,
			},
		}, nil
	}

	return &Data_Value{
		ValueType: &Data_Value_NegIntValue{
			NegIntValue: val,
		},
	}, nil
}

func encodeDouble(val float64) (*Data_Value, error) {
	return &Data_Value{
		ValueType: &Data_Value_DoubleValue{
			DoubleValue: val,
		},
	}, nil
}

func encodeString(val string) (*Data_Value, error) {
	return &Data_Value{
		ValueType: &Data_Value_StringValue{
			StringValue: val,
		},
	}, nil
}

func encodeBool(val bool) (*Data_Value, error) {
	return &Data_Value{
		ValueType: &Data_Value_BoolValue{
			BoolValue: val,
		},
	}, nil
}

func encodeJSON(val interface{}) (*Data_Value, error) {
	encoded, err := json.Marshal(val)
	return &Data_Value{
		ValueType: &Data_Value_JsonValue{
			JsonValue: string(encoded),
		},
	}, err
}

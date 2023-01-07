package payuni

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func StructToParamsMap(data interface{}) map[string]string {
	params := map[string]string{}
	iVal := reflect.ValueOf(data)
	iTyp := reflect.TypeOf(data)
	if iVal.Kind() == reflect.Ptr {
		iVal = iVal.Elem()
	}
	if iTyp.Kind() == reflect.Ptr {
		iTyp = iTyp.Elem()
	}

	//stringerInterface := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	marshalerInterface := reflect.TypeOf((*json.Marshaler)(nil)).Elem()
	//typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {

		f := iVal.Field(i)
		ft := iTyp.Field(i)

		if f.Kind() == reflect.Ptr && f.IsNil() {
			continue
		}
		if f.Kind() == reflect.Ptr {
			f = f.Elem()

		}

		if ft.Anonymous {
			nestedParams := StructToParamsMap(f.Interface())
			for key, val := range nestedParams {
				params[key] = val
			}
			continue
		}

		var v string
		//f.Anonymous
		switch realVal := f.Interface().(type) {
		case int, int8, int16, int32, int64:
			tag := ft.Tag.Get("json")
			if strings.Contains(tag, ",omitempty") && f.Int() == 0 {
				continue
			}
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			tag := ft.Tag.Get("json")
			if strings.Contains(tag, ",omitempty") && f.Uint() == 0 {
				continue
			}
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			tag := ft.Tag.Get("json")
			if strings.Contains(tag, ",omitempty") && f.Float() == 0 {
				continue
			}
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			tag := ft.Tag.Get("json")
			if strings.Contains(tag, ",omitempty") && f.Float() == 0 {
				continue
			}
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			tag := ft.Tag.Get("json")
			if strings.Contains(tag, ",omitempty") && string(realVal) == "" {
				continue
			}
			v = string(realVal)
		case string:
			tag := ft.Tag.Get("json")
			if strings.Contains(tag, ",omitempty") && string(realVal) == "" {
				continue
			}
			v = realVal
		//case time.Time:
		//	v = realVal.Format(time.RFC3339)
		//case base.ECPayDateTime:
		//	v = realVal.String()
		default:
			switch {
			case f.Type().Implements(marshalerInterface):
				data, err := f.Interface().(json.Marshaler).MarshalJSON()
				if err != nil {
					panic(err)
				}
				unquoteData, err := strconv.Unquote(string(data))
				if err != nil {
					v = string(data)
				}
				v = unquoteData
			//case f.Type().Implements(stringerInterface):
			//	v = f.Interface().(fmt.Stringer).String()
			default:
				switch f.Kind() {
				case reflect.String:
					v = f.String()
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					v = strconv.FormatInt(f.Int(), 10)
				default:
					panic(fmt.Sprintf("Unknown type %T during transfer struct to map!\n", f.Interface()))

				}
			}
		}
		params[ft.Name] = v
	}
	return params
}

func ParamsMapToURLEncode(data map[string]string) string {
	var uri url.URL
	q := uri.Query()
	for key, value := range data {
		q.Set(key, value)
	}
	queryStr := q.Encode()
	fmt.Println(queryStr)
	return queryStr
}

type LowerStringSlice []string

func (p LowerStringSlice) Len() int           { return len(p) }
func (p LowerStringSlice) Less(i, j int) bool { return strings.ToLower(p[i]) < strings.ToLower(p[j]) }
func (p LowerStringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type NewebPayValues struct {
	url.Values
}

func NewValuesFromMap(values map[string]string) *NewebPayValues {
	v := NewebPayValues{Values: url.Values{}}
	for key, value := range values {

		v.Set(key, value)
	}
	return &v
}

func (v NewebPayValues) ToMap() map[string]string {
	result := make(map[string]string)
	for key, val := range v.Values {
		result[key] = val[0]
	}
	return result
}

func (v NewebPayValues) Encode() string {
	if v.Values == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v.Values))
	for k := range v.Values {
		keys = append(keys, k)
	}
	sort.Sort(LowerStringSlice(keys))
	for _, k := range keys {
		vs := v.Values[k]
		//keyEscaped := url.QueryEscape(k)
		keyEscaped := k
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	return buf.String()
}

func StrToMap(in string) map[string]string {
	res := make(map[string]string)
	array := strings.Split(in, ",")
	temp := make([]string, 2)
	for _, val := range array {
		temp = strings.Split(string(val), ":")
		res[temp[0]] = temp[1]
	}
	return res
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

package helper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

func ArrayIntToString(a []int, delimeter string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delimeter, -1), "[]")
}

func ToInt64(t interface{}) int64 {
	switch t := t.(type) { // This is a type switch.
	case int64:
		return t
	case int32:
		return int64(t)
	case int:
		return int64(t)
	case float32:
		return int64(t)
	case float64:
		return int64(t)
	case string:
		res, _ := strconv.ParseInt(t, 10, 64)
		return res
	default:
		return 0
	}
}

func ToInt(t interface{}) int {
	return int(ToInt64(t))
}

func ToInt32(t interface{}) int32 {
	return int32(ToInt64(t))
}

func ToFloat64(str string) float64 {
	res, _ := strconv.ParseFloat(str, 64)
	return res
}

func Serialize(msg interface{}) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(msg)
	return b.Bytes(), err
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// CheckDeadline is check if context has cancelled
func CheckDeadline(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func NonZeroCols(m any, nonZeroVal bool) []string {
	maps := StructToMap(m, nonZeroVal)

	keys := make([]string, 0)
	for k := range maps {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

func StructToMap(m any, nonZeroVal bool) map[string]any {
	v := reflect.ValueOf(m)

	// if pointer get the underlying element≤
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("not struct")
	}

	t := v.Type()
	out := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := field.Name
		val := v.FieldByName(name)
		// we can't access the value of unexported fields
		if field.PkgPath != "" {
			continue
		}

		if nonZeroVal {
			zero := reflect.Zero(val.Type()).Interface()
			current := val.Interface()

			if reflect.DeepEqual(current, zero) {
				continue
			}
		}

		out[field.Name] = val.Interface()
	}

	return out
}

func NowStrUTC() string {
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		time.Now().UTC().Year(), time.Now().UTC().Month(), time.Now().UTC().Day(),
		time.Now().UTC().Hour(), time.Now().UTC().Minute(), time.Now().UTC().Second())
}

func InArray(val interface{}, array interface{}) (found bool) {
	values := reflect.ValueOf(array)

	if reflect.TypeOf(array).Kind() == reflect.Slice || values.Len() > 0 {
		for i := 0; i < values.Len(); i++ {
			if reflect.DeepEqual(val, values.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}

func Dump(array any) {
	s, _ := json.MarshalIndent(array, "", "\t")

	fmt.Println("-------------")
	fmt.Println(string(s))
	fmt.Println("-------------")
}

func String(v string) *string { return &v }

func RemoveFirstChar(input string) string {
	if len(input) <= 1 {
		return ""
	}
	return input[1:]
}

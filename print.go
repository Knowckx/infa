package infa

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/pelletier/go-toml/v2"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

func Printf(format string, a ...interface{}) (n int, err error) {
	format = fmt.Sprintf("%s\n", format)
	return fmt.Printf(format, a...)
}

func PrintJson(in interface{}) {
	res, err := json.Marshal(in)
	if err != nil {
		log.Error().Stack().Err(err).Send()
		return
	}
	Printf(string(res))
}

func PrintYaml(in interface{}) {
	res, err := yaml.Marshal(in)
	if err != nil {
		log.Error().Stack().Err(err).Send()
		return
	}
	Printf(string(res))
}

func PrintToml(in interface{}) {
	res, err := toml.Marshal(in)
	if err != nil {
		log.Error().Stack().Err(err).Send()
		return
	}
	Printf(string(res))
}

func Print(x interface{}) {
	Printf("Input Type:%T", x)
	displayPath("value", reflect.ValueOf(x))
}

func displayPath(path string, v reflect.Value) {
	inter := v.Interface()
	switch inter.(type) {
	case time.Time:
		Printf("%s = %s", path, inter.(time.Time).String())
		return
	}

	Printf(v.Kind().String())

	switch v.Kind() {
	case reflect.Invalid:
		Printf("%s = invalid", path)
	case reflect.Slice, reflect.Array:
		leng := v.Len()
		if leng > 5 {
			Printf("%s Slice.len = %d only show top 5!", path, leng)
			leng = 5
		}
		for i := 0; i < leng; i++ {
			displayPath(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			displayPath(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			fieldPath := fmt.Sprintf("%s[%+v]", path, key)
			displayPath(fieldPath, v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fieldPath := fmt.Sprintf("*%s", path)
			displayPath(fieldPath, v.Elem())
		}
	default:
		Printf("%s = %s", path, formatAtom(v))
	}
}

// base type
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func:
		return v.Type().String() + "0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + "value ????"
	}
}

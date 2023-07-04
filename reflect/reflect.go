package reflect

import (
	"reflect"

	"github.com/jinzhu/gorm"
)

type Struct struct {
	Name   string
	DbName string
	Type   reflect.Kind
	Tag    reflect.StructTag
	Fields []Struct
}

func Parse(v reflect.Type) *Struct {
	if v == nil {
		return nil
	}
	for v.Kind() == reflect.Slice || v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	str := Struct{Type: v.Kind()}
	if nm := v.Name(); nm != "" {
		str.Name = nm
		str.DbName = gorm.TheNamingStrategy.TableName(nm)
	}
	if str.Type != reflect.Struct {
		return &str
	}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		st := Struct{
			Type:   f.Type.Kind(),
			Name:   f.Name,
			DbName: gorm.TheNamingStrategy.TableName(f.Name),
			Tag:    f.Tag,
		}
		if f.Type.Kind() == reflect.Struct {
			if s := Parse(f.Type); s != nil {
				st.Fields = s.Fields
			}
		}
		str.Fields = append(str.Fields, st)
	}
	return &str
}

func ToInterface(v reflect.Value) interface{} {
	switch v.Kind() {
	case reflect.Bool:
		return v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint()
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Complex64, reflect.Complex128:
		return v.Complex()
	case reflect.String:
		return v.String()
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct, reflect.Interface:
		return v.Interface()
	default:
		return nil
	}
}

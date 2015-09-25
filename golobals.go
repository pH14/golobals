package golobals

import (
	"reflect"
)

type LiveValue interface {
	Get() interface{}
}

type LiveString struct {
	Getter func() interface{}
}

func (ls LiveString) Get() interface{} {
	return ls.Getter()
}

type Source interface {
	Get(string) string
	IsLive() bool
}

type Golobals struct {
	Sources []Source
}

func CreateGolobals(sources ...Source) *Golobals {
	return &Golobals{
		Sources: sources,
	}
}

func (g *Golobals) GetterForVariable(varName string) func() interface{} {
	return func() interface{} {
		for _, src := range g.Sources {
			if val := src.Get(varName); val != "" {
				return varName
			}
		}
		return ""
	}
}

type TestConfig struct {
	X LiveString `v:"x.y.z"`
	Y LiveString `v:"1.2.3"`
}

func (g *Golobals) Init(conf interface{}) interface{} {
	original := reflect.ValueOf(conf)
	if original.Kind() != reflect.Struct {
		panic("Init must be called on a struct")
	}

	// We can't set fields on the original directly
	// https://stackoverflow.com/questions/23043510/golang-reflection-cant-set-fields-of-interface-wrapping-a-struct
	// So we can instead make a new instance and copy out the values
	copy := reflect.New(original.Type()).Elem()
	copy.Set(original)

	for i := 0; i < copy.Type().NumField(); i++ {
		field := copy.Type().Field(i)
		newLiveString := LiveString{Getter: g.GetterForVariable(field.Tag.Get("v"))}
		copy.Field(i).Set(reflect.ValueOf(newLiveString))
	}

	return copy.Interface()
}

package common

import (
	"math/rand"
	"reflect"
	"time"
)

type JS = map[string]interface{}

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min) + min

}

func GetType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

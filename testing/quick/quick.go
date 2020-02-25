// This file has automatically been generated on Wed Feb 26 02:10:15 +05 2020.
// DO NOT EDIT.
package quick

import (
	"math/rand"
	"reflect"
	"testing/quick"
	_ "unsafe"
)

//go:linkname Value testing/quick.Value
//go:noescape
func Value(t reflect.Type, rand *rand.Rand) (reflect.Value, bool)

//go:linkname checkequalerrorerror quick.sub_checkequalerrorerror
func checkequalerrorerror(s *quick.CheckEqualError) string {
	return s.Error()
}

//go:linkname CheckEqualErrorError quick.sub_checkequalerrorerror
//go:noescape
func CheckEqualErrorError(s *quick.CheckEqualError) string

//go:linkname checkerrorerror quick.sub_checkerrorerror
func checkerrorerror(s *quick.CheckError) string {
	return s.Error()
}

//go:linkname CheckErrorError quick.sub_checkerrorerror
//go:noescape
func CheckErrorError(s *quick.CheckError) string

//go:linkname setuperrorerror quick.sub_setuperrorerror
func setuperrorerror(s quick.SetupError) string {
	return s.Error()
}

//go:linkname SetupErrorError quick.sub_setuperrorerror
//go:noescape
func SetupErrorError(s quick.SetupError) string

//go:linkname Check testing/quick.Check
//go:noescape
func Check(f interface{}, config *quick.Config) error

//go:linkname CheckEqual testing/quick.CheckEqual
//go:noescape
func CheckEqual(f, g interface{}, config *quick.Config) error

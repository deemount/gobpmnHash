package gobpmn_hash

import (
	"crypto/rand"
	"fmt"
	"hash/fnv"
	"reflect"
	"strings"

	"github.com/deemount/gobpmnHash/internals/utils"
	"github.com/deemount/gobpmnModels/pkg/core"
	gobpmn_reflection "github.com/deemount/gobpmnReflection"
)

type (
	// Def ...
	Def core.DefinitionsRepository

	// Reflection ...
	Reflection struct {
		Suffix string
	}
)

// Hash ...
func (r *Reflection) Hash() string {
	if r.Suffix == "" {
		result, _ := r.hash()
		r.Suffix = result.Suffix
	}
	return r.Suffix
}

// Inject itself reflects a given struct and inject
// signed fields with hash values.
// There are two conditions to assign fields of a struct:
// a) The struct has anonymous fields
// b) The struct has no anymous fields
// It also counts the element in their specification to know
// how much elements of each package needs to be mapped later then.
func (r *Reflection) Inject(p interface{}) interface{} {

	ref := gobpmn_reflection.NewReflect(p)
	ref.Interface().Allocate().Maps().Call()

	switch true {
	// anonymous fields are reflected
	case len(ref.Anonym) > 0:
		length := len(ref.Anonym)

		// create anonymMap and hashMap
		anonymMap := make(map[int][]interface{}, length)
		hashMap := make(map[string][]interface{}, length)

		// walk through the map with names of anonymous fields
		for index, field := range ref.Anonym {

			// get the reflected value of field
			n := ref.Temporary.FieldByName(field)

			// append to anonymMap the name of anonymous field
			anonymMap[index] = append(anonymMap[index], n.Type().Name())

			// create the field map and the hash slice
			fieldMap := make(map[int][]interface{}, n.NumField())
			hashSlice := make([]interface{}, n.NumField())

			// walk through the values of fields assigned to the interface {}
			for i := 0; i < n.NumField(); i++ {

				// get the name of field and append to fieldMap the name of field
				name := n.Type().Field(i).Name
				fieldMap[i] = append(fieldMap[i], name)

				// set by kind of reflected value above
				switch n.Field(i).Kind() {

				// kind is a bool
				case reflect.Bool:

					// only the first field, which IsExecutable, is set to true.
					// Means, only one process in a collaboration can be executed
					// at runtime this can be changed in the future, if the engine
					// fits for more execution options
					hashSlice = r.injectConfig(name, i, hashSlice, n)

				// kind is a struct
				case reflect.Struct:

					// if the field Suffix is empty, generate hash value and
					// start to inject by each index of the given structs. Then,
					// check next element, generate hash value and inject the field
					// Suffix again
					hashSlice = r.injectCurrentField(i, hashSlice, n)
					hashSlice = r.injectNextField(i, hashSlice, n)

				}

			}

			// merge the hashSlice with the hashMap
			utils.MergeStringSliceToMap(hashMap, n.Type().Name(), hashSlice)

		}

	// zero anonymous fields are reflected
	case len(ref.Anonym) == 0:

		// walk through the map with names of reflection fields
		for _, reflectionField := range ref.Rflct {
			// get the reflected name of reflectionField
			nonAnonymReflectionField := ref.Temporary.FieldByName(reflectionField)
			// generate hash value and inject the field Suffix
			hash, _ := r.hash()
			nonAnonymReflectionField.Set(reflect.ValueOf(hash))
		}

		// walk through the map with names of boolean fields
		for _, configField := range ref.Config {
			// get the reflected value of field
			nonAnonymConfigField := ref.Temporary.FieldByName(configField)
			// only the first field, which IsExecutable is set to true
			nonAnonymConfigField.SetBool(true)
		}
	}

	p = ref.Set()

	return p

}

// Create receives the definitions repository by the app in p argument
// and calls the main elements to set the maps, including process parameters
// n of process. The method contains the reflected process definition (p interface{})
// and calls it by the reflected method name.
// This method hides specific setters (SetProcess, SetCollaboration, SetDiagram)
// in the example process by building the model with reflection.
func (r *Reflection) Create(p interface{}) {
	// el is the interface {}
	el := reflect.ValueOf(&p).Elem()

	// Allocate a temporary variable with type of the struct.
	// el.Elem() is the value contained in the interface
	definitions := reflect.New(el.Elem().Type()).Elem() // *core.Definitions
	definitions.Set(el.Elem())                          // reflected process definitions el will be assigned to the core definitions

	// set collaboration, process and diagram
	collaboration := definitions.MethodByName("SetCollaboration")
	collaboration.Call([]reflect.Value{})

	process := definitions.MethodByName("SetProcess")
	process.Call([]reflect.Value{reflect.ValueOf(2)}) // r.Process represents number of processes

	diagram := definitions.MethodByName("SetDiagram")
	diagram.Call([]reflect.Value{reflect.ValueOf(1)}) // 1 represents number of diagrams
}

/*
 * @private
 */

// hash generates a hash value by using the crypto/rand package
// and the hash/fnv package to generate a 32-bit FNV-1a hash.
// If the error is not nil, it means that the hash value could not be generated.
// The suffix is used to generate a unique ID for each element of a process.
func (r Reflection) hash() (Reflection, error) {

	n := 8
	b := make([]byte, n)
	c := fnv.New32a()

	if _, err := rand.Read(b); err != nil {
		return Reflection{}, err
	}
	s := fmt.Sprintf("%x", b)

	if _, err := c.Write([]byte(s)); err != nil {
		return Reflection{}, err
	}
	defer c.Reset()

	result := Reflection{Suffix: fmt.Sprintf("%x", string(c.Sum(nil)))}

	return result, nil
}

// injectConfig sets the bool type.
func (r *Reflection) injectConfig(name string, index int, slice []interface{}, field reflect.Value) []interface{} {
	if strings.Contains(name, "IsExecutable") && index == 0 {
		field.Field(0).SetBool(true)
		slice[index] = bool(true)
	} else {
		slice[index] = bool(false)
	}
	return slice
}

// injectCurrentField injects the current field with a hash value
func (r *Reflection) injectCurrentField(index int, slice []interface{}, field reflect.Value) []interface{} {
	strHash := fmt.Sprintf("%s", field.Field(index).FieldByName("Suffix"))
	if strHash == "" {
		hash, _ := r.hash()
		slice[index] = hash.Suffix
		field.Field(index).Set(reflect.ValueOf(hash))
	}
	return slice
}

// injectNextField injects the next field with a hash value
func (r *Reflection) injectNextField(index int, slice []interface{}, field reflect.Value) []interface{} {
	if index+1 < field.NumField() {
		nexthash, _ := r.hash()
		slice[index+1] = nexthash.Suffix
		field.Field(index + 1).Set(reflect.ValueOf(nexthash))
	}
	return slice
}
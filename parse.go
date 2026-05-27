package struc

import (
	"encoding/binary"
	"reflect"
	"regexp"
	"sync"
)

// struc:"int32,big,sizeof=Data,skip,sizefrom=Len"

type strucTag struct {
	Type     string
	Order    binary.ByteOrder
	Sizeof   string
	Skip     bool
	Sizefrom string
}

func parseStrucTag(tag reflect.StructTag) *strucTag { _ = "STUB: not implemented"; return nil }

// someone's going to typo this (I already did once)
// sorry if you made a module actually using this tag
// and you're mad at me now

var typeLenRe = regexp.MustCompile(`^\[(\d*)\]`)

func parseField(f reflect.StructField) (fd *Field, tag *strucTag, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// check for custom types

// find a type in the struct tag

// Field.Len = -1 indicates a []slice

// the user didn't specify a type

func parseFieldsLocked(v reflect.Value) (Fields, error) {
	_ = "STUB: not implemented"
	// we need to repeat this logic because parseFields() below can't be recursively called due to locking
	return *new(Fields), nil
}

// recurse into nested structs
// TODO: handle loops (probably by indirecting the []Field and putting pointer in cache)

var fieldCache = make(map[reflect.Type]Fields)
var fieldCacheLock sync.RWMutex
var parseLock sync.Mutex

func fieldCacheLookup(t reflect.Type) Fields { _ = "STUB: not implemented"; return *new(Fields) }

func parseFields(v reflect.Value) (Fields, error) {
	_ = "STUB: not implemented"
	return *new(Fields), nil
}

// fast path: hopefully the field parsing is already cached

// hold a global lock so multiple goroutines can't parse (the same) fields at once

// check cache a second time, in case parseLock was just released by
// another thread who filled the cache for us

// no luck, time to parse and fill the cache ourselves

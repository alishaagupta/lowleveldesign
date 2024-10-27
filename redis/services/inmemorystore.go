package services

import (
	"fmt"
	"lld/redis/utils"
	"reflect"
	"strings"
)

// var storage map[string]models.Attribute

// var attributeTypeMap map[string]interface{}

// key -> value
// "sde_bootcamp" : { "title": "SDE-Bootcamp", "price": 30000.00, "enrolled": false, "estimated_time": 30 }
// attribute key -> title, attribute value SDE-Bootcamp
// map[string]something
// map[string]map[string]something

type Store struct {
	keyValueStore    map[string]map[string]interface{}
	attributeTypeMap map[string]interface{}
}

func CreateStorage() *Store {
	// storage = make(map[string]models.Attribute)
	return &Store{
		make(map[string]map[string]interface{}),
		make(map[string]interface{}),
	}
}

func (store *Store) Put(key string, attributes [][]interface{}) {
	isValid := store.validateAttrs(attributes)

	if !isValid {
		fmt.Println("Data Type Error")
		return
	}

	val, ok := store.keyValueStore[key]

	if !ok {
		attrMap := make(map[string]interface{})
		for _, j := range attributes {
			attrKey := j[0].(string)
			attrVal := j[1]

			attrMap[attrKey] = attrVal
			store.attributeTypeMap[attrKey] = reflect.TypeOf(attrVal)
		}
		store.keyValueStore[key] = attrMap
	} else {
		for _, attribute := range attributes {
			val[attribute[0].(string)] = attribute[1]
			store.attributeTypeMap[attribute[0].(string)] = reflect.TypeOf(attribute[1])
		}

	}

}

func (store *Store) validateAttrs(attributes [][]interface{}) bool {

	isValid := true

	for _, attribute := range attributes {

		val, ok := store.attributeTypeMap[attribute[0].(string)]

		if !ok {
			continue
		} else {
			if reflect.TypeOf(attribute[1]) != val {
				isValid = false
				break
			}
		}

	}

	return isValid

}

func (store *Store) Keys() []string {
	arr := make([]string, 10)
	for key := range store.keyValueStore {
		arr = append(arr, key)

	}
	return arr
}

func (store *Store) Get(key string) string {
	val, ok := store.keyValueStore[key]

	if !ok {
		fmt.Println("No entry found for " + key)
	}

	return store.processValue(val)

}

func (store *Store) processValue(dict map[string]interface{}) string {
	var sb strings.Builder

	for attrKey, attrVal := range dict {
		sb.WriteString(attrKey)
		sb.WriteString(":")
		sb.WriteString(utils.ConvertToString(attrVal))
		sb.WriteString(" ,")

	}

	return sb.String()
}

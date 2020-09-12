package structures

import (
	"sync"
)

// SMap is the structure that have the results
var SMap sync.Map

// AddElement add an element to the map
func AddElement(key string, value string) {
	results, ok := SMap.Load(key)
	if ok == true {
		tempArray := results.([]string)
		tempArray = append(tempArray, value)
		SMap.Store(key, tempArray)
	} else {
		var newArray []string
		newArray = make([]string, 0)
		newArray = append(newArray, value)
		SMap.Store(key, newArray)
	}
}

// GetMap returns the map
func GetMap() *sync.Map {
	return &SMap
}

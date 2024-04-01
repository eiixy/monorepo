package helpers

import (
	"reflect"
	"testing"
)

func TestMapMerge(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"c": 3, "d": 4}
	expected := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	result := MapMerge(m1, m2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestToStringMap(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	p := Person{"John", 30}
	expected := map[string]interface{}{"Age": 30, "Name": "John"}
	result := ToStringMap(p)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestToStruct(t *testing.T) {
	m := map[string]interface{}{"Name": "John", "Age": 30}
	var p struct {
		Name string
		Age  int
	}
	expected := struct {
		Name string
		Age  int
	}{"John", 30}
	err := ToStruct(m, &p)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(p, expected) {
		t.Errorf("Expected %v, but got %v", expected, p)
	}
}

func TestToMapE(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	p := Person{"John", 30}
	expected := map[string]interface{}{"Name": "John", "Age": 30}
	result, err := ToMapE(p)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestMapGet(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	expected := 1
	result := MapGet(m, "a", 0)
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
	expected = 0
	result = MapGet(m, "c", 0)
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

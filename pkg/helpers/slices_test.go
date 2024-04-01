package helpers

import (
	"fmt"
	"testing"
)

type Gender uint8

const (
	Male   Gender = 1
	Female Gender = 2
)

func (g Gender) String() string {
	switch g {
	case Male:
		return "male"
	case Female:
		return "female"
	default:
		return "unknown"
	}
}

type Test struct {
	Id   uint
	Name string
	Type uint8
	Attr struct {
		Age    uint8
		Gender Gender
	}
}

type Attr struct {
	Age    uint8
	Gender Gender
}

func TestSliceConv(t *testing.T) {
	items := []Test{
		{
			Id:   1,
			Name: "TEST-1",
			Attr: Attr{Age: 18, Gender: Male},
		}, {
			Id:   2,
			Name: "TEST-2",
			Attr: Attr{Age: 19, Gender: Female},
		},
		{
			Id:   3,
			Name: "TEST-3",
			Attr: Attr{Age: 22, Gender: Male},
		},
	}
	ids := SliceConv(items, func(item Test) uint { return item.Id })

	names := SliceConv(items, func(item Test) string {
		return item.Name
	})

	ages := SliceConv(items, func(item Test) uint8 {
		return item.Attr.Age
	})

	type Info struct {
		Name   string
		Age    uint8
		Gender string
	}
	maps := SliceConvMap(items, func(item Test) (uint, Info) {
		return item.Id, Info{
			Name:   item.Name,
			Age:    item.Attr.Age,
			Gender: item.Attr.Gender.String(),
		}
	})
	fmt.Println(ids, names, ages)
	fmt.Println(maps)
}

func TestSliceChunk(t *testing.T) {
	err := SliceChunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, 3, func(n []int) error {
		fmt.Println(n)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func TestSliceGroup(t *testing.T) {
	items := []Test{
		{
			Id:   1,
			Type: 1,
			Name: "TEST-1",
			Attr: Attr{Age: 18, Gender: Male},
		}, {
			Id:   2,
			Type: 1,
			Name: "TEST-2",
			Attr: Attr{Age: 19, Gender: Female},
		},
		{
			Id:   3,
			Type: 2,
			Name: "TEST-3",
			Attr: Attr{Age: 22, Gender: Male},
		},
		{
			Id:   4,
			Type: 3,
			Name: "TEST-4",
			Attr: Attr{Age: 22, Gender: Male},
		},
	}
	maps := SliceGroup(items, func(item Test) (uint8, string) {
		return item.Type, item.Name
	})
	fmt.Println(maps)
}

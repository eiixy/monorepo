package helpers

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func SliceUint32ToUint8(s []uint32) []uint8 {
	return SliceConv(s, func(t uint32) uint8 {
		return uint8(t)
	})
}

func SliceUint32ToUint(s []uint32) []uint {
	return SliceConv(s, func(t uint32) uint {
		return uint(t)
	})
}

func SliceUint32ToInt(s []uint32) []int {
	return SliceConv(s, func(t uint32) int {
		return int(t)
	})
}

func SliceUint32ToInt64(s []uint32) []int64 {
	return SliceConv(s, func(t uint32) int64 {
		return int64(t)
	})
}

func SliceInt64ToUint32(s []int64) []uint32 {
	return SliceConv(s, func(t int64) uint32 {
		return uint32(t)
	})
}

func SliceIntToUint(s []int) []uint {
	return SliceConv(s, func(t int) uint {
		return uint(t)
	})
}

func SliceUintToUint32(s []uint) []uint32 {
	return SliceConv(s, func(t uint) uint32 {
		return uint32(t)
	})
}

func SliceIntToUint32(s []int) []uint32 {
	return SliceConv(s, func(t int) uint32 {
		return uint32(t)
	})
}

func SliceUint8ToUint32(s []uint8) []uint32 {
	return SliceConv(s, func(t uint8) uint32 {
		return uint32(t)
	})
}

// SliceUnique 获取唯一值
func SliceUnique[T comparable](items []T) []T {
	maps := map[T]struct{}{}
	for _, item := range items {
		if _, ok := maps[item]; !ok {
			maps[item] = struct{}{}
		}
	}

	result, i := make([]T, len(maps)), 0
	for key := range maps {
		result[i] = key
		i++
	}
	return result
}

// SliceFilter 筛选切片
func SliceFilter[T any](items []T, fn func(T) bool) []T {
	var result []T
	for _, item := range items {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

// SliceFind 查找符合条件的 item
func SliceFind[T any](items []T, fn func(T) bool) (T, bool) {
	var result T
	for _, item := range items {
		if fn(item) {
			result = item
			return result, true
		}
	}
	return result, false
}

// SliceConv 将切片转换为指定切片类型
func SliceConv[T any, K any](items []T, fn func(T) K) []K {
	var result = make([]K, len(items))
	for i, item := range items {
		result[i] = fn(item)
	}
	return result
}

// SliceConvMap 将切片转换为指定map类型
func SliceConvMap[T any, K comparable, V any](items []T, fn func(T) (K, V)) map[K]V {
	var result = make(map[K]V)
	for _, item := range items {
		key, val := fn(item)
		result[key] = val
	}
	return result
}

// SliceGroup 将切片进行分组
func SliceGroup[T any, K comparable, V any](items []T, fn func(T) (K, V)) map[K][]V {
	var result = make(map[K][]V)
	for _, item := range items {
		key, val := fn(item)
		if _, ok := result[key]; !ok {
			result[key] = []V{}
		}
		result[key] = append(result[key], val)
	}
	return result
}

// SliceIndex 返回目标 `item` 出现的的第一个索引位置，
// 或者在没有匹配值时返回 -1。
func SliceIndex[T comparable](items []T, t T) int {
	for i, item := range items {
		if item == t {
			return i
		}
	}
	return -1
}

// SliceInclude 如果目标 `item` 在这个切片中则返回 `true`。
func SliceInclude[T comparable](items []T, item T) bool {
	return SliceIndex(items, item) >= 0
}

func In[T comparable](item T, items ...T) bool {
	for _, val := range items {
		if item == val {
			return true
		}
	}
	return false
}

// SliceChunk chunk
func SliceChunk[T any](items []T, chunk int, fn func([]T) error) error {
	size := len(items)
	for i := 0; i < size; i = i + chunk {
		end := i + chunk
		if end > size {
			end = size
		}
		err := fn(items[i:end])
		if err != nil {
			return err
		}
	}
	return nil
}

// Min 取最小值
func Min[T constraints.Ordered](items ...T) T {
	if len(items) == 0 {
		var zero T
		return zero
	}
	min := items[0]
	for _, item := range items {
		if min > item {
			min = item
		}
	}
	return min
}

// Max 取最小值
func Max[T constraints.Ordered](items ...T) T {
	if len(items) == 0 {
		var zero T
		return zero
	}
	min := items[0]
	for _, item := range items {
		if min < item {
			min = item
		}
	}
	return min
}

// SliceSum 获取切片总和
func SliceSum[T any, V constraints.Integer | constraints.Float](items []T, fn func(T) V) V {
	var sum V
	for _, item := range items {
		sum += fn(item)
	}
	return sum
}

// Intersection 获取两个切片中的交集
func Intersection[T comparable](a, b []T) []T {
	var result []T
	for _, t := range a {
		if slices.Contains(b, t) {
			result = append(result, t)
		}
	}
	return result
}

// SliceMerge 将切片进行合并
func SliceMerge[T any](items []T, s ...[]T) []T {
	items = slices.Clone(items)
	for _, ts := range s {
		items = append(items, slices.Clone(ts)...)
	}
	return items
}

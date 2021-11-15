package ministats

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math"
	"sort"
	"sync"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Data contains dataset with samples
type Data struct {
	items    dataset
	capacity int
	mu       *sync.RWMutex
}

type dataset []uint64

func (d dataset) Len() int           { return len(d) }
func (d dataset) Less(i, j int) bool { return d[i] < d[j] }
func (d dataset) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

// ////////////////////////////////////////////////////////////////////////////////// //

// NewData creates new Data struct
func NewData(capacity int) *Data {
	return &Data{capacity: capacity, mu: &sync.RWMutex{}}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Add appends new value to dataset
func (d *Data) Add(v uint64) {
	d.mu.Lock()

	d.items = append(d.items, v)

	if d.capacity > 0 && len(d.items) > d.capacity {
		d.items = d.items[len(d.items)-d.capacity:]
	}

	d.mu.Unlock()
}

// Reset removes all data from dataset
func (d *Data) Reset() {
	d.mu.Lock()
	d.items = d.items[:0:0]
	d.mu.Unlock()
}

// Len returns number of items in dataset
func (d *Data) Len() int {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return len(d.items)
}

// Cap returns dataset capacity
func (d *Data) Cap() int {
	return d.capacity
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Min returns the smallest value in dataset
func (d *Data) Min() uint64 {
	d.mu.RLock()

	if len(d.items) == 0 {
		d.mu.RUnlock()
		return 0
	}

	v, _, _ := calcBasic(d.items)

	d.mu.RUnlock()

	return v
}

// Max returns the largest value in dataset
func (d *Data) Max() uint64 {
	d.mu.RLock()

	if len(d.items) == 0 {
		d.mu.RUnlock()
		return 0
	}

	_, v, _ := calcBasic(d.items)

	d.mu.RUnlock()

	return v
}

// Mean returns average value
func (d *Data) Mean() uint64 {
	d.mu.RLock()

	if len(d.items) == 0 {
		d.mu.RUnlock()
		return 0
	}

	_, _, v := calcBasic(d.items)

	d.mu.RUnlock()

	return v
}

// StdDevP returns standard deviation population from dataset
func (d *Data) StdDevP() uint64 {
	d.mu.RLock()

	if len(d.items) == 0 {
		d.mu.RUnlock()
		return 0
	}

	defer d.mu.RUnlock()

	return calcStdDev(d.items, true)
}

// StdDevS returns standard deviation sample from dataset
func (d *Data) StdDevS() uint64 {
	d.mu.RLock()

	if len(d.items) == 0 {
		d.mu.RUnlock()
		return 0
	}

	defer d.mu.RUnlock()

	return calcStdDev(d.items, false)
}

// Percentile returns slice with values from the dataset
// for given percentages values
func (d *Data) Percentile(percentages ...float64) []uint64 {
	d.mu.RLock()

	result := make([]uint64, len(percentages))

	if len(d.items) == 0 {
		d.mu.RUnlock()
		return result
	}

	l := float64(len(d.items))
	sd := createSortedCopy(d.items)

	for i, p := range percentages {
		p = math.Min(math.Max(p, 0), 100)
		pi := int(math.Ceil((p / 100.0) * l))

		if pi == 0 {
			result[i] = sd[0]
		} else {
			result[i] = sd[pi-1]
		}
	}

	d.mu.RUnlock()

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// calcBasic calculates min, max and mean
func calcBasic(items dataset) (uint64, uint64, uint64) {
	var min, max uint64 = items[0], items[0]
	var sum uint64

	for i := range items {
		min = minU64(min, items[i])
		max = maxU64(max, items[i])
		sum += items[i]
	}

	return min, max, sum / uint64(len(items))
}

// calcStdDev calculates statndart deviation (population/sample)
func calcStdDev(items dataset, isPopulation bool) uint64 {
	_, _, mean := calcBasic(items)
	m := int64(mean)

	var vv int64
	var vr float64

	for i := range items {
		v := int64(items[i])
		vv += (v - m) * (v - m)
	}

	if isPopulation {
		vr = float64(vv) / float64(len(items))
	} else {
		vr = float64(vv) / float64(len(items)-1)
	}

	return uint64(math.Pow(vr, 0.5))
}

// minU64 returns the smaller of x or y
func minU64(v1, v2 uint64) uint64 {
	if v1 < v2 {
		return v1
	}

	return v2
}

// maxU64 returns the larger of x or y
func maxU64(v1, v2 uint64) uint64 {
	if v1 > v2 {
		return v1
	}

	return v2
}

// createSortedCopy creates sorted copy of dataset
func createSortedCopy(d dataset) dataset {
	result := make(dataset, len(d))

	copy(result, d)
	sort.Sort(result)

	return result
}

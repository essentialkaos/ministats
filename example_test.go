package ministats

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleNewData() {
	// New data struct with dataset capacity 3
	d := NewData(3)

	d.Add(1)
	d.Add(10)
	d.Add(5)

	fmt.Printf("Min: %d\n", d.Min())
	fmt.Printf("Max: %d\n", d.Max())

	d.Add(3)
	d.Add(8)

	fmt.Printf("Min: %d\n", d.Min())
	fmt.Printf("Max: %d\n", d.Max())

	// Output:
	// Min: 1
	// Max: 10
	// Min: 3
	// Max: 8
}

func ExampleData_Add() {
	d := NewData(3)

	d.Add(1)
	d.Add(10)
	d.Add(5)

	fmt.Printf("Min: %d\n", d.Min())
	fmt.Printf("Max: %d\n", d.Max())

	// Output:
	// Min: 1
	// Max: 10
}

func ExampleData_Reset() {
	d := NewData(3)

	d.Add(5)

	fmt.Printf("Min: %d\n", d.Min())

	d.Reset()

	fmt.Printf("Min: %d\n", d.Min())

	// Output:
	// Min: 5
	// Min: 0
}

func ExampleData_Len() {
	d := NewData(10)

	d.Add(5)
	d.Add(15)
	d.Add(3)

	fmt.Printf("Len: %d\n", d.Len())

	// Output:
	// Len: 3
}

func ExampleData_Cap() {
	d := NewData(10)

	d.Add(5)
	d.Add(15)
	d.Add(3)

	fmt.Printf("Len: %d\n", d.Len())
	fmt.Printf("Cap: %d\n", d.Cap())

	// Output:
	// Len: 3
	// Cap: 10
}

func ExampleData_Min() {
	d := NewData(3)

	d.Add(1)
	d.Add(10)
	d.Add(5)

	fmt.Printf("Min: %d\n", d.Min())

	// Output:
	// Min: 1
}

func ExampleData_Max() {
	d := NewData(3)

	d.Add(1)
	d.Add(10)
	d.Add(5)

	fmt.Printf("Max: %d\n", d.Max())

	// Output:
	// Max: 10
}

func ExampleData_Mean() {
	d := NewData(3)

	d.Add(1)
	d.Add(3)
	d.Add(8)

	fmt.Printf("Mean: %d\n", d.Mean())

	// Output:
	// Mean: 4
}

func ExampleData_StdDevP() {
	d := NewData(10)

	d.Add(3)
	d.Add(17)
	d.Add(36)
	d.Add(20)
	d.Add(9)

	fmt.Printf("StdDev: %d\n", d.StdDevP())

	// Output:
	// StdDev: 11
}

func ExampleData_StdDevS() {
	d := NewData(10)

	d.Add(3)
	d.Add(17)
	d.Add(36)
	d.Add(20)
	d.Add(9)

	fmt.Printf("StdDev: %d\n", d.StdDevS())

	// Output:
	// StdDev: 12
}

func ExampleData_Percentile() {
	d := NewData(10)

	d.Add(1)
	d.Add(34)
	d.Add(6)
	d.Add(38)
	d.Add(3)
	d.Add(8)
	d.Add(2)
	d.Add(1)
	d.Add(5)
	d.Add(2)

	p := d.Percentile(0, 90, 99)

	fmt.Printf("Percentile(0): %d\n", p[0])
	fmt.Printf("Percentile(90): %d\n", p[1])
	fmt.Printf("Percentile(99): %d\n", p[2])

	// Output:
	// Percentile(0): 1
	// Percentile(90): 34
	// Percentile(99): 38
}

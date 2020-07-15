package china_division

import (
	"fmt"
)

func ExampleGetFullName() {
	fmt.Println(GetFullName(`511181`))

	// Output:
	// 四川省 乐山市 峨眉山市
}

func ExampleGetJsonChildren() {
	fmt.Println(len(GetJsonChildren(`510000`)) > 2)

	// Output:
	// true
}
func ExampleGetChildren() {
	fmt.Println(len(GetChildren(`510000`)) != 0)
	fmt.Println(len(GetChildren(`511181`)) == 0)

	// Output:
	// true
	// true
}

func ExampleGetJsonProvinces() {
	fmt.Println(len(GetJsonProvinces()) > 2)

	// Output:
	// true
}
func ExampleGetProvinces() {
	fmt.Println(len(GetProvinces()) != 0)

	// Output:
	// true
}

func ExampleGetJsonCities() {
	fmt.Println(len(GetJsonCities(`510000`)) > 2)

	// Output:
	// true
}
func ExampleGetCities() {
	fmt.Println(len(GetCities(`510000`)) != 0)

	// Output:
	// true
}

func ExampleGetJsonAreas() {
	fmt.Println(len(GetJsonAreas(`511100`)) > 2)

	// Output:
	// true
}
func ExampleGetAreas() {
	fmt.Println(len(GetAreas(`511100`)) != 0)

	// Output:
	// true
}

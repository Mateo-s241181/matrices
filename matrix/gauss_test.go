package matrix

import "fmt"

func ExampleMatrix_Normalize() {
	m := Matrix{
		{-1, -2, -3},
		{0, 5, 10},
		{0, 7, 2},
	}

	m.Normalize(0)
	fmt.Println(m)

	m.Normalize(1)
	fmt.Println(m)

	m.Normalize(2)
	fmt.Println(m)

	// Output:
	// [[1 2 3] [0 5 10] [0 7 2]]
	// [[1 2 3] [0 1 2] [0 7 2]]
	// [[1 2 3] [0 1 2] [0 3.5 1]]
}

func ExampleMatrix_EliminateBelow() {
	m := Matrix{
		{1, 2, 3},
		{2, 5, 4},
		{0, 6, -3},
	}

	m.EliminateBelow(0)
	fmt.Println(m)

	// Output:
	// [[1 2 3] [0 -0.5 1] [0 6 -3]]

}

func ExampleMatrix_EliminateAbove() {
	m := Matrix{
		{-3, 6, -3},
		{4, 5, 0},
		{3, 2, 1},
	}

	m.EliminateAbove(2)
	fmt.Println(m)

	// Output:
	// [[2 4 0] [4 5 0] [3 2 1]]
}

func ExampleMatrix_UpperTriangular() {
	m := Matrix{
		{1, 2, 3},
		{2, 5, 4},
		{0, 6, -3},
	}

	m.UpperTriangular()
	fmt.Println(m)

	// Output:
	// [[1 2 3] [0 1 -2] [0 0 1]]
}

func ExampleMatrix_LowerTriangular() {
	m := Matrix{
		{-3, 6, -3},
		{4, 5, 2},
		{3, 2, 1},
	}

	m.LowerTriangular()
	fmt.Println(m)

	// Output:
	// [[1 0 0] [-2 1 0] [3 2 1]]
}

func ExampleMatrix_Gauss() {
	m := Matrix{
		{1, 2, 1},
		{2, 1, 0},
		{1, 3, 2},
	}

	m.Gauss()
	fmt.Println(m)

	// Output:
	// [[1 0 0] [0 1 0] [0 0 1]]
}

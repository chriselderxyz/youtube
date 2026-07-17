package math

import (
	"testing"
)

func setup()    {}
func teardown() {}

func TestMain(m *testing.M) {
	setup()
	defer teardown()
	m.Run()
}

func TestAdd(t *testing.T) {

	// Categorization Using ENV Variables
	// if os.Getenv("UNIT") != true {
	// 	t.Skip("skipping")
	// }

	// Basic Test
	got := Add(2, 3)
	if got != 5 {
		// t.Fail()
		// t.Error(got)
		t.Errorf("Expected 5, got %d", got)
		// t.FailNow()
		// t.Fatal(got)
		// t.Fatalf("Expected 5, got %d", got)
	}

	// Expect(got).ToBe(5) - No Native Assertions in Go
}

func TestAddWithSubtests(t *testing.T) {
	// setup code
	t.Run("test 1", func(t *testing.T) {
		// subtest
		got := Add(2, 3)
		if got != 5 {
			t.Errorf("Expected 5, got %d", got)
		}
	})

	t.Run("test 2", func(t *testing.T) {
		// subtest
		got := Add(3, 5)
		if got != 8 {
			t.Errorf("Expected 8, got %d", got)
		}
	})
}

func TestMax(t *testing.T) {

	// Mark test to run in parallel
	// Parallel tests are run together, after sequential tests finish
	t.Parallel()

	// Table Driven Test
	// Uses a slice of test cases you can range over
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "First is Larger",
			a:    10,
			b:    4,
			want: 10,
		},
		{
			name: "second number is larger",
			a:    3,
			b:    8,
			want: 8,
		},
		{
			name: "numbers are equal",
			a:    5,
			b:    5,
			want: 5,
		},
		{
			name: "negative numbers",
			a:    -2,
			b:    -7,
			want: -2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Max(test.a, test.b)
			if got != test.want {
				t.Errorf(
					"Max(%d, %d) = %d; want %d",
					test.a,
					test.b,
					got,
					test.want,
				)
			}
		})
	}
}

func TestLong(t *testing.T) {
	
	// Using Short Mode With -short Flag
	// go test -short -v ./...
	if testing.Short() {
		t.Skip("Short mode skipping")
	}
}

package handy_test

import (
	"testing"

	"github.com/podhmo/handy"
)

func TestEqual(t *testing.T) {
	add := func(x, y int) int {
		return x + y
	}

	handy.Must(t, handy.Equal(30).Actual(add(10, 20)))
	handy.Should(t, handy.Equal(30).Actual(add(10, 20)))

	handy.Must(t, handy.NotEqual(31).Actual(add(10, 20)))
	handy.Should(t, handy.NotEqual(31).Actual(add(10, 20)))
}

func TestDeepEqual(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "foo", Age: 20}
	p2 := Person{Name: "foo", Age: 20}
	handy.Must(t, handy.DeepEqual(p).Actual(p))
	handy.Must(t, handy.DeepEqual(&p).Actual(&p))
	handy.Should(t, handy.DeepEqual(p).Actual(p2))

	handy.Should(t, handy.NotDeepEqual(p).Actual(&p))
}

func TestJSONEqual(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	type Person2 struct {
		Name string
		Age  int
	}
	type Person3 struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := Person{Name: "foo", Age: 20}
	p1 := Person{Name: "foo", Age: 20}
	p2 := Person2{Name: "foo", Age: 20}
	p3 := Person3{Name: "foo", Age: 20}

	handy.Must(t, handy.JSONEqual(p).Actual(p))
	handy.Must(t, handy.JSONEqual(&p).Actual(p))
	handy.Must(t, handy.JSONEqual(p).Actual(&p))
	handy.Must(t, handy.JSONEqual(&p).Actual(&p))
	handy.Should(t, handy.JSONEqual(p).Actual(p1))
	handy.Should(t, handy.JSONEqual(p).Actual(p2))

	handy.Must(t, handy.NotJSONEqual(nil).Actual(&p))
	handy.Must(t, handy.NotJSONEqual(&p).Actual(nil))
	handy.Should(t, handy.NotJSONEqual(p).Actual(p3))
	handy.Should(t, handy.NotJSONEqual(p3).Actual(p))
}

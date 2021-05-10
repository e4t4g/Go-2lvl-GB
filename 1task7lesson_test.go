package main

import "testing"

func TestPrintStruct(t *testing.T) {

	type Data struct {
		Name string
		Surname string
		Age int
	}

	type testStr struct {
		in1 *Data
		in2 map[string]interface{}
		want    error
	}

	mapTest := []testStr{
		{	&Data{},
			map[string]interface{}{
				"Name": "Alex",
				"Surname": "K",
				"Age": 28,
		},
		nil,
		},
	}

	for _, test := range mapTest {
		if got := PrintStruct(test.in1, test.in2); got != nil {
			t.Errorf("Got %v, want %v", got, test.want)
		}
	}

}
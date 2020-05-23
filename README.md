# picksort

``` go
	type Person struct {
		name string
		age  int
	}

	persons := []Person{
		{
			name: "jack",
			age:  55,
		},
		{
			name: "lucy",
			age:  11,
		},
		{
			name: "bob",
			age:  33,
		},
	}

	personIfs := []interface{}{}

	for _, p := range persons {
		personIfs = append(personIfs, p)
	}
	PickInt(func(person interface{}) int {
		return person.(Person).age
	}).Sort(personIfs, func(i, j int) bool {
		return i < j
	})

	if personIfs[0].(Person).name != "lucy" || personIfs[1].(Person).name != "bob" || personIfs[2].(Person).name != "jack" {
		t.Fail()
	}
  ```

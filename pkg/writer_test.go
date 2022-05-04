package pkg

import "testing"

func TestToGoIdentierName(t *testing.T) {

	//given
	table := []struct {
		name         string
		expectedName string
	}{
		{"Something", "something"},
		{"SomeThing", "someThing"},
		{"SOMEthing", "something"},
		{"SOmeThing", "someThing"},
	}
	for _, row := range table {

		result := toGoIdentifierName(row.name)

		if result != row.expectedName {
			t.Errorf("got wrong value. expected %v but got %v", row.expectedName, result)
		}

	}

}

package algorithms

import (
	"reflect"
	"testing"
)

func TestUniqueNames(t *testing.T) {
	want := []string{"Ava", "Olivia", "Emma", "Sophia"}
	if got := UniqueNames([]string{"Ava", "Emma", "Olivia"}, []string{"Olivia", "Sophia", "Emma"}); !reflect.DeepEqual(got, want) {
		t.Errorf("UniqueNames() = %v, want %v", got, want)
	}

}

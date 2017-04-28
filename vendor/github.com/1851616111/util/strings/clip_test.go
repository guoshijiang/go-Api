package strings

import (
	"reflect"
	"testing"
)

func TestClip(t *testing.T) {
	target := `"{2,1,3}"`
	if ret := Clip(&target, "\"{", ",", "}\""); !reflect.DeepEqual(ret, []string{"2", "1", "3"}) {
		t.Fatal(ret)
	}
	target = `2,1,3`
	if ret := Clip(&target, ``, `,`, ``); !reflect.DeepEqual(ret, []string{"2", "1", "3"}) {
		t.Fatal(ret)
	}

	target = `,,`
	if ret := Clip(&target, ``, `,`, ``); len(ret) != 3 {
		t.Fatal(ret)
	}

	if "" != `` {
		t.Fatal()
	}
}

package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items:  []string{"Matt"},
			result: "Hello, Matt!",
		},
		{
			items:  []string{"Matt, Anne"},
			result: "Hello, Matt, Anne!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("got %s, gave %v, wanted %s", s,
				st.items, st.result)
		}
	}

}

package hello

import "testing"

func TestHelloNew(t *testing.T) {
	t.Run("hello test ", func(t *testing.T) {
		got := Hello("syam")
		want := "Hello, syam"

		if got != want {
			t.Errorf("Test case failed. Got %q. Want %q", got, want)
		}
	})

	t.Run("say, Hello, world when nothing is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("Test case failed. Got %q. Want %q", got, want)
		}
	})
}

// just a comment

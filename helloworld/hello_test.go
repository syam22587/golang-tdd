package hello

import "testing"

func TestHelloNew(t *testing.T) {
	t.Run("hello test ", func(t *testing.T) {
		got := Hello("syam")
		want := "Hello, syam"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say, Hello, world when nothing is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}

// assertCorrectMessage helper function
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Test case failed. Got %q. Want %q", got, want)
	}
}

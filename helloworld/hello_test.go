package hello

import "testing"

func TestHelloNew(t *testing.T) {
	t.Run("hello test ", func(t *testing.T) {
		got := Hello("syam", "Spanish")
		want := "Hola, syam"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say in french", func(t *testing.T) {
		got := Hello("syam", "French")
		want := "Bonjour, syam"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say, Hello, world when nothing is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say english Hello with voleti, when name is supllied but not language", func(t *testing.T) {
		got := Hello("voleti", "")
		want := "Hello, voleti"
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

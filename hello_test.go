package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Ahoy, world!"
	got := Hello()

	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

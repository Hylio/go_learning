package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// 使用helper报错会报在调用该函数的测试函数中
		// 而不会包在辅助函数
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("hello in spanish", func(t *testing.T) {
		got := Hello("spanish", "tom")
		want := "Hola, tom"

		assertCorrectMessage(t, got, want)
	})

	t.Run("hello to nil", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
}

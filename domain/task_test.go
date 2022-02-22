package domain

import "testing"

func TestTask(t *testing.T) {
	task := Task{
		Name: "my-task",
	}
	_ = task
	// want := "Hello, world."
	// if got := Hello(); got != want {
	// 	t.Errorf("Hello() = %q, want %q", got, want)
	// }
}

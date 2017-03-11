package main

import (
  "os"
  "testing"
)

func TestShake(t *testing.T) {
  expectedFailure := ShakeFailureMessage
  os.Args = []string{"/path/name"}
  actual := Shake()

  if actual != expectedFailure {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expectedFailure, actual)
  }

  os.Args = []string{"/path/name", "Will I win the lottery?"}
  actual = Shake()

  if actual == expectedFailure {
    t.Errorf("Test failed, did not expect: '%s'", expectedFailure)
  }
}

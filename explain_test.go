package main

import (
	"testing"
)

func TestExplain(t *testing.T) {
	Explain("job").ToYaml()
}

package struct_metadata_test

import (
	"github.com/nodejayes/struct-metadata"
	"testing"
)

func init() {
	struct_metadata.Metadata[testStruct](testStructMeta{
		Value: "a string value",
	})
}

type testStructMeta struct {
	Value string
}

type testStruct struct{}

func TestGetMetadata(t *testing.T) {
	meta := struct_metadata.GetMetadata[testStruct, testStructMeta]()
	if meta.Value != "a string value" {
		t.Errorf("expect `a string value` in metadata")
	}
}

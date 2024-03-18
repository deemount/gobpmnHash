package gobpmn_hash_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	gobpmn_hash "github.com/deemount/gobpmnHash"
	"github.com/deemount/gobpmnModels/pkg/core"
)

/*

Note:
Test Process and TestEmbeddedProcess are two structs copied from the
gobpmnByExample package.

*/

// TestEmbeddedProcess is a struct that contains fields of type
// core.DefinitionsRepository and two anonymous embedded structs.
// It is a copy of the gobpmnByExample/example/01 package.
type TestEmbeddedProcess struct {
	Def core.DefinitionsRepository
	TestEmbeddedPool
	TestEmbeddedTenant
}

// TestEmbeddedPool is a struct that contains fields of type
// gobpmn_hash.Injection and bool.
// It is embedded in the TestEmbeddedProcess struct.
type TestEmbeddedPool struct {
	XYZIsExecutable bool
	XYZProcess      gobpmn_hash.Injection
}

// TestEmbeddedTenant is a struct that contains a field of type
// gobpmn_hash.Injection.
// It is embedded in the TestEmbeddedProcess struct.
type TestEmbeddedTenant struct {
	XYZStartEvent gobpmn_hash.Injection
}

// TestProcess is a struct that contains fields of type
// core.DefinitionsRepository, gobpmn_hash.Injection, and bool.
// It is a copy of the gobpmnByExample/example/02 package.
type TestProcess struct {
	Def            core.DefinitionsRepository
	IsExecutable   bool
	Process        gobpmn_hash.Injection
	StartEvent     gobpmn_hash.Injection
	FromStartEvent gobpmn_hash.Injection
	Task           gobpmn_hash.Injection
	FromTask       gobpmn_hash.Injection
	EndEvent       gobpmn_hash.Injection
}

// TestInjectConfig tests the injection of a bool type
// into the field IsExecutable of the struct TestProcess
// and TestEmbeddedPool.
// The field IsExecutable is set to false as default, but
// if reflected the field is set to true.
// In gobpmnHash, a Config consists of a reflected struct field
// which has a bool type.
// This test should pass if the field IsExecutable is expected
// set to true or false
func TestInjectConfig(t *testing.T) {
	t.Run("p(TestProcess{}).IsExecutable expected true",
		func(t *testing.T) {
			hash := new(gobpmn_hash.Injection)
			p := hash.Inject(TestProcess{}).(TestProcess)
			if assert.NotNil(t, p) {
				assert.Equal(t, bool(true), p.IsExecutable)
			}
		})
	t.Run("p(TestProcess{}).IsExecutable expected false",
		func(t *testing.T) {
			hash := new(gobpmn_hash.Injection)
			p := hash.Inject(TestProcess{}).(TestProcess)
			if assert.NotNil(t, p) {
				assert.NotEqual(t, bool(false), p.IsExecutable)
			}
		})
	t.Run("p(TestEmbeddedProcess{}).IsExecutable expected true", func(t *testing.T) {
		hash := new(gobpmn_hash.Injection)
		p := hash.Inject(TestEmbeddedProcess{}).(TestEmbeddedProcess)
		if assert.NotNil(t, p) {
			assert.Equal(t, bool(true), p.XYZIsExecutable)
		}
	})
	t.Run("p(TestEmbeddedProcess{}).IsExecutable expected false", func(t *testing.T) {
		hash := new(gobpmn_hash.Injection)
		p := hash.Inject(TestEmbeddedProcess{}).(TestEmbeddedProcess)
		if assert.NotNil(t, p) {
			assert.NotEqual(t, bool(false), p.XYZIsExecutable)
		}
	})
}

func BenchmarkInject(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hash := new(gobpmn_hash.Injection)
		_ = hash.Inject(TestProcess{}).(TestProcess)
	}
}

func BenchmarkInjectEmbedded(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hash := new(gobpmn_hash.Injection)
		_ = hash.Inject(TestEmbeddedProcess{}).(TestEmbeddedProcess)
	}
}

package gobpmn_hash_test

import (
	"testing"

	gobpmn_hash "github.com/deemount/gobpmnHash"
	"github.com/deemount/gobpmnModels/pkg/core"
)

type Injection struct {
	Suffix string
}

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

type TestEmbeddedProcess struct {
	Def core.DefinitionsRepository
	TestEmbeddedPool
	TestEmbeddedTenant
}

type TestEmbeddedPool struct {
	XYZIsExecutable bool
	XYZProcess      gobpmn_hash.Injection
}

type TestEmbeddedTenant struct {
	XYZStartEvent gobpmn_hash.Injection
}

func TestInject(t *testing.T) {
	t.Log("TestInjectEmbedded")
	var hash gobpmn_hash.Injection
	p := hash.Inject(TestProcess{}).(TestProcess)
	p.Def = core.NewDefinitions()
	t.Logf("p: %#+v", p)
}

func TestInjectEmbedded(t *testing.T) {
	t.Log("TestInjectEmbedded")
	var hash gobpmn_hash.Injection
	p := hash.Inject(TestEmbeddedProcess{}).(TestEmbeddedProcess)
	p.Def = core.NewDefinitions()
	t.Logf("p: %#+v", p)
}

func BenchmarkInject(b *testing.B) {
	var hash gobpmn_hash.Injection
	for n := 0; n < b.N; n++ {
		p := hash.Inject(TestProcess{}).(TestProcess)
		p.Def = core.NewDefinitions()
		b.Logf("p: %#+v", p)
	}
}

func BenchmarkInjectEmbedded(b *testing.B) {
	var hash gobpmn_hash.Injection
	for n := 0; n < b.N; n++ {
		p := hash.Inject(TestEmbeddedProcess{}).(TestEmbeddedProcess)
		p.Def = core.NewDefinitions()
		b.Logf("p: %#+v", p)
	}
}

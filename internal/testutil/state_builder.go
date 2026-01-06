package testutil

import "github.com/lucasgjanot/go-gator-feed/internal/runtime"

type StateBuilder struct {
	DB      *FakeDatabase
	Config *FakeConfig
	Output  *FakeOutput
}

func NewState() *StateBuilder {
	return &StateBuilder{
		DB:      NewFakeDatabase(),
		Config: NewFakeConfig(),
		Output:  &FakeOutput{},
	}
}

func (b *StateBuilder) Build() *runtime.State {
	return &runtime.State{
		Database: b.DB,
		Config:  b.Config,
		Output:   b.Output,
	}
}

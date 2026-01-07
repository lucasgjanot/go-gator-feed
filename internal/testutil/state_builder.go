package testutil

import "github.com/lucasgjanot/go-gator-feed/internal/runtime"


type StateBuilder struct {
	Database *FakeDatabase
	Config *FakeConfig
	Output *FakeOutput
}

func NewState() *StateBuilder {
	return &StateBuilder{
		Database: NewFakeDatabase(),
		Config: NewFakeConfig(),
		Output: &FakeOutput{},
	}
}

func (b *StateBuilder) Build() *runtime.State {
	return &runtime.State{
		Database: runtime.Database{
			User: b.Database,
			Feed: b.Database,
		},
		Config: b.Config,
		Output: b.Output,
	}
}

package testutil

type FakeConfig struct {
	CurrentUser string
	SetUserErr  error
}

func NewFakeConfig() *FakeConfig {
	return &FakeConfig{}
}

func (f *FakeConfig) SetUser(name string) error {
	if f.SetUserErr != nil {
		return f.SetUserErr
	}
	f.CurrentUser = name
	return nil
}

func (f *FakeConfig) GetCurrentUser() string {
	return f.CurrentUser
}
package testdata

func init() {}

type initializeFirstTest struct{}

func (i *initializeFirstTest) Initialize() error {
	return nil
}

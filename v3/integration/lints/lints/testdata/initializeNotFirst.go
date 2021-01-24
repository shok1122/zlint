package testdata

func initializeNotFirst() {}

type initializeNotFirstTest struct{}

func (i *initializeNotFirstTest) Initialize() error {
	return nil
}

func init() {}

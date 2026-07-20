package zaptest

type TestingT interface {
	Logf(string, ...interface{})

	Errorf(string, ...interface{})

	Fail()

	Failed() bool

	Name() string

	FailNow()
}

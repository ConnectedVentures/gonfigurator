package gonfigurator

// AUTOGENERATED BY MOQ
// github.com/matryer/moq

// LoaderMock is a mock implementation of Loader.
//
//     func TestSomethingThatUsesLoader(t *testing.T) {
//
//         // make and configure a mocked Loader
//         mockedLoader := &LoaderMock{ 
//             ParseFunc: func(in1 string, in2 interface{}) error {
// 	               panic("TODO: mock out the Parse function")
//             },
//         }
//
//         // TODO: use mockedLoader in code that requires Loader
//     
//     }
type LoaderMock struct {
	// ParseFunc mocks the Parse function.
	ParseFunc func(in1 string, in2 interface{}) error
}

// Parse calls ParseFunc.
func (mock *LoaderMock) Parse(in1 string, in2 interface{}) error {
	if mock.ParseFunc == nil {
		panic("moq: LoaderMock.ParseFunc is nil but was just called")
	}
	return mock.ParseFunc(in1, in2)
}
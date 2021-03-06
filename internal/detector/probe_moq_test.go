// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package detector

import (
	"sync"
)

// Ensure, that ProbeMock does implement Probe.
// If this is not the case, regenerate this file with moq.
var _ Probe = &ProbeMock{}

// ProbeMock is a mock implementation of Probe.
//
//     func TestSomethingThatUsesProbe(t *testing.T) {
//
//         // make and configure a mocked Probe
//         mockedProbe := &ProbeMock{
//             RunFunc: func(host string) bool {
// 	               panic("mock out the Run method")
//             },
//         }
//
//         // use mockedProbe in code that requires Probe
//         // and then make assertions.
//
//     }
type ProbeMock struct {
	// RunFunc mocks the Run method.
	RunFunc func(host string) bool

	// calls tracks calls to the methods.
	calls struct {
		// Run holds details about calls to the Run method.
		Run []struct {
			// Host is the host argument value.
			Host string
		}
	}
	lockRun sync.RWMutex
}

// Run calls RunFunc.
func (mock *ProbeMock) Run(host string) bool {
	if mock.RunFunc == nil {
		panic("ProbeMock.RunFunc: method is nil but Probe.Run was just called")
	}
	callInfo := struct {
		Host string
	}{
		Host: host,
	}
	mock.lockRun.Lock()
	mock.calls.Run = append(mock.calls.Run, callInfo)
	mock.lockRun.Unlock()
	return mock.RunFunc(host)
}

// RunCalls gets all the calls that were made to Run.
// Check the length with:
//     len(mockedProbe.RunCalls())
func (mock *ProbeMock) RunCalls() []struct {
	Host string
} {
	var calls []struct {
		Host string
	}
	mock.lockRun.RLock()
	calls = mock.calls.Run
	mock.lockRun.RUnlock()
	return calls
}

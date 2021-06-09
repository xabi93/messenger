// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package publisher

import (
	"context"
	"github.com/xabi93/messenger"
	"sync"
)

// Ensure, that SourceMock does implement Source.
// If this is not the case, regenerate this file with moq.
var _ Source = &SourceMock{}

// SourceMock is a mock implementation of Source.
//
// 	func TestSomethingThatUsesSource(t *testing.T) {
//
// 		// make and configure a mocked Source
// 		mockedSource := &SourceMock{
// 			MessagesFunc: func(ctx context.Context, batch int64) ([]*messenger.Message, error) {
// 				panic("mock out the Messages method")
// 			},
// 			PublishedFunc: func(ctx context.Context, msg ...*messenger.Message) error {
// 				panic("mock out the Published method")
// 			},
// 		}
//
// 		// use mockedSource in code that requires Source
// 		// and then make assertions.
//
// 	}
type SourceMock struct {
	// MessagesFunc mocks the Messages method.
	MessagesFunc func(ctx context.Context, batch int64) ([]*messenger.Message, error)

	// PublishedFunc mocks the Published method.
	PublishedFunc func(ctx context.Context, msg ...*messenger.Message) error

	// calls tracks calls to the methods.
	calls struct {
		// Messages holds details about calls to the Messages method.
		Messages []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Batch is the batch argument value.
			Batch int64
		}
		// Published holds details about calls to the Published method.
		Published []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Msg is the msg argument value.
			Msg []*messenger.Message
		}
	}
	lockMessages  sync.RWMutex
	lockPublished sync.RWMutex
}

// Messages calls MessagesFunc.
func (mock *SourceMock) Messages(ctx context.Context, batch int64) ([]*messenger.Message, error) {
	callInfo := struct {
		Ctx   context.Context
		Batch int64
	}{
		Ctx:   ctx,
		Batch: batch,
	}
	mock.lockMessages.Lock()
	mock.calls.Messages = append(mock.calls.Messages, callInfo)
	mock.lockMessages.Unlock()
	if mock.MessagesFunc == nil {
		var (
			messagesOut []*messenger.Message
			errOut      error
		)
		return messagesOut, errOut
	}
	return mock.MessagesFunc(ctx, batch)
}

// MessagesCalls gets all the calls that were made to Messages.
// Check the length with:
//     len(mockedSource.MessagesCalls())
func (mock *SourceMock) MessagesCalls() []struct {
	Ctx   context.Context
	Batch int64
} {
	var calls []struct {
		Ctx   context.Context
		Batch int64
	}
	mock.lockMessages.RLock()
	calls = mock.calls.Messages
	mock.lockMessages.RUnlock()
	return calls
}

// Published calls PublishedFunc.
func (mock *SourceMock) Published(ctx context.Context, msg ...*messenger.Message) error {
	callInfo := struct {
		Ctx context.Context
		Msg []*messenger.Message
	}{
		Ctx: ctx,
		Msg: msg,
	}
	mock.lockPublished.Lock()
	mock.calls.Published = append(mock.calls.Published, callInfo)
	mock.lockPublished.Unlock()
	if mock.PublishedFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.PublishedFunc(ctx, msg...)
}

// PublishedCalls gets all the calls that were made to Published.
// Check the length with:
//     len(mockedSource.PublishedCalls())
func (mock *SourceMock) PublishedCalls() []struct {
	Ctx context.Context
	Msg []*messenger.Message
} {
	var calls []struct {
		Ctx context.Context
		Msg []*messenger.Message
	}
	mock.lockPublished.RLock()
	calls = mock.calls.Published
	mock.lockPublished.RUnlock()
	return calls
}

// Ensure, that PublisherMock does implement Publisher.
// If this is not the case, regenerate this file with moq.
var _ Publisher = &PublisherMock{}

// PublisherMock is a mock implementation of Publisher.
//
// 	func TestSomethingThatUsesPublisher(t *testing.T) {
//
// 		// make and configure a mocked Publisher
// 		mockedPublisher := &PublisherMock{
// 			PublishFunc: func(ctx context.Context, msg *messenger.Message) error {
// 				panic("mock out the Publish method")
// 			},
// 		}
//
// 		// use mockedPublisher in code that requires Publisher
// 		// and then make assertions.
//
// 	}
type PublisherMock struct {
	// PublishFunc mocks the Publish method.
	PublishFunc func(ctx context.Context, msg *messenger.Message) error

	// calls tracks calls to the methods.
	calls struct {
		// Publish holds details about calls to the Publish method.
		Publish []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Msg is the msg argument value.
			Msg *messenger.Message
		}
	}
	lockPublish sync.RWMutex
}

// Publish calls PublishFunc.
func (mock *PublisherMock) Publish(ctx context.Context, msg *messenger.Message) error {
	callInfo := struct {
		Ctx context.Context
		Msg *messenger.Message
	}{
		Ctx: ctx,
		Msg: msg,
	}
	mock.lockPublish.Lock()
	mock.calls.Publish = append(mock.calls.Publish, callInfo)
	mock.lockPublish.Unlock()
	if mock.PublishFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.PublishFunc(ctx, msg)
}

// PublishCalls gets all the calls that were made to Publish.
// Check the length with:
//     len(mockedPublisher.PublishCalls())
func (mock *PublisherMock) PublishCalls() []struct {
	Ctx context.Context
	Msg *messenger.Message
} {
	var calls []struct {
		Ctx context.Context
		Msg *messenger.Message
	}
	mock.lockPublish.RLock()
	calls = mock.calls.Publish
	mock.lockPublish.RUnlock()
	return calls
}

// Ensure, that ReporterMock does implement Reporter.
// If this is not the case, regenerate this file with moq.
var _ Reporter = &ReporterMock{}

// ReporterMock is a mock implementation of Reporter.
//
// 	func TestSomethingThatUsesReporter(t *testing.T) {
//
// 		// make and configure a mocked Reporter
// 		mockedReporter := &ReporterMock{
// 			EndFunc: func()  {
// 				panic("mock out the End method")
// 			},
// 			ErrorFunc: func(err error)  {
// 				panic("mock out the Error method")
// 			},
// 			FailedFunc: func(msg messenger.Message, err error)  {
// 				panic("mock out the Failed method")
// 			},
// 			InitFunc: func(totalMsgs int)  {
// 				panic("mock out the Init method")
// 			},
// 			SuccessFunc: func(msg messenger.Message)  {
// 				panic("mock out the Success method")
// 			},
// 		}
//
// 		// use mockedReporter in code that requires Reporter
// 		// and then make assertions.
//
// 	}
type ReporterMock struct {
	// EndFunc mocks the End method.
	EndFunc func()

	// ErrorFunc mocks the Error method.
	ErrorFunc func(err error)

	// FailedFunc mocks the Failed method.
	FailedFunc func(msg messenger.Message, err error)

	// InitFunc mocks the Init method.
	InitFunc func(totalMsgs int)

	// SuccessFunc mocks the Success method.
	SuccessFunc func(msg messenger.Message)

	// calls tracks calls to the methods.
	calls struct {
		// End holds details about calls to the End method.
		End []struct {
		}
		// Error holds details about calls to the Error method.
		Error []struct {
			// Err is the err argument value.
			Err error
		}
		// Failed holds details about calls to the Failed method.
		Failed []struct {
			// Msg is the msg argument value.
			Msg messenger.Message
			// Err is the err argument value.
			Err error
		}
		// Init holds details about calls to the Init method.
		Init []struct {
			// TotalMsgs is the totalMsgs argument value.
			TotalMsgs int
		}
		// Success holds details about calls to the Success method.
		Success []struct {
			// Msg is the msg argument value.
			Msg messenger.Message
		}
	}
	lockEnd     sync.RWMutex
	lockError   sync.RWMutex
	lockFailed  sync.RWMutex
	lockInit    sync.RWMutex
	lockSuccess sync.RWMutex
}

// End calls EndFunc.
func (mock *ReporterMock) End() {
	callInfo := struct {
	}{}
	mock.lockEnd.Lock()
	mock.calls.End = append(mock.calls.End, callInfo)
	mock.lockEnd.Unlock()
	if mock.EndFunc == nil {
		return
	}
	mock.EndFunc()
}

// EndCalls gets all the calls that were made to End.
// Check the length with:
//     len(mockedReporter.EndCalls())
func (mock *ReporterMock) EndCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockEnd.RLock()
	calls = mock.calls.End
	mock.lockEnd.RUnlock()
	return calls
}

// Error calls ErrorFunc.
func (mock *ReporterMock) Error(err error) {
	callInfo := struct {
		Err error
	}{
		Err: err,
	}
	mock.lockError.Lock()
	mock.calls.Error = append(mock.calls.Error, callInfo)
	mock.lockError.Unlock()
	if mock.ErrorFunc == nil {
		return
	}
	mock.ErrorFunc(err)
}

// ErrorCalls gets all the calls that were made to Error.
// Check the length with:
//     len(mockedReporter.ErrorCalls())
func (mock *ReporterMock) ErrorCalls() []struct {
	Err error
} {
	var calls []struct {
		Err error
	}
	mock.lockError.RLock()
	calls = mock.calls.Error
	mock.lockError.RUnlock()
	return calls
}

// Failed calls FailedFunc.
func (mock *ReporterMock) Failed(msg messenger.Message, err error) {
	callInfo := struct {
		Msg messenger.Message
		Err error
	}{
		Msg: msg,
		Err: err,
	}
	mock.lockFailed.Lock()
	mock.calls.Failed = append(mock.calls.Failed, callInfo)
	mock.lockFailed.Unlock()
	if mock.FailedFunc == nil {
		return
	}
	mock.FailedFunc(msg, err)
}

// FailedCalls gets all the calls that were made to Failed.
// Check the length with:
//     len(mockedReporter.FailedCalls())
func (mock *ReporterMock) FailedCalls() []struct {
	Msg messenger.Message
	Err error
} {
	var calls []struct {
		Msg messenger.Message
		Err error
	}
	mock.lockFailed.RLock()
	calls = mock.calls.Failed
	mock.lockFailed.RUnlock()
	return calls
}

// Init calls InitFunc.
func (mock *ReporterMock) Init(totalMsgs int) {
	callInfo := struct {
		TotalMsgs int
	}{
		TotalMsgs: totalMsgs,
	}
	mock.lockInit.Lock()
	mock.calls.Init = append(mock.calls.Init, callInfo)
	mock.lockInit.Unlock()
	if mock.InitFunc == nil {
		return
	}
	mock.InitFunc(totalMsgs)
}

// InitCalls gets all the calls that were made to Init.
// Check the length with:
//     len(mockedReporter.InitCalls())
func (mock *ReporterMock) InitCalls() []struct {
	TotalMsgs int
} {
	var calls []struct {
		TotalMsgs int
	}
	mock.lockInit.RLock()
	calls = mock.calls.Init
	mock.lockInit.RUnlock()
	return calls
}

// Success calls SuccessFunc.
func (mock *ReporterMock) Success(msg messenger.Message) {
	callInfo := struct {
		Msg messenger.Message
	}{
		Msg: msg,
	}
	mock.lockSuccess.Lock()
	mock.calls.Success = append(mock.calls.Success, callInfo)
	mock.lockSuccess.Unlock()
	if mock.SuccessFunc == nil {
		return
	}
	mock.SuccessFunc(msg)
}

// SuccessCalls gets all the calls that were made to Success.
// Check the length with:
//     len(mockedReporter.SuccessCalls())
func (mock *ReporterMock) SuccessCalls() []struct {
	Msg messenger.Message
} {
	var calls []struct {
		Msg messenger.Message
	}
	mock.lockSuccess.RLock()
	calls = mock.calls.Success
	mock.lockSuccess.RUnlock()
	return calls
}

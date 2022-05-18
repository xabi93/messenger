// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package sqs_test

import (
	"context"
	servicesqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	publishsqs "github.com/xabi93/messenger/publish/sqs"
	"sync"
)

// Ensure, that ClientMock does implement publishsqs.Client.
// If this is not the case, regenerate this file with moq.
var _ publishsqs.Client = &ClientMock{}

// ClientMock is a mock implementation of publishsqs.Client.
//
// 	func TestSomethingThatUsesClient(t *testing.T) {
//
// 		// make and configure a mocked publishsqs.Client
// 		mockedClient := &ClientMock{
// 			GetQueueUrlFunc: func(ctx context.Context, params *servicesqs.GetQueueUrlInput, optFns ...func(*servicesqs.Options)) (*servicesqs.GetQueueUrlOutput, error) {
// 				panic("mock out the GetQueueUrl method")
// 			},
// 			SendMessageFunc: func(ctx context.Context, params *servicesqs.SendMessageInput, optFns ...func(*servicesqs.Options)) (*servicesqs.SendMessageOutput, error) {
// 				panic("mock out the SendMessage method")
// 			},
// 		}
//
// 		// use mockedClient in code that requires publishsqs.Client
// 		// and then make assertions.
//
// 	}
type ClientMock struct {
	// GetQueueUrlFunc mocks the GetQueueUrl method.
	GetQueueUrlFunc func(ctx context.Context, params *servicesqs.GetQueueUrlInput, optFns ...func(*servicesqs.Options)) (*servicesqs.GetQueueUrlOutput, error)

	// SendMessageFunc mocks the SendMessage method.
	SendMessageFunc func(ctx context.Context, params *servicesqs.SendMessageInput, optFns ...func(*servicesqs.Options)) (*servicesqs.SendMessageOutput, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetQueueUrl holds details about calls to the GetQueueUrl method.
		GetQueueUrl []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *servicesqs.GetQueueUrlInput
			// OptFns is the optFns argument value.
			OptFns []func(*servicesqs.Options)
		}
		// SendMessage holds details about calls to the SendMessage method.
		SendMessage []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *servicesqs.SendMessageInput
			// OptFns is the optFns argument value.
			OptFns []func(*servicesqs.Options)
		}
	}
	lockGetQueueUrl sync.RWMutex
	lockSendMessage sync.RWMutex
}

// GetQueueUrl calls GetQueueUrlFunc.
func (mock *ClientMock) GetQueueUrl(ctx context.Context, params *servicesqs.GetQueueUrlInput, optFns ...func(*servicesqs.Options)) (*servicesqs.GetQueueUrlOutput, error) {
	callInfo := struct {
		Ctx    context.Context
		Params *servicesqs.GetQueueUrlInput
		OptFns []func(*servicesqs.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockGetQueueUrl.Lock()
	mock.calls.GetQueueUrl = append(mock.calls.GetQueueUrl, callInfo)
	mock.lockGetQueueUrl.Unlock()
	if mock.GetQueueUrlFunc == nil {
		var (
			getQueueUrlOutputOut *servicesqs.GetQueueUrlOutput
			errOut               error
		)
		return getQueueUrlOutputOut, errOut
	}
	return mock.GetQueueUrlFunc(ctx, params, optFns...)
}

// GetQueueUrlCalls gets all the calls that were made to GetQueueUrl.
// Check the length with:
//     len(mockedClient.GetQueueUrlCalls())
func (mock *ClientMock) GetQueueUrlCalls() []struct {
	Ctx    context.Context
	Params *servicesqs.GetQueueUrlInput
	OptFns []func(*servicesqs.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *servicesqs.GetQueueUrlInput
		OptFns []func(*servicesqs.Options)
	}
	mock.lockGetQueueUrl.RLock()
	calls = mock.calls.GetQueueUrl
	mock.lockGetQueueUrl.RUnlock()
	return calls
}

// SendMessage calls SendMessageFunc.
func (mock *ClientMock) SendMessage(ctx context.Context, params *servicesqs.SendMessageInput, optFns ...func(*servicesqs.Options)) (*servicesqs.SendMessageOutput, error) {
	callInfo := struct {
		Ctx    context.Context
		Params *servicesqs.SendMessageInput
		OptFns []func(*servicesqs.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockSendMessage.Lock()
	mock.calls.SendMessage = append(mock.calls.SendMessage, callInfo)
	mock.lockSendMessage.Unlock()
	if mock.SendMessageFunc == nil {
		var (
			sendMessageOutputOut *servicesqs.SendMessageOutput
			errOut               error
		)
		return sendMessageOutputOut, errOut
	}
	return mock.SendMessageFunc(ctx, params, optFns...)
}

// SendMessageCalls gets all the calls that were made to SendMessage.
// Check the length with:
//     len(mockedClient.SendMessageCalls())
func (mock *ClientMock) SendMessageCalls() []struct {
	Ctx    context.Context
	Params *servicesqs.SendMessageInput
	OptFns []func(*servicesqs.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *servicesqs.SendMessageInput
		OptFns []func(*servicesqs.Options)
	}
	mock.lockSendMessage.RLock()
	calls = mock.calls.SendMessage
	mock.lockSendMessage.RUnlock()
	return calls
}

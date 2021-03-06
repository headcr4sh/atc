// Code generated by counterfeiter. DO NOT EDIT.
package authfakes

import (
	"sync"

	"github.com/concourse/atc/auth"
)

type FakeCSRFTokenGenerator struct {
	GenerateTokenStub        func() (string, error)
	generateTokenMutex       sync.RWMutex
	generateTokenArgsForCall []struct{}
	generateTokenReturns     struct {
		result1 string
		result2 error
	}
	generateTokenReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCSRFTokenGenerator) GenerateToken() (string, error) {
	fake.generateTokenMutex.Lock()
	ret, specificReturn := fake.generateTokenReturnsOnCall[len(fake.generateTokenArgsForCall)]
	fake.generateTokenArgsForCall = append(fake.generateTokenArgsForCall, struct{}{})
	fake.recordInvocation("GenerateToken", []interface{}{})
	fake.generateTokenMutex.Unlock()
	if fake.GenerateTokenStub != nil {
		return fake.GenerateTokenStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.generateTokenReturns.result1, fake.generateTokenReturns.result2
}

func (fake *FakeCSRFTokenGenerator) GenerateTokenCallCount() int {
	fake.generateTokenMutex.RLock()
	defer fake.generateTokenMutex.RUnlock()
	return len(fake.generateTokenArgsForCall)
}

func (fake *FakeCSRFTokenGenerator) GenerateTokenReturns(result1 string, result2 error) {
	fake.GenerateTokenStub = nil
	fake.generateTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCSRFTokenGenerator) GenerateTokenReturnsOnCall(i int, result1 string, result2 error) {
	fake.GenerateTokenStub = nil
	if fake.generateTokenReturnsOnCall == nil {
		fake.generateTokenReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.generateTokenReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCSRFTokenGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateTokenMutex.RLock()
	defer fake.generateTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCSRFTokenGenerator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ auth.CSRFTokenGenerator = new(FakeCSRFTokenGenerator)

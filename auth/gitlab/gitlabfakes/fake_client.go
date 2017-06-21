// Code generated by counterfeiter. DO NOT EDIT.
package gitlabfakes

import (
	"net/http"
	"sync"

	"github.com/concourse/atc/auth/gitlab"
)

type FakeClient struct {
	GroupsStub        func(*http.Client) ([]string, error)
	groupsMutex       sync.RWMutex
	groupsArgsForCall []struct {
		arg1 *http.Client
	}
	groupsReturns struct {
		result1 []string
		result2 error
	}
	groupsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) Groups(arg1 *http.Client) ([]string, error) {
	fake.groupsMutex.Lock()
	ret, specificReturn := fake.groupsReturnsOnCall[len(fake.groupsArgsForCall)]
	fake.groupsArgsForCall = append(fake.groupsArgsForCall, struct {
		arg1 *http.Client
	}{arg1})
	fake.recordInvocation("Groups", []interface{}{arg1})
	fake.groupsMutex.Unlock()
	if fake.GroupsStub != nil {
		return fake.GroupsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.groupsReturns.result1, fake.groupsReturns.result2
}

func (fake *FakeClient) GroupsCallCount() int {
	fake.groupsMutex.RLock()
	defer fake.groupsMutex.RUnlock()
	return len(fake.groupsArgsForCall)
}

func (fake *FakeClient) GroupsArgsForCall(i int) *http.Client {
	fake.groupsMutex.RLock()
	defer fake.groupsMutex.RUnlock()
	return fake.groupsArgsForCall[i].arg1
}

func (fake *FakeClient) GroupsReturns(result1 []string, result2 error) {
	fake.GroupsStub = nil
	fake.groupsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GroupsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.GroupsStub = nil
	if fake.groupsReturnsOnCall == nil {
		fake.groupsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.groupsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.groupsMutex.RLock()
	defer fake.groupsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ gitlab.Client = new(FakeClient)

// This file was generated by counterfeiter
package fake

import (
	"sync"

	"github.com/phogolabs/oak/sqlmigr"
)

type MigrationRunner struct {
	RunStub        func(item *sqlmigr.Item) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		item *sqlmigr.Item
	}
	runReturns struct {
		result1 error
	}
	RevertStub        func(item *sqlmigr.Item) error
	revertMutex       sync.RWMutex
	revertArgsForCall []struct {
		item *sqlmigr.Item
	}
	revertReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MigrationRunner) Run(item *sqlmigr.Item) error {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		item *sqlmigr.Item
	}{item})
	fake.recordInvocation("Run", []interface{}{item})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(item)
	}
	return fake.runReturns.result1
}

func (fake *MigrationRunner) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *MigrationRunner) RunArgsForCall(i int) *sqlmigr.Item {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].item
}

func (fake *MigrationRunner) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *MigrationRunner) Revert(item *sqlmigr.Item) error {
	fake.revertMutex.Lock()
	fake.revertArgsForCall = append(fake.revertArgsForCall, struct {
		item *sqlmigr.Item
	}{item})
	fake.recordInvocation("Revert", []interface{}{item})
	fake.revertMutex.Unlock()
	if fake.RevertStub != nil {
		return fake.RevertStub(item)
	}
	return fake.revertReturns.result1
}

func (fake *MigrationRunner) RevertCallCount() int {
	fake.revertMutex.RLock()
	defer fake.revertMutex.RUnlock()
	return len(fake.revertArgsForCall)
}

func (fake *MigrationRunner) RevertArgsForCall(i int) *sqlmigr.Item {
	fake.revertMutex.RLock()
	defer fake.revertMutex.RUnlock()
	return fake.revertArgsForCall[i].item
}

func (fake *MigrationRunner) RevertReturns(result1 error) {
	fake.RevertStub = nil
	fake.revertReturns = struct {
		result1 error
	}{result1}
}

func (fake *MigrationRunner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.revertMutex.RLock()
	defer fake.revertMutex.RUnlock()
	return fake.invocations
}

func (fake *MigrationRunner) recordInvocation(key string, args []interface{}) {
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

var _ sqlmigr.ItemRunner = new(MigrationRunner)

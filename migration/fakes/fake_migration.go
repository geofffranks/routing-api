// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/routing-api/db"
	"code.cloudfoundry.org/routing-api/migration"
)

type FakeMigration struct {
	RunStub        func(*db.SqlDB) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 *db.SqlDB
	}
	runReturns struct {
		result1 error
	}
	VersionStub        func() int
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns     struct {
		result1 int
	}
}

func (fake *FakeMigration) Run(arg1 *db.SqlDB) error {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 *db.SqlDB
	}{arg1})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1)
	} else {
		return fake.runReturns.result1
	}
}

func (fake *FakeMigration) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeMigration) RunArgsForCall(i int) *db.SqlDB {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].arg1
}

func (fake *FakeMigration) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMigration) Version() int {
	fake.versionMutex.Lock()
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	} else {
		return fake.versionReturns.result1
	}
}

func (fake *FakeMigration) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeMigration) VersionReturns(result1 int) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 int
	}{result1}
}

var _ migration.Migration = new(FakeMigration)
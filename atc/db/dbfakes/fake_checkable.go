// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"context"
	"sync"
	"time"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
)

type FakeCheckable struct {
	CheckEveryStub        func() string
	checkEveryMutex       sync.RWMutex
	checkEveryArgsForCall []struct {
	}
	checkEveryReturns struct {
		result1 string
	}
	checkEveryReturnsOnCall map[int]struct {
		result1 string
	}
	CheckPlanStub        func(atc.Version, time.Duration, db.ResourceTypes, atc.Source) atc.CheckPlan
	checkPlanMutex       sync.RWMutex
	checkPlanArgsForCall []struct {
		arg1 atc.Version
		arg2 time.Duration
		arg3 db.ResourceTypes
		arg4 atc.Source
	}
	checkPlanReturns struct {
		result1 atc.CheckPlan
	}
	checkPlanReturnsOnCall map[int]struct {
		result1 atc.CheckPlan
	}
	CheckTimeoutStub        func() string
	checkTimeoutMutex       sync.RWMutex
	checkTimeoutArgsForCall []struct {
	}
	checkTimeoutReturns struct {
		result1 string
	}
	checkTimeoutReturnsOnCall map[int]struct {
		result1 string
	}
	CreateBuildStub        func(context.Context, bool) (db.Build, bool, error)
	createBuildMutex       sync.RWMutex
	createBuildArgsForCall []struct {
		arg1 context.Context
		arg2 bool
	}
	createBuildReturns struct {
		result1 db.Build
		result2 bool
		result3 error
	}
	createBuildReturnsOnCall map[int]struct {
		result1 db.Build
		result2 bool
		result3 error
	}
	CurrentPinnedVersionStub        func() atc.Version
	currentPinnedVersionMutex       sync.RWMutex
	currentPinnedVersionArgsForCall []struct {
	}
	currentPinnedVersionReturns struct {
		result1 atc.Version
	}
	currentPinnedVersionReturnsOnCall map[int]struct {
		result1 atc.Version
	}
	HasWebhookStub        func() bool
	hasWebhookMutex       sync.RWMutex
	hasWebhookArgsForCall []struct {
	}
	hasWebhookReturns struct {
		result1 bool
	}
	hasWebhookReturnsOnCall map[int]struct {
		result1 bool
	}
	LastCheckEndTimeStub        func() time.Time
	lastCheckEndTimeMutex       sync.RWMutex
	lastCheckEndTimeArgsForCall []struct {
	}
	lastCheckEndTimeReturns struct {
		result1 time.Time
	}
	lastCheckEndTimeReturnsOnCall map[int]struct {
		result1 time.Time
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	PipelineStub        func() (db.Pipeline, bool, error)
	pipelineMutex       sync.RWMutex
	pipelineArgsForCall []struct {
	}
	pipelineReturns struct {
		result1 db.Pipeline
		result2 bool
		result3 error
	}
	pipelineReturnsOnCall map[int]struct {
		result1 db.Pipeline
		result2 bool
		result3 error
	}
	PipelineIDStub        func() int
	pipelineIDMutex       sync.RWMutex
	pipelineIDArgsForCall []struct {
	}
	pipelineIDReturns struct {
		result1 int
	}
	pipelineIDReturnsOnCall map[int]struct {
		result1 int
	}
	PipelineInstanceVarsStub        func() atc.InstanceVars
	pipelineInstanceVarsMutex       sync.RWMutex
	pipelineInstanceVarsArgsForCall []struct {
	}
	pipelineInstanceVarsReturns struct {
		result1 atc.InstanceVars
	}
	pipelineInstanceVarsReturnsOnCall map[int]struct {
		result1 atc.InstanceVars
	}
	PipelineNameStub        func() string
	pipelineNameMutex       sync.RWMutex
	pipelineNameArgsForCall []struct {
	}
	pipelineNameReturns struct {
		result1 string
	}
	pipelineNameReturnsOnCall map[int]struct {
		result1 string
	}
	PipelineRefStub        func() atc.PipelineRef
	pipelineRefMutex       sync.RWMutex
	pipelineRefArgsForCall []struct {
	}
	pipelineRefReturns struct {
		result1 atc.PipelineRef
	}
	pipelineRefReturnsOnCall map[int]struct {
		result1 atc.PipelineRef
	}
	ResourceConfigScopeIDStub        func() int
	resourceConfigScopeIDMutex       sync.RWMutex
	resourceConfigScopeIDArgsForCall []struct {
	}
	resourceConfigScopeIDReturns struct {
		result1 int
	}
	resourceConfigScopeIDReturnsOnCall map[int]struct {
		result1 int
	}
	SourceStub        func() atc.Source
	sourceMutex       sync.RWMutex
	sourceArgsForCall []struct {
	}
	sourceReturns struct {
		result1 atc.Source
	}
	sourceReturnsOnCall map[int]struct {
		result1 atc.Source
	}
	TagsStub        func() atc.Tags
	tagsMutex       sync.RWMutex
	tagsArgsForCall []struct {
	}
	tagsReturns struct {
		result1 atc.Tags
	}
	tagsReturnsOnCall map[int]struct {
		result1 atc.Tags
	}
	TeamIDStub        func() int
	teamIDMutex       sync.RWMutex
	teamIDArgsForCall []struct {
	}
	teamIDReturns struct {
		result1 int
	}
	teamIDReturnsOnCall map[int]struct {
		result1 int
	}
	TeamNameStub        func() string
	teamNameMutex       sync.RWMutex
	teamNameArgsForCall []struct {
	}
	teamNameReturns struct {
		result1 string
	}
	teamNameReturnsOnCall map[int]struct {
		result1 string
	}
	TypeStub        func() string
	typeMutex       sync.RWMutex
	typeArgsForCall []struct {
	}
	typeReturns struct {
		result1 string
	}
	typeReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCheckable) CheckEvery() string {
	fake.checkEveryMutex.Lock()
	ret, specificReturn := fake.checkEveryReturnsOnCall[len(fake.checkEveryArgsForCall)]
	fake.checkEveryArgsForCall = append(fake.checkEveryArgsForCall, struct {
	}{})
	fake.recordInvocation("CheckEvery", []interface{}{})
	fake.checkEveryMutex.Unlock()
	if fake.CheckEveryStub != nil {
		return fake.CheckEveryStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkEveryReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) CheckEveryCallCount() int {
	fake.checkEveryMutex.RLock()
	defer fake.checkEveryMutex.RUnlock()
	return len(fake.checkEveryArgsForCall)
}

func (fake *FakeCheckable) CheckEveryCalls(stub func() string) {
	fake.checkEveryMutex.Lock()
	defer fake.checkEveryMutex.Unlock()
	fake.CheckEveryStub = stub
}

func (fake *FakeCheckable) CheckEveryReturns(result1 string) {
	fake.checkEveryMutex.Lock()
	defer fake.checkEveryMutex.Unlock()
	fake.CheckEveryStub = nil
	fake.checkEveryReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) CheckEveryReturnsOnCall(i int, result1 string) {
	fake.checkEveryMutex.Lock()
	defer fake.checkEveryMutex.Unlock()
	fake.CheckEveryStub = nil
	if fake.checkEveryReturnsOnCall == nil {
		fake.checkEveryReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.checkEveryReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) CheckPlan(arg1 atc.Version, arg2 time.Duration, arg3 db.ResourceTypes, arg4 atc.Source) atc.CheckPlan {
	fake.checkPlanMutex.Lock()
	ret, specificReturn := fake.checkPlanReturnsOnCall[len(fake.checkPlanArgsForCall)]
	fake.checkPlanArgsForCall = append(fake.checkPlanArgsForCall, struct {
		arg1 atc.Version
		arg2 time.Duration
		arg3 db.ResourceTypes
		arg4 atc.Source
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("CheckPlan", []interface{}{arg1, arg2, arg3, arg4})
	fake.checkPlanMutex.Unlock()
	if fake.CheckPlanStub != nil {
		return fake.CheckPlanStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkPlanReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) CheckPlanCallCount() int {
	fake.checkPlanMutex.RLock()
	defer fake.checkPlanMutex.RUnlock()
	return len(fake.checkPlanArgsForCall)
}

func (fake *FakeCheckable) CheckPlanCalls(stub func(atc.Version, time.Duration, db.ResourceTypes, atc.Source) atc.CheckPlan) {
	fake.checkPlanMutex.Lock()
	defer fake.checkPlanMutex.Unlock()
	fake.CheckPlanStub = stub
}

func (fake *FakeCheckable) CheckPlanArgsForCall(i int) (atc.Version, time.Duration, db.ResourceTypes, atc.Source) {
	fake.checkPlanMutex.RLock()
	defer fake.checkPlanMutex.RUnlock()
	argsForCall := fake.checkPlanArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeCheckable) CheckPlanReturns(result1 atc.CheckPlan) {
	fake.checkPlanMutex.Lock()
	defer fake.checkPlanMutex.Unlock()
	fake.CheckPlanStub = nil
	fake.checkPlanReturns = struct {
		result1 atc.CheckPlan
	}{result1}
}

func (fake *FakeCheckable) CheckPlanReturnsOnCall(i int, result1 atc.CheckPlan) {
	fake.checkPlanMutex.Lock()
	defer fake.checkPlanMutex.Unlock()
	fake.CheckPlanStub = nil
	if fake.checkPlanReturnsOnCall == nil {
		fake.checkPlanReturnsOnCall = make(map[int]struct {
			result1 atc.CheckPlan
		})
	}
	fake.checkPlanReturnsOnCall[i] = struct {
		result1 atc.CheckPlan
	}{result1}
}

func (fake *FakeCheckable) CheckTimeout() string {
	fake.checkTimeoutMutex.Lock()
	ret, specificReturn := fake.checkTimeoutReturnsOnCall[len(fake.checkTimeoutArgsForCall)]
	fake.checkTimeoutArgsForCall = append(fake.checkTimeoutArgsForCall, struct {
	}{})
	fake.recordInvocation("CheckTimeout", []interface{}{})
	fake.checkTimeoutMutex.Unlock()
	if fake.CheckTimeoutStub != nil {
		return fake.CheckTimeoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkTimeoutReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) CheckTimeoutCallCount() int {
	fake.checkTimeoutMutex.RLock()
	defer fake.checkTimeoutMutex.RUnlock()
	return len(fake.checkTimeoutArgsForCall)
}

func (fake *FakeCheckable) CheckTimeoutCalls(stub func() string) {
	fake.checkTimeoutMutex.Lock()
	defer fake.checkTimeoutMutex.Unlock()
	fake.CheckTimeoutStub = stub
}

func (fake *FakeCheckable) CheckTimeoutReturns(result1 string) {
	fake.checkTimeoutMutex.Lock()
	defer fake.checkTimeoutMutex.Unlock()
	fake.CheckTimeoutStub = nil
	fake.checkTimeoutReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) CheckTimeoutReturnsOnCall(i int, result1 string) {
	fake.checkTimeoutMutex.Lock()
	defer fake.checkTimeoutMutex.Unlock()
	fake.CheckTimeoutStub = nil
	if fake.checkTimeoutReturnsOnCall == nil {
		fake.checkTimeoutReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.checkTimeoutReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) CreateBuild(arg1 context.Context, arg2 bool) (db.Build, bool, error) {
	fake.createBuildMutex.Lock()
	ret, specificReturn := fake.createBuildReturnsOnCall[len(fake.createBuildArgsForCall)]
	fake.createBuildArgsForCall = append(fake.createBuildArgsForCall, struct {
		arg1 context.Context
		arg2 bool
	}{arg1, arg2})
	fake.recordInvocation("CreateBuild", []interface{}{arg1, arg2})
	fake.createBuildMutex.Unlock()
	if fake.CreateBuildStub != nil {
		return fake.CreateBuildStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createBuildReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCheckable) CreateBuildCallCount() int {
	fake.createBuildMutex.RLock()
	defer fake.createBuildMutex.RUnlock()
	return len(fake.createBuildArgsForCall)
}

func (fake *FakeCheckable) CreateBuildCalls(stub func(context.Context, bool) (db.Build, bool, error)) {
	fake.createBuildMutex.Lock()
	defer fake.createBuildMutex.Unlock()
	fake.CreateBuildStub = stub
}

func (fake *FakeCheckable) CreateBuildArgsForCall(i int) (context.Context, bool) {
	fake.createBuildMutex.RLock()
	defer fake.createBuildMutex.RUnlock()
	argsForCall := fake.createBuildArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCheckable) CreateBuildReturns(result1 db.Build, result2 bool, result3 error) {
	fake.createBuildMutex.Lock()
	defer fake.createBuildMutex.Unlock()
	fake.CreateBuildStub = nil
	fake.createBuildReturns = struct {
		result1 db.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCheckable) CreateBuildReturnsOnCall(i int, result1 db.Build, result2 bool, result3 error) {
	fake.createBuildMutex.Lock()
	defer fake.createBuildMutex.Unlock()
	fake.CreateBuildStub = nil
	if fake.createBuildReturnsOnCall == nil {
		fake.createBuildReturnsOnCall = make(map[int]struct {
			result1 db.Build
			result2 bool
			result3 error
		})
	}
	fake.createBuildReturnsOnCall[i] = struct {
		result1 db.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCheckable) CurrentPinnedVersion() atc.Version {
	fake.currentPinnedVersionMutex.Lock()
	ret, specificReturn := fake.currentPinnedVersionReturnsOnCall[len(fake.currentPinnedVersionArgsForCall)]
	fake.currentPinnedVersionArgsForCall = append(fake.currentPinnedVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("CurrentPinnedVersion", []interface{}{})
	fake.currentPinnedVersionMutex.Unlock()
	if fake.CurrentPinnedVersionStub != nil {
		return fake.CurrentPinnedVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.currentPinnedVersionReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) CurrentPinnedVersionCallCount() int {
	fake.currentPinnedVersionMutex.RLock()
	defer fake.currentPinnedVersionMutex.RUnlock()
	return len(fake.currentPinnedVersionArgsForCall)
}

func (fake *FakeCheckable) CurrentPinnedVersionCalls(stub func() atc.Version) {
	fake.currentPinnedVersionMutex.Lock()
	defer fake.currentPinnedVersionMutex.Unlock()
	fake.CurrentPinnedVersionStub = stub
}

func (fake *FakeCheckable) CurrentPinnedVersionReturns(result1 atc.Version) {
	fake.currentPinnedVersionMutex.Lock()
	defer fake.currentPinnedVersionMutex.Unlock()
	fake.CurrentPinnedVersionStub = nil
	fake.currentPinnedVersionReturns = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeCheckable) CurrentPinnedVersionReturnsOnCall(i int, result1 atc.Version) {
	fake.currentPinnedVersionMutex.Lock()
	defer fake.currentPinnedVersionMutex.Unlock()
	fake.CurrentPinnedVersionStub = nil
	if fake.currentPinnedVersionReturnsOnCall == nil {
		fake.currentPinnedVersionReturnsOnCall = make(map[int]struct {
			result1 atc.Version
		})
	}
	fake.currentPinnedVersionReturnsOnCall[i] = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeCheckable) HasWebhook() bool {
	fake.hasWebhookMutex.Lock()
	ret, specificReturn := fake.hasWebhookReturnsOnCall[len(fake.hasWebhookArgsForCall)]
	fake.hasWebhookArgsForCall = append(fake.hasWebhookArgsForCall, struct {
	}{})
	fake.recordInvocation("HasWebhook", []interface{}{})
	fake.hasWebhookMutex.Unlock()
	if fake.HasWebhookStub != nil {
		return fake.HasWebhookStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.hasWebhookReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) HasWebhookCallCount() int {
	fake.hasWebhookMutex.RLock()
	defer fake.hasWebhookMutex.RUnlock()
	return len(fake.hasWebhookArgsForCall)
}

func (fake *FakeCheckable) HasWebhookCalls(stub func() bool) {
	fake.hasWebhookMutex.Lock()
	defer fake.hasWebhookMutex.Unlock()
	fake.HasWebhookStub = stub
}

func (fake *FakeCheckable) HasWebhookReturns(result1 bool) {
	fake.hasWebhookMutex.Lock()
	defer fake.hasWebhookMutex.Unlock()
	fake.HasWebhookStub = nil
	fake.hasWebhookReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCheckable) HasWebhookReturnsOnCall(i int, result1 bool) {
	fake.hasWebhookMutex.Lock()
	defer fake.hasWebhookMutex.Unlock()
	fake.HasWebhookStub = nil
	if fake.hasWebhookReturnsOnCall == nil {
		fake.hasWebhookReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.hasWebhookReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCheckable) LastCheckEndTime() time.Time {
	fake.lastCheckEndTimeMutex.Lock()
	ret, specificReturn := fake.lastCheckEndTimeReturnsOnCall[len(fake.lastCheckEndTimeArgsForCall)]
	fake.lastCheckEndTimeArgsForCall = append(fake.lastCheckEndTimeArgsForCall, struct {
	}{})
	fake.recordInvocation("LastCheckEndTime", []interface{}{})
	fake.lastCheckEndTimeMutex.Unlock()
	if fake.LastCheckEndTimeStub != nil {
		return fake.LastCheckEndTimeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.lastCheckEndTimeReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) LastCheckEndTimeCallCount() int {
	fake.lastCheckEndTimeMutex.RLock()
	defer fake.lastCheckEndTimeMutex.RUnlock()
	return len(fake.lastCheckEndTimeArgsForCall)
}

func (fake *FakeCheckable) LastCheckEndTimeCalls(stub func() time.Time) {
	fake.lastCheckEndTimeMutex.Lock()
	defer fake.lastCheckEndTimeMutex.Unlock()
	fake.LastCheckEndTimeStub = stub
}

func (fake *FakeCheckable) LastCheckEndTimeReturns(result1 time.Time) {
	fake.lastCheckEndTimeMutex.Lock()
	defer fake.lastCheckEndTimeMutex.Unlock()
	fake.LastCheckEndTimeStub = nil
	fake.lastCheckEndTimeReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeCheckable) LastCheckEndTimeReturnsOnCall(i int, result1 time.Time) {
	fake.lastCheckEndTimeMutex.Lock()
	defer fake.lastCheckEndTimeMutex.Unlock()
	fake.LastCheckEndTimeStub = nil
	if fake.lastCheckEndTimeReturnsOnCall == nil {
		fake.lastCheckEndTimeReturnsOnCall = make(map[int]struct {
			result1 time.Time
		})
	}
	fake.lastCheckEndTimeReturnsOnCall[i] = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeCheckable) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nameReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeCheckable) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeCheckable) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) Pipeline() (db.Pipeline, bool, error) {
	fake.pipelineMutex.Lock()
	ret, specificReturn := fake.pipelineReturnsOnCall[len(fake.pipelineArgsForCall)]
	fake.pipelineArgsForCall = append(fake.pipelineArgsForCall, struct {
	}{})
	fake.recordInvocation("Pipeline", []interface{}{})
	fake.pipelineMutex.Unlock()
	if fake.PipelineStub != nil {
		return fake.PipelineStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.pipelineReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCheckable) PipelineCallCount() int {
	fake.pipelineMutex.RLock()
	defer fake.pipelineMutex.RUnlock()
	return len(fake.pipelineArgsForCall)
}

func (fake *FakeCheckable) PipelineCalls(stub func() (db.Pipeline, bool, error)) {
	fake.pipelineMutex.Lock()
	defer fake.pipelineMutex.Unlock()
	fake.PipelineStub = stub
}

func (fake *FakeCheckable) PipelineReturns(result1 db.Pipeline, result2 bool, result3 error) {
	fake.pipelineMutex.Lock()
	defer fake.pipelineMutex.Unlock()
	fake.PipelineStub = nil
	fake.pipelineReturns = struct {
		result1 db.Pipeline
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCheckable) PipelineReturnsOnCall(i int, result1 db.Pipeline, result2 bool, result3 error) {
	fake.pipelineMutex.Lock()
	defer fake.pipelineMutex.Unlock()
	fake.PipelineStub = nil
	if fake.pipelineReturnsOnCall == nil {
		fake.pipelineReturnsOnCall = make(map[int]struct {
			result1 db.Pipeline
			result2 bool
			result3 error
		})
	}
	fake.pipelineReturnsOnCall[i] = struct {
		result1 db.Pipeline
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCheckable) PipelineID() int {
	fake.pipelineIDMutex.Lock()
	ret, specificReturn := fake.pipelineIDReturnsOnCall[len(fake.pipelineIDArgsForCall)]
	fake.pipelineIDArgsForCall = append(fake.pipelineIDArgsForCall, struct {
	}{})
	fake.recordInvocation("PipelineID", []interface{}{})
	fake.pipelineIDMutex.Unlock()
	if fake.PipelineIDStub != nil {
		return fake.PipelineIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pipelineIDReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) PipelineIDCallCount() int {
	fake.pipelineIDMutex.RLock()
	defer fake.pipelineIDMutex.RUnlock()
	return len(fake.pipelineIDArgsForCall)
}

func (fake *FakeCheckable) PipelineIDCalls(stub func() int) {
	fake.pipelineIDMutex.Lock()
	defer fake.pipelineIDMutex.Unlock()
	fake.PipelineIDStub = stub
}

func (fake *FakeCheckable) PipelineIDReturns(result1 int) {
	fake.pipelineIDMutex.Lock()
	defer fake.pipelineIDMutex.Unlock()
	fake.PipelineIDStub = nil
	fake.pipelineIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeCheckable) PipelineIDReturnsOnCall(i int, result1 int) {
	fake.pipelineIDMutex.Lock()
	defer fake.pipelineIDMutex.Unlock()
	fake.PipelineIDStub = nil
	if fake.pipelineIDReturnsOnCall == nil {
		fake.pipelineIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.pipelineIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeCheckable) PipelineInstanceVars() atc.InstanceVars {
	fake.pipelineInstanceVarsMutex.Lock()
	ret, specificReturn := fake.pipelineInstanceVarsReturnsOnCall[len(fake.pipelineInstanceVarsArgsForCall)]
	fake.pipelineInstanceVarsArgsForCall = append(fake.pipelineInstanceVarsArgsForCall, struct {
	}{})
	fake.recordInvocation("PipelineInstanceVars", []interface{}{})
	fake.pipelineInstanceVarsMutex.Unlock()
	if fake.PipelineInstanceVarsStub != nil {
		return fake.PipelineInstanceVarsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pipelineInstanceVarsReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) PipelineInstanceVarsCallCount() int {
	fake.pipelineInstanceVarsMutex.RLock()
	defer fake.pipelineInstanceVarsMutex.RUnlock()
	return len(fake.pipelineInstanceVarsArgsForCall)
}

func (fake *FakeCheckable) PipelineInstanceVarsCalls(stub func() atc.InstanceVars) {
	fake.pipelineInstanceVarsMutex.Lock()
	defer fake.pipelineInstanceVarsMutex.Unlock()
	fake.PipelineInstanceVarsStub = stub
}

func (fake *FakeCheckable) PipelineInstanceVarsReturns(result1 atc.InstanceVars) {
	fake.pipelineInstanceVarsMutex.Lock()
	defer fake.pipelineInstanceVarsMutex.Unlock()
	fake.PipelineInstanceVarsStub = nil
	fake.pipelineInstanceVarsReturns = struct {
		result1 atc.InstanceVars
	}{result1}
}

func (fake *FakeCheckable) PipelineInstanceVarsReturnsOnCall(i int, result1 atc.InstanceVars) {
	fake.pipelineInstanceVarsMutex.Lock()
	defer fake.pipelineInstanceVarsMutex.Unlock()
	fake.PipelineInstanceVarsStub = nil
	if fake.pipelineInstanceVarsReturnsOnCall == nil {
		fake.pipelineInstanceVarsReturnsOnCall = make(map[int]struct {
			result1 atc.InstanceVars
		})
	}
	fake.pipelineInstanceVarsReturnsOnCall[i] = struct {
		result1 atc.InstanceVars
	}{result1}
}

func (fake *FakeCheckable) PipelineName() string {
	fake.pipelineNameMutex.Lock()
	ret, specificReturn := fake.pipelineNameReturnsOnCall[len(fake.pipelineNameArgsForCall)]
	fake.pipelineNameArgsForCall = append(fake.pipelineNameArgsForCall, struct {
	}{})
	fake.recordInvocation("PipelineName", []interface{}{})
	fake.pipelineNameMutex.Unlock()
	if fake.PipelineNameStub != nil {
		return fake.PipelineNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pipelineNameReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) PipelineNameCallCount() int {
	fake.pipelineNameMutex.RLock()
	defer fake.pipelineNameMutex.RUnlock()
	return len(fake.pipelineNameArgsForCall)
}

func (fake *FakeCheckable) PipelineNameCalls(stub func() string) {
	fake.pipelineNameMutex.Lock()
	defer fake.pipelineNameMutex.Unlock()
	fake.PipelineNameStub = stub
}

func (fake *FakeCheckable) PipelineNameReturns(result1 string) {
	fake.pipelineNameMutex.Lock()
	defer fake.pipelineNameMutex.Unlock()
	fake.PipelineNameStub = nil
	fake.pipelineNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) PipelineNameReturnsOnCall(i int, result1 string) {
	fake.pipelineNameMutex.Lock()
	defer fake.pipelineNameMutex.Unlock()
	fake.PipelineNameStub = nil
	if fake.pipelineNameReturnsOnCall == nil {
		fake.pipelineNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.pipelineNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) PipelineRef() atc.PipelineRef {
	fake.pipelineRefMutex.Lock()
	ret, specificReturn := fake.pipelineRefReturnsOnCall[len(fake.pipelineRefArgsForCall)]
	fake.pipelineRefArgsForCall = append(fake.pipelineRefArgsForCall, struct {
	}{})
	fake.recordInvocation("PipelineRef", []interface{}{})
	fake.pipelineRefMutex.Unlock()
	if fake.PipelineRefStub != nil {
		return fake.PipelineRefStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pipelineRefReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) PipelineRefCallCount() int {
	fake.pipelineRefMutex.RLock()
	defer fake.pipelineRefMutex.RUnlock()
	return len(fake.pipelineRefArgsForCall)
}

func (fake *FakeCheckable) PipelineRefCalls(stub func() atc.PipelineRef) {
	fake.pipelineRefMutex.Lock()
	defer fake.pipelineRefMutex.Unlock()
	fake.PipelineRefStub = stub
}

func (fake *FakeCheckable) PipelineRefReturns(result1 atc.PipelineRef) {
	fake.pipelineRefMutex.Lock()
	defer fake.pipelineRefMutex.Unlock()
	fake.PipelineRefStub = nil
	fake.pipelineRefReturns = struct {
		result1 atc.PipelineRef
	}{result1}
}

func (fake *FakeCheckable) PipelineRefReturnsOnCall(i int, result1 atc.PipelineRef) {
	fake.pipelineRefMutex.Lock()
	defer fake.pipelineRefMutex.Unlock()
	fake.PipelineRefStub = nil
	if fake.pipelineRefReturnsOnCall == nil {
		fake.pipelineRefReturnsOnCall = make(map[int]struct {
			result1 atc.PipelineRef
		})
	}
	fake.pipelineRefReturnsOnCall[i] = struct {
		result1 atc.PipelineRef
	}{result1}
}

func (fake *FakeCheckable) ResourceConfigScopeID() int {
	fake.resourceConfigScopeIDMutex.Lock()
	ret, specificReturn := fake.resourceConfigScopeIDReturnsOnCall[len(fake.resourceConfigScopeIDArgsForCall)]
	fake.resourceConfigScopeIDArgsForCall = append(fake.resourceConfigScopeIDArgsForCall, struct {
	}{})
	fake.recordInvocation("ResourceConfigScopeID", []interface{}{})
	fake.resourceConfigScopeIDMutex.Unlock()
	if fake.ResourceConfigScopeIDStub != nil {
		return fake.ResourceConfigScopeIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.resourceConfigScopeIDReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) ResourceConfigScopeIDCallCount() int {
	fake.resourceConfigScopeIDMutex.RLock()
	defer fake.resourceConfigScopeIDMutex.RUnlock()
	return len(fake.resourceConfigScopeIDArgsForCall)
}

func (fake *FakeCheckable) ResourceConfigScopeIDCalls(stub func() int) {
	fake.resourceConfigScopeIDMutex.Lock()
	defer fake.resourceConfigScopeIDMutex.Unlock()
	fake.ResourceConfigScopeIDStub = stub
}

func (fake *FakeCheckable) ResourceConfigScopeIDReturns(result1 int) {
	fake.resourceConfigScopeIDMutex.Lock()
	defer fake.resourceConfigScopeIDMutex.Unlock()
	fake.ResourceConfigScopeIDStub = nil
	fake.resourceConfigScopeIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeCheckable) ResourceConfigScopeIDReturnsOnCall(i int, result1 int) {
	fake.resourceConfigScopeIDMutex.Lock()
	defer fake.resourceConfigScopeIDMutex.Unlock()
	fake.ResourceConfigScopeIDStub = nil
	if fake.resourceConfigScopeIDReturnsOnCall == nil {
		fake.resourceConfigScopeIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.resourceConfigScopeIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeCheckable) Source() atc.Source {
	fake.sourceMutex.Lock()
	ret, specificReturn := fake.sourceReturnsOnCall[len(fake.sourceArgsForCall)]
	fake.sourceArgsForCall = append(fake.sourceArgsForCall, struct {
	}{})
	fake.recordInvocation("Source", []interface{}{})
	fake.sourceMutex.Unlock()
	if fake.SourceStub != nil {
		return fake.SourceStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sourceReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) SourceCallCount() int {
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	return len(fake.sourceArgsForCall)
}

func (fake *FakeCheckable) SourceCalls(stub func() atc.Source) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = stub
}

func (fake *FakeCheckable) SourceReturns(result1 atc.Source) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = nil
	fake.sourceReturns = struct {
		result1 atc.Source
	}{result1}
}

func (fake *FakeCheckable) SourceReturnsOnCall(i int, result1 atc.Source) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = nil
	if fake.sourceReturnsOnCall == nil {
		fake.sourceReturnsOnCall = make(map[int]struct {
			result1 atc.Source
		})
	}
	fake.sourceReturnsOnCall[i] = struct {
		result1 atc.Source
	}{result1}
}

func (fake *FakeCheckable) Tags() atc.Tags {
	fake.tagsMutex.Lock()
	ret, specificReturn := fake.tagsReturnsOnCall[len(fake.tagsArgsForCall)]
	fake.tagsArgsForCall = append(fake.tagsArgsForCall, struct {
	}{})
	fake.recordInvocation("Tags", []interface{}{})
	fake.tagsMutex.Unlock()
	if fake.TagsStub != nil {
		return fake.TagsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.tagsReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) TagsCallCount() int {
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	return len(fake.tagsArgsForCall)
}

func (fake *FakeCheckable) TagsCalls(stub func() atc.Tags) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = stub
}

func (fake *FakeCheckable) TagsReturns(result1 atc.Tags) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = nil
	fake.tagsReturns = struct {
		result1 atc.Tags
	}{result1}
}

func (fake *FakeCheckable) TagsReturnsOnCall(i int, result1 atc.Tags) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = nil
	if fake.tagsReturnsOnCall == nil {
		fake.tagsReturnsOnCall = make(map[int]struct {
			result1 atc.Tags
		})
	}
	fake.tagsReturnsOnCall[i] = struct {
		result1 atc.Tags
	}{result1}
}

func (fake *FakeCheckable) TeamID() int {
	fake.teamIDMutex.Lock()
	ret, specificReturn := fake.teamIDReturnsOnCall[len(fake.teamIDArgsForCall)]
	fake.teamIDArgsForCall = append(fake.teamIDArgsForCall, struct {
	}{})
	fake.recordInvocation("TeamID", []interface{}{})
	fake.teamIDMutex.Unlock()
	if fake.TeamIDStub != nil {
		return fake.TeamIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.teamIDReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) TeamIDCallCount() int {
	fake.teamIDMutex.RLock()
	defer fake.teamIDMutex.RUnlock()
	return len(fake.teamIDArgsForCall)
}

func (fake *FakeCheckable) TeamIDCalls(stub func() int) {
	fake.teamIDMutex.Lock()
	defer fake.teamIDMutex.Unlock()
	fake.TeamIDStub = stub
}

func (fake *FakeCheckable) TeamIDReturns(result1 int) {
	fake.teamIDMutex.Lock()
	defer fake.teamIDMutex.Unlock()
	fake.TeamIDStub = nil
	fake.teamIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeCheckable) TeamIDReturnsOnCall(i int, result1 int) {
	fake.teamIDMutex.Lock()
	defer fake.teamIDMutex.Unlock()
	fake.TeamIDStub = nil
	if fake.teamIDReturnsOnCall == nil {
		fake.teamIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.teamIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeCheckable) TeamName() string {
	fake.teamNameMutex.Lock()
	ret, specificReturn := fake.teamNameReturnsOnCall[len(fake.teamNameArgsForCall)]
	fake.teamNameArgsForCall = append(fake.teamNameArgsForCall, struct {
	}{})
	fake.recordInvocation("TeamName", []interface{}{})
	fake.teamNameMutex.Unlock()
	if fake.TeamNameStub != nil {
		return fake.TeamNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.teamNameReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) TeamNameCallCount() int {
	fake.teamNameMutex.RLock()
	defer fake.teamNameMutex.RUnlock()
	return len(fake.teamNameArgsForCall)
}

func (fake *FakeCheckable) TeamNameCalls(stub func() string) {
	fake.teamNameMutex.Lock()
	defer fake.teamNameMutex.Unlock()
	fake.TeamNameStub = stub
}

func (fake *FakeCheckable) TeamNameReturns(result1 string) {
	fake.teamNameMutex.Lock()
	defer fake.teamNameMutex.Unlock()
	fake.TeamNameStub = nil
	fake.teamNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) TeamNameReturnsOnCall(i int, result1 string) {
	fake.teamNameMutex.Lock()
	defer fake.teamNameMutex.Unlock()
	fake.TeamNameStub = nil
	if fake.teamNameReturnsOnCall == nil {
		fake.teamNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.teamNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) Type() string {
	fake.typeMutex.Lock()
	ret, specificReturn := fake.typeReturnsOnCall[len(fake.typeArgsForCall)]
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct {
	}{})
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if fake.TypeStub != nil {
		return fake.TypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.typeReturns
	return fakeReturns.result1
}

func (fake *FakeCheckable) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeCheckable) TypeCalls(stub func() string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = stub
}

func (fake *FakeCheckable) TypeReturns(result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) TypeReturnsOnCall(i int, result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	if fake.typeReturnsOnCall == nil {
		fake.typeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.typeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCheckable) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkEveryMutex.RLock()
	defer fake.checkEveryMutex.RUnlock()
	fake.checkPlanMutex.RLock()
	defer fake.checkPlanMutex.RUnlock()
	fake.checkTimeoutMutex.RLock()
	defer fake.checkTimeoutMutex.RUnlock()
	fake.createBuildMutex.RLock()
	defer fake.createBuildMutex.RUnlock()
	fake.currentPinnedVersionMutex.RLock()
	defer fake.currentPinnedVersionMutex.RUnlock()
	fake.hasWebhookMutex.RLock()
	defer fake.hasWebhookMutex.RUnlock()
	fake.lastCheckEndTimeMutex.RLock()
	defer fake.lastCheckEndTimeMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.pipelineMutex.RLock()
	defer fake.pipelineMutex.RUnlock()
	fake.pipelineIDMutex.RLock()
	defer fake.pipelineIDMutex.RUnlock()
	fake.pipelineInstanceVarsMutex.RLock()
	defer fake.pipelineInstanceVarsMutex.RUnlock()
	fake.pipelineNameMutex.RLock()
	defer fake.pipelineNameMutex.RUnlock()
	fake.pipelineRefMutex.RLock()
	defer fake.pipelineRefMutex.RUnlock()
	fake.resourceConfigScopeIDMutex.RLock()
	defer fake.resourceConfigScopeIDMutex.RUnlock()
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	fake.teamIDMutex.RLock()
	defer fake.teamIDMutex.RUnlock()
	fake.teamNameMutex.RLock()
	defer fake.teamNameMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCheckable) recordInvocation(key string, args []interface{}) {
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

var _ db.Checkable = new(FakeCheckable)

// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"autoscaler/models"
	"autoscaler/scalingengine"
	"sync"
)

type FakeScalingEngine struct {
	ScaleStub        func(appId string, trigger *models.Trigger) (int, error)
	scaleMutex       sync.RWMutex
	scaleArgsForCall []struct {
		appId   string
		trigger *models.Trigger
	}
	scaleReturns struct {
		result1 int
		result2 error
	}
	scaleReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	ComputeNewInstancesStub        func(currentInstances int, adjustment string) (int, error)
	computeNewInstancesMutex       sync.RWMutex
	computeNewInstancesArgsForCall []struct {
		currentInstances int
		adjustment       string
	}
	computeNewInstancesReturns struct {
		result1 int
		result2 error
	}
	computeNewInstancesReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	SetActiveScheduleStub        func(appId string, schedule *models.ActiveSchedule) error
	setActiveScheduleMutex       sync.RWMutex
	setActiveScheduleArgsForCall []struct {
		appId    string
		schedule *models.ActiveSchedule
	}
	setActiveScheduleReturns struct {
		result1 error
	}
	setActiveScheduleReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveActiveScheduleStub        func(appId string, scheduleId string) error
	removeActiveScheduleMutex       sync.RWMutex
	removeActiveScheduleArgsForCall []struct {
		appId      string
		scheduleId string
	}
	removeActiveScheduleReturns struct {
		result1 error
	}
	removeActiveScheduleReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeScalingEngine) Scale(appId string, trigger *models.Trigger) (int, error) {
	fake.scaleMutex.Lock()
	ret, specificReturn := fake.scaleReturnsOnCall[len(fake.scaleArgsForCall)]
	fake.scaleArgsForCall = append(fake.scaleArgsForCall, struct {
		appId   string
		trigger *models.Trigger
	}{appId, trigger})
	fake.recordInvocation("Scale", []interface{}{appId, trigger})
	fake.scaleMutex.Unlock()
	if fake.ScaleStub != nil {
		return fake.ScaleStub(appId, trigger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.scaleReturns.result1, fake.scaleReturns.result2
}

func (fake *FakeScalingEngine) ScaleCallCount() int {
	fake.scaleMutex.RLock()
	defer fake.scaleMutex.RUnlock()
	return len(fake.scaleArgsForCall)
}

func (fake *FakeScalingEngine) ScaleArgsForCall(i int) (string, *models.Trigger) {
	fake.scaleMutex.RLock()
	defer fake.scaleMutex.RUnlock()
	return fake.scaleArgsForCall[i].appId, fake.scaleArgsForCall[i].trigger
}

func (fake *FakeScalingEngine) ScaleReturns(result1 int, result2 error) {
	fake.ScaleStub = nil
	fake.scaleReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeScalingEngine) ScaleReturnsOnCall(i int, result1 int, result2 error) {
	fake.ScaleStub = nil
	if fake.scaleReturnsOnCall == nil {
		fake.scaleReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.scaleReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeScalingEngine) ComputeNewInstances(currentInstances int, adjustment string) (int, error) {
	fake.computeNewInstancesMutex.Lock()
	ret, specificReturn := fake.computeNewInstancesReturnsOnCall[len(fake.computeNewInstancesArgsForCall)]
	fake.computeNewInstancesArgsForCall = append(fake.computeNewInstancesArgsForCall, struct {
		currentInstances int
		adjustment       string
	}{currentInstances, adjustment})
	fake.recordInvocation("ComputeNewInstances", []interface{}{currentInstances, adjustment})
	fake.computeNewInstancesMutex.Unlock()
	if fake.ComputeNewInstancesStub != nil {
		return fake.ComputeNewInstancesStub(currentInstances, adjustment)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.computeNewInstancesReturns.result1, fake.computeNewInstancesReturns.result2
}

func (fake *FakeScalingEngine) ComputeNewInstancesCallCount() int {
	fake.computeNewInstancesMutex.RLock()
	defer fake.computeNewInstancesMutex.RUnlock()
	return len(fake.computeNewInstancesArgsForCall)
}

func (fake *FakeScalingEngine) ComputeNewInstancesArgsForCall(i int) (int, string) {
	fake.computeNewInstancesMutex.RLock()
	defer fake.computeNewInstancesMutex.RUnlock()
	return fake.computeNewInstancesArgsForCall[i].currentInstances, fake.computeNewInstancesArgsForCall[i].adjustment
}

func (fake *FakeScalingEngine) ComputeNewInstancesReturns(result1 int, result2 error) {
	fake.ComputeNewInstancesStub = nil
	fake.computeNewInstancesReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeScalingEngine) ComputeNewInstancesReturnsOnCall(i int, result1 int, result2 error) {
	fake.ComputeNewInstancesStub = nil
	if fake.computeNewInstancesReturnsOnCall == nil {
		fake.computeNewInstancesReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.computeNewInstancesReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeScalingEngine) SetActiveSchedule(appId string, schedule *models.ActiveSchedule) error {
	fake.setActiveScheduleMutex.Lock()
	ret, specificReturn := fake.setActiveScheduleReturnsOnCall[len(fake.setActiveScheduleArgsForCall)]
	fake.setActiveScheduleArgsForCall = append(fake.setActiveScheduleArgsForCall, struct {
		appId    string
		schedule *models.ActiveSchedule
	}{appId, schedule})
	fake.recordInvocation("SetActiveSchedule", []interface{}{appId, schedule})
	fake.setActiveScheduleMutex.Unlock()
	if fake.SetActiveScheduleStub != nil {
		return fake.SetActiveScheduleStub(appId, schedule)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setActiveScheduleReturns.result1
}

func (fake *FakeScalingEngine) SetActiveScheduleCallCount() int {
	fake.setActiveScheduleMutex.RLock()
	defer fake.setActiveScheduleMutex.RUnlock()
	return len(fake.setActiveScheduleArgsForCall)
}

func (fake *FakeScalingEngine) SetActiveScheduleArgsForCall(i int) (string, *models.ActiveSchedule) {
	fake.setActiveScheduleMutex.RLock()
	defer fake.setActiveScheduleMutex.RUnlock()
	return fake.setActiveScheduleArgsForCall[i].appId, fake.setActiveScheduleArgsForCall[i].schedule
}

func (fake *FakeScalingEngine) SetActiveScheduleReturns(result1 error) {
	fake.SetActiveScheduleStub = nil
	fake.setActiveScheduleReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScalingEngine) SetActiveScheduleReturnsOnCall(i int, result1 error) {
	fake.SetActiveScheduleStub = nil
	if fake.setActiveScheduleReturnsOnCall == nil {
		fake.setActiveScheduleReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setActiveScheduleReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeScalingEngine) RemoveActiveSchedule(appId string, scheduleId string) error {
	fake.removeActiveScheduleMutex.Lock()
	ret, specificReturn := fake.removeActiveScheduleReturnsOnCall[len(fake.removeActiveScheduleArgsForCall)]
	fake.removeActiveScheduleArgsForCall = append(fake.removeActiveScheduleArgsForCall, struct {
		appId      string
		scheduleId string
	}{appId, scheduleId})
	fake.recordInvocation("RemoveActiveSchedule", []interface{}{appId, scheduleId})
	fake.removeActiveScheduleMutex.Unlock()
	if fake.RemoveActiveScheduleStub != nil {
		return fake.RemoveActiveScheduleStub(appId, scheduleId)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeActiveScheduleReturns.result1
}

func (fake *FakeScalingEngine) RemoveActiveScheduleCallCount() int {
	fake.removeActiveScheduleMutex.RLock()
	defer fake.removeActiveScheduleMutex.RUnlock()
	return len(fake.removeActiveScheduleArgsForCall)
}

func (fake *FakeScalingEngine) RemoveActiveScheduleArgsForCall(i int) (string, string) {
	fake.removeActiveScheduleMutex.RLock()
	defer fake.removeActiveScheduleMutex.RUnlock()
	return fake.removeActiveScheduleArgsForCall[i].appId, fake.removeActiveScheduleArgsForCall[i].scheduleId
}

func (fake *FakeScalingEngine) RemoveActiveScheduleReturns(result1 error) {
	fake.RemoveActiveScheduleStub = nil
	fake.removeActiveScheduleReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScalingEngine) RemoveActiveScheduleReturnsOnCall(i int, result1 error) {
	fake.RemoveActiveScheduleStub = nil
	if fake.removeActiveScheduleReturnsOnCall == nil {
		fake.removeActiveScheduleReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeActiveScheduleReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeScalingEngine) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.scaleMutex.RLock()
	defer fake.scaleMutex.RUnlock()
	fake.computeNewInstancesMutex.RLock()
	defer fake.computeNewInstancesMutex.RUnlock()
	fake.setActiveScheduleMutex.RLock()
	defer fake.setActiveScheduleMutex.RUnlock()
	fake.removeActiveScheduleMutex.RLock()
	defer fake.removeActiveScheduleMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeScalingEngine) recordInvocation(key string, args []interface{}) {
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

var _ scalingengine.ScalingEngine = new(FakeScalingEngine)

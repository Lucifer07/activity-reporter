// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	activityreporter "github.com/Lucifer07/activity-reporter/activityreporter"
	mock "github.com/stretchr/testify/mock"
)

// publisher is an autogenerated mock type for the publisher type
type publisher struct {
	mock.Mock
}

// NotifyObserverLike provides a mock function with given fields: ownerPhoto
func (_m *publisher) NotifyObserverLike(ownerPhoto *activityreporter.User) {
	_m.Called(ownerPhoto)
}

// NotifyObserverUploadPhoto provides a mock function with given fields:
func (_m *publisher) NotifyObserverUploadPhoto() {
	_m.Called()
}

// UpdateFollower provides a mock function with given fields: _a0
func (_m *publisher) UpdateFollower(_a0 activityreporter.Observer) {
	_m.Called(_a0)
}

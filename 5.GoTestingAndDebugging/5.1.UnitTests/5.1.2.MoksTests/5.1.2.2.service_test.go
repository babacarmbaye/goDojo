package myservice

import (
	"debugging/5.1.UnitTests/5.1.2.MoksTests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoSomething(t *testing.T) {
	// Create an instance of our test object
	apiMock := new(mocks.ExternalAPIMock)

	// Setup expectations
	apiMock.On("FetchData", "foo").Return("bar", nil)

	// Create an instance of MyService using the mock
	myService := New(apiMock)

	// Call the method we want to test
	result, err := myService.DoSomething("foo")

	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, "bar", result)
	apiMock.AssertExpectations(t)
}

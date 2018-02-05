package util

import (
	"testing"
	"reflect"

	"github.com/stretchr/testify/assert"
)

func TestCreateContext(t *testing.T){
	contextName := "test-context"
	Create(contextName)

	context := getContext(contextName)

	assert.Equal(t, contextName, context.Name)

	RemoveContext(contextName)
}

func TestSaveGetContext(t *testing.T) {
	testContext := Context{
		Name: "testing",
		Containers: []Container{
			Container{
				Name: "test-container",
				Image: "test:latest",
				Port: "101:10101",
			},
		},
	}

	saveContext(testContext)

	actualContext := getContext(testContext.Name)
	
	assert.True(t, reflect.DeepEqual(testContext, actualContext))

	RemoveContext(testContext.Name)
}
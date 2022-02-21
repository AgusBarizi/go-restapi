package test

import (
	"github.com/stretchr/testify/assert"
	"m3gaplazma/go-restapi/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializeService(false)
	//fmt.Println(err)
	//fmt.Println(simpleService.SimpleRepository)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializeService(true)
	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
}

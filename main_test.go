package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Important pre for testing")
	exitVal := m.Run()
	fmt.Println("Important post for testing")

	os.Exit(exitVal)
}

func TestNeedRootAccess(t *testing.T) {
	NeedRootAccess()
	_, err := os.Stat("/you_can_remove_it")
	assert.NotNil(t, err)
	fmt.Println("TestNeedRootAccess(): success")
}

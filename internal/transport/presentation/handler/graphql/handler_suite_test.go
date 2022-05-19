package graphql_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Case struct {
	Context    string
	SetUp      func(t *testing.T)
	StatusCode int
	WantError  bool
	TearDown   func(t *testing.T)
}

type Cases []Case

type ReturnArgs [][]interface{}

type TestSuite struct {
	suite.Suite
	Cases Cases
}
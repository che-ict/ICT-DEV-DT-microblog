package repositories_test

import (
	"github.com/che-ict/DEV-DT-Microblog/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser_NoDuplicates(t *testing.T) {
	err1 := repositories.CreateUser("flipdebeer", "flip4life", "Flip")
	err2 := repositories.CreateUser("flipdebeer", "flip4life", "Flip")
	assert.NoError(t, err1)
	assert.Error(t, err2)
}

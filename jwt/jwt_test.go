package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenParse(t *testing.T) {
	secret := "12345"
	token, err := GenToken(time.Second, "test", secret)
	assert.NoError(t, err)
	j, err := ParseToken(token, secret)
	assert.NoError(t, err)
	assert.Equal(t, j.Issuer, "test")
}

func TestGenParse_Fail(t *testing.T) {
	secret := "12345"
	token, err := GenToken(time.Second, "test", secret)
	assert.NoError(t, err)
	time.Sleep(time.Second * 2)
	j, err := ParseToken(token, secret)
	assert.Error(t, err)
	assert.Nil(t, j)
	//assert.Equal(t, j.ExpiresAt , "test")
}

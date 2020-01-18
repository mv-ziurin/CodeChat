package validate

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	assert := assert.New(t)

	assert.True(ValidateEmail("email@mail.ru"))
	assert.True(ValidateEmail("email@192.168.0.1"))
	assert.True(ValidateEmail("hello_world@mail.ru"))

	// see please https://en.wikipedia.org/wiki/Email_address#Examples
	assert.True(ValidateEmail("a@b"))

	assert.False(ValidateEmail("sometext"))
	assert.False(ValidateEmail("email@@mail.ru"))
	assert.False(ValidateEmail("q@."))
}

func TestValidatePasswordForRegister(t *testing.T) {
	assert := assert.New(t)

	assert.True(ValidatePasswordForRegister("H@lloworld00"))

	assert.False(ValidatePasswordForRegister("test"))
	assert.False(ValidatePasswordForRegister("bad_password "))
	assert.False(ValidatePasswordForRegister("Helloworld"))
	assert.False(ValidatePasswordForRegister("Helloworld00"))
}

func TestUsername(t *testing.T) {
	assert := assert.New(t)

	assert.True(ValidateUsername("helloler"))

	assert.False(ValidateUsername("na"))
	assert.False(ValidateUsername("test~bad"))
	assert.False(ValidateUsername("longlonglonglonglonglonglonglonglonglonglonglonglonglong"))
}

func TestPassword(t *testing.T) {
	assert := assert.New(t)

	assert.True(ValidatePassword("goodpassword"))
	assert.True(ValidatePassword("b"))

	assert.False(ValidatePassword("longlonglonglonglonglonglonglonglonglonglonglonglonglong"))
}
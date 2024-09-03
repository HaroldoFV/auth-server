package domain

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewUser(t *testing.T) {
	t.Run("Valid User", func(t *testing.T) {
		user, err := NewUser("John Doe", "j@j.com", "12345678")
		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.NotEmpty(t, user.id)
		assert.NotEmpty(t, user.name)
		assert.Equal(t, "John Doe", user.name)
		assert.Equal(t, "j@j.com", user.email)
	})

	t.Run("Invalid Name", func(t *testing.T) {
		user, err := NewUser("", "j@j.com", "12345678")
		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "name cannot be empty", err.Error())

	})

	t.Run("Name Too Long", func(t *testing.T) {
		longName := string(make([]byte, 101))
		_, err := NewUser(longName, "j@j.com", "12345678")
		require.EqualError(t, err, "name cannot be longer than 100 characters")
	})

	t.Run("Invalid Email", func(t *testing.T) {
		user, err := NewUser("John Doe", "", "12345678")
		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "email cannot be empty", err.Error())

		user, err = NewUser("John Doe", "aaa@", "12345678")
		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "invalid email format", err.Error())
	})

	t.Run("Validate Password", func(t *testing.T) {
		user, err := NewUser("John Doe", "j@j.com", "12345678")
		assert.Nil(t, err)
		assert.True(t, user.ValidatePassword("12345678"))
		assert.False(t, user.ValidatePassword("12345"))
		assert.NotEqual(t, "12345", user.password)
	})

	t.Run("Invalid Password", func(t *testing.T) {
		user, err := NewUser("John Doe", "j@j.com", "")
		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "password must be at least 8 characters long", err.Error())
	})
}

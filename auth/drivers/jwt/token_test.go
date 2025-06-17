package jwt_auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	{
		token, err := NewToken(test_data.raw_valid_token, test_data.pkl)
		assert.NoError(t, err)
		assert.NotNil(t, token)
	}
	{
		token, err := NewToken(test_data.raw_expired_token, test_data.pkl)
		assert.Error(t, err)
		assert.Nil(t, token)
	}
	{
		token, err := NewToken(test_data.raw_invalid_token, test_data.pkl)
		assert.Error(t, err)
		assert.Nil(t, token)
	}
}

func TestIsValid(t *testing.T) {
	{
		valid_token, err := NewToken(test_data.raw_valid_token, test_data.pkl)
		assert.NoError(t, err)
		assert.True(t, valid_token.IsValid())
	}
	{
		invalid_token := Token{}
		assert.False(t, invalid_token.IsValid())
	}
}

func TestGetUserId(t *testing.T) {
	{
		valid_token, err := NewToken(test_data.raw_valid_token, test_data.pkl)
		assert.NoError(t, err)
		uid, err := valid_token.GetUserId()
		assert.NoError(t, err)
		assert.Equal(t, uid, int64(0))
	}
	{
		invalid_token := Token{}
		_, err := invalid_token.GetUserId()
		assert.Error(t, err)
	}
}

package local_pem_loader

import (
	"crypto/ed25519"
	"path/filepath"
	"testing"

	"github.com/qcrg/silver_broccoli/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockConfig struct {
	mock.Mock
}

func (t *MockConfig) GetFilePath() string {
	args := t.Called()
	return args.String(0)
}

func new_mock_config(key_name string) *MockConfig {
	cfg := MockConfig{}
	cfg.On("GetFilePath").Return(get_key_path(key_name))
	return &cfg
}

func new_default_key_loader(t *testing.T) KeyLoader {
	kl, err := NewKeyLoader(new_mock_config("ed25519_0.pub"))
	require.NoError(t, err)
	return *kl.(*KeyLoader)
}

func get_key_path(name string) string {
	return filepath.Join(utils.GetProjectDir(), "testdata/keys", name)
}

func TestNewKeyLoader(t *testing.T) {
	{
		_, err := NewKeyLoader(new_mock_config("iufvhpewqiujrncqjewcjqwc.pub"))
		assert.Error(t, err)
	}

	{
		kl, err := NewKeyLoader(new_mock_config("ed25519_0.pub"))
		assert.NoError(t, err)
		assert.NotNil(t, kl)
	}

	{
		_, err := NewKeyLoader(new_mock_config("empty"))
		assert.Error(t, err)
	}

	{
		_, err := NewKeyLoader(new_mock_config("invalid"))
		assert.Error(t, err)
	}
}

func TestKey(t *testing.T) {
	kl := new_default_key_loader(t)

	key, err := kl.Key(nil)
	assert.NoError(t, err)
	assert.NotNil(t, key)
	assert.IsType(t, ed25519.PublicKey{}, key)
}

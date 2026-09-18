package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func tempRoot(t *testing.T) string {
	t.Helper()
	root, err := os.MkdirTemp("", "garage-test-*")
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := os.RemoveAll(root); err != nil {
			t.Errorf("cleanup %s: %v", root, err)
		}
	})
	return root
}

func TestMakeDir(t *testing.T) {
	t.Run("creates missing directory", func(t *testing.T) {
		dir := filepath.Join(tempRoot(t), "cover")
		require.NoError(t, MakeDir(dir))

		info, err := os.Stat(dir)
		require.NoError(t, err)
		assert.True(t, info.IsDir())
	})

	t.Run("creates nested directory", func(t *testing.T) {
		dir := filepath.Join(tempRoot(t), "javbus", "cover")
		require.NoError(t, MakeDir(dir))

		info, err := os.Stat(dir)
		require.NoError(t, err)
		assert.True(t, info.IsDir())
	})

	t.Run("succeeds when directory already exists", func(t *testing.T) {
		dir := filepath.Join(tempRoot(t), "cover")
		require.NoError(t, MakeDir(dir))
		require.NoError(t, MakeDir(dir))
	})

	t.Run("errors when path is an existing file", func(t *testing.T) {
		file := filepath.Join(tempRoot(t), "not-a-dir")
		require.NoError(t, os.WriteFile(file, []byte("x"), 0o644))

		err := MakeDir(file)
		require.EqualError(t, err, "file name exist, but not dir")
	})
}

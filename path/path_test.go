package path

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetAppPath tests the GetAppPath function.
func TestGetAppPath(t *testing.T) {
	// In a test environment, GetAppPath should return the current working directory.
	path, err := GetAppPath()
	require.NoError(t, err, "GetAppPath should not return an error during tests")

	wd, err := os.Getwd()
	require.NoError(t, err, "os.Getwd should not fail")

	assert.Equal(t, wd, path, "In a test environment, GetAppPath should return the same as os.Getwd()")
}

// TestGetProjRootPath tests the GetProjRootPath function.
func TestGetProjRootPath(t *testing.T) {
	// Since we are in the 'path' package of the 'infa' project,
	// GetProjRootPath should correctly identify the project root.
	rootPath, err := GetProjRootPath(InfaProjName)
	require.NoError(t, err, "GetProjRootPath should not return an error")

	// The result should end with the project name.
	require.True(t, strings.HasSuffix(rootPath, InfaProjName), "The project root path should end with the project name")

	// To be more robust, we can check if a known file (like go.mod) exists relative to the found root.
	goModPath := filepath.Join(rootPath, "go.mod")
	_, err = os.Stat(goModPath)
	assert.NoError(t, err, "go.mod should exist at the calculated project root")
}

// TestGetProjRootPath_NoProjName tests that GetProjRootPath returns the app path when projName is not found.
func TestGetProjRootPath_NoProjName(t *testing.T) {
	// Use a project name that doesn't exist in the path.
	path, err := GetProjRootPath("a-project-that-does-not-exist")
	require.NoError(t, err)

	// It should fall back to the app path, which in a test is the working directory.
	wd, _ := os.Getwd()
	assert.Equal(t, wd, path, "Should return working directory when project name is not found")
}

// TestLocFilePath tests locating a file within the project.
func TestLocFilePath(t *testing.T) {
	// We'll try to locate this test file itself.
	filePath, err := LocFilePath(InfaProjName, "path", "path_test.go")
	require.NoError(t, err)

	_, err = os.Stat(filePath)
	assert.NoError(t, err, "The file path located by LocFilePath should exist")
}

// TestGetInfaPath tests the GetInfaPath convenience function.
func TestGetInfaPath(t *testing.T) {
	// We'll try to locate the go.mod file using GetInfaPath.
	filePath, err := GetInfaPath("go.mod")
	require.NoError(t, err)

	_, err = os.Stat(filePath)
	assert.NoError(t, err, "The file path located by GetInfaPath should exist")
}

// TestGetFilePath tests the basic GetFilePath utility.
func TestGetFilePath(t *testing.T) {
	expected := filepath.Join("a", "b", "c")
	actual := GetFilePath("a", "b", "c")
	assert.Equal(t, expected, actual)
}

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

//go:build !windows
// +build !windows

package globpath

import (
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompileAndMatch(t *testing.T) {
	dir := getTestdataDir()
	// test super asterisk
	g1, err := Compile(dir + "/**")
	require.NoError(t, err)
	// test single asterisk
	g2, err := Compile(dir + "/*.log")
	require.NoError(t, err)
	// test no meta characters (file exists)
	g3, err := Compile(dir + "/log1.log")
	require.NoError(t, err)
	// test file that doesn't exist
	g4, err := Compile(dir + "/i_dont_exist.log")
	require.NoError(t, err)
	// test super asterisk that doesn't exist
	g5, err := Compile(dir + "/dir_doesnt_exist/**")
	require.NoError(t, err)

	matches := g1.Match()
	assert.Len(t, matches, 6)
	matches = g2.Match()
	assert.Len(t, matches, 2)
	matches = g3.Match()
	assert.Len(t, matches, 1)
	matches = g4.Match()
	assert.Len(t, matches, 0)
	matches = g5.Match()
	assert.Len(t, matches, 0)
}

func TestFindRootDir(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"/var/log/telegraf.conf", "/var/log"},
		{"/home/**", "/home"},
		{"/home/*/**", "/home"},
		{"/lib/share/*/*/**.txt", "/lib/share"},
	}

	for _, test := range tests {
		actual := findRootDir(test.input)
		assert.Equal(t, test.output, actual)
	}
}

func TestFindNestedTextFile(t *testing.T) {
	dir := getTestdataDir()
	// test super asterisk
	g1, err := Compile(dir + "/**.txt")
	require.NoError(t, err)

	matches := g1.Match()
	assert.Len(t, matches, 1)
}

func getTestdataDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return strings.Replace(filename, "globpath_test.go", "testdata", 1)
}

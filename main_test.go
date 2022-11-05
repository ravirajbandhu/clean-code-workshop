package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func Test_traverseDir(t *testing.T) {
	type args struct {
		hashes     map[string]string
		duplicates map[string]string
		dupeSize   *int64
		entries    []os.FileInfo
		directory  string
	}
	type expectedValues struct {
		hashCount     int
		duplicateCount int
		dupeSize   int64
	}

	path := "/Users/sarthakgupta/go/src/clean-code-workshop/src"
	dirFiles, err := ioutil.ReadDir(path)
	size := int64(0)
	assert.NoError(t, err)
	tests := []struct {
		name string
		args args
		expectedValues expectedValues
	}{
		// TODO: Add test cases.
		{
			name: "No file entries in the directory",
			args: args{
				hashes:     nil,
				duplicates: nil,
				dupeSize:   nil,
				entries:    nil,
				directory:  "",
			},
			expectedValues: expectedValues{
				hashCount: 0,
				duplicateCount: 0,
				dupeSize: 0,
			},
		},

		{
			name: "Check Duplicates",
			args: args{
				hashes:     map[string]string{},
				duplicates: map[string]string{},
				dupeSize:   &size,
				entries:    dirFiles,
				directory:  path,
			},
			expectedValues: expectedValues{
				hashCount: 2,
				duplicateCount: 1,
				dupeSize: int64(22),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			traverseDir(tt.args.hashes, tt.args.duplicates, tt.args.dupeSize, tt.args.entries, tt.args.directory)
			//fmt.Println(*tt.args.dupeSize)
			//assert.Nil(t, tt.args.dupeSize)
			if tt.args.dupeSize != nil {
				assert.Equal(t, tt.expectedValues.dupeSize, *tt.args.dupeSize)
			}
			assert.Equal(t, tt.expectedValues.duplicateCount, len(tt.args.duplicates))
			assert.Equal(t, tt.expectedValues.hashCount, len(tt.args.hashes))
		})
	}
}

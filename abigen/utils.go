package abigen

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func Exit(msg string, err error) {
	if err != nil {
		fmt.Println(msg+":", err)
	} else {
		fmt.Println(msg)
	}
	os.Exit(1)
}

// GetProjectRoot returns the root of the chainlink project
func GetProjectRoot() (rootPath string) {
	root, err := os.Getwd()
	if err != nil {
		Exit("could not get current working directory while seeking project root",
			err)
	}
	for root != "/" { // Walk up path to find dir containing go.mod
		if _, err := os.Stat(filepath.Join(root, "go.mod")); os.IsNotExist(err) {
			root = filepath.Dir(root)
		} else {
			return root
		}
	}
	Exit("could not find project root", nil)
	panic("can't get here")
}

func TempDir(dirPrefix string) (string, func()) {
	tmpDir, err := ioutil.TempDir("", dirPrefix+"-contractWrapper")
	if err != nil {
		Exit("failed to create temporary working directory", err)
	}
	return tmpDir, func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			fmt.Println("failure while cleaning up temporary working directory:", err)
		}
	}
}

func DeepCopyLog(l types.Log) types.Log {
	var cpy types.Log
	cpy.Address = l.Address
	if l.Topics != nil {
		cpy.Topics = make([]common.Hash, len(l.Topics))
		copy(cpy.Topics, l.Topics)
	}
	if l.Data != nil {
		cpy.Data = make([]byte, len(l.Data))
		copy(cpy.Data, l.Data)
	}
	cpy.BlockNumber = l.BlockNumber
	cpy.TxHash = l.TxHash
	cpy.TxIndex = l.TxIndex
	cpy.BlockHash = l.BlockHash
	cpy.Index = l.Index
	cpy.Removed = l.Removed
	return cpy
}

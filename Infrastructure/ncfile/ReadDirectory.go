package ncfile

import (
	"fmt"
	"io/ioutil"
	"os"

	"nc-script-converter/Domain/alterationncscript"
)

type NcScriptDir struct{}

func NewNcScriptDir() *alterationncscript.DirViewer {
	var obj alterationncscript.DirViewer = &NcScriptDir{}
	return &obj
}

func (n *NcScriptDir) FetchDir(path string) ([]string, error) {
	if len(path) == 0 {
		return nil, fmt.Errorf("引数が空です")
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("ディレクトリ取得に失敗しました error:%v", err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			// ディレクトリは無視
			continue
		}
		paths = append(paths, file.Name())
	}

	return paths, nil
}

func (n *NcScriptDir) DirExist(path string) bool {
	if len(path) <= 0 {
		return false
	} else if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return false
	}
	return true
}

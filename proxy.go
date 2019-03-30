// プロキシファイルに#を加える関数と、抜き取る関数
package main

import (
	"io/ioutil"
	"strings"
)

func proxyAddComment(filename string) (err error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//"kanazawa-it.ac.jp"があるが、 "# kanazawa-it.ac.jp"ではないとき
	if (strings.Contains(string(input), "kanazawa-it.ac.jp")) {
		if !strings.Contains(string(input), "# kanazawa-it.ac.jp") {
			output := strings.Replace(string(input), "kanazawa-it.ac.jp", "# kanazawa-it.ac.jp", 1)
			err = ioutil.WriteFile(filename, []byte(output), 0666)
		}
	}
	if err != nil {
		return
	}
	return
}

func proxySubComment(filename string) (err error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	if strings.Contains(string(input), "# kanazawa-it.ac.jp") {
		output := strings.Replace(string(input), "# kanazawa-it.ac.jp", "kanazawa-it.ac.jp", 1)
		err = ioutil.WriteFile(filename, []byte(output), 0666)
	}
	if err != nil {
		return
	}
	return
}

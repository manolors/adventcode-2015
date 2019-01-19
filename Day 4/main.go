package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	s := string(file[:len(file)])
	fmt.Println("Read:", s)
	if err != nil {
		panic(err)
	}
	keepSearching := true
	i := -1

	for keepSearching {
		i = i + 1
		checkSum := md5.Sum([]byte(s + strconv.Itoa(i)))
		checkSumString := hex.EncodeToString(checkSum[:])

		if strings.HasPrefix(checkSumString, "00000") {
			keepSearching = false
		}
	}

	fmt.Println("Found at index:", i)
}

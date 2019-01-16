package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// Will add more stuff when needed.
var content = `#include <bits/stdc++.h>
using namespace std;

typedef long long ll;
#define REP(i, a, b) for (int i = a; i <= b; i++)

int main() {
}
`

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please tell me the file name")
	}
	fileName, path := os.Args[1], path.Dir(os.Args[1])
	if !strings.HasSuffix(fileName, ".cpp") {
		fileName += ".cpp"
	}
	if path != "." {
		os.Mkdir(path, os.ModePerm)
	}

	err := ioutil.WriteFile(fileName, []byte(content), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s created \n", fileName)
}

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

typedef double f64;
typedef long long i64;
typedef int i32;
typedef pair<i32, i32> pi32;
typedef unsigned long long u64;
typedef unsigned int u32;
typedef vector<i32> vi32;
typedef deque<i32> di32;

#define REP(i, a, b) for (auto i = a; i < b; i++)
#define REPA(i, a, b, acc) for (auto i = a; i < b; i += acc)
#define PB push_back
#define PF push_front

i32 main() {
  ios::sync_with_stdio(false); // Makes IO faster, remove this line if C style scanf/printf needed.
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

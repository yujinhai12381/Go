/*
* File Name: wordFreq.go
* Desc: 
* Author:yujinhai
* mail: yujinhai12381@126.com
* Created Time: 四  3/ 7 21:00:04 2019
*/
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "Go/wordcount"
)

func main() {
    if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
        fmt.Printf("usage: %s <file1> [<file2> [... <fileN>]]\n",
            filepath.Base(os.Args[0]))
        os.Exit(1)
    }

    wordcounter := make(wordcount.WordCount)
    // for _, filename := range os.Args[1:] {
    //  wordcount.UpdateFreq(filename)
    // }
    wordcounter.WordFreqCounter(os.Args[1:])

    wordcounter.SortReport()
}


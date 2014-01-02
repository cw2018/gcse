package main

import (
	"fmt"
	"log"
	
	"github.com/daviddengcn/gcse"
	"github.com/daviddengcn/sophie"
)

func main() {
//	path := "data/docs"
	path := "data/docs-updated"
	kvDir := sophie.KVDirInput{
		Fs: sophie.LocalFS,
		Path: path,
	}
	
	cnt, err := kvDir.PartCount()
	if err != nil {
		log.Fatalf("kvDir.PartCount failed: %v", err)
	}
	
	totalEntries := 0
	for i := 0; i < cnt; i ++ {
		it, err := kvDir.Iterator(i)
		if err != nil {
			log.Fatalf("kvDir.Collector(%d) failed: %v", i, err)
		}
		
		var key sophie.RawString
		var val gcse.DocInfo
		for {
			if err := it.Next(&key, &val); err != nil {
				if err == sophie.EOF {
					break;
				}
				log.Fatalf("it.Next failed %v", err)
			}
			totalEntries++
		}
		
		it.Close()
	}
	
	fmt.Printf("Total %d files, %d entries.\n", cnt, totalEntries)
}
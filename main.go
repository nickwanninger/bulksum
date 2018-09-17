package main



import (
	"fmt"
	"os"
	"crypto/sha256"
	"io"
	"runtime"
	"path"
	"github.com/ivpusic/grpool"
)



func shasum(path string) {

}




func main() {

	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	if len(os.Args) <= 1 {
		fmt.Println("usage: bulksum [file...]")
		os.Exit(1)
	}


	pool := grpool.NewPool(numCPUs, 10)
	defer pool.Release()



	files := os.Args[1:]

	pool.WaitCount(len(files))


	for _, f := range files {
		file := f
		pool.JobQueue <- func () {

			defer pool.JobDone()
			h := sha256.New()

			f, err := os.Open(file)
			if err != nil {
				return
			}
			defer f.Close()


			if _, err := io.Copy(h, f); err != nil {
				return
			}
			fmt.Printf("%x  %s\n", h.Sum(nil), path.Base(file))

		}
	}



	pool.WaitAll()

}

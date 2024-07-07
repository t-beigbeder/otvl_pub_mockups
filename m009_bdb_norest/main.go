package main

import (
	"fmt"
	"github.com/tidwall/buntdb"
	_ "github.com/tidwall/buntdb"
	"os"
	"os/exec"
	"time"
)

func exErr(err error) {
	fmt.Fprintf(os.Stderr, "error %v", err)
	os.Exit(-1)
}

func subexErr(err error) {
	fmt.Fprintf(os.Stdout, "error %v", err)
	os.Exit(-1)
}

func bdbrun() error {
	db, err := buntdb.Open("/tmp/bdbrun.bdb")
	if err != nil {
		subexErr(err)
	}
	defer db.Close()
	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set("mykey", "myvalue", nil)
		return err
	})
	if err != nil {
		subexErr(err)
	}
	time.Sleep(10 * time.Millisecond)
	if err = db.Shrink(); err != nil {
		subexErr(err)
	}
	return nil
}

func main() {
	if len(os.Args) == 1 {
		for i := 1; i <= 200; i++ {
			ep, err := os.Executable()
			if err != nil {
				exErr(err)
			}
			if i%100 == 0 {
				fmt.Printf("exec #%d %s %s\n", i, ep, "buntdb_norest")
			}
			cmd := exec.Command(ep, "buntdb_norest")
			out, err := cmd.Output()
			if err != nil {
				exErr(err)
			}
			fmt.Printf(string(out))
		}
		os.Exit(0)
	}
	if err := bdbrun(); err != nil {
		fmt.Fprintf(os.Stdout, "in bdbrun: %v\n", err)
	}
	os.Exit(0)
}

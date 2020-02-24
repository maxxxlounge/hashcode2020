package main

import (
	"github.com/maxxxlounge/hashcode2020/hashcode/massimo/processor"
	"github.com/pkg/errors"
	"log"
)

func ProcessFile(filename string) error {
	rte, err := processor.ParseInput(filename)
	if err != nil {
		err = errors.Wrap(err, "error parsing input")
		return err
	}

	err = rte.Analyze()
	if err != nil {
		err = errors.Wrap(err, "error elaborating data")
		return err
	}
	/*
		s, err := rte.Elaborate()
		if err != nil {
			err = errors.Wrap(err, "error elaborating data")
			return err
			}

		err = s.ParseOutput(filename)
		if err != nil {
			err = errors.Wrap(err, "error writing output")
			return err
		}*/
	return nil
}

func main() {
	var err error

	/*err = ProcessFile("../a.txt")
	if err != nil {
		err = errors.Wrap(err, "error processing a.txt")
		log.Fatal(err.Error())
	}
	err = ProcessFile("../b.txt")
	if err != nil {
		err = errors.Wrap(err, "error processing b.txt")
		log.Fatal(err.Error())
	}
	err = ProcessFile("../c.txt")
	if err != nil {
		err = errors.Wrap(err, "error processing c.txt")
		log.Fatal(err.Error())
	}
	err = ProcessFile("../d.txt")
	if err != nil {
		err = errors.Wrap(err, "error processing d.txt")
		log.Fatal(err.Error())
	}
	err = ProcessFile("../e.txt")
	if err != nil {
		err = errors.Wrap(err, "error processing e.txt")
		log.Fatal(err.Error())
	}*/
	err = ProcessFile("f.txt")
	if err != nil {
		err = errors.Wrap(err, "error processing f.txt")
		log.Fatal(err.Error())
	}

}

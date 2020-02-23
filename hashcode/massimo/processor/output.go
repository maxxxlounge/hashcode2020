package processor

import (
	"os"
	"strconv"
	"strings"
)

func (s *Solution) ParseOutput(filename string) error {
	filename = strings.Replace(filename, ".txt", ".txt.out", -1)
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	ll := strconv.Itoa(len(s.Libraries)) + "\n"
	for _, out := range s.Libraries {
		ll += strconv.Itoa(out.ID) + " " + strconv.Itoa(len(out.Books)) + "\n"
		for _, b := range out.Books {
			ll += strconv.Itoa(b) + " "
		}
		ll += "\n"
	}

	ll = strings.Replace(ll, " \n", "\n", -1)

	_, err = outFile.Write([]byte(ll))
	if err != nil {
		return err
	}
	return nil
}

package processor

import "fmt"

func (d *ReadyToProcess) Analyze() error {
	for _, l := range d.Libraries {
		fmt.Printf("SignupDay: %v\n", l.SighUpDay)
		fmt.Printf("MaxBookScan: %v\n", l.MaxBookScanPerDay)
		fmt.Printf("BookCount: %v\n", len(l.Books))
		fmt.Printf("Days for scan all book: %v\n", len(l.Books)/l.MaxBookScanPerDay)
		fmt.Printf("Point book sum: %v\n", l.GetSumBookPoint())
		fmt.Printf("\n")
	}
	return nil
}

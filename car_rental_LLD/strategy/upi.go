package strategy

import "fmt"

type Upi struct{}

func (u *Upi) Pay(amt float64) bool {
	fmt.Printf("Paid %.2f via UPI\n", amt)
	return true
}

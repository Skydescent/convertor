package conventor

import "fmt"

func Conventor(decimal int64) string {
	binary := fmt.Sprintf("%b", decimal)

	fmt.Println("decimal = ", decimal, " binary = ", binary)

	return binary
}

package main

import "fmt"

// Struct type with a method
type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

// Method with struct type receiver
func (a author) show() {
	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.particles)
	fmt.Println("Salary: ", a.salary)
}

// Struct type receiver example with a pointer receiver method
type pointerReceiver struct {
	name       string
	email      string
	isVerified bool
	pass       string
	id         int
}

// Method with pointer receiver to update the pass field
func (b *pointerReceiver) updatePass(newPass string) {
	b.pass = newPass
}

// Struct type receiver example with a method that demonstrates both pointer and value receiver
func (b pointerReceiver) displayInfo() {
	fmt.Printf("\nPointer Receiver Info:\nName: %v\nID: %v\nVerified: %v\nPass: %v\nEmail: %v\n",
		b.name, b.id, b.isVerified, b.pass, b.email)
}

// Non-struct type receiver example
type data int

// Method with non-struct type receiver
func (d1 data) multiplication(d2 data) data {
	return d1 * d2
}

func main() {
	// Example usage of struct type receiver
	res := author{
		name:      "azmi-84",
		branch:    "SE",
		particles: 84,
		salary:    84,
	}
	res.show()

	// Example usage of non-struct type receiver
	value1 := data(23)
	value2 := data(98)
	resultForNonStructMethod := value1.multiplication(value2)
	fmt.Println("Final result of multiplication:", resultForNonStructMethod)

	// Example usage of pointer receiver and updating fields using pointer method
	info := pointerReceiver{
		name:       "azmi-84",
		email:      "azmi-84@go.io",
		isVerified: true,
		pass:       "azmi-84",
		id:         84,
	}

	// Using the pointer receiver method to update the pass field
	info.updatePass("newPass123")
	info.displayInfo()

	// Demonstrating how pointer and value receivers can access the struct
	after := &info
	after.displayInfo()
}

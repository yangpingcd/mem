package mem

import (
	"fmt"
	"testing"	
)

func TestMem(*testing.T) {
	fmt.Println("total:", GetTotal())
	fmt.Println("available:", GetAvailable())	
}

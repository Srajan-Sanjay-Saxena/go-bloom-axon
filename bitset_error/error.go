package bitset_error

import (
	"fmt"
)

type BitSetError struct {
	status uint32 
	message string
}


func (bse *BitSetError) Error() string {
	return fmt.Sprintf("status : %d , message : %s" , bse.status , bse.message);
}

var (
	ErrIndexOutOfBounds = &BitSetError{status: 400, message: "Index out of bounds"}
)

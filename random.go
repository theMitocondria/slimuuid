package slimuuid
import (
	"github.com/google/uuid"
)

func ID () ( string , error ) {
	curr , err := uuid.NewRandom()
	if err != nil {
		return "" , err
	}
	ID := curr.String()
	return ID , nil 
}
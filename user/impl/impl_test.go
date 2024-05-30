package impl_test

import (
	"fmt"
	"testing"
	"vblog/user"
)

func TestUserFormat(t *testing.T) {
	fmt.Println(user.CreateNewUser().String())
}

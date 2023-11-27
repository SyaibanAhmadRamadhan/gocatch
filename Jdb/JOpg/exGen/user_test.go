package exGen

import (
	"fmt"
	"testing"

	"github.com/SyaibanAhmadRamadhan/jolly/Jsql"
)

func TestNewUser(t *testing.T) {
	user := NewUser()

	testSetColumn(t, user)
	testDeleteColumnSuccessfully(t, user)
	testDeleteInvalidColumn(t, user)

	user.ResetQueryColumnFields()
}

func testSetColumn(t *testing.T, user *User) {
	err := user.SetColumn(user.FieldID(), user.FieldEmail(), user.FieldCreatedAt())
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	user.SetID("rama")
	createdAt := "rama"
	user.SetCreatedAt(Jsql.NewNullString(&createdAt))

	printUser(user)
}

func testDeleteColumnSuccessfully(t *testing.T, user *User) {
	err := user.DeleteColumnFromQueryColumnFields(user.FieldID())
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	printUser(user)
}

func testDeleteInvalidColumn(t *testing.T, user *User) {
	err := user.DeleteColumnFromQueryColumnFields(user.FieldEmail(), "asal", user.FieldID())
	if err == nil {
		t.Errorf("Expected error, but got none")
	}

	printUser(user)
}

func printUser(user *User) {
	// fmt.Println(user)
	fmt.Println(user.QueryColumnFieldToStrings())
}

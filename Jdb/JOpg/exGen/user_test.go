package exGen

import (
	"fmt"
	"testing"

	"github.com/SyaibanAhmadRamadhan/jolly/Jsql"
)

func TestNewUser(t *testing.T) {
	user := NewUser()

	qSetArg(t, user)
	testSetColumn(t, user)
	testDeleteColumnSuccessfully(t, user)
	testDeleteInvalidColumn(t, user)

	user.ResetQColumnFields()
	user.ResetQFilterNamedArgs()

	fmt.Println(user)
}

func testSetColumn(t *testing.T, user *User) {
	err := user.SetColumn(user.FieldID(), user.FieldEmail(), user.FieldCreatedAt())
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	user.SetID("rama")
	createdAt := "rama"
	user.SetCreatedAt(createdAt)

	printUser(user)
}

func testDeleteColumnSuccessfully(t *testing.T, user *User) {
	err := user.DeleteColumnFromQColumnFields(user.FieldID())
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	fmt.Println(user.QFilterNamedArgs.ToQuery(true, Jsql.Pgx))
	printUser(user)
}

func testDeleteInvalidColumn(t *testing.T, user *User) {
	err := user.DeleteColumnFromQColumnFields(user.FieldEmail(), "asal", user.FieldID())
	if err == nil {
		t.Errorf("Expected error, but got none")
	}

	printUser(user)
}

func printUser(user *User) {
	// fmt.Println(user)
	fmt.Println(user.QueryColumnFieldToStrings())
}

func qSetArg(t *testing.T, user *User) {
	user.SetArgFieldID("rama", Jsql.Equals, Jsql.And, "custom")

	fmt.Println(user.QFilterNamedArgs.ToQuery(true, Jsql.Pgx))
}

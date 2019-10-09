//Package common holds common Github elements
package common

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

type Label struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

// Issue represents a github issue
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url,omitempty"`
	Title     string `json:"title"`
	State     string
	Milestone NullInt  `json:"milestone,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
	Labels    []Label  `json:"labels,omitempty"`
	User      *User
	CreatedAt time.Time `json:"created_at,omitempty"`
	Body      string    `json:"body,omitempty"` // in Markdown format
}

// User represents a github user
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// StringSliceFlag allows capture of a comma separated list of strings
type StringSliceFlag []string

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (s *StringSliceFlag) String() string {
	return fmt.Sprint(*s)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
func (s *StringSliceFlag) Set(value string) error {
	// If we wanted to allow the flag to be set multiple times,
	// accumulating values, we would delete this if statement.
	// That would permit usages such as
	//	-deltaT 10s -deltaT 15s
	// and other combinations.
	if len(*s) > 0 {
		return errors.New("interval flag already set")
	}

	for _, v := range strings.Split(value, ",") {
		*s = append(*s, v)
	}
	return nil
}

// NullInt noqa
type NullInt struct {
	Valid bool
	Int   int
}

// MarshalJSON from the json.Marshaler interface
func (v *NullInt) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON from the json.Unmarshaler interface
func (v *NullInt) UnmarshalJSON(data []byte) error {
	var x *int
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Int = *x
		v.Valid = true
	} else {
		v.Valid = false
	}
	return nil
}

func Credentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Github Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Github Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	password := string(bytePassword)

	return strings.TrimSpace(username), strings.TrimSpace(password)
}

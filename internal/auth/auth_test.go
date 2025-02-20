package auth

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAuth(t *testing.T) {
    validTk := "a;lsdkjfa;lskdjf;aksdjf"
    valid := "ApiKey " + validTk
    invalid := "KeyApi " + validTk
    invalid_null := " " + validTk

    tests := map[string]struct {
        header string
        want    string
    }{
        "valid":            {header: valid,         want: validTk},
        "invalid_header":   {header: invalid,       want: ""},
        "invalid_null":     {header: invalid_null,  want: ""},
        "null_token":       {header: "",            want: ""},
    }

    head, errH := http.Head("https://boot.dev")
    if errH != nil {
        t.Fatalf("Couldn't get head")
    }

    for name, tc := range tests {
        t.Run(name, func(t *testing.T) {

            http.Header.Set(head.Header, "Authorization", tc.header)

            got, err := GetAPIKey(head.Header)
            if err != nil {
                fmt.Printf("\t- on %s, error: %v\n", name, err)
            }
            if !reflect.DeepEqual(tc.want, got) {
                t.Fatalf("expected: %#v, got: %#v", tc.want, got)
            }
        })
    }
}

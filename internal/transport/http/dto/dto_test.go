package dto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//TestCreateUserDTO...
func TestValidateOK(t *testing.T) {
	cases := []struct {
		name string
		in   *CreateUserDTO
	}{
		{
			name: "ok",
			in: &CreateUserDTO{
				Name: "stas",
				Age:  "31",
			},
		},
	}
	for _, casse := range cases {

		err := casse.in.Validate()
		if err != nil {
			t.Error()
		}
		require.NoError(t, err)

	}

}
func TestValidationErr(t *testing.T) {
	cases := []struct {
		name string
		in   *CreateUserDTO
		exp  string
	}{
		{
			name: "WrongName",
			in: &CreateUserDTO{
				Name: "123",
				Age:  "21",
			},
			exp: "name: must contain English letters only.",
		}, {
			name: "WrongAge",
			in: &CreateUserDTO{
				Name: "name",
				Age:  "asd",
			},
			exp: "age: must contain digits only.",
		}, {
			name: "EmptyFields",
			in: &CreateUserDTO{
				Name: "",
				Age:  "",
			},
			exp: "age: cannot be blank; name: cannot be blank.",
		},
	}

	for _, test := range cases {
		err := test.in.Validate()
		require.Error(t, err)
		require.Equal(t, test.exp, err.Error())
	}
}

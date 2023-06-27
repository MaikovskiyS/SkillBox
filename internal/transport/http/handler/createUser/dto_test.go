package createHandler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//TestCreateUserDTO...
func TestValidateOK(t *testing.T) {
	cases := []struct {
		name string
		in   DTO
	}{
		{
			name: "ok",
			in: DTO{
				Name: "stas",
				Age:  "31",
			},
		},
	}
	for _, cases := range cases {

		err := cases.in.Validate()
		if err != nil {
			t.Error()
		}
		require.NoError(t, err)

	}

}
func TestValidationErr(t *testing.T) {
	cases := []struct {
		name string
		in   DTO
		exp  string
	}{
		{
			name: "WrongName",
			in: DTO{
				Name: "123",
				Age:  "21",
			},
			exp: "name: must contain English letters only.",
		}, {
			name: "WrongAge",
			in: DTO{
				Name: "name",
				Age:  "asd",
			},
			exp: "age: must contain digits only.",
		}, {
			name: "EmptyFields",
			in: DTO{
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

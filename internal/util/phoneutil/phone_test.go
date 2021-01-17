//nolint:scopelint
package phoneutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalizePhone(t *testing.T) {
	tests := []struct {
		data string
		want string
	}{
		{
			data: "7916&80*80909",
			want: "79168080909",
		},
		{
			data: "7(916)80-80-909",
			want: "79168080909",
		},
		{
			data: `+7 916\ 80-80-909`,
			want: "79168080909",
		},
		{
			data: "+7 916 80-80-909 ",
			want: "79168080909",
		},
		{
			data: "",
			want: "",
		},
		{
			data: "+375_7916_8=0-80-9/09/",
			want: "37579168080909",
		},
	}

	for _, tt := range tests {
		t.Run("normalize", func(t *testing.T) {
			result := NormalizePhone(tt.data)
			require.Equal(t, tt.want, result)
		})
	}
}

func TestCheckPhone(t *testing.T) {
	tests := []struct {
		data string
		want bool
	}{
		{
			data: "79096667777a",
			want: false,
		},
		{
			data: "790(96667777",
			want: false,
		},
		{
			data: "79b096667777",
			want: false,
		},
		{
			data: "79096667777",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run("check phone", func(t *testing.T) {
			result := HasOnlyDigits(tt.data)
			require.Equal(t, tt.want, result)
		})
	}
}

package uwu

import "testing"

func TestUwUify(t *testing.T) {
	tests := []struct {
		name string
		original string
		want string
	}{
		{
			name: "Converts r's and l's to w",
			original: "r, l",
			want: "w, w",
		},
		{
			name: "Converts R's and L's to W",
			original: "R, L",
			want: "W, W",
		},
		{
			name: "Converts th to d and Th to D",
			original: "th Th",
			want: "d D",
		},
		{
			name: "Does not convert t's if next letter is not 'h'",
			original: "td",
			want: "td",
		},
		{
			name: "Does not convert t if it is the last character",
			original: "t",
			want: "t",
		},
		{
			name: "Inserts a 'w' in between a vowel and a 't'",
			original: "at et it ot ut",
			want: "awt ewt iwt owt uwt",
		},
		{
			name: "Does not insert a 'w' in between a vowel and consonants other than 't'",
			original: "ac eb im on uq",
			want: "ac eb im on uq",
		},
		{
			name: "Does not convert if vowel is the last character",
			original: "a",
			want: "a",
		},
		{
			name: "Inserts ' UwU' when there is an exclamation mark",
			original: "!",
			want: " UwU",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UwUify(tt.original); got != tt.want {
				t.Errorf("UwUify() = %v, want %v", got, tt.want)
			}
		})
	}
}

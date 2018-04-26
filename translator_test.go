package gotranslate

import (
	"testing"

	"golang.org/x/text/language"
)

func Test_translationRequest(t *testing.T) {
	type args struct {
		text string
		from language.Tag
		to   language.Tag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"translate_hello", args{"hello", language.English, language.Chinese}, "你好", false},
		{"translate_hello_auto", args{"hello", language.Und, language.Chinese}, "你好", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := translationRequest(tt.args.text, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("translationRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("translationRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

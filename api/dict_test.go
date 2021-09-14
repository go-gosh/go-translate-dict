package api

import (
	"testing"
)

func TestTranslate(t *testing.T) {
	type args struct {
		query string
		opts  []Option
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"en to zh", args{"China", nil}, false},
		{"zh to en", args{"中国", []Option{WithTo(LangTypeEN)}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Translate(tt.args.query, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%s", got)
		})
	}
}

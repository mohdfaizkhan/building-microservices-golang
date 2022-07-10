package data

import (
	"bytes"
	"testing"
)

func TestProducts_ToJSON(t *testing.T) {
	tests := []struct {
		name    string
		p       *Products
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := tt.p.ToJSON(w); (err != nil) != tt.wantErr {
				t.Errorf("Products.ToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Products.ToJSON() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

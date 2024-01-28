package eval

import "testing"

func TestKotoba_Eval(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name    string
		k       *Kotoba
		args    args
		want    string
		wantErr error
	}{
		{
			name:    "integer",
			k:       &Kotoba{},
			args:    args{expr: `1`},
			want:    `1`,
			wantErr: nil,
		},
		{
			name:    "string",
			k:       &Kotoba{},
			args:    args{expr: `"hello"`},
			want:    `hello`,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Kotoba{}
			got, err := k.Eval(tt.args.expr)
			if err != tt.wantErr {
				t.Errorf("Kotoba.Eval() %q %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Kotoba.Eval() = %v, want %v", got, tt.want)
			}
		})
	}
}

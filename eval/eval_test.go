package eval

import "testing"

func TestKotoba_Eval(t *testing.T) {
	type args struct {
		expr []any
	}
	tests := []struct {
		name    string
		k       *Kotoba
		args    args
		want    any
		wantErr error
	}{
		{
			name:    "integer",
			k:       &Kotoba{},
			args:    args{expr: []any{1}},
			want:    1,
			wantErr: nil,
		},
		{
			name:    "string",
			k:       &Kotoba{},
			args:    args{expr: []any{`"hello"`}},
			want:    `hello`,
			wantErr: nil,
		},
		{
			name:    "add",
			k:       &Kotoba{},
			args:    args{expr: []any{`+`, 1, 2}},
			want:    3,
			wantErr: nil,
		},
		{
			name:    "concat",
			k:       &Kotoba{},
			args:    args{expr: []any{`+`, `"hello"`, `"world"`}},
			want:    `helloworld`,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Kotoba{}
			got, err := k.Eval(tt.args.expr...)
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

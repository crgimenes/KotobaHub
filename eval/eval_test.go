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
			name:    "sub",
			k:       &Kotoba{},
			args:    args{expr: []any{`-`, 1, 2}},
			want:    -1,
			wantErr: nil,
		},
		{
			name:    "concat",
			k:       &Kotoba{},
			args:    args{expr: []any{`+`, `"hello"`, `"world"`}},
			want:    `helloworld`,
			wantErr: nil,
		},

		{
			name:    "nested add",
			k:       &Kotoba{},
			args:    args{expr: []any{`+`, 1, []any{`+`, 2, 3}}},
			want:    6,
			wantErr: nil,
		},
		{
			name:    "nested add 2",
			k:       &Kotoba{},
			args:    args{expr: []any{`+`, []any{`+`, 1, 2}, 3}},
			want:    6,
			wantErr: nil,
		},
		{
			name:    "nested add 3",
			k:       &Kotoba{},
			args:    args{expr: []any{`+`, []any{`+`, 1, 2}, []any{`+`, 3, 4}}},
			want:    10,
			wantErr: nil,
		},
		{
			name:    "nested add 4",
			k:       &Kotoba{},
			args:    args{expr: []any{`+`, []any{`+`, 1, 2}, []any{`+`, 3, []any{`+`, 4, []any{`+`, 5, 6}}}}},
			want:    21,
			wantErr: nil,
		},
		{
			name: "nested add 5",
			k:    &Kotoba{},
			args: args{
				expr: []any{
					`+`,
					[]any{`+`,
						[]any{`+`,
							[]any{`+`, 1, 2},
							[]any{`+`, 3, 4}},
						[]any{`+`, 5, 6}}}},
			want: 21,
		},
		{
			name:    "nested sub",
			k:       &Kotoba{},
			args:    args{expr: []any{`-`, 1, []any{`-`, 2, 3}}},
			want:    2,
			wantErr: nil,
		},
		{
			name:    "multi",
			k:       &Kotoba{},
			args:    args{expr: []any{`*`, 2, 3}},
			want:    6,
			wantErr: nil,
		},
		{
			name:    "div",
			k:       &Kotoba{},
			args:    args{expr: []any{`/`, 6, 3}},
			want:    2,
			wantErr: nil,
		},
		{
			name:    "divide by zero",
			k:       &Kotoba{},
			args:    args{expr: []any{`/`, 6, 0}},
			want:    nil,
			wantErr: ERR_DIV_BY_ZERO,
		},
		{
			name:    "set variable",
			k:       &Kotoba{},
			args:    args{expr: []any{`set`, `x`, 1}},
			want:    1,
			wantErr: nil,
		},
		{
			name:    "get variable",
			k:       &Kotoba{},
			args:    args{expr: []any{`get`, `true`}},
			want:    true,
			wantErr: nil,
		},
		{
			name:    "variable not found",
			k:       &Kotoba{},
			args:    args{expr: []any{`get`, `x`}},
			want:    nil,
			wantErr: ERR_VARIABLE_NOT_FOUND,
		},
	}
	for _, tt := range tests {
		k := New()
		t.Run(tt.name, func(t *testing.T) {
			got, err := k.Eval(nil, tt.args.expr...)
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

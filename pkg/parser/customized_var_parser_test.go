package parser

import (
	"testing"
)

func TestParseCustVars(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    CustVars
		wantErr bool
	}{
		{
			name: "Get customized vars from vars.yml",
			args: args{
				path: "./valid_vars_test.yml",
			},
			want: CustVars{
				Vars: []map[string]CustVar{
					map[string]CustVar{
						"public_subnets": CustVar{
							Type:    "list",
							Desc:    "subnets for public",
							Default: `["sub1", "sub2"]`,
						},
					},
					map[string]CustVar{
						"tags": CustVar{
							Type: "map",
							Default: `{
								Name = "Terraform"
							}`,
						},
					},
					map[string]CustVar{
						"cidrs": CustVar{
							Type:    "list",
							Default: `["10.0.0.0/16", "10.1.0.0/16"]`,
						},
					},
					map[string]CustVar{
						"t1-var2": CustVar{
							Desc: "var for t1",
							Type: "string",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Get error when reading yml file",
			args: args{
				path: "./ad-hoc.yml",
			},
			want:    CustVars{},
			wantErr: true,
		},
		{
			name: "Get error when parsing yml file",
			args: args{
				path: "./invalid_vars_test.yml",
			},
			want:    CustVars{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCustVars(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCustVars() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.equal(tt.want) {
				t.Errorf("ParseCustVars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustVars_equal(t *testing.T) {
	type fields struct {
		Vars []map[string]CustVar
	}
	type args struct {
		other CustVars
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "CustVars is equal",
			fields: fields{
				Vars: []map[string]CustVar{
					map[string]CustVar{
						"public_subnets": CustVar{
							Type:    "list",
							Desc:    "subnets for public",
							Default: `["sub1", "sub2"]`,
						},
					},
				},
			},
			args: args{
				other: CustVars{
					Vars: []map[string]CustVar{
						map[string]CustVar{
							"public_subnets": CustVar{
								Type:    "list",
								Desc:    "subnets for public",
								Default: `["sub1", "sub2"]`,
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "CustVars lens is not equal",
			fields: fields{
				Vars: []map[string]CustVar{},
			},
			args: args{
				other: CustVars{
					Vars: []map[string]CustVar{
						map[string]CustVar{
							"public_subnets": CustVar{
								Type:    "list",
								Desc:    "subnets for public",
								Default: `["sub1", "sub2"]`,
							},
						},
					},
				},
			},
			want: false,
		},
		{
			name: "CustVars keys is not equal",
			fields: fields{
				Vars: []map[string]CustVar{
					map[string]CustVar{
						"public_subnets": CustVar{
							Type:    "list",
							Desc:    "subnets for public",
							Default: `["sub1", "sub2"]`,
						},
					},
				},
			},
			args: args{
				other: CustVars{
					Vars: []map[string]CustVar{
						map[string]CustVar{
							"hello": CustVar{
								Type:    "list",
								Default: `["sub1", "sub2"]`,
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CustVars{
				Vars: tt.fields.Vars,
			}
			if got := c.equal(tt.args.other); got != tt.want {
				t.Errorf("CustVars.equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

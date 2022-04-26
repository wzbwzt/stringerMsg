package main

import (
	"go/ast"
	"testing"
)

func Test_getComment(t *testing.T) {
	type args struct {
		name  string
		group *ast.CommentGroup
	}
	demos := []struct {
		name string
		args args
		want string
	}{
		{
			name: "demo1",
			args: args{
				name:  "ErrCodeNotExist",
				group: &ast.CommentGroup{List: []*ast.Comment{{Text: "//ErrCodeNotExist not exist"}}},
			},
			want: "not exist",
		}, {

			name: "demo2",
			args: args{
				name:  "ErrCodeNotExist",
				group: &ast.CommentGroup{List: []*ast.Comment{{Text: "// ErrCodeNotExist not exist"}}},
			},
			want: "not exist",
		}, {

			name: "demo3",
			args: args{
				name:  "ErrCodeNotExist",
				group: &ast.CommentGroup{List: []*ast.Comment{{Text: "//   not exist"}}},
			},
			want: "not exist",
		},
	}
	for _, demo := range demos {
		t.Run(demo.name, func(t *testing.T) {
			if res := getComment(demo.args.name, demo.args.group); res != demo.want {
				t.Errorf("getComment();get[%s],but want[%s]", res, demo.want)
			}
		})
	}

}

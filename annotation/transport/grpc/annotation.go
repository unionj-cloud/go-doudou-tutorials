/**
* Generated by go-doudou v2.0.0.
* Don't edit!
 */
package grpc

import "github.com/unionj-cloud/go-doudou/v2/framework/rest"

var MethodAnnotationStore = rest.AnnotationStore{
	"GetUserRpc": {
		{
			Name: "@role",
			Params: []string{
				"USER",
				"ADMIN",
			},
		},
	},
	"GetAdminRpc": {
		{
			Name: "@role",
			Params: []string{
				"ADMIN",
			},
		},
	},
}

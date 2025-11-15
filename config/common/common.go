// Inspired from https://github.com/crossplane-contrib/provider-upjet-aws/blob/main/config/cluster/common/common.go

package common

import (
	"github.com/crossplane/crossplane-runtime/v2/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
)

func PrivateNetworkOpenStackIdExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		r, err := paved.GetString("status.atProvider.regionsAttributes[0].openstackid")
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		return r
	}
}

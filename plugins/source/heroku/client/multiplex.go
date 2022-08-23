package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func NoMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	return []schema.ClientMeta{meta.(*Client)}
}

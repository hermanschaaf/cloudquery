package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func TeamMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, t := range client.Teams {
		l = append(l, client.WithTeam(t))
	}
	return l
}

package app

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
)

func getCredFromMetadata(ctx context.Context) (string, bool) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if uID, ok := md["cred"]; ok {
			return strings.Join(uID, ","), true
		}
	}
	return "", false
}

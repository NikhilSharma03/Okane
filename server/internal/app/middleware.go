package app

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
)

func getCredFromMetadata(ctx context.Context, mdtype string) (string, bool) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if uID, ok := md[mdtype]; ok {
			return strings.Join(uID, ","), true
		}
	}
	return "", false
}

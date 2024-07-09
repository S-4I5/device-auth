package auth

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"slices"
	"strings"
	"user-service/pkg/auth_v1"
)

func (a *AuthenticationServiceImplementation) ValidateToken(ctx context.Context, req *auth_v1.ValidateTokenRequest) (*auth_v1.ValidateTokenResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "no auth token provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "no auth token provided")
	}

	issuerToken, _ := strings.CutPrefix(values[0], "Bearer ")

	claims, err := a.authService.VerifyUserToken(req.Token)
	if err != nil {
		return nil, err
	}

	aud, err := claims.GetAudience()
	if err != nil {
		return nil, err
	}

	if !slices.Contains(aud, issuerToken) {
		return nil, status.Errorf(codes.InvalidArgument, "not valid for current client")
	}

	userId, _ := claims.GetSubject()

	return &auth_v1.ValidateTokenResponse{UserId: userId}, nil
}

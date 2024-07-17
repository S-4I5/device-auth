package scope

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"slices"
)

type Validator interface {
	IsAllowed(methodName string, scopes []Scope) error
}

var (
	errScopeNotProvided     = status.Error(codes.PermissionDenied, "client does not have needed scope")
	errMethodScopeNotStated = status.Error(codes.Internal, "method scope not stated")
)

type validator struct {
	methodMaps []MethodMap
}

func NewValidator(maps ...MethodMap) *validator {
	return &validator{methodMaps: maps}
}

func (v *validator) IsAllowed(methodName string, scopes []Scope) error {
	for _, methodMap := range v.methodMaps {
		methodScope, ok := methodMap.GetMethodScope(methodName)
		if !ok {
			continue
		}

		if !slices.Contains(scopes, methodScope) {
			return errScopeNotProvided
		}
		return nil
	}
	return errMethodScopeNotStated
}

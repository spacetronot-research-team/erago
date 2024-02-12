package ctxutil

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	ErrCtxNil = errors.New("err ctx is nil")
)

type ctxKey string

var (
	KeyModuleName  ctxKey = "erago@MODULE"
	KeyDomain      ctxKey = "erago@DOMAIN"
	KeyProjectPath ctxKey = "erago@PROJECT_PATH"
	KeyProjectName ctxKey = "erago@PROJECT_NAME"
)

func WithProjectName(ctx context.Context, projectName string) context.Context {
	return set(ctx, KeyProjectName, projectName)
}

func GetProjectName(ctx context.Context) string {
	return getString(ctx, KeyProjectName)
}

func WithModuleName(ctx context.Context, module string) context.Context {
	return set(ctx, KeyModuleName, module)
}

func GetModuleName(ctx context.Context) string {
	return getString(ctx, KeyModuleName)
}

func WithDomain(ctx context.Context, domain string) context.Context {
	return set(ctx, KeyDomain, domain)
}

func GetDomain(ctx context.Context) string {
	return getString(ctx, KeyDomain)
}

func GetDomainShort(ctx context.Context) string {
	domain := GetDomain(ctx)
	domainSplitted := strings.Split(domain, " ")
	short := ""

	for _, word := range domainSplitted {
		short += strings.ToLower(string(word[0]))
	}

	return short
}

func set(ctx context.Context, key ctxKey, value string) context.Context {
	if ctx == nil {
		logrus.Warn(ErrCtxNil)
		return context.WithValue(context.Background(), key, value)
	}
	return context.WithValue(ctx, key, value)
}

func getString(ctx context.Context, key ctxKey) string {
	if ctx == nil {
		logrus.Warn(ErrCtxNil)
		return ""
	}

	valueAny := ctx.Value(key)
	if valueAny == nil {
		logrus.WithFields(logrus.Fields{"key": key}).Warn("err ctx does not have the given key")
		return ""
	}

	return fmt.Sprintf("%v", valueAny)
}

package ctxutil

import (
	"context"
	"path/filepath"
)

var (
	KeyControllerDirPath ctxKey = "erago@CONTROLLER_DIR_PATH"
	KeyServiceDirPath    ctxKey = "erago@SERVICE_DIR_PATH"
	KeyRepositoryDirPath ctxKey = "erago@REPOSITORY_DIR_PATH"
	KeyRouterDirPath     ctxKey = "erago@ROUTER_DIR_PATH"
	KeyDocsDirPath       ctxKey = "erago@DOCS_DIR_PATH"
	KeyDatabaseDirPath   ctxKey = "erago@DATABASE_DIR_PATH"
	KeyCMDDirPath        ctxKey = "erago@CMD_DIR_PATH"
)

func SetAllDirPath(ctx context.Context, projectPath string) context.Context {
	ctx = withProjectPath(ctx, projectPath)
	ctx = withControllerDirPath(ctx, filepath.Join(projectPath, "internal", "controller", "http"))
	ctx = withServiceDirPath(ctx, filepath.Join(projectPath, "internal", "service"))
	ctx = withRepositoryDirPath(ctx, filepath.Join(projectPath, "internal", "repository"))
	ctx = withRouterDirPath(ctx, filepath.Join(projectPath, "internal", "router"))
	ctx = withDocsDirPath(ctx, filepath.Join(projectPath, "docs"))
	ctx = withCMDDirPath(ctx, filepath.Join(projectPath, "cmd"))
	ctx = withDatabaseDirPath(ctx, filepath.Join(projectPath, "database"))
	return ctx
}

func withProjectPath(ctx context.Context, projectPath string) context.Context {
	return set(ctx, KeyProjectPath, projectPath)
}

func GetProjectPath(ctx context.Context) string {
	return getString(ctx, KeyProjectPath)
}

func withControllerDirPath(ctx context.Context, controllerDirPath string) context.Context {
	return set(ctx, KeyControllerDirPath, controllerDirPath)
}

func GetControllerDirPath(ctx context.Context) string {
	return getString(ctx, KeyControllerDirPath)
}

func withServiceDirPath(ctx context.Context, serviceDirPath string) context.Context {
	return set(ctx, KeyServiceDirPath, serviceDirPath)
}

func GetServiceDirPath(ctx context.Context) string {
	return getString(ctx, KeyServiceDirPath)
}

func withRepositoryDirPath(ctx context.Context, repositoryDirPath string) context.Context {
	return set(ctx, KeyRepositoryDirPath, repositoryDirPath)
}

func GetRepositoryDirPath(ctx context.Context) string {
	return getString(ctx, KeyRepositoryDirPath)
}

func withRouterDirPath(ctx context.Context, routerDirPath string) context.Context {
	return set(ctx, KeyRouterDirPath, routerDirPath)
}

func GetRouterDirPath(ctx context.Context) string {
	return getString(ctx, KeyRouterDirPath)
}

func withDocsDirPath(ctx context.Context, docsDirPath string) context.Context {
	return set(ctx, KeyDocsDirPath, docsDirPath)
}

func GetDocsDirPath(ctx context.Context) string {
	return getString(ctx, KeyDocsDirPath)
}

func withCMDDirPath(ctx context.Context, cmdDirPath string) context.Context {
	return set(ctx, KeyCMDDirPath, cmdDirPath)
}

func GetCMDDirPath(ctx context.Context) string {
	return getString(ctx, KeyCMDDirPath)
}

func withDatabaseDirPath(ctx context.Context, databaseDirPath string) context.Context {
	return set(ctx, KeyDatabaseDirPath, databaseDirPath)
}

func GetDatabaseDirPath(ctx context.Context) string {
	return getString(ctx, KeyDatabaseDirPath)
}

func GetDatabaseMigrateDirPath(ctx context.Context) string {
	databaseDir := GetDatabaseDirPath(ctx)
	return filepath.Join(databaseDir, "migrate")
}

func GetDatabaseSchemaMigrationDirPath(ctx context.Context) string {
	databaseDir := GetDatabaseDirPath(ctx)
	return filepath.Join(databaseDir, "schema_migration")
}

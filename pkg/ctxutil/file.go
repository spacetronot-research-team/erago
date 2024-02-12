package ctxutil

import (
	"context"
	"path/filepath"
)

func GetCMDMainFilePath(ctx context.Context) string {
	cmdDir := GetCMDDirPath(ctx)
	return filepath.Join(cmdDir, "main.go")
}

func GetDatabaseOpenFilePath(ctx context.Context) string {
	databaseDir := GetDatabaseDirPath(ctx)
	return filepath.Join(databaseDir, "open.go")
}

func GetDatabaseMigrateUpFilePath(ctx context.Context) string {
	databaseMigrateDir := GetDatabaseMigrateDirPath(ctx)
	return filepath.Join(databaseMigrateDir, "up.go")
}

func GetDatabaseSchemaMigrationFileExamplePath(ctx context.Context) string {
	dbSchemaMigrationDir := GetDatabaseSchemaMigrationDirPath(ctx)
	return filepath.Join(dbSchemaMigrationDir, "20240114192700-example.sql")
}

func GetInjectionFilePath(ctx context.Context) string {
	routerDirPath := GetRouterDirPath(ctx)
	return filepath.Join(routerDirPath, "injection.go")
}

func GetRouterFilePath(ctx context.Context) string {
	routerDirPath := GetRouterDirPath(ctx)
	return filepath.Join(routerDirPath, "router.go")
}

func GetDocsErrorsJSONFilePath(ctx context.Context) string {
	docsDirPath := GetDocsDirPath(ctx)
	return filepath.Join(docsDirPath, "errors.json")
}

func GetReadmeFilePath(ctx context.Context) string {
	projectPath := GetProjectPath(ctx)
	return filepath.Join(projectPath, "README.md")
}

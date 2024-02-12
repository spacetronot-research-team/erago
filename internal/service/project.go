package service

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/erago/internal/repository"
	"github.com/spacetronot-research-team/erago/pkg/ctxutil"
)

//go:generate mockgen -source=project.go -destination=mockservice/project.go -package=mockservice

type Project interface {
	// CreateProject create new project.
	CreateProject(ctx context.Context)
}

type projectService struct {
	projectRepository repository.Project
	domainService     Domain
}

func NewProjectService(projectRepository repository.Project, domainService Domain) Project {
	return &projectService{
		projectRepository: projectRepository,
		domainService:     domainService,
	}
}

// CreateProject create new project.
func (ps *projectService) CreateProject(ctx context.Context) {
	logrus.Info("create project start")

	if ps.isProjectAlreadyExist(ctx) {
		logrus.Fatal("err project already exist")
	}

	ps.createProjectDir(ctx)

	ps.changeDirToProjectDir(ctx)

	projectPath, err := os.Getwd()
	if err != nil {
		logrus.Fatal(err)
	}
	ctx = ctxutil.SetAllDirPath(ctx, projectPath)

	ps.createGitIgnoreFile(ctx)

	ps.createEnvFile(ctx)

	ps.runGoModInit(ctx)

	ps.mkdirControllerHTTP(ctx)

	ps.mkdirService(ctx)

	ps.mkdirRepository(ctx)

	ps.createRouterTemplate(ctx)

	ps.createDocsErrorJSONFile(ctx)

	ps.createDatabaseTemplate(ctx)

	ps.createCMDTemplate(ctx)

	ps.createReadmeFile(ctx)

	ps.domainService.CreateDomain(ctx)

	logrus.Info("create project finish, go to your project:\n\tcd ", ctxutil.GetProjectPath(ctx), "\nsetup .env file then run:\n\tgo run cmd/main.go") //nolint:lll
}

func (ps *projectService) isProjectAlreadyExist(ctx context.Context) bool {
	projectName := ctxutil.GetProjectName(ctx)
	_, err := os.Stat(projectName)
	if err == nil {
		return true
	} else if !os.IsNotExist(err) {
		logrus.Fatal(err)
	}
	return false
}

func (ps *projectService) createProjectDir(ctx context.Context) {
	logrus.Info("create project dir start")
	projectName := ctxutil.GetProjectName(ctx)
	if err := os.MkdirAll(projectName, os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err mkdir projectPath: %v", err))
	}
	logrus.Info("create project dir finish")
}

func (ps *projectService) changeDirToProjectDir(ctx context.Context) {
	logrus.Info("change dir to project dir start")
	projectName := ctxutil.GetProjectName(ctx)
	if err := os.Chdir(projectName); err != nil {
		logrus.Fatal(fmt.Errorf("err change dir to project dir: %v", err))
	}
	logrus.Info("change dir to project dir finish")
}

func (ps *projectService) createGitIgnoreFile(ctx context.Context) {
	logrus.Info("create git ignore file start")

	gitIgnoreTemplate := ps.projectRepository.GetGitIgnoreTemplate(ctx)

	if err := os.WriteFile(".gitignore", []byte(gitIgnoreTemplate), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err create git ignore file: %v", err))
	}

	logrus.Info("create git ignore file finish")
}

func (ps *projectService) createEnvFile(ctx context.Context) { //nolint:unparam
	logrus.Info("create env file start")
	envTemplate := `DB_HOST="localhost"
DB_USER="user"
DB_PASSWORD="password"
DB_NAME="boilerplate_catalog"
DB_PORT="5432"
SSL_MODE="disable"
TZ="Asia/Jakarta"
`
	if err := os.WriteFile(".env", []byte(envTemplate), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err create .env file: %v", err))
	}
	logrus.Info("create env file finish")
}

func (ps *projectService) runGoModInit(ctx context.Context) {
	logrus.Info("run go mod init start")

	moduleName := ctxutil.GetModuleName(ctx)

	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Dir = "."
	stdout, err := cmd.Output()
	if err != nil {
		logrus.Fatal(fmt.Errorf("err run go mod init: %v", err))
	}

	if string(stdout) != "" {
		fmt.Println(string(stdout)) //nolint:forbidigo
	}

	logrus.Info("run go mod init finish")
}

func (ps *projectService) mkdirControllerHTTP(ctx context.Context) {
	logrus.Info("mkdir controller http start")

	if err := os.MkdirAll(ctxutil.GetControllerDirPath(ctx), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err mkdir controller http: %v", err))
	}

	logrus.Info("mkdir controller http finish")
}

func (ps *projectService) mkdirService(ctx context.Context) {
	logrus.Info("mkdir service start")

	if err := os.MkdirAll(ctxutil.GetServiceDirPath(ctx), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err mkdir service: %v", err))
	}

	logrus.Info("mkdir service finish")
}

func (ps *projectService) mkdirRepository(ctx context.Context) {
	logrus.Info("mkdir repository start")

	if err := os.MkdirAll(ctxutil.GetRepositoryDirPath(ctx), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err mkdir repository  %v", err))
	}

	logrus.Info("mkdir repository finish")
}

func (ps *projectService) createRouterTemplate(ctx context.Context) {
	logrus.Info("create router template start")

	ps.mkdirRouter(ctx)

	ps.createRouterFile(ctx)

	logrus.Info("create router template finish")
}

func (ps *projectService) mkdirRouter(ctx context.Context) {
	logrus.Info("mkdir router start")

	if err := os.MkdirAll(ctxutil.GetRouterDirPath(ctx), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err mkdir router: %v", err))
	}

	logrus.Info("mkdir router finish")
}

func (ps *projectService) createRouterFile(ctx context.Context) {
	logrus.Info("create router file start")

	routerTemplate, err := ps.projectRepository.GetRouterTemplate(ctx)
	if err != nil {
		logrus.Fatal(fmt.Errorf("err repo get router template: %v", err))
	}

	if err := os.WriteFile(ctxutil.GetRouterFilePath(ctx), []byte(routerTemplate), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err create router file: %v", err))
	}

	logrus.Info("create router file finish")
}

func (ps *projectService) createDocsErrorJSONFile(ctx context.Context) {
	logrus.Info("create docs errors.json start")

	if err := os.MkdirAll(ctxutil.GetDocsDirPath(ctx), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err create docs folder: %v", err))
	}

	if err := os.WriteFile(ctxutil.GetDocsErrorsJSONFilePath(ctx), []byte("{}"), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err write errors.json file: %v", err))
	}

	logrus.Info("create docs errors.json finish")
}

func (ps *projectService) createDatabaseTemplate(ctx context.Context) {
	logrus.Info("create database template start")

	ps.mkdirDatabase(ctx)

	ps.createDatabaseOpenFile(ctx)

	ps.mkdirDatabaseMigrate(ctx)

	ps.createDatabaseMigrateUpFile(ctx)

	ps.mkdirDatabaseSchemaMigration(ctx)

	ps.createDatabaseSchemaMigrationFileExample(ctx)

	logrus.Info("create database template finish")
}

func (ps *projectService) mkdirDatabase(ctx context.Context) {
	logrus.Info("mkdir database start")

	if err := os.MkdirAll(ctxutil.GetDatabaseDirPath(ctx), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err mkdir database: %v", err))
	}

	logrus.Info("mkdir database finish")
}

func (ps *projectService) createDatabaseOpenFile(ctx context.Context) {
	logrus.Info("create database open file start")

	var databaseOpenTemplate = `package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSL_MODE")
	tz := os.Getenv("TZ")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, sslMode, tz,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("fail initialize db session: %v", err)
	}

	return db, err
}
`

	if err := os.WriteFile(ctxutil.GetDatabaseOpenFilePath(ctx), []byte(databaseOpenTemplate), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err create database open file: %v", err))
	}

	logrus.Info("create database open file finish")
}

func (ps *projectService) mkdirDatabaseMigrate(ctx context.Context) {
	logrus.Info("mkdir database migrate start")

	if err := os.MkdirAll(ctxutil.GetDatabaseMigrateDirPath(ctx), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err mkdir database migrate: %v", err))
	}

	logrus.Info("mkdir database migrate finish")
}

func (ps *projectService) createDatabaseMigrateUpFile(ctx context.Context) {
	logrus.Info("create database migrate up file start")

	var databaseMigrateUpTemplate = `package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSL_MODE")
	tz := os.Getenv("TZ")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, sslMode, tz,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("fail initialize db session: %v", err)
	}

	return db, err
}
`

	if err := os.WriteFile(ctxutil.GetDatabaseMigrateUpFilePath(ctx), []byte(databaseMigrateUpTemplate), os.ModePerm); err != nil { //nolint:lll
		logrus.Fatal(fmt.Errorf("err create database migrate up file: %v", err))
	}

	logrus.Info("create database migrate up file finish")
}

func (ps *projectService) mkdirDatabaseSchemaMigration(ctx context.Context) {
	logrus.Info("mkdir database schema migration start")

	if err := os.MkdirAll(ctxutil.GetDatabaseSchemaMigrationDirPath(ctx), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err mkdir database schema migration: %v", err))
	}

	logrus.Info("mkdir database schema migration finish")
}

func (ps *projectService) createDatabaseSchemaMigrationFileExample(ctx context.Context) {
	logrus.Info("create database schema migration file example start")

	var databaseSchemaMigrationFileTemplate = `-- +migrate Up
SELECT
  *
from
  erago;

-- +migrate Down`

	if err := os.WriteFile(ctxutil.GetDatabaseSchemaMigrationFileExamplePath(ctx), []byte(databaseSchemaMigrationFileTemplate), os.ModePerm); err != nil { //nolint:lll
		logrus.Fatal(fmt.Errorf("err create database schema migration file example: %v", err))
	}

	logrus.Info("create database schema migration file example finish")
}

func (ps *projectService) createCMDTemplate(ctx context.Context) {
	logrus.Info("create cmd main template start")

	ps.mkdirCMD(ctx)

	ps.createCMDMainFile(ctx)

	logrus.Info("create cmd main template finish")
}

func (ps *projectService) mkdirCMD(ctx context.Context) {
	logrus.Info("mkdir cmd start")

	if err := os.MkdirAll(ctxutil.GetCMDDirPath(ctx), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err mkdir cmd: %v", err))
	}

	logrus.Info("mkdir cmd finish")
}

func (ps *projectService) createCMDMainFile(ctx context.Context) {
	logrus.Info("create cmd main file start")

	cmdMainTemplate, err := ps.projectRepository.GetCMDMainTemplate(ctx)
	if err != nil {
		logrus.Fatal(fmt.Errorf("err repo get cmd main template: %v", err))
	}

	if err := os.WriteFile(ctxutil.GetCMDMainFilePath(ctx), []byte(cmdMainTemplate), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err create cmd main file: %v", err))
	}

	logrus.Info("create cmd main file finish")
}

func (ps *projectService) createReadmeFile(ctx context.Context) {
	logrus.Info("create readme file start")

	readmeTemplate, err := ps.projectRepository.GetReadmeTemplate(ctx)
	if err != nil {
		logrus.Fatal(fmt.Errorf("err repo get readme template: %v", err))
	}

	if err := os.WriteFile(ctxutil.GetReadmeFilePath(ctx), []byte(readmeTemplate), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err create readme file: %v", err))
	}

	logrus.Info("create readme file finish")
}

package templ

func GitIgnoreTemplate() string {
	var gitIgnoreTemplate string

	gitIgnoreTemplate += "# Created by https://www.toptal.com/developers/gitignore/api/go\n"
	gitIgnoreTemplate += "# Edit at https://www.toptal.com/developers/gitignore?templates=go\n"
	gitIgnoreTemplate += "\n"
	gitIgnoreTemplate += "### Go ###\n"
	gitIgnoreTemplate += "# If you prefer the allow list template instead of the deny list, see community template:\n"
	gitIgnoreTemplate += "# https://github.com/github/gitignore/blob/main/community/Golang/Go.AllowList.gitignore\n"
	gitIgnoreTemplate += "#\n"
	gitIgnoreTemplate += "# Binaries for programs and plugins\n"
	gitIgnoreTemplate += "*.exe\n"
	gitIgnoreTemplate += "*.exe~\n"
	gitIgnoreTemplate += "*.dll\n"
	gitIgnoreTemplate += "*.so\n"
	gitIgnoreTemplate += "*.dylib\n"
	gitIgnoreTemplate += "\n"
	gitIgnoreTemplate += "# Test binary, built with `go test -c`\n"
	gitIgnoreTemplate += "*.test\n"
	gitIgnoreTemplate += "\n"
	gitIgnoreTemplate += "# Output of the go coverage tool, specifically when used with LiteIDE\n"
	gitIgnoreTemplate += "*.out\n"
	gitIgnoreTemplate += "\n"
	gitIgnoreTemplate += "# Dependency directories (remove the comment below to include it)\n"
	gitIgnoreTemplate += "# vendor/\n"
	gitIgnoreTemplate += "\n"
	gitIgnoreTemplate += "# Go workspace file\n"
	gitIgnoreTemplate += "go.work\n"
	gitIgnoreTemplate += "\n"
	gitIgnoreTemplate += "# End of https://www.toptal.com/developers/gitignore/api/go\n"

	return gitIgnoreTemplate
}

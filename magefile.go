//go:build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	// Default target to run when none is specified
	Default = Dev

	// Version info
	Version = "v0.0.1" // TODO: Automate versioning

	// Go commands
	goCmd     = "go"
	goRun     = sh.RunCmd(goCmd, "run")
	goBuild   = sh.RunCmd(goCmd, "build")
	goTest    = sh.RunCmd(goCmd, "test")
	goTidy    = sh.RunCmd(goCmd, "mod", "tidy")
	goInstall = sh.RunCmd(goCmd, "install")

	// Other tools
	airCmd          = "air"
	templCmd        = "templ"
	sqlcCmd         = "sqlc"
	atlasCmd        = "atlas"
	golangciLintCmd = "golangci-lint"
	npmCmd          = "npm"
	dockerCmd       = "docker"
)

// ldflags returns the linker flags for building the binaries.
func ldflags() string {
	return fmt.Sprintf("-X main.version=%s", Version)
}

// -----------------------------------------------------------------------------
// Development & Execution Targets
// -----------------------------------------------------------------------------

// Dev starts the development server with live reload and watches for CSS changes.
func Dev() error {
	mg.SerialDeps(Generate.All, Build.All)
	fmt.Println("Starting server with live reload and watching for CSS changes...")
	go sh.Run(npmCmd, "run", "css:watch")
	return sh.Run(airCmd)
}

// -----------------------------------------------------------------------------
// Build Targets
// -----------------------------------------------------------------------------

type Build mg.Namespace

// All builds all application binaries.
func (Build) All() error {
	mg.SerialDeps(Build.Server, Build.CLI)
	return nil
}

// Server builds the main web server binary.
func (Build) Server() error {
	fmt.Println("Building server...")
	return goBuild("-ldflags", ldflags(), "-o", "bin/server", "./cmd/server")
}

// CLI builds the command-line interface binary.
func (Build) CLI() error {
	fmt.Println("Building CLI...")
	return goBuild("-ldflags", ldflags(), "-o", "bin/cli", "./cmd/cli")
}

// -----------------------------------------------------------------------------
// Code Generation & Formatting
// -----------------------------------------------------------------------------

type Generate mg.Namespace

// All runs all code generators.
func (Generate) All() error {
	mg.SerialDeps(Generate.CSS, Generate.Templ, Generate.SQLC)
	return nil
}

// CSS generates the Tailwind CSS file.
func (Generate) CSS() error {
	fmt.Println("Generating Tailwind CSS...")
	return sh.Run(npmCmd, "run", "css:build")
}

// Templ generates Go code from templ components.
func (Generate) Templ() error {
	fmt.Println("Generating templ components...")
	return sh.Run(templCmd, "generate")
}

// SQLC generates Go code from SQL queries.
func (Generate) SQLC() error {
	fmt.Println("Generating sqlc code...")
	return sh.Run(sqlcCmd, "generate")
}

// Format runs all code formatters.
func Format() error {
	fmt.Println("Formatting code...")
	return sh.Run(npmCmd, "run", "format")
}

// -----------------------------------------------------------------------------
// Database Targets
// -----------------------------------------------------------------------------

type DB mg.Namespace

// Migrate runs database migrations using Atlas.
func (DB) Migrate() error {
	fmt.Println("Running database migrations...")
	return sh.Run(atlasCmd, "migrate", "apply", "--env", "local")
}

// -----------------------------------------------------------------------------
// Quality & CI/CD Targets
// -----------------------------------------------------------------------------

type Check mg.Namespace

// All runs all quality checks.
func (Check) All() error {
	mg.SerialDeps(Check.Lint, Check.Test, Check.Vuln)
	return nil
}

// Lint runs the linter.
func (Check) Lint() error {
	fmt.Println("Running linter...")
	return sh.Run(golangciLintCmd, "run", "./...")
}

// Test runs all unit tests.
func (Check) Test() error {
	fmt.Println("Running tests...")
	return goTest("-v", "-race", "-cover", "./...")
}

// Cover runs tests and displays coverage in the browser.
func (Check) Cover() error {
	fmt.Println("Running tests and displaying coverage...")
	if err := goTest("-coverprofile=coverage.out", "./..."); err != nil {
		return err
	}
	return sh.Run(goCmd, "tool", "cover", "-html=coverage.out")
}

// Vuln scans for vulnerabilities.
func (Check) Vuln() error {
	fmt.Println("Scanning for vulnerabilities...")
	return sh.Run("govulncheck", "./...")
}

// -----------------------------------------------------------------------------
// Release Targets
// -----------------------------------------------------------------------------

type Release mg.Namespace

// All cross-compiles release binaries for all platforms.
func (Release) All() error {
	mg.SerialDeps(Release.Linux, Release.Windows, Release.Darwin)
	return nil
}

// Linux cross-compiles release binaries for Linux.
func (Release) Linux() error {
	return release("linux", "amd64")
}

// Windows cross-compiles release binaries for Windows.
func (Release) Windows() error {
	return release("windows", "amd64")
}

// Darwin cross-compiles release binaries for macOS.
func (Release) Darwin() error {
	return release("darwin", "amd64")
}

// release is a helper function to cross-compile binaries.
func release(goos, goarch string) error {
	fmt.Printf("Building release for %s/%s...\n", goos, goarch)
	env := map[string]string{"GOOS": goos, "GOARCH": goarch}

	// Build server
	serverOut := filepath.Join("bin", fmt.Sprintf("server-%s-%s", goos, goarch))
	if goos == "windows" {
		serverOut += ".exe"
	}
	if err := sh.RunWith(env, goCmd, "build", "-ldflags", ldflags(), "-o", serverOut, "./cmd/server"); err != nil {
		return err
	}

	// Build CLI
	cliOut := filepath.Join("bin", fmt.Sprintf("cli-%s-%s", goos, goarch))
	if goos == "windows" {
		cliOut += ".exe"
	}
	return sh.RunWith(env, goCmd, "build", "-ldflags", ldflags(), "-o", cliOut, "./cmd/cli")
}

// -----------------------------------------------------------------------------
// Docker Targets
// -----------------------------------------------------------------------------

type Docker mg.Namespace

// Up starts the Docker containers.
func (Docker) Up() error {
	fmt.Println("Starting Docker containers...")
	return sh.Run(dockerCmd, "compose", "up", "-d")
}

// Down stops the Docker containers.
func (Docker) Down() error {
	fmt.Println("Stopping Docker containers...")
	return sh.Run(dockerCmd, "compose", "down")
}

// Logs tails the logs of the Docker containers.
func (Docker) Logs() error {
	fmt.Println("Tailing Docker logs...")
	return sh.Run(dockerCmd, "compose", "logs", "-f")
}

// -----------------------------------------------------------------------------
// Housekeeping & Dependency Management
// -----------------------------------------------------------------------------

// Tidy tidies the go.mod file.
func Tidy() error {
	fmt.Println("Tidying go.mod...")
	return goTidy()
}

// Clean removes all build artifacts and generated files.
func Clean() {
	fmt.Println("Cleaning up...")
	os.RemoveAll("bin")
	os.RemoveAll("public/css")
	os.RemoveAll("coverage.out")
}

// Deps installs all necessary development tools.
func Deps() error {
	fmt.Println("Installing dev dependencies...")
	if err := goInstall("github.com/cosmtrek/air@latest"); err != nil {
		return err
	}
	if err := goInstall("github.com/a-h/templ/cmd/templ@latest"); err != nil {
		return err
	}
	if err := goInstall("github.com/sqlc-dev/sqlc/cmd/sqlc@latest"); err != nil {
		return err
	}
	if err := goInstall("ariga.io/atlas/cmd/atlas@latest"); err != nil {
		return err
	}
	if err := goInstall("github.com/golangci/golangci-lint/cmd/golangci-lint@latest"); err != nil {
		return err
	}
	if err := goInstall("golang.org/x/vuln/cmd/govulncheck@latest"); err != nil {
		return err
	}
	if err := sh.Run(npmCmd, "install"); err != nil {
		return err
	}
	return nil
}

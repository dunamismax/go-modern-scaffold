//go:build mage

package main

import (
	"fmt"
	os"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Default sets the default target for Mage.
var Default = Dev

// Aliases for namespaces
var (
	Build = BuildNS{}
	DB    = DatabaseNS{}
	Gen   = GenerateNS{}
	Dev   = DevNS{}
)

// -----------------------------------------------------------------------------
// Build Targets
// -----------------------------------------------------------------------------

type BuildNS mg.Namespace

// All builds all application binaries and assets.
func (BuildNS) All() error {
	mg.Deps(Build.Server, Build.CLI, Gen.CSS)
	return nil
}

// Server builds the main web server binary.
func (BuildNS) Server() error {
	mg.Deps(Gen.All)
	fmt.Println("🚀 Building server...")
	return sh.Run("go", "build", "-o", "bin/server", "./cmd/server")
}

// CLI builds the command-line interface binary.
func (BuildNS) CLI() error {
	mg.Deps(Gen.All)
	fmt.Println("🚀 Building CLI...")
	return sh.Run("go", "build", "-o", "bin/cli", "./cmd/cli")
}

// -----------------------------------------------------------------------------
// Generation Targets
// -----------------------------------------------------------------------------

type GenerateNS mg.Namespace

// All runs all code generators.
func (GenerateNS) All() {
	mg.Deps(Gen.Templ, Gen.SQLC)
}

// CSS builds the Tailwind CSS.
func (GenerateNS) CSS() error {
	fmt.Println("🎨 Building CSS...")
	return sh.Run("npm", "run", "build:css")
}

// Templ generates Go code from Templ components.
func (GenerateNS) Templ() error {
	fmt.Println("✨ Generating Templ components...")
	return sh.Run("go", "run", "github.com/a-h/templ/cmd/templ", "generate")
}

// SQLC generates Go code from SQL queries.
func (GenerateNS) SQLC() error {
	fmt.Println("🗃️ Generating SQLC code...")
	return sh.Run("go", "run", "github.com/sqlc-dev/sqlc/cmd/sqlc", "generate")
}

// -----------------------------------------------------------------------------
// Development & Execution Targets
// -----------------------------------------------------------------------------

type DevNS mg.Namespace

// All starts all development services.
func (DevNS) All() error {
	fmt.Println("🏃 Starting all dev services...")
	mg.Deps(Gen.All)
	return sh.Run("go", "run", "github.com/air-verse/air", "-c", "air.toml")
}

// CSS starts the Tailwind CSS watcher.
func (DevNS) CSS() error {
	fmt.Println("🎨 Watching CSS...")
	return sh.Run("npm", "run", "dev:css")
}

// -----------------------------------------------------------------------------
// Database Targets
// -----------------------------------------------------------------------------

type DatabaseNS mg.Namespace

// NewMigration creates a new database migration file.
func (DatabaseNS) NewMigration(name string) error {
	fmt.Printf("🆕 Creating migration: %s...\n", name)
	return sh.Run("go", "run", "github.com/pressly/goose/v3/cmd/goose", "-dir", "internal/database/migrations", "create", name, "sql")
}

// Migrate runs all pending database migrations.
func (DatabaseNS) Migrate() error {
	fmt.Println("Applying database migrations...")
	return sh.Run("go", "run", "github.com/pressly/goose/v3/cmd/goose", "-dir", "internal/database/migrations", "up")
}

// Rollback rolls back the last database migration.
func (DatabaseNS) Rollback() error {
	fmt.Println("Rolling back last migration...")
	return sh.Run("go", "run", "github.com/pressly/goose/v3/cmd/goose", "-dir", "internal/database/migrations", "down")
}

// Status shows the status of database migrations.
func (DatabaseNS) Status() error {
	fmt.Println("Checking migration status...")
	return sh.Run("go", "run", "github.com/pressly/goose/v3/cmd/goose", "-dir", "internal/database/migrations", "status")
}

// -----------------------------------------------------------------------------
// Quality & CI/CD Targets
// -----------------------------------------------------------------------------

type Check mg.Namespace

// All runs all quality checks.
func (Check) All() {
	mg.Deps(Check.Format, Check.Vet, Check.Test, Check.Vuln)
}

// Format checks if the code is formatted correctly.
func (Check) Format() error {
	fmt.Println("Checking format...")
	return sh.Run("npx", "prettier", "--check", ".")
}

// Vet runs go vet to find common issues.
func (Check) Vet() error {
	fmt.Println("🔬 Running go vet...")
	return sh.Run("go", "vet", "./...")
}

// Test runs all unit tests.
func (Check) Test() error {
	fmt.Println("🧪 Running tests...")
	return sh.Run("go", "test", "-v", "-cover", "./...")
}

// Vuln scans for vulnerabilities.
func (Check) Vuln() error {
	fmt.Println("Scanning for vulnerabilities...")
	return sh.Run("go", "run", "golang.org/x/vuln/cmd/govulncheck", "./...")
}

// -----------------------------------------------------------------------------
// Housekeeping Targets
// -----------------------------------------------------------------------------

// Install installs all dependencies.
func Install() error {
	mg.Deps(Tidy)
	fmt.Println("Installing frontend dependencies...")
	return sh.Run("npm", "install")
}

// Format formats the source code.
func Format() error {
	fmt.Println("💅 Formatting code...")
	if err := sh.Run("npx", "prettier", "--write", "."); err != nil {
		return err
	}
	return sh.Run("go", "fmt", "./...")
}

// Tidy tidies the go.mod file.
func Tidy() error {
	fmt.Println("🧹 Tidying go.mod...")
	return sh.Run("go", "mod", "tidy")
}

// Clean removes all build artifacts.
func Clean() {
	fmt.Println("🔥 Cleaning up...")
	os.RemoveAll(filepath.Join("bin"))
	os.RemoveAll(filepath.Join("tmp"))
	os.RemoveAll(filepath.Join("dist"))
	os.RemoveAll(filepath.Join("coverage.out"))
}

// Release creates a new release using GoReleaser.
func Release() error {
	mg.Deps(Tidy)
	fmt.Println("Creating release...")
	return sh.Run("goreleaser", "release", "--clean")
}

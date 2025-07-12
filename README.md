<p align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1920px-Go_Logo_Blue.svg.png" alt="The Go programming language logo." width="150"/>
</p>

<p align="center">
  <a href="https://github.com/dunamismax/go-modern-scaffold">
    <img src="https://readme-typing-svg.demolab.com/?font=Fira+Code&size=24&pause=1000&color=00ADD8&center=true&vCenter=true&width=800&lines=The+Go-Modern+Stack;Official+Reference+Implementation;Go+%2B+Fiber+%2B+HTMX+%2B+Tailwind+CSS;Templ%2C+sqlc%2C+and+Mage;Interactive%2C+Performant%2C+and+Beautiful." alt="Typing SVG" />
  </a>
</p>

<p align="center">
  <a href="https://github.com/dunamismax/go-modern-scaffold/actions/workflows/ci.yml"><img src="https://github.com/dunamismax/go-modern-scaffold/actions/workflows/ci.yml/badge.svg" alt="Build Status"></a>
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Go-1.22+-00ADD8.svg" alt="Go Version"></a>
  <a href="https://img.shields.io/github/license/dunamismax/go-modern-scaffold"><img src="https://img.shields.io/github/license/dunamismax/go-modern-scaffold" alt="License"></a>
  <a href="https://img.shields.io/github/repo-size/dunamismax/go-modern-scaffold"><img src="https://img.shields.io/github/repo-size/dunamismax/go-modern-scaffold" alt="Repo Size"></a>
  <a href="https://github.com/dunamismax/go-modern-scaffold/pulls"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg" alt="PRs Welcome"></a>
  <a href="https://github.com/dunamismax/go-modern-scaffold/stargazers"><img src="https://img.shields.io/github/stars/dunamismax/go-modern-scaffold" alt="GitHub Stars"></a>
</p>

---

## About This Project

This monorepo is the official reference implementation for **The Go-Modern Stack**, a complete architecture for building feature-rich, visually appealing, and highly interactive web and command-line applications with Go.

The primary goal is to provide a structured and scalable environment that strategically combines powerful, community-vetted libraries with core Go idioms for a productive and ergonomic development experience. The result is a stack that prioritizes speed, type safety, and modern design without compromising on the core strengths of Go.

---

<details>
<summary><h3>The Go-Modern Stack (Click to Expand)</h3></summary>

---

This stack is architected for developers aiming to build feature-rich, visually appealing, and highly interactive web and command-line applications with Go. It strategically combines powerful, community-vetted libraries with core Go idioms for a productive and ergonomic development experience. The result is a stack that prioritizes speed, type safety, and modern design without compromising on the core strengths of Go.

---

### **Frontend: Rich, Interactive & Beautifully Styled**

This frontend architecture is designed for building modern, animated user experiences with maximum velocity. It combines a Go-native templating engine with a complete ecosystem for styling, component-building, and interactivity, enabling the rapid development of dynamic and polished interfaces.

- [**Templ**](https://templ.guide/)
    - **Role:** Type-Safe, Component-Based HTML Templating.
    - **Description:** A modern templating language that generates Go code from your components, providing compile-time type safety for your HTML. It allows you to build encapsulated, reusable UI elements with pure Go logic, eliminating runtime errors common in string-based templates.
- [**Tailwind CSS**](https://tailwindcss.com/docs/installation)
    - **Role:** Utility-First CSS Framework.
    - **Description:** Enables incredibly fast UI development by providing low-level utility classes that can be composed to build any design directly within your HTML. This removes the need for writing custom CSS and ensures a consistent, scalable styling system.
- [**DaisyUI**](https://daisyui.com/)
    - **Role:** Component Library for Tailwind CSS.
    - **Description:** A plugin for Tailwind CSS that provides a rich set of pre-built, themeable components like buttons, cards, and modals. It drastically speeds up development by allowing you to use simple class names (e.g., `btn-primary`) instead of long strings of utilities.
- [**HTMX**](https://htmx.org/docs/)
    - **Role:** HTML-driven Interactivity.
    - **Description:** Adds dynamic, AJAX-powered interactivity to your application with simple HTML attributes. It allows server-side rendered applications to feel as responsive as a single-page app (SPA) without writing complex JavaScript.
- [**Hyperscript**](https://hyperscript.org/)
    - **Role:** Expressive, Event-Driven Scripting.
    - **Description:** A scripting language designed for modern web development that lives directly in your HTML. It uses an intuitive, English-like syntax to handle user events and DOM manipulations, making it a natural companion to HTMX for client-side interactivity.
- [**HTMX CSS Transitions**](https://htmx.org/examples/animation/)
    - **Role:** Native, Lightweight Animations.
    - **Description:** Leverage HTMX's built-in support for CSS transitions to create smooth animations between page states. By default, HTMX adds classes during its lifecycle, allowing you to easily apply fades, slides, and other effects with pure CSS.
- [**Animate.css**](https://animate.style/)
    - **Role:** Drop-in CSS Animation Library.
    - **Description:** A library of ready-to-use, cross-browser CSS animations. It provides an extensive collection of effects that can be easily triggered by Hyperscript or HTMX events to add polish and visual feedback to your interface.

---

### **Backend: Ergonomic, Performant & Well-Structured**

This backend is built for developer productivity and robust performance, leveraging a high-speed web framework and best-in-class libraries for logging, configuration, and validation.

- [**Fiber**](https://docs.gofiber.io/)
    - **Role:** High-Performance Web Framework.
    - **Description:** An Express.js-inspired web framework built on top of Fasthttp, Go's fastest HTTP engine. It is designed for high-performance and zero memory allocation, providing a developer-friendly API for building APIs and web services rapidly.
- [**slog**](https://pkg.go.dev/log/slog)
    - **Role:** Structured, Level-Based Logging.
    - **Description:** The official structured logging package in Go's standard library. It enables the creation of machine-readable, key-value pair logs with severity levels, which is essential for effective parsing, filtering, and analysis in modern observability platforms.
- [**Viper**](https://github.com/spf13/viper)
    - **Role:** Complete Configuration Management.
    - **Description:** A comprehensive configuration solution for Go applications. Viper can manage configuration from various sources—including YAML, JSON, and TOML files, environment variables, and remote key-value stores—unifying them into a single, accessible interface.
- [**Validator**](https://pkg.go.dev/github.com/go-playground/validator/v10)
    - **Role:** Struct-Tag Based Data Validation.
    - **Description:** The de-facto standard for data validation in Go. It enables declarative validation on struct fields using simple tags (e.g., `validate:"required,email"`), integrating seamlessly with frameworks like Fiber to ensure data integrity.

---

### **TUI (Terminal User Interface): Beautiful & Interactive Command-Line Apps**

For building polished and modern command-line applications, the [**Charm Bracelet**](https://charm.sh/) ecosystem provides a complete and elegant solution.

- [**Bubble Tea**](https://github.com/charmbracelet/bubbletea)
    - **Role:** Stateful TUI Framework.
    - **Description:** Brings The Elm Architecture (a functional, model-view-update pattern) to terminal applications, making it ideal for building complex, interactive, and stateful TUIs that are easy to reason about and maintain.
- [**Bubbles**](https://github.com/charmbracelet/bubbles)
    - **Role:** Reusable TUI Components.
    - **Description:** A library of common, ready-to-use TUI components—such as spinners, text inputs, and tables—that are designed to work with Bubble Tea, dramatically accelerating the development of sophisticated interfaces.
- [**Lipgloss**](https://github.com/charmbracelet/lipgloss)
    - **Role:** Declarative Terminal Styling.
    - **Description:** Offers a fluent, expressive API for styling terminal text. It makes it simple to define colors, layouts, borders, and margins, enabling you to design beautiful and readable TUIs with ease.

---

### **Database & Caching: Type-Safe, Performant & Scalable**

This data layer is optimized for performance and maintainability by pairing direct SQL control with generated, type-safe Go code and a modern, declarative schema migration tool.

- [**sqlc**](https://docs.sqlc.dev/)
    - **Role:** Type-Safe SQL Code Generation.
    - **Description:** Generates fully type-safe, idiomatic Go code from your SQL schema and queries. This allows you to write raw SQL for maximum control and performance while benefiting from compile-time safety, eliminating an entire class of runtime database errors.
- [**Atlas**](https://atlasgo.io/)
    - **Role:** Database Schema Migrations.
    - **Description:** A modern, language-agnostic tool for managing and migrating database schemas. Atlas can automatically generate migration plans by comparing the desired schema (defined in HCL, SQL, or ORM) to the database's current state, streamlining schema evolution with a declarative workflow.
- [**Ristretto**](https://github.com/dgraph-io/ristretto)
    - **Role:** High-Performance In-Process Cache.
    - **Description:** A fast, concurrent, and memory-bounded in-process cache from Dgraph. It is designed to achieve high hit ratios with low memory overhead, making it an excellent choice for performance-critical caching within a single application instance.
- [**go-redis**](https://redis.io/docs/clients/go/)
    - **Role:** Redis Client for Distributed Caching.
    - **Description:** The premier Go client for Redis, providing a high-performance interface for all Redis features. It is essential for implementing a distributed cache, which is critical for scaling applications that require shared state or session management.

---

### **Development Workflow: Automated, Rapid & High-Quality**

A modern toolchain using best-in-class tools to automate common tasks, ensure code quality, and maintain a fast and efficient developer feedback loop.

- [**Mage**](https://magefile.org/)
    - **Role:** Task Runner / Build System.
    - **Description:** An elegant, Make-like tool that allows you to define build tasks (like compiling, testing, or linting) as simple Go functions within a `magefile.go`. This provides a clean, cross-platform, and idiomatic way to automate your project's workflow.
- [**Air**](https://github.com/cosmtrek/air)
    - **Role:** Live Reloading for Development.
    - **Description:** A powerful command-line utility that watches for file changes in your project and automatically recompiles and restarts your application. Air provides a rapid, real-time feedback loop that is essential for productive web development.
- [**GolangCI-Lint**](https://golangci-lint.run/)
    - **Role:** Code Quality Linter Aggregator.
    - **Description:** A fast, configurable linter that runs many Go linters in parallel. It analyzes source code for stylistic issues, bugs, and complexities, enforcing code quality and consistency across the entire project.
- [**Prettier Tailwind CSS Plugin**](https://github.com/tailwindlabs/prettier-plugin-tailwindcss)
    - **Role:** Automatic Class Sorting.
    - **Description:** An official Prettier plugin that automatically sorts your Tailwind CSS classes in a consistent, logical order. This keeps your HTML clean and readable, significantly improving maintainability with zero manual effort.

---

### **Testing: Idiomatic & Dependency-Free**

This stack relies exclusively on Go's powerful, built-in testing framework to ensure code quality and correctness without external dependencies.

- [**Go `testing` Package**](https://pkg.go.dev/testing)
    - **Role:** Core Testing Framework.
    - **Description:** Go's standard library for writing unit, integration, and benchmark tests. Assertions are handled with simple `if` statements and `t.Errorf`, which keeps tests explicit, clear, and easy to maintain while avoiding third-party dependencies.

</details>

---

<p align="center">
  <img src="https://user-images.githubusercontent.com/3185864/32058716-5ee9b512-ba38-11e7-978a-287eb2a62743.png" alt="Gopher Mage." width="150"/>
</p>

## Getting Started

### Prerequisites

- **Go 1.22+**
- **Node.js & npm**
- **Mage**
- **Docker**

### Installation & Usage

1. **Clone the repository:**

   ```bash
   git clone https://github.com/dunamismax/go-modern-scaffold.git
   cd go-modern-scaffold
   ```

2. **Install Dependencies & Tools:**
   This one-time command installs all Go and Node.js dependencies and the required Go development tools.

   ```bash
   mage deps
   ```

3. **Configure Your Environment:**
   Copy the example environment file and customize it as needed for your database and server settings.

   ```bash
   cp .env.example .env
   ```

4. **Start the Database:**
   This command starts the PostgreSQL database container using Docker.

   ```bash
   mage docker:up
   ```

5. **Run Migrations:**
   This command applies all pending database migrations.

   ```bash
   mage db:migrate
   ```

6. **Run the Development Server:**
   This is the primary command for local development. It starts the web server with live reloading and watches for CSS changes.

   ```bash
   mage dev
   ```

   The application will be available at `http://localhost:3000`.

---

## Mage Commands

[Mage](https://magefile.org/) is used to automate common development tasks. The build script is the `magefile.go` at the root of the project. You can list all available commands by running `mage -l`.

### Primary Workflow Commands

- `mage dev`: (Default) Starts the development server with live reload and watches for CSS changes.
- `mage check:all`: **(CI)** Runs all quality checks: lint, test, and vulnerability scan.
- `mage format`: Automatically formats all Go and frontend code.
- `mage build:all`: Builds all application binaries for the current platform.
- `mage release:all`: Cross-compiles release binaries for all platforms.
- `mage docker:up`: Starts the Docker containers for the database.
- `mage docker:down`: Stops the Docker containers.

### Individual Commands

- **Building (`build:`)**
  - `mage build:server`: Builds only the main web server binary.
  - `mage build:cli`: Builds only the command-line interface binary.
- **Code Generation (`generate:`)**
  - `mage generate:all`: Runs all code generators (CSS, Templ, SQLC).
  - `mage generate:css`: Generates the Tailwind CSS file.
  - `mage generate:templ`: Generates Go code from templ components.
  - `mage generate:sqlc`: Generates Go code from SQL queries.
- **Quality Checks (`check:`)**
  - `mage check:lint`: Runs the linter.
  - `mage check:test`: Runs all unit tests.
  - `mage check:cover`: Runs tests and displays coverage in the browser.
  - `mage check:vuln`: Scans for known vulnerabilities.
- **Database (`db:`)**
  - `mage db:migrate`: Applies all pending database migrations.
- **Docker (`docker:`)**
  - `mage docker:up`: Starts the Docker containers.
  - `mage docker:down`: Stops the Docker containers.
  - `mage docker:logs`: Tails the logs of the Docker containers.
- **Housekeeping**
  - `mage clean`: Removes all build artifacts and generated files.
  - `mage tidy`: Tidies the `go.mod` and `go.sum` files.
  - `mage deps`: Installs all necessary development tools.

---

## Project Structure

- **`cmd/`**: Application entry points.
  - **`server/`**: The main Fiber web server.
  - **`cli/`**: The Bubble Tea command-line application.
- **`internal/`**: Private application code.
  - **`cache/`**: Ristretto and Redis caching logic.
  - **`config/`**: Viper configuration management.
  - **`db/`**: Database connection logic, sqlc queries, and models.
    - **`migrations/`**: Atlas schema migrations.
  - **`web/`**: Fiber handlers, Templ components, and CSS styles.
- **`public/`**: Compiled, publicly-served static assets (CSS, JS).
- **`magefile.go`**: The build script for the project, written in Go.
- **`.github/workflows/`**: GitHub Actions CI/CD pipelines.

---

## Contributing

Contributions are welcome! Please feel free to fork the repository, create a feature branch, and open a pull request.

---

### Support My Work

If you find my work on this stack valuable, consider supporting me. It helps me dedicate more time to creating and maintaining open-source projects.

<p align="center">
  <a href="https://coff.ee/dunamismax" target="_blank">
    <img src="https://raw.githubusercontent.com/egonelbre/gophers/master/.thumb/animation/buy-morning-coffee-3x.gif" alt="Buy Me a Coffee" />
  </a>
</p>

---

### Let's Connect

<p align="center">
  <a href="https://twitter.com/dunamismax" target="_blank"><img src="https://img.shields.io/badge/Twitter-%231DA1F2.svg?&style=for-the-badge&logo=twitter&logoColor=white" alt="Twitter"></a>
  <a href="https://bsky.app/profile/dunamismax.bsky.social" target="_blank"><img src="https://img.shields.io/badge/Bluesky-blue?style=for-the-badge&logo=bluesky&logoColor=white" alt="Bluesky"></a>
  <a href="https://reddit.com/user/dunamismax" target="_blank"><img src="https://img.shields.io/badge/Reddit-%23FF4500.svg?&style=for-the-badge&logo=reddit&logoColor=white" alt="Reddit"></a>
  <a href="https://discord.com/users/dunamismax" target="_blank"><img src="https://img.shields.io/badge/Discord-dunamismax-7289DA.svg?style=for-the-badge&logo=discord&logoColor=white" alt="Discord"></a>
  <a href="https://signal.me/#p/+dunamismax.66" target="_blank"><img src="https://img.shields.io/badge/Signal-dunamismax.66-3A76F0.svg?style=for-the-badge&logo=signal&logoColor=white" alt="Signal"></a>
</p>

---

<p align="center">
    <img src="https://raw.githubusercontent.com/egonelbre/gophers/refs/heads/master/.thumb/animation/2bit-sprite/demo.gif" alt="Gopher Sprite Animation" />
</p>

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
# The Go-Modern Stack

This stack is architected for developers aiming to build feature-rich, visually appealing, and highly interactive web and command-line applications with Go. It strategically combines powerful, community-vetted libraries with core Go idioms for a productive and ergonomic development experience. The result is a stack that prioritizes speed, type safety, and modern design without compromising on the core strengths of Go.

---

### **Frontend: Rich, Interactive & Rapidly Styled**

This frontend architecture is designed for building modern user experiences with minimal complexity. It combines a Go-native templating engine with industry-standard styling and interactivity libraries, enabling the rapid development of dynamic interfaces.

- [**Templ**](https://templ.guide/)
    - **Role:** Type-Safe, Component-Based HTML Templating.
    - **Description:** A modern templating language that generates Go code from your components, providing compile-time type safety for your HTML. It allows you to build encapsulated, reusable UI elements with pure Go logic, eliminating runtime errors common in string-based templates.
- [**Tailwind CSS**](https://tailwindcss.com/docs/installation)
    - **Role:** Utility-First CSS Framework.
    - **Description:** Enables incredibly fast UI development by providing low-level utility classes that can be composed to build any design directly within your HTML. This removes the need for writing custom CSS and ensures a consistent, scalable styling system.
- [**HTMX**](https://htmx.org/docs/)
    - **Role:** HTML-driven Interactivity.
    - **Description:** Adds dynamic, AJAX-powered interactivity to your application with simple HTML attributes. It allows server-side rendered applications to feel as responsive as a single-page app (SPA) without writing complex JavaScript.
- [**Alpine.js**](https://alpinejs.dev/start-here)
    - **Role:** Lightweight, Declarative JavaScript.
    - **Description:** A minimal framework for composing client-side behaviors where they are needed most. It is the perfect tool for handling small "islands" of interactivity, such as dropdowns or modals, without the overhead of a larger JavaScript framework.

---

### **Backend: Ergonomic, Performant & Well-Structured**

This backend is built for developer productivity and robust performance, leveraging a full-featured web framework and best-in-class libraries for configuration and validation.

- [**Echo**](https://echo.labstack.com/guide/)
    - **Role:** High-Performance Web Framework.
    - **Description:** A minimalist yet powerful web framework known for its exceptional performance and developer-friendly API. It offers a robust router, extensive middleware, and streamlined data binding, simplifying the creation of scalable REST APIs and web services.
- [**Godotenv**](https://github.com/joho/godotenv)
    - **Role:** Environment Variable Management.
    - **Description:** A helper library that loads environment variables from a `.env` file into the application during development. This simplifies configuration by keeping secrets and environment-specific settings out of source control.
- [**Validator**](https://pkg.go.dev/github.com/go-playground/validator/v10)
    - **Role:** Struct-Tag Based Data Validation.
    - **Description:** The de-facto standard for data validation in Go. It enables declarative validation on struct fields using simple tags (e.g., `validate:"required,email"`), integrating seamlessly with frameworks like Echo to ensure data integrity.

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

This data layer is optimized for performance and maintainability by pairing direct SQL control with generated, type-safe Go code.

- [**sqlc**](https://docs.sqlc.dev/)
    - **Role:** Type-Safe SQL Code Generation.
    - **Description:** Generates fully type-safe, idiomatic Go code from your SQL schema and queries. This allows you to write raw SQL for maximum control and performance while benefiting from compile-time safety, eliminating an entire class of runtime database errors.
- [**Goose**](https://github.com/pressly/goose)
    - **Role:** Database Schema Migrations.
    - **Description:** A robust, CLI-centric tool for managing database schema evolution. You write migrations in plain SQL or Go, giving you precise control over database changes and making your migration process reliable and version-controlled.
- [**Ristretto**](https://github.com/dgraph-io/ristretto)
    - **Role:** High-Performance In-Process Cache.
    - **Description:** A fast, concurrent, and memory-bounded in-process cache from Dgraph. It is designed to achieve high hit ratios with low memory overhead, making it an excellent choice for performance-critical caching within a single application instance.
- [**go-redis**](https://redis.io/docs/clients/go/)
    - **Role:** Redis Client for Distributed Caching.
    - **Description:** The premier Go client for Redis, providing a high-performance interface for all Redis features. It is essential for implementing a distributed cache, which is critical for scaling applications that require shared state or session management.

---

### **Development Workflow: Automated & Rapid**

A modern toolchain using best-in-class Go tools to automate common tasks, ensuring a fast and efficient developer feedback loop.

- [**Mage**](https://magefile.org/)
    - **Role:** Task Runner / Build System.
    - **Description:** An elegant, Make-like tool that allows you to define build tasks (like compiling, testing, or linting) as simple Go functions within a `magefile.go`. This provides a clean, cross-platform, and idiomatic way to automate your project's workflow.
- [**Air**](https://github.com/cosmtrek/air)
    - **Role:** Live Reloading for Development.
    - **Description:** A powerful command-line utility that watches for file changes in your project and automatically recompiles and restarts your application. Air provides a rapid, real-time feedback loop that is essential for productive web development.

---

### **Testing: Idiomatic & Dependency-Free**

This stack relies exclusively on Go's powerful, built-in testing framework to ensure code quality and correctness without external dependencies.

- [**Go `testing` Package**](https://pkg.go.dev/testing)
    - **Role:** Core Testing Framework.
    - **Description:** Go's standard library for writing unit, integration, and benchmark tests. Assertions are handled with simple `if` statements and `t.Errorf`, which keeps tests explicit, clear, and easy to maintain while avoiding third-party dependencies.

---

## Getting Started

### Prerequisites

- Go 1.22+
- Node.js & npm
- Mage

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/dunamismax/go-modern-scaffold.git
    cd go-modern-scaffold
    ```

2.  **Install dependencies:**
    ```bash
    mage install
    ```

3.  **Set up environment variables:**
    Copy the `.env.example` to `.env` and fill in your database connection details.
    ```bash
    cp .env.example .env
    ```

4.  **Run database migrations:**
    ```bash
    mage db:migrate
    ```

### Running the Applications

-   **Web App:**
    To run the web server with live reloading for both Go and Tailwind CSS, run the following commands in separate terminals:
    ```bash
    # Terminal 1: Start the Go server
    mage dev
    ```
    ```bash
    # Terminal 2: Start the Tailwind CSS watcher
    mage dev:css
    ```
    The web application will be available at `http://localhost:8080`.

-   **TUI App:**
    To run the TUI application, build and run the binary:
    ```bash
    mage build:cli
    ./bin/cli
    ```
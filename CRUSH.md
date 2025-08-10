## Build, Lint, and Test

- **Build & Run:** `go run main.go`
- **Run Tailwind CSS:** `npm run dev`
- **Test:** This project does not currently have any tests. 

## Code Style

- **Go:**
    - Run `go fmt` to format code before committing.
    - Use standard Go naming conventions (e.g., `camelCase` for variables, `PascalCase` for exported functions).
    - Handle errors explicitly; avoid using `_` to discard them unless absolutely necessary.
- **Tailwind CSS:**
    - Use the `npm run dev` command to watch for changes and rebuild the CSS.

## Project Structure

- `main.go`: Main application entry point.
- `assets/`: Static assets like CSS, images, and fonts.
- `index.html`: Main HTML file.
- `tailwind.config.js`: Tailwind CSS configuration.
- `go.mod`: Go module definition.
- `package.json`: Node.js dependencies and scripts.

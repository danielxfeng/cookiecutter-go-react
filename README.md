# cookiecutter-go-react

Cookiecutter template for a full stack React (Vite) frontend with an optional Go backend (Chi or Gin).

Both frontend options use the shared workspace package at packages/schemas.

**Quickstart**
```bash
cookiecutter https://github.com/danielxfeng/cookiecutter-go-react.git --checkout main
```

**Options**
- `project_name`
- `author_name`
- `github_username`
- `description`
- `license`
- `frontend_framework` (`react-router` or `react-start`)
- `go_backend` (`gin`, `chi`, or `none`)
- `init_git` (`yes` or `no`)

**After Generation**
```bash
cd <your-project>
pnpm install
pnpm --filter frontend dev   # if frontend_framework=react-router
pnpm --filter start-app dev  # if frontend_framework=react-start
```

**Pre-commit**

Generated projects use Husky + lint-staged.

- Frontend files trigger package precommit checks for the selected frontend stack.
- Go files trigger backend lint/format checks for the selected backend stack.

If you chose the Go backend:
```bash
pnpm --filter backend-go dev # for gin
pnpm --filter backend-chi dev # for chi
```

If you did not initialize git during generation:
```bash
git init
git add .
git commit -m "Initial commit"
```

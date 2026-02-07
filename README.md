# cookiecutter-go-react

Cookiecutter template for a full stack React (Vite) frontend with an optional Go (Gin) backend.

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
- `include_go_backend` (`yes` or `no`)

**After Generation**
```bash
cd <your-project>
pnpm install
pnpm --filter frontend dev
```

If you chose the Go backend:
```bash
pnpm --filter backend-go dev
```

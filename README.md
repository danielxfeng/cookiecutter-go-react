# cookiecutter-go-react

Cookiecutter template for a full stack React (Vite) frontend with an optional Go backend (Gin or Chi).

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
- `go_backend` (`gin`, `chi`, or `none`)
- `init_git` (`yes` or `no`)

**After Generation**
```bash
cd <your-project>
pnpm install
pnpm --filter frontend dev
```

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

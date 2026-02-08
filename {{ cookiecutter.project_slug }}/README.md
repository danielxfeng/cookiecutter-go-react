# {{ cookiecutter.project_name }}

{{ cookiecutter.description }}

**Stack**
- React (Vite) frontend
{% if cookiecutter.go_backend == "gin" -%}
- Go backend (Gin)
{% elif cookiecutter.go_backend == "chi" -%}
- Go backend (Chi)
{% endif -%}

**Prereqs**
- Node.js 20+
- pnpm 10+
{% if cookiecutter.go_backend != "none" -%}
- Go 1.22+
{% endif -%}

**Install**
```bash
pnpm install
```

{% if cookiecutter.init_git == "no" -%}
**Git**
```bash
git init
git add .
git commit -m "Initial commit"
```

{% endif -%}
**Run**
Frontend dev server:
```bash
pnpm --filter frontend dev
```

{% if cookiecutter.go_backend == "gin" -%}
Backend dev server:
```bash
pnpm --filter backend-go dev
```
{% elif cookiecutter.go_backend == "chi" -%}
Backend dev server:
```bash
pnpm --filter backend-chi dev
```
{% endif -%}

**Test**
```bash
pnpm -r run test
```

**Build**
```bash
pnpm -r run build
```

{% if cookiecutter.go_backend != "none" -%}
**Go Module**

The Go module path is:
```
{% if cookiecutter.go_backend == "gin" -%}
github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/apps/backend-go
{% elif cookiecutter.go_backend == "chi" -%}
github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/apps/backend-chi
{% endif -%}
```
{% endif -%}

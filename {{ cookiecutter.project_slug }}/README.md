# {{ cookiecutter.project_name }}

{{ cookiecutter.description }}

**Stack**
- React (Vite) frontend
{% if cookiecutter.include_go_backend == "yes" -%}
- Go backend (Gin)
{% endif -%}

**Prereqs**
- Node.js 20+
- pnpm 10+
{% if cookiecutter.include_go_backend == "yes" -%}
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

{% if cookiecutter.include_go_backend == "yes" -%}
Backend dev server:
```bash
pnpm --filter backend-go dev
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

{% if cookiecutter.include_go_backend == "yes" -%}
**Go Module**

The Go module path is:
```
github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/apps/backend-go
```
{% endif -%}

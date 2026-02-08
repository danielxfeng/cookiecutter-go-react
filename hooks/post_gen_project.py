#!/usr/bin/env python3
import json
import os
import shutil
import subprocess
from pathlib import Path


def _truthy(value: str) -> bool:
    return value.strip().lower() in {"y", "yes", "true", "1"}


def _remove_backend_go() -> None:
    backend_dir = Path("apps/backend-go")
    if backend_dir.exists():
        shutil.rmtree(backend_dir)


def _remove_backend_chi() -> None:
    backend_dir = Path("apps/backend-chi")
    if backend_dir.exists():
        shutil.rmtree(backend_dir)


def _strip_go_lint_staged(selected_backend: str) -> None:
    package_json_path = Path("package.json")
    if not package_json_path.exists():
        return

    data = json.loads(package_json_path.read_text())
    lint_staged = data.get("lint-staged")
    if not isinstance(lint_staged, dict):
        package_json_path.write_text(json.dumps(data, indent=2) + "\n")
        return

    if selected_backend != "gin":
        lint_staged.pop("apps/backend-go/**/*.go", None)
    if selected_backend != "chi":
        lint_staged.pop("apps/backend-chi/**/*.go", None)
    data["lint-staged"] = lint_staged
    package_json_path.write_text(json.dumps(data, indent=2) + "\n")


def _remove_backend_workflows(selected_backend: str) -> None:
    workflows_dir = Path(".github/workflows")
    if not workflows_dir.exists():
        return

    if selected_backend != "gin":
        backend_go = workflows_dir / "ci-backend-go.yml"
        if backend_go.exists():
            backend_go.unlink()

    if selected_backend != "chi":
        backend_chi = workflows_dir / "ci-backend-chi.yml"
        if backend_chi.exists():
            backend_chi.unlink()


if __name__ == "__main__":
    selected_backend = "{{ cookiecutter.go_backend }}"
    if selected_backend == "none":
        _remove_backend_go()
        _remove_backend_chi()
    elif selected_backend == "gin":
        _remove_backend_chi()
    elif selected_backend == "chi":
        _remove_backend_go()

    _strip_go_lint_staged(selected_backend)
    _remove_backend_workflows(selected_backend)

    init_git = "{{ cookiecutter.init_git }}"
    if _truthy(init_git):
        subprocess.run(["git", "init"], check=False)

    print("\nNext steps:")
    print("  cd {{ cookiecutter.project_slug }}")
    print("  pnpm install")
    print("  pnpm --filter frontend dev")
    if selected_backend == "gin":
        print("  pnpm --filter backend-go dev")
    elif selected_backend == "chi":
        print("  pnpm --filter backend-chi dev")

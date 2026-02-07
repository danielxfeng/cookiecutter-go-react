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


def _strip_go_lint_staged() -> None:
    package_json_path = Path("package.json")
    if not package_json_path.exists():
        return

    data = json.loads(package_json_path.read_text())
    lint_staged = data.get("lint-staged")
    if not isinstance(lint_staged, dict):
        package_json_path.write_text(json.dumps(data, indent=2) + "\n")
        return

    lint_staged.pop("apps/backend-go/**/*.go", None)
    data["lint-staged"] = lint_staged
    package_json_path.write_text(json.dumps(data, indent=2) + "\n")


if __name__ == "__main__":
    include_go_backend = "{{ cookiecutter.include_go_backend }}"
    if not _truthy(include_go_backend):
        _remove_backend_go()
        _strip_go_lint_staged()

    init_git = "{{ cookiecutter.init_git }}"
    if _truthy(init_git):
        subprocess.run(["git", "init"], check=False)

    print("\nNext steps:")
    print("  cd {{ cookiecutter.project_slug }}")
    print("  pnpm install")
    print("  pnpm --filter frontend dev")
    if _truthy(include_go_backend):
        print("  pnpm --filter backend-go dev")

# Angular Failure Patterns

`intelligence/engineering/language-specific/angular/failure/` stores Angular-specific failure
patterns and diagnostic knowledge.

## Scope

This directory is responsible for:

- Angular workspace configuration pitfalls (`angular.json`, `tsconfig` path mappings, library projects)
- Build and codegen interactions that fail silently rather than loudly
- Angular runtime or tooling behaviour that differs from other ecosystems

## Relationship to Other Layers

- `intelligence/engineering/analytical-reasoning/failure/` stores cross-language or
  analysis-technique failures; this directory stores failures specific to Angular
- `intelligence/engineering/ui/` stores UI/layout reasoning that is not framework-bound
- `enforcement/failure-learning-system.md` defines the generic failure learning framework

## Current Atoms

| Atom | Description | Source |
|------|-------------|--------|
| [`angular-library-rename-config-only-sites.md`](angular-library-rename-config-only-sites.md) | Angular library rename has declaration sites no type checker reads — `angular.json`, `package.json` scripts and codegen `output`; a stale codegen output path recreates the deleted library directory silently | validated during a library re-bucketing (2026-08-05) |

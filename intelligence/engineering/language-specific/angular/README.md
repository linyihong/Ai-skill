# Angular Intelligence

Angular-specific intelligence atoms for engineering patterns, failures, and best practices.

## Why a framework sits under `language-specific/`

Angular is a framework, not a language. It lives here because the axis this directory actually
separates is "knowledge that only applies inside one ecosystem" from cross-language analytical
technique — and an Angular workspace quirk fails that test the same way a Java standard-library
quirk does. Splitting frameworks into their own tree would duplicate the taxonomy for no gain at
one atom. Revisit if framework atoms outgrow language ones.

## Scope

- Angular workspace and build-configuration pitfalls
- Angular-specific failure patterns and solutions

## Current Atoms

- [`angular-library-rename-config-only-sites.md`](failure/angular-library-rename-config-only-sites.md) — a library rename is declared in several places the type checker never reads, so a partial rename builds green and a stale codegen `output` silently recreates the directory that was deleted

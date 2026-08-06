# Angular Library Rename: The Sites the Compiler Cannot Catch

> **Framework**: Angular workspace (framework-specific; see parent README on why it sits under `language-specific/`)
> **Layer**: `intelligence/engineering/language-specific/angular/failure/`
> **Classification reason**: Root cause is where an Angular workspace declares a library — several declaration sites are JSON that no type checker reads, so a partially-applied rename compiles and tests green while still being wrong.

## Problem

Renaming or moving a library in an Angular workspace looks like a mechanical find-and-replace. A green build afterwards does not mean it is complete: the workspace declares a library in several places, and only some of them are checked by anything.

## Root Cause

The declaration sites split into two groups:

**Checked — a miss fails the build, so it self-reports:**

| Site | What breaks |
| --- | --- |
| `tsconfig.base.json` `paths` | import of the package name fails to resolve |
| `tsconfig.lib.base.json` `paths` (dist mapping) | library-to-library import fails |
| `apps/<app>/tsconfig.spec.json` `paths` | specs fail to resolve |
| import specifiers in app and test code | compile error |

**Unchecked — a miss is silent:**

| Site | What actually happens |
| --- | --- |
| `angular.json` project entry (`root`, `sourceRoot`, `ng-package.json`, both `tsConfig` paths) | `ng build <name>` fails only when that library is built — not during a normal app build |
| `package.json` `build:<name>` script **and the aggregate script** (`build:with-libs`) — the name appears twice in one file | the aggregate build keeps invoking a project name that no longer exists, or silently omits the library |
| codegen `output` paths (e.g. `ng-openapi-gen/*.json`) | **worst case: the next codegen run recreates the deleted directory** and writes generated clients into it. Nothing fails. A directory that was deliberately removed comes back with plausible content in it. |
| library's own `package.json` `name`, `README` | published/consumed name disagrees with the workspace |

The codegen row is the one that bites, because the failure mode is resurrection rather than an error, and it appears one command later — usually in someone else's session.

## Example

Moving `libs/erp/kaizen-erp-work-order` → `libs/wms/kaizen-wms-work-order-supply`:

```jsonc
// web/ng-openapi-gen/wms.m-wos.json — still points at the old directory
{
  "input":  "../openapi/wms/modules/m-wos.openapi.json",   // correct, untouched
  "output": "libs/erp/kaizen-erp-work-order/src/lib/api/generated"  // stale — recreates libs/erp/
}
```

A bulk rename loop can also skip files without failing. Applying `perl -pi -e` over a file list built by `grep -rl` picked up the library sources but silently missed `tsconfig.base.json`, `package.json` and `angular.json` — the loop reported nothing and the build still passed, because those three are in the unchecked group.

## Fix

Do not trust the rename loop or the build. Verify by absence:

```bash
# Must return nothing. Exclude build caches, which legitimately hold the old name.
grep -rn "old-lib-name\|libs/old/path" web/ tests/ docs/ \
  | grep -v node_modules | grep -v "/dist/" | grep -v "\.angular"
```

Then build the moved library by name (`ng build <new-name>`), not just the app — that is the only thing that exercises the `angular.json` entry.

## Applies When

- Renaming, moving, or re-bucketing a library in an Angular workspace
- Any workspace where library identity is declared in both TypeScript config and untyped JSON (Nx and plain Angular CLI both qualify)
- A codegen tool writes into the library directory

## Does Not Apply When

- Renaming only a symbol or file inside a library — the compiler covers that fully
- The workspace has a single application and no library projects

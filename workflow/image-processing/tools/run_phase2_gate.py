#!/usr/bin/env python3
"""Phase 2 gate: required negative fixtures must not become accepted."""

from __future__ import annotations

import argparse
import json
import subprocess
import sys
from pathlib import Path

# Phase 2 minimum; Phase 3 expands via run_phase3_gate.py
REQUIRED = ("neg.alpha-opaque", "neg.residue-halo")


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--repo", type=Path, default=None)
    parser.add_argument("--skip-generate", action="store_true")
    args = parser.parse_args()

    repo = args.repo or Path(__file__).resolve().parents[3]
    tools = repo / "workflow" / "image-processing" / "tools"
    inputs = repo / "workflow" / "image-processing" / "fixtures" / "inputs"
    reports = repo / "workflow" / "image-processing" / "fixtures" / "reports"
    reports.mkdir(parents=True, exist_ok=True)

    py = sys.executable
    if not args.skip_generate:
        subprocess.check_call(
            [py, str(tools / "generate_negative_fixtures.py"), "--repo", str(repo)]
        )

    results = []
    failed = False
    for name in REQUIRED:
        png = inputs / f"{name}.png"
        if not png.is_file():
            print(f"missing fixture {png}", file=sys.stderr)
            failed = True
            continue
        out = reports / f"{name}.verify.json"
        proc = subprocess.run(
            [
                py,
                str(tools / "verify_independent.py"),
                str(png),
                "--profile",
                name,
                "--json-out",
                str(out),
                "--expect-block",
            ],
            capture_output=True,
            text=True,
        )
        results.append(
            {
                "fixture": name,
                "exit_code": proc.returncode,
                "stdout": proc.stdout,
                "stderr": proc.stderr,
            }
        )
        if proc.returncode != 0:
            failed = True
            print(f"BLOCKED? NO — {name} (exit {proc.returncode})", file=sys.stderr)
        else:
            print(f"OK blocked false success: {name}")

    summary = {
        "gate": "phase2.negative_fixtures",
        "required": list(REQUIRED),
        "passed": not failed,
        "results": [
            {
                "fixture": r["fixture"],
                "exit_code": r["exit_code"],
                "report": json.loads(r["stdout"])
                if r["stdout"].strip().startswith("{")
                else None,
            }
            for r in results
        ],
    }
    summary_path = reports / "phase2-gate-latest.json"
    summary_path.write_text(json.dumps(summary, indent=2) + "\n", encoding="utf-8")
    print(f"wrote {summary_path}")
    if failed:
        print("Phase 2 Gate FAILED", file=sys.stderr)
        return 1
    print("Phase 2 Gate PASSED")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

#!/usr/bin/env python3
"""Phase 3 gate: full negative matrix + rembg synthetic dogfood IDR."""

from __future__ import annotations

import argparse
import json
import subprocess
import sys
from pathlib import Path

NEGATIVES = (
    "neg.alpha-opaque",
    "neg.residue-halo",
    "neg.interior-damage",
    "neg.over-erode",
    "neg.consumer-layout",
)


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--repo", type=Path, default=None)
    parser.add_argument("--skip-generate", action="store_true")
    parser.add_argument("--skip-rembg", action="store_true")
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
        subprocess.check_call(
            [py, str(tools / "generate_dogfood_sources.py"), "--repo", str(repo)]
        )

    failed = False
    neg_results = []
    for name in NEGATIVES:
        png = inputs / f"{name}.png"
        out = reports / f"{name}.verify.json"
        cmd = [
            py,
            str(tools / "verify_independent.py"),
            str(png),
            "--profile",
            name,
            "--json-out",
            str(out),
            "--expect-block",
        ]
        proc = subprocess.run(cmd, capture_output=True, text=True)
        ok = proc.returncode == 0
        if not ok:
            failed = True
            print(f"NEG FAIL {name}: {proc.stderr}", file=sys.stderr)
        else:
            print(f"OK blocked: {name}")
        neg_results.append(
            {
                "fixture": name,
                "exit_code": proc.returncode,
                "report": json.loads(proc.stdout)
                if proc.stdout.strip().startswith("{")
                else None,
            }
        )

    rembg_meta = None
    if not args.skip_rembg:
        proc = subprocess.run(
            [py, str(tools / "run_rembg_candidate.py"), "--repo", str(repo)],
            capture_output=True,
            text=True,
        )
        if proc.returncode != 0:
            failed = True
            print(proc.stderr, file=sys.stderr)
            print(proc.stdout, file=sys.stderr)
        else:
            rembg_meta = json.loads(proc.stdout) if proc.stdout.strip().startswith("{") else {
                "raw": proc.stdout
            }
            print(f"OK rembg dogfood: {rembg_meta}")

    summary = {
        "gate": "phase3.dogfood_and_negatives",
        "negatives_required": list(NEGATIVES),
        "negatives_passed": all(r["exit_code"] == 0 for r in neg_results),
        "rembg_dogfood": rembg_meta,
        "passed": not failed,
        "results": neg_results,
        "learning": {
            "guard_promotion": "forbidden_without_human_review",
            "note": "Failure patterns stay candidate until human-reviewed",
        },
    }
    path = reports / "phase3-gate-latest.json"
    path.write_text(json.dumps(summary, indent=2) + "\n", encoding="utf-8")
    print(f"wrote {path}")
    if failed:
        print("Phase 3 Gate FAILED", file=sys.stderr)
        return 1
    print("Phase 3 Gate PASSED")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

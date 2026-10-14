#!/usr/bin/env python3
"""Phase 4 gate: B/C/D fragments exist; D pos/neg consumer checks; loop file unchanged."""

from __future__ import annotations

import argparse
import json
import subprocess
import sys
from pathlib import Path

FRAGMENTS = (
    "contracts/task-class-b.yaml",
    "contracts/task-class-c.yaml",
    "contracts/task-class-d.yaml",
)


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--repo", type=Path, default=None)
    parser.add_argument("--skip-generate", action="store_true")
    args = parser.parse_args()
    repo = args.repo or Path(__file__).resolve().parents[3]
    root = repo / "workflow" / "image-processing"
    tools = root / "tools"
    inputs = root / "fixtures" / "inputs"
    reports = root / "fixtures" / "reports"
    reports.mkdir(parents=True, exist_ok=True)
    py = sys.executable
    failed = False

    # 1) fragments present
    missing = [p for p in FRAGMENTS if not (root / p).is_file()]
    if missing:
        print(f"missing fragments: {missing}", file=sys.stderr)
        failed = True
    else:
        print("OK fragments present")

    # 2) lifecycle still has shared steps 0-9 marker
    flow = (root / "execution-flow.md").read_text(encoding="utf-8")
    if "9. Finality + IDR Closeout" not in flow or "0b. Lazy-load" not in flow:
        print("execution-flow missing shared loop or 0b lazy-load", file=sys.stderr)
        failed = True
    else:
        print("OK shared loop + 0b lazy-load")

    if not args.skip_generate:
        subprocess.check_call(
            [py, str(tools / "generate_negative_fixtures.py"), "--repo", str(repo)]
        )

    # 3) D negative still blocked
    neg = inputs / "neg.consumer-layout.png"
    proc = subprocess.run(
        [
            py,
            str(tools / "verify_independent.py"),
            str(neg),
            "--profile",
            "neg.consumer-layout",
            "--json-out",
            str(reports / "neg.consumer-layout.verify.json"),
            "--expect-block",
        ],
        capture_output=True,
        text=True,
    )
    if proc.returncode != 0:
        failed = True
        print("D neg consumer-layout not blocked", file=sys.stderr)
    else:
        print("OK D neg blocked")

    # 4) D positive accepted (canvas match + alpha ok)
    pos = inputs / "pos.d-canvas-328.png"
    proc = subprocess.run(
        [
            py,
            str(tools / "verify_independent.py"),
            str(pos),
            "--profile",
            "pos.d-canvas-328",
            "--json-out",
            str(reports / "pos.d-canvas-328.verify.json"),
        ],
        capture_output=True,
        text=True,
    )
    report = json.loads(proc.stdout) if proc.stdout.strip().startswith("{") else {}
    if report.get("finality") != "accepted":
        failed = True
        print(f"D pos not accepted: {report}", file=sys.stderr)
        print(proc.stderr, file=sys.stderr)
    else:
        print("OK D pos accepted")

    summary = {
        "gate": "phase4.bcd_fragments",
        "fragments": list(FRAGMENTS),
        "passed": not failed,
        "d_pos_finality": report.get("finality"),
    }
    out = reports / "phase4-gate-latest.json"
    out.write_text(json.dumps(summary, indent=2) + "\n", encoding="utf-8")
    print(f"wrote {out}")
    if failed:
        print("Phase 4 Gate FAILED", file=sys.stderr)
        return 1
    print("Phase 4 Gate PASSED")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

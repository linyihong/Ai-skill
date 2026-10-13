#!/usr/bin/env python3
"""
Execute seg.generic via rembg on a synthetic dogfood source, then independent verify.

Writes an IDR-shaped JSON (not self-PASS). License checklist must be considered
before production use — see adapters/license-checklist.md.
"""

from __future__ import annotations

import argparse
import hashlib
import json
import time
from pathlib import Path


def sha256_bytes(data: bytes) -> str:
    return hashlib.sha256(data).hexdigest()


def relpath(path: Path, repo: Path) -> str:
    try:
        return str(path.resolve().relative_to(repo.resolve()))
    except ValueError:
        return path.name


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--repo", type=Path, default=None)
    parser.add_argument(
        "--source",
        type=Path,
        default=None,
        help="Defaults to dogfood.object-hard.png",
    )
    args = parser.parse_args()
    repo = args.repo or Path(__file__).resolve().parents[3]
    tools = Path(__file__).resolve().parent
    dogfood = repo / "workflow" / "image-processing" / "fixtures" / "dogfood"
    sources = dogfood / "sources"
    outputs = dogfood / "outputs"
    reports = dogfood / "reports"
    outputs.mkdir(parents=True, exist_ok=True)
    reports.mkdir(parents=True, exist_ok=True)

    source = args.source or (sources / "dogfood.object-hard.png")
    if not source.is_file():
        raise SystemExit(f"missing source {source}; run generate_dogfood_sources.py")

    from rembg import remove  # local import — optional heavy dep

    raw = source.read_bytes()
    t0 = time.perf_counter()
    out_bytes = remove(raw)
    elapsed = time.perf_counter() - t0

    out_path = outputs / f"{source.stem}.rembg.png"
    out_path.write_bytes(out_bytes)
    out_path.with_suffix(".sha256").write_text(
        sha256_bytes(out_bytes) + "\n", encoding="utf-8"
    )

    # Independent verify (separate process semantics: call module functions)
    import verify_independent as vi

    verify = vi.verify_path(out_path, profile="dogfood.rembg")
    # Soft review for synthetic dogfood without ground-truth quality
    if verify.get("finality") == "accepted":
        verify["finality"] = "accepted_with_review"
        verify["disposition"] = "queue_human_review"
        verify["reason"] = "synthetic_dogfood_no_gt_human_review_required"
        for c in verify["checks"]:
            if c["check_id"] == "alpha.not_fully_opaque" and c["status"] == "pass":
                c["review_required"] = True

    idr = {
        "contract_ref": {
            "task_class": "A",
            "background": "transparent",
            "fidelity_preset": "strict",
            "note": "synthetic dogfood; not product art",
        },
        "analysis": {
            "subject_count": "single",
            "edge_type": "hard",
            "content_type": "illustration",
            "risk_flags": [],
            "confidence": "high",
        },
        "feasible": ["seg.generic", "post.deterministic"],
        "selected": {
            "candidate_ids": ["seg.generic"],
            "rationale": "Phase 3 baseline_candidate rembg bind",
        },
        "rejected": [
            {
                "candidate_id": "gen.edit",
                "reason": "fidelity strict / not in feasible set",
            }
        ],
        "execution": {
            "adapter": "rembg.remove",
            "params_summary": "default session",
            "output_path": relpath(out_path, repo),
            "output_hash": sha256_bytes(out_bytes),
            "elapsed_sec": round(elapsed, 3),
            "source_path": relpath(source, repo),
            "source_hash": sha256_bytes(raw),
        },
        "verification": verify.get("checks", []),
        "repairs": [],
        "finality": verify.get("finality"),
        "disposition": verify.get("disposition"),
        "notes": (
            "Guard promotion forbidden without human review. "
            "Compare only against other candidates in future runs."
        ),
        "metrics": {
            "elapsed_sec": round(elapsed, 3),
            "output_bytes": len(out_bytes),
            "cost": "local_cpu_unknown",
        },
    }

    idr_path = reports / f"{source.stem}.rembg.idr.json"
    idr_path.write_text(json.dumps(idr, indent=2) + "\n", encoding="utf-8")
    verify_path = reports / f"{source.stem}.rembg.verify.json"
    verify_path.write_text(json.dumps(verify, indent=2) + "\n", encoding="utf-8")
    print(json.dumps({"idr": relpath(idr_path, repo), "finality": idr["finality"], "elapsed_sec": idr["metrics"]["elapsed_sec"]}, indent=2))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

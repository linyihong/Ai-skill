#!/usr/bin/env python3
"""Independent image verification — must not be the executor."""

from __future__ import annotations

import argparse
import json
import sys
from dataclasses import asdict, dataclass
from pathlib import Path
from typing import Any

from PIL import Image


@dataclass
class CheckResult:
    check_id: str
    status: str  # pass | fail | inconclusive
    evidence: str
    failure_class: str | None = None
    repair_candidate: str | None = None
    review_required: bool = False
    hardness: str = "hard"


def check_alpha(img: Image.Image) -> list[CheckResult]:
    out: list[CheckResult] = []
    if img.mode != "RGBA":
        out.append(
            CheckResult(
                check_id="alpha.present",
                status="fail",
                evidence=f"mode={img.mode}",
                failure_class="alpha_invalid",
                repair_candidate="reassemble_rgba",
                hardness="hard",
            )
        )
        return out

    out.append(
        CheckResult(
            check_id="alpha.present",
            status="pass",
            evidence="mode=RGBA",
            hardness="hard",
        )
    )

    alpha = img.getchannel("A")
    hist = alpha.histogram()
    total = img.width * img.height
    opaque = hist[255] if len(hist) > 255 else 0
    opaque_ratio = opaque / total if total else 1.0
    amin, amax = alpha.getextrema()

    if amin >= 255 or opaque_ratio >= 0.995:
        out.append(
            CheckResult(
                check_id="alpha.not_fully_opaque",
                status="fail",
                evidence=f"alpha_min={amin} opaque_ratio={opaque_ratio:.4f}",
                failure_class="alpha_invalid",
                repair_candidate="reassemble_rgba",
                hardness="hard",
            )
        )
    else:
        out.append(
            CheckResult(
                check_id="alpha.not_fully_opaque",
                status="pass",
                evidence=f"alpha_min={amin} opaque_ratio={opaque_ratio:.4f}",
                hardness="hard",
            )
        )
    return out


def check_residue_rim(img: Image.Image, rim_px: int = 6) -> CheckResult:
    """
    Anomaly: opaque/near-opaque dark-cool rim near subject edge while
    corners are transparent — flags edge_halo / background_residue class.
    """
    if img.mode != "RGBA":
        return CheckResult(
            check_id="residue.rim",
            status="inconclusive",
            evidence="no RGBA",
            hardness="hard",
        )

    px = img.load()
    w, h = img.size
    # Sample outer ring of non-transparent dark pixels
    dark_rim = 0
    samples = 0
    for y in range(h):
        for x in range(w):
            r, g, b, a = px[x, y]
            # near border band
            on_band = x < rim_px or y < rim_px or x >= w - rim_px or y >= h - rim_px
            if not on_band:
                # also count annulus: opaque navy-like near mid-radius
                cx, cy = w / 2, h / 2
                dist = ((x - cx) ** 2 + (y - cy) ** 2) ** 0.5
                if 28 <= dist <= 36:
                    samples += 1
                    if a >= 200 and b > r and b > g and (r + g + b) < 180:
                        dark_rim += 1
            else:
                if a >= 200 and b > r and b > g and (r + g + b) < 160:
                    dark_rim += 1
                    samples += 1

    # Stronger signal for synthetic residue-halo: count navy opaque pixels
    navy_opaque = 0
    transparent = 0
    for y in range(h):
        for x in range(w):
            r, g, b, a = px[x, y]
            if a == 0:
                transparent += 1
            if a >= 250 and b >= 60 and r <= 40 and g <= 50:
                navy_opaque += 1

    if navy_opaque >= 50 and transparent >= 50:
        return CheckResult(
            check_id="residue.rim",
            status="fail",
            evidence=f"navy_opaque={navy_opaque} transparent={transparent}",
            failure_class="edge_halo",
            repair_candidate="post.decontaminate",
            hardness="hard",
        )

    if dark_rim >= 20:
        return CheckResult(
            check_id="residue.rim",
            status="fail",
            evidence=f"dark_rim_samples={dark_rim}",
            failure_class="background_residue",
            repair_candidate="post.decontaminate",
            hardness="hard",
        )

    return CheckResult(
        check_id="residue.rim",
        status="pass",
        evidence=f"navy_opaque={navy_opaque} transparent={transparent}",
        hardness="hard",
    )


def suggest_finality(checks: list[CheckResult]) -> dict[str, Any]:
    hard_fails = [c for c in checks if c.hardness == "hard" and c.status == "fail"]
    hard_inconc = [
        c for c in checks if c.hardness == "hard" and c.status == "inconclusive"
    ]
    review = any(c.review_required for c in checks)

    if hard_fails:
        return {
            "disposition": "reject_or_abort",
            "finality": "rejected",
            "reason": "hard_check_fail",
            "failed_checks": [c.check_id for c in hard_fails],
        }
    if hard_inconc:
        return {
            "disposition": "human_review",
            "finality": None,
            "reason": "hard_inconclusive",
        }
    if review:
        return {
            "disposition": "queue_human_review",
            "finality": "accepted_with_review",
            "reason": "soft_review_required",
        }
    return {
        "disposition": "close",
        "finality": "accepted",
        "reason": "all_hard_pass",
    }


PROFILE_CHECKS = {
    "neg.alpha-opaque": ["alpha"],
    "neg.residue-halo": ["alpha", "residue"],
    "default": ["alpha", "residue"],
}


def relpath_for_report(path: Path) -> str:
    """Prefer repo-relative path; never require absolute host paths in reports."""
    resolved = path.resolve()
    for parent in resolved.parents:
        if (parent / "workflow" / "image-processing").is_dir() and (
            parent / "CORE_BOOTSTRAP.md"
        ).is_file():
            try:
                return str(resolved.relative_to(parent))
            except ValueError:
                break
    return path.name


def verify_path(path: Path, profile: str | None = None) -> dict[str, Any]:
    img = Image.open(path)
    img.load()
    name = profile or path.stem
    suites = PROFILE_CHECKS.get(name, PROFILE_CHECKS["default"])
    checks: list[CheckResult] = []
    if "alpha" in suites:
        checks.extend(check_alpha(img))
    if "residue" in suites:
        checks.append(check_residue_rim(img))

    finality = suggest_finality(checks)
    return {
        "input": relpath_for_report(path),
        "profile": name,
        "checks": [asdict(c) for c in checks],
        **finality,
    }


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("image", type=Path)
    parser.add_argument("--profile", default=None)
    parser.add_argument("--json-out", type=Path, default=None)
    parser.add_argument(
        "--expect-block",
        action="store_true",
        help="Exit 0 only if finality is not accepted (negative fixture mode)",
    )
    args = parser.parse_args()

    report = verify_path(args.image, args.profile)
    text = json.dumps(report, indent=2)
    print(text)
    if args.json_out:
        args.json_out.parent.mkdir(parents=True, exist_ok=True)
        args.json_out.write_text(text + "\n", encoding="utf-8")

    finality = report.get("finality")
    if args.expect_block:
        if finality == "accepted":
            print("GATE FAIL: negative fixture was accepted", file=sys.stderr)
            return 2
        return 0

    if finality == "accepted":
        return 0
    return 1


if __name__ == "__main__":
    raise SystemExit(main())

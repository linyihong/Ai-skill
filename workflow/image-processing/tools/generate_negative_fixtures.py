#!/usr/bin/env python3
"""Synthesize Phase 2 negative fixtures (desensitized, no product art)."""

from __future__ import annotations

import argparse
import hashlib
import json
from pathlib import Path

from PIL import Image, ImageDraw


def sha256_file(path: Path) -> str:
    h = hashlib.sha256()
    h.update(path.read_bytes())
    return h.hexdigest()


def write_png(path: Path, img: Image.Image) -> str:
    path.parent.mkdir(parents=True, exist_ok=True)
    img.save(path, format="PNG")
    digest = sha256_file(path)
    path.with_suffix(".sha256").write_text(digest + "\n", encoding="utf-8")
    return digest


def make_alpha_opaque(size: int = 64) -> Image.Image:
    """RGBA with alpha channel present but fully opaque (false transparency)."""
    img = Image.new("RGBA", (size, size), (40, 120, 200, 255))
    draw = ImageDraw.Draw(img)
    draw.ellipse((12, 12, size - 12, size - 12), fill=(220, 180, 60, 255))
    return img


def make_residue_halo(size: int = 96) -> Image.Image:
    """
    Subject roughly cut out but navy rim remains opaque/semi-opaque on
    what should be transparent background — classic false success.
    """
    img = Image.new("RGBA", (size, size), (0, 0, 0, 0))
    draw = ImageDraw.Draw(img)
    cx = cy = size // 2
    # Contaminating rim (navy) outside true subject
    draw.ellipse((cx - 34, cy - 34, cx + 34, cy + 34), fill=(10, 20, 80, 255))
    # Subject core
    draw.ellipse((cx - 26, cy - 26, cx + 26, cy + 26), fill=(230, 190, 70, 255))
    # Outer field left partially transparent — but rim annulus stays
    # Punch true background corners already transparent; leave rim ring.
    return img


FIXTURES = {
    "neg.alpha-opaque": make_alpha_opaque,
    "neg.residue-halo": make_residue_halo,
}


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "--out-dir",
        type=Path,
        default=None,
        help="Defaults to <repo>/workflow/image-processing/fixtures/inputs",
    )
    parser.add_argument("--repo", type=Path, default=None)
    args = parser.parse_args()

    if args.out_dir is None:
        repo = args.repo or Path(__file__).resolve().parents[3]
        out_dir = repo / "workflow" / "image-processing" / "fixtures" / "inputs"
    else:
        out_dir = args.out_dir

    manifest: dict[str, str] = {}
    for name, factory in FIXTURES.items():
        path = out_dir / f"{name}.png"
        digest = write_png(path, factory())
        manifest[name] = digest
        print(f"wrote {path} sha256={digest}")

    manifest_path = out_dir / "manifest.json"
    manifest_path.write_text(json.dumps(manifest, indent=2) + "\n", encoding="utf-8")
    print(f"wrote {manifest_path}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

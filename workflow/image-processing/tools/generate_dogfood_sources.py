#!/usr/bin/env python3
"""Synthetic dogfood sources — category proxies (not product art)."""

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


def object_hard(size: int = 128) -> Image.Image:
    img = Image.new("RGB", (size, size), (180, 200, 220))
    d = ImageDraw.Draw(img)
    d.ellipse((24, 24, size - 24, size - 24), fill=(40, 140, 90))
    d.rectangle((50, 70, 78, 100), fill=(30, 100, 70))
    return img


def illust_glow(size: int = 128) -> Image.Image:
    img = Image.new("RGB", (size, size), (8, 16, 48))
    d = ImageDraw.Draw(img)
    cx = cy = size // 2
    d.ellipse((cx - 40, cy - 40, cx + 40, cy + 40), fill=(20, 40, 100))
    d.ellipse((cx - 28, cy - 28, cx + 28, cy + 28), fill=(240, 200, 80))
    return img


def ui_overlay(size: int = 128) -> Image.Image:
    img = object_hard(size).convert("RGB")
    d = ImageDraw.Draw(img)
    d.rectangle((4, 4, size - 4, 18), fill=(255, 255, 255))
    d.line((10, size - 12, size - 10, size - 12), fill=(255, 255, 255), width=2)
    return img


def multi_subject(size: int = 128) -> Image.Image:
    img = Image.new("RGB", (size, size), (210, 210, 210))
    d = ImageDraw.Draw(img)
    d.ellipse((10, 40, 55, 95), fill=(200, 80, 80))
    d.ellipse((70, 40, 115, 95), fill=(80, 80, 200))
    return img


def soft_fringe(size: int = 128) -> Image.Image:
    """Approximate soft edge with concentric alpha-ish colors (RGB source)."""
    img = Image.new("RGB", (size, size), (240, 240, 240))
    d = ImageDraw.Draw(img)
    cx = cy = size // 2
    for r, c in ((42, (120, 120, 140)), (34, (160, 140, 180)), (26, (90, 60, 40))):
        d.ellipse((cx - r, cy - r, cx + r, cy + r), fill=c)
    return img


SOURCES = {
    "dogfood.object-hard": object_hard,
    "dogfood.illust-glow": illust_glow,
    "dogfood.ui-overlay": ui_overlay,
    "dogfood.multi-subject": multi_subject,
    "dogfood.soft-fringe": soft_fringe,
}


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--repo", type=Path, default=None)
    args = parser.parse_args()
    repo = args.repo or Path(__file__).resolve().parents[3]
    out = repo / "workflow" / "image-processing" / "fixtures" / "dogfood" / "sources"
    manifest = {}
    for name, factory in SOURCES.items():
        path = out / f"{name}.png"
        digest = write_png(path, factory())
        manifest[name] = digest
        print(f"wrote {path} sha256={digest}")
    (out / "manifest.json").write_text(json.dumps(manifest, indent=2) + "\n", encoding="utf-8")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

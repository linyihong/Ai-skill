#!/usr/bin/env python3
"""Synthesize negative fixtures (desensitized, no product art)."""

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
    img = Image.new("RGBA", (size, size), (40, 120, 200, 255))
    draw = ImageDraw.Draw(img)
    draw.ellipse((12, 12, size - 12, size - 12), fill=(220, 180, 60, 255))
    return img


def make_residue_halo(size: int = 96) -> Image.Image:
    img = Image.new("RGBA", (size, size), (0, 0, 0, 0))
    draw = ImageDraw.Draw(img)
    cx = cy = size // 2
    draw.ellipse((cx - 34, cy - 34, cx + 34, cy + 34), fill=(10, 20, 80, 255))
    draw.ellipse((cx - 26, cy - 26, cx + 26, cy + 26), fill=(230, 190, 70, 255))
    return img


def make_interior_pair(size: int = 96) -> tuple[Image.Image, Image.Image]:
    """Source keeps dark interior; bad output punches it to transparent."""
    src = Image.new("RGBA", (size, size), (0, 0, 0, 0))
    draw = ImageDraw.Draw(src)
    cx = cy = size // 2
    draw.ellipse((cx - 30, cy - 30, cx + 30, cy + 30), fill=(230, 190, 70, 255))
    # Interior "tooth/hole" dark region that must remain opaque subject
    draw.rectangle((cx - 8, cy - 4, cx + 8, cy + 10), fill=(40, 30, 20, 255))

    bad = src.copy()
    # Punch interior to transparent (interior_damage)
    ImageDraw.Draw(bad).rectangle((cx - 8, cy - 4, cx + 8, cy + 10), fill=(0, 0, 0, 0))
    return src, bad


def make_over_erode_pair(size: int = 96) -> tuple[Image.Image, Image.Image]:
    """Reference large subject; bad output heavily eroded but clean edges."""
    ref = Image.new("RGBA", (size, size), (0, 0, 0, 0))
    draw = ImageDraw.Draw(ref)
    cx = cy = size // 2
    draw.ellipse((cx - 36, cy - 36, cx + 36, cy + 36), fill=(200, 160, 80, 255))

    bad = Image.new("RGBA", (size, size), (0, 0, 0, 0))
    ImageDraw.Draw(bad).ellipse((cx - 12, cy - 12, cx + 12, cy + 12), fill=(200, 160, 80, 255))
    return ref, bad


def make_consumer_layout(size: int = 64) -> Image.Image:
    """Valid-looking cutout but wrong canvas for slot-symbol 328²."""
    img = Image.new("RGBA", (size, size), (0, 0, 0, 0))
    draw = ImageDraw.Draw(img)
    draw.ellipse((8, 8, size - 8, size - 8), fill=(210, 170, 70, 255))
    return img


def make_pos_d_canvas(size: int = 328) -> Image.Image:
    """Deterministic D positive: correct canvas + true transparency."""
    img = Image.new("RGBA", (size, size), (0, 0, 0, 0))
    draw = ImageDraw.Draw(img)
    cx = cy = size // 2
    r = 90
    draw.ellipse((cx - r, cy - r, cx + r, cy + r), fill=(210, 170, 70, 255))
    return img


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--out-dir", type=Path, default=None)
    parser.add_argument("--repo", type=Path, default=None)
    args = parser.parse_args()

    if args.out_dir is None:
        repo = args.repo or Path(__file__).resolve().parents[3]
        out_dir = repo / "workflow" / "image-processing" / "fixtures" / "inputs"
    else:
        out_dir = args.out_dir
    ref_dir = out_dir / "refs"
    ref_dir.mkdir(parents=True, exist_ok=True)

    manifest: dict[str, str] = {}

    simple = {
        "neg.alpha-opaque": make_alpha_opaque,
        "neg.residue-halo": make_residue_halo,
        "neg.consumer-layout": make_consumer_layout,
        "pos.d-canvas-328": make_pos_d_canvas,
    }
    for name, factory in simple.items():
        path = out_dir / f"{name}.png"
        manifest[name] = write_png(path, factory())
        print(f"wrote {path} sha256={manifest[name]}")

    src, bad = make_interior_pair()
    manifest["neg.interior-damage.ref"] = write_png(ref_dir / "neg.interior-damage.ref.png", src)
    manifest["neg.interior-damage"] = write_png(out_dir / "neg.interior-damage.png", bad)
    print(f"wrote interior pair sha256={manifest['neg.interior-damage']}")

    ref, eroded = make_over_erode_pair()
    manifest["neg.over-erode.ref"] = write_png(ref_dir / "neg.over-erode.ref.png", ref)
    manifest["neg.over-erode"] = write_png(out_dir / "neg.over-erode.png", eroded)
    print(f"wrote over-erode pair sha256={manifest['neg.over-erode']}")

    # Consumer expectation sidecars (neg wrong size; pos correct size)
    for stem in ("neg.consumer-layout", "pos.d-canvas-328"):
        (out_dir / f"{stem}.profile.json").write_text(
            json.dumps(
                {
                    "consumer_profile": "slot-symbol",
                    "canvas": [328, 328],
                },
                indent=2,
            )
            + "\n",
            encoding="utf-8",
        )

    manifest_path = out_dir / "manifest.json"
    manifest_path.write_text(json.dumps(manifest, indent=2) + "\n", encoding="utf-8")
    print(f"wrote {manifest_path}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

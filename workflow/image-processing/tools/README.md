# Tools — independent verification & dogfood

機械工具；**不是** Selection Actor。

| Script | Role |
| --- | --- |
| [`generate_negative_fixtures.py`](generate_negative_fixtures.py) | 合成 `neg.*`（含 refs／consumer profile） |
| [`generate_dogfood_sources.py`](generate_dogfood_sources.py) | 合成類別代理源圖 |
| [`verify_independent.py`](verify_independent.py) | 獨立驗收 → checks + finality |
| [`run_rembg_candidate.py`](run_rembg_candidate.py) | `seg.generic`←rembg + IDR（可選依賴） |
| [`run_phase2_gate.py`](run_phase2_gate.py) | 最少負向（opaque／halo） |
| [`run_phase3_gate.py`](run_phase3_gate.py) | 五類負向 + rembg dogfood |

```text
python3 workflow/image-processing/tools/run_phase3_gate.py --repo <AI_SKILL_REPO>
```

依賴：Pillow（必要）；rembg（Phase 3 dogfood）。授權見 [`../adapters/license-checklist.md`](../adapters/license-checklist.md)。

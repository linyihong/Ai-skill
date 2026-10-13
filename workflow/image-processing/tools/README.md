# Tools — independent verification（Phase 2）

機械工具；**不是** Selection Actor。

| Script | Role |
| --- | --- |
| [`generate_negative_fixtures.py`](generate_negative_fixtures.py) | 合成 `neg.*` PNG + 寫 hash sidecar |
| [`verify_independent.py`](verify_independent.py) | 讀輸出 PNG → check 列表 + finality 建議 |
| [`run_phase2_gate.py`](run_phase2_gate.py) | 生成（可選）+ 對負向 fixture 跑 verify；擋住假成功才 exit 0 |

依賴：Pillow（見 [`../adapters/license-checklist.md`](../adapters/license-checklist.md)）。不依賴 rembg。

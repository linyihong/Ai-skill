# Evidence — Product multi-locale gaps（address title／honorific gate／registry）

Status: observed  
Date: 2026-09-29  
Linked: [`14-locale-layered-failure-registry.md`](../14-locale-layered-failure-registry.md)、address-title dogfood [`04`](../04-dogfood-case-address-title-id-ID.md)

## Scope

Sanitized mechanical audit across product target langs（ja／id／ms／en／tr／ar／ko／th／vi／pt／es）. No live LLM、no host paths、no episode dump.

## Findings

1. **Address-title Candidate Space incomplete**  
   Product `address_title` registry for 小姐 covers id／en／ja／ms／es／tr／fil／pt／fr.  
   Missing ar／ko／th／vi → structured path collapses to bare surname（title lost）.

2. **English honorific leak gate was Latin-script-only**  
   `Ms./Mr./Miss/Mrs.` on ar／ko／th previously passed locale consistency while Finality／failure detector already flagged leak.  
   Product fixed: honorific check applies to **all non-en** targets；other English-leak heuristics remain Latin-script-scoped.

3. **Locale failure-registry layer still ja／id only**  
   ms／tr／ar have temporal knowledge folders but no `failure_patterns` manifestation files；ko／th／vi／pt／es absent. Matches plan 14 “add locale file + binding；do not change main chain”.

4. **Kinship Candidate seeds sparse**  
   Seeds present for ja／en／id／ms only；other langs empty → Selection lacks I23 space unless LLM invents.

## Plan impact

- Does **not** close Phase 4.  
- Confirms expanding locales must extend **knowledge／locale registry／binding**，not Selection prompt case lists.  
- Open work: evidence-backed title candidates for ko／vi／th／ar；minimal MS／TR failure manifestations when dogfood exists.

## Sanitization

No absolute paths、SSH hosts、API keys、or raw episode text.

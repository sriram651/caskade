# Decisions already made — do not relitigate

One short entry per closed design decision, with the reason. The point of this
file is that a decision settled in August is not reopened in November.

| Date | Decision | Why |
|---|---|---|
| 2026-08-04 | Bitcask KV store over finishing `p2p-chat` | Chunks cleanly; every piece leaves something concrete on disk; has an obvious finish line |
| 2026-09-09 | Finish line is the storage engine, not the server | RESP, LSM and Raft are optional expansions. "Done" must be reachable |
| 2026-09-09 | Daily 15-minute floor, 45-minute weekday ceiling | Same ~3h/week as before, spread out. Daily keeps re-entry cost near zero; short keeps daily survivable |
| 2026-09-09 | Chunk list detailed only for Phases 0–2 | Later phases get written when the real pace is known, not guessed |
| 2026-09-09 | Named `caskade`, module `github.com/sriram651/caskade` | Reads as a named project rather than an exercise. Keeps the Bitcask barrel metaphor: writes cascade down an append-only cask. Agents renamed `cask-*` to match |
| 2026-09-09 | `Version` in `cmd/caskdemo` is a `var`, not a `const` | Deliberate: a `const` cannot be stamped at build time, a `var` can, via `go build -ldflags "-X main.Version=..."`. Name is MixedCaps per Go convention, not `VERSION` |
| 2026-09-09 | The Aug 2026 plan file was folded into `PLAN.md` and deleted, not archived | It was ~95% duplicate, and its pacing rules ("~3h/week, not a floor to beat", "no streaks, the project waits") now contradict the daily floor. A superseded instruction sitting in the repo is worse than no record |
| 2026-09-11 | Phase 0 Go (chunks 0.3–0.5) lives in a throwaway `internal/warmup` package, labelled scratch, deleted when Phase 0 ends | A real type now would lock in API choices before the record format exists. Phase 0 is Go recall, not design |

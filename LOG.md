# Daily log

One row per day, newest last. Minutes are recorded because the **45-minute
weekday ceiling** matters (weekends can run longer, on request), not to be
maximised. `??` means not recorded; `--` means no work that day.

A skipped day gets a row too. This is a record, not a scoreboard: an honest
gap is more useful than a flattering one.

Floor: 15 minutes, every day. Reading counts. Writing a note counts.

Keep each summary to one line. Detail belongs in `NOTES.md` and the commits.

| Date | Mins | What | Summary |
|---|---|---|---|
| 2026-09-09 Wed | ~55 | setup + 0.1 | Scaffold, rename to caskade, `go.mod`, caskdemo prints version. Over ceiling (setup day) |
| 2026-09-10 Thu | 35 | 0.2 | Paper Qs 1 and 3 answered; 2/4/5 marked provisional. Lots of distractions |
| 2026-09-11 Fri | ~55–60 | 0.3 | `internal/warmup` Writer + first test. Over ceiling; rust was test-writing and receivers |
| 2026-09-12 Sat | -- | skipped | |
| 2026-09-13 Sun | ~55 | 0.4a | `ErrNegativeDelta` + `errors.Is` test, vs a 15–20m estimate. Late night (to 23:12); breaks skipped. 4th chunk in a row at 2–3x estimate |
| 2026-09-14 → 10-04 | -- | away | 21 days. Cause: motivation dropped on throwaway warm-up drills that never touched Bitcask |
| 2026-10-05 Mon | -- | skipped | Agreed the re-plan with cask-lead; no work because of the day job |
| 2026-10-06 Tue | ?? | restart + 1.1a | Phase 0 closed, Phase 1 halved (9ba29d9); warmup deleted (e0b14f9); record layout table (60182f1). Over ceiling |
| 2026-10-07 Wed | ?? | 1.1b | Width reasons in `docs/format.md` (c776afa); most of the time on understanding CRC |
| 2026-10-08 Thu | ?? | planning | No code. Set up a fixed weekday slot (see cadence below); 1.2a moved to Fri 9 Oct, not an extra chunk |
| 2026-10-09 Fri | -- | skipped | 1.2a had been moved here; no work happened |
| 2026-10-10 → 10-11 | -- | away | Planned away weekend (1 of 2 in the trial) |
| 2026-10-12 → 10-13 | -- | skipped | |
| 2026-10-14 Wed | ?? | 1.2a (start) | Handover from cask-lead; `HeaderSize` declared (var, then untyped const after cask-coach); doc comment drafted |
| 2026-10-15 → 10-16 | -- | skipped | |
| 2026-10-17 → 10-18 | -- | away | Planned away weekend (2 of 2 in the trial) |
| 2026-10-19 → 10-22 | -- | skipped | |
| 2026-10-23 Fri | ?? | 1.2a (finish) | Committed (c32b9e0). `record.test.go` broke build/vet/test/gofmt, renamed `record_test.go`; print-only test given a real check, red check done; doc comment fixed twice. **Exception** to reading-only Friday; reason: finish 1.2a green and committed |

## Cadence changes

Read the rows after each date against the rule that was in force.

- **From 2026-10-08 (two-week trial, review Mon 26 Oct):** caskade Mon–Thu
  13:45–14:15 with a hard stop at 14:15. Fridays are reading only. DDIA, from
  the Ch. 3 hash-index section, counts toward the floor; log those days as
  `reading`. The next two weekends are away, with no chunks planned.
  - *Note for the 26 Oct review:* no DDIA reading happened during the trial.
    The owner's phone copy of DDIA is lost and needs replacing so it can be
    read while travelling. Until it is, reading-only Fridays and
    reading-counts-toward-floor have nothing to run on.

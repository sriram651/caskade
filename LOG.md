# Daily log

One line per day, newest last. Minutes are recorded because the **45-minute
weekday ceiling** matters (weekends can run longer, on request) — not to be maximised.

A skipped day gets a line too. This is a record, not a scoreboard: an honest
gap is more useful than a flattering one.

Floor: 15 minutes, every day. Reading counts. Writing a note counts.

```
date        mins  what          detail
```

2026-09-09  ~55m  setup + 0.1   scaffold + rename to caskade + chunk 0.1 (go.mod,
                                caskdemo prints version, build/vet/fmt green).
                                Setup day, over the ceiling. Minutes are
                                wall-clock — correct if it's off.
2026-09-10  35m   0.2 + docs   Paper Qs 1 and 3 answered in own prose — exit
                                criterion met. Qs 2/4/5 also attempted, out of
                                scope, each marked provisional to redo in 0.6.
                                Docs-sync commit for 0.1. No Go touched. 35m is
                                wall-clock "with a lot of distractions" — under
                                the ceiling, attention was less than the number.
2026-09-11  ~55-60m 0.3         Over the ceiling. Time went on getting back
                                into pointer/value receivers (receiver vs
                                parameter), reading the chunk's requirement,
                                self-check questions, then both files.
                                Throwaway internal/warmup package (see
                                decisions.md). Warm-up with cask-coach; writing
                                test functions is the main rust. Wrote Writer,
                                NewWriter, pointer-receiver
                                MoveOffset, one test; two test iterations after
                                feedback (idiomatic failure checks, Name check,
                                got/want messages). Reviewer: done, build/vet/
                                test/gofmt clean. Scope: test also exercises
                                MoveOffset (3 checks vs 1 field check in the
                                handover), accepted.

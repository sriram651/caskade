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
2026-09-12  --    skipped
2026-09-13  ~55m  0.4a          Sunday, no ceiling. 67m elapsed (22:05-23:12),
                                ~10-12 of it away from the desk, so ~55m on the
                                chunk against a 15-20m estimate. A 1h pomodoro
                                with two 5m breaks was set up; neither break was
                                taken. Landed: ErrNegativeDelta, package-level,
                                built once with errors.New; MoveOffset returns
                                it on a negative delta without touching the
                                offset, nil otherwise; test matches with
                                errors.Is. MoveOffset at 100% statement
                                coverage. Reviewer passed it with four findings,
                                all explicitly non-blocking, all carried to
                                0.4b. Commit also carried NOTES.md, PLAN.md and
                                docs/decisions.md.
                                Where the time went, since it is not the obvious
                                answer: not the Go. The sentinel was right first
                                try and unprompted, including the subtle parts
                                (package-level identity, guard returning before
                                the mutation), which the reviewer called the
                                substance of the chunk. The time went to a
                                vocabulary gap ("sentinel" was a term that had
                                never been handed over), one reasoning slip (an
                                inverted invariant asserting the offset should
                                equal the rejected value), and three rounds of
                                test-assertion correction in which the move each
                                time was to delete the assertion rather than
                                re-derive it. Deleting is the cheap way to make
                                red go away without holding the invariant in
                                your head; that is a fatigue signature at 23:00,
                                not a difficulty one.
                                Pacing: fourth chunk in a row at 2-3x its
                                estimate (~55m, 35m, ~55-60m, ~55m). Tonight's
                                15-20m estimate was set after the 2x pattern had
                                already been written into PLAN.md earlier the
                                same evening, so writing it down did not change
                                it. Twice observed in PLAN.md, four times in
                                this log, zero estimates revised. The next
                                correction has to be a smaller number on a
                                chunk, not another paragraph.
                                Finished 23:12 on a Sunday; Monday's 45m ceiling
                                gets paid out of that.
2026-09-14  --    skipped       14 Sep - 4 Oct: away, 21 days, one line for
  to                            the range. Cause, as named by the owner:
2026-10-04                      motivation dropped on the throwaway Phase 0
                                warm-up drills, which never touched Bitcask.
                                That is a plan-shape signal, not a discipline
                                one, and it was answered with a re-plan.
2026-10-05  --    skipped       Came back and agreed with cask-lead to re-plan,
                                then no work because of the day job.
2026-10-06  ??m   restart+1.1a  Tue, one session. Minutes not recorded; clearly well over
                                the 45m weekday ceiling. Owner said yes to the
                                re-plan: Phase 0 closed at 0.4a, its lessons
                                folded into Phase 1, Phase 1 chunks halved
                                (9ba29d9). Step 0: owner deleted internal/warmup
                                with build/vet/test green (e0b14f9). Chunk 1.1a
                                landed (60182f1, committed and pushed by owner):
                                record layout table in docs/format.md. crc u32,
                                tstamp u64 (Unix ms), ksz u16, value_sz u32,
                                key and value variable; 18-byte fixed header,
                                big-endian. Most of the time went to catching up
                                after three weeks and guided reasoning on field
                                widths: width vs policy limit, no padding for
                                variable fields, uint32-ms overflow (~49.7
                                days), the 128 KB value cap as a Put policy on
                                top of a uint32 width, endianness. Two
                                save-related false starts; two factual fixes in
                                the doc (a swapped endianness example, and a
                                header-size note that said "all bytes").
                                Ceiling: counted as an over-ceiling day, with
                                the reason recorded (one-off re-plan + three-
                                week catch-up). Not a pre-agreed exception: no
                                reason was stated before going past 45m, so it
                                goes in the exception count as over, not as
                                approved. The one-off reason does not carry
                                forward; the next session is a normal 45m day.

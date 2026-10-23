# Session handover

Two lines, written **before** the session ends: what just landed, and what was
about to happen next. This is the first file to read on the way back in.

---

**10 Sep 2026**
Landed: chunk **0.2** — paper questions **1** and **3** answered in prose in
`docs/paper-notes.md`. Questions **2, 4 and 5** were also attempted, ahead of
reading pages 4–7; each carries a `_provisional_` marker and is to be **redone
from scratch in 0.6**, not trusted or edited. Also a docs-sync commit for 0.1
(`PLAN.md` status, a dead `NOTES.md` line, the `var`-not-`const` reasoning
moved into `docs/decisions.md`). No Go touched; build/vet/gofmt still green.
Next: chunk **0.3** — the first real Go of the project. A struct with both an
exported and an unexported field, a constructor function, and one method on a
pointer receiver. Exit criterion: **one test asserts one field after
construction**. One field, one assertion — that is the whole chunk.
Unfamiliar bits (pointer vs. value receivers, why a constructor rather than a
literal, `_test.go` naming) are `cask-coach` questions, not blockers.

**11 Sep 2026**
Landed: chunk **0.3**, the first Go of the project; `cask-reviewer` says done.
`internal/warmup/warmup.go` has `Writer` (exported `Name`, unexported `int64`
`offset`), a `NewWriter` constructor, and a pointer-receiver `MoveOffset`.
`warmup_test.go` has one test in package `warmup`. It uses the same package on
purpose, so the test can see `offset`. It makes three checks: `Name` after
construction, then `offset` after `MoveOffset(10)` and again after
`MoveOffset(27)`, expecting 37. The handover asked for one field check and no
method test, so the method test went beyond the chunk; accepted as-is. Rusty
today: *writing a test function*, not file mechanics or pointer receivers. The
test's inverted if/else got reworked into the idiomatic "check for the failure,
report it" shape during the session. Build/vet/test/gofmt green.
Next: chunk **0.4**, errors, in the same `internal/warmup/` package. A sentinel
error, a function that returns it, and a caller that checks it. Exit criterion:
**the test passes using `errors.Is`, not `==`.** It'll be handed over in detail
at the start of the next session. From 0.3's style points, which stay the
owner's call: 0.4 will be printing an error value, so `Fatalf`/`Errorf` with a
format verb is right there if wanted. The choice between `Fatal` (stop) and
`Error` (keep going) on independent checks also comes up again.

**13 Sep 2026**
Landed: chunk **0.4a**, after a two-day gap and late on a Sunday. 0.4 was split
into 0.4a/0.4b at the start of the sitting because the observed pace is running
about double the estimates. `internal/warmup/warmup.go` now has
`ErrNegativeDelta`, built once at package level with `errors.New`, and
`MoveOffset` returns it on a negative delta without touching the offset, `nil`
otherwise. The test checks it with `errors.Is`. Exit criterion met;
build/vet/test/gofmt all green. **Two assertion gaps were left open on purpose,
not by accident.** First: the two good-path calls (`warmup_test.go:17` and
`:23`) discard the returned error. Converting them to nil-checks was proposed
and **declined** — that is now a closed decision (`docs/decisions.md`,
2026-09-13), scoped to this throwaway package, with the assert-the-error
convention starting in Phase 1 instead. Do not reopen it here. Second: nothing
asserts that the offset is unchanged after a rejected call, so a `MoveOffset`
that returned the sentinel *and* corrupted the offset would still pass. That
one **is** picked up in 0.4b. Trade-off, stated plainly: green-and-committed was
chosen over more assertions at 22:45 on a Sunday. That is the right call under
the pacing rules — the cost is one known hole, written down here and in
`PLAN.md` rather than discovered by surprise later.
Next: chunk **0.4b**, same package, two steps in order. (1) Assert the offset is
**still 37** after the rejected `MoveOffset(-14)` call — one check, and it goes
first so it cannot get squeezed out if step 2 runs long. Note that an earlier
draft asserted `-14` there, an inverted invariant that only passed while the
guard was broken; 37 is the correct value. (2) Add a caller that wraps
`ErrNegativeDelta` with some context, via `fmt.Errorf` and the `%w` verb, and
assert `errors.Is` still matches through the wrap. Exit criterion: both
assertions pass. Not in scope: `errors.As`, custom error types, `errors.Join`,
table tests, and the good-path nil-checks. Worth doing once for the lesson: try
`==` against the wrapped error, watch it fail, then put `errors.Is` back.

**6 Oct 2026** (restart after a three-week gap; 5 Oct agreed the re-plan,
then skipped because of work)
Landed: three commits. `e0b14f9` deleted `internal/warmup` and closed Phase 0,
still green. `9ba29d9` is the re-plan (0.4b superseded; see `docs/decisions.md`,
2026-10-06). `60182f1` is chunk **1.1a**, the record layout table in
`docs/format.md`: `crc` u32 (set by `crc32.ChecksumIEEE`), `tstamp` u64 Unix
milliseconds, `ksz` u16, `value_sz` u32, then `key` and `value` (variable).
Fixed header 18 bytes, big-endian. Exit criterion met.
**Reasoning already worked out for 1.1b, so it isn't lost:**
- `tstamp`: a u32 of milliseconds overflows after ~49.7 days (Feb 1970). A u64
  lasts ~584 million years.
- `ksz`: a 64 KB maximum is ample. The width is a ceiling and only costs the
  field's own 2 bytes, because variable fields aren't padded. The keydir's RAM
  cost argues for a lower *policy* key limit later, not a smaller width.
- `value_sz`: u32 on disk, with a 128 KB value cap enforced in `Put` as policy
  (owner's choice), so the cap can change without changing the format.
- Byte order: the reason now in `format.md` explains why writer and reader must
  *agree*, not why *big* over little. The real reason is that big-endian reads
  naturally in a hex dump.
- Carry-forward for 1.2b: `time.Time.UnixMilli` returns a signed type, so the
  encoder needs an explicit conversion into the u64 field.
Next: chunk **1.1b**, `docs/format.md` only, no Go. One sentence justifying
each width (the four fixed fields), and sharpen the byte-order reason to say
why big rather than little. Exit criterion: every fixed-width row has a reason,
and the byte-order line says why big-endian. ~15 min. Not in scope: the
`UnixMilli` conversion (1.2b), choosing the policy key limit, enforcing the
value cap (that comes with `Put`), and file layout (2.1).

**7 Oct 2026**
Landed: chunk **1.1b** (`c776afa`). `docs/format.md` now has a "Why these
widths" list with one reason per fixed field, plus why big-endian: it reads in
natural order, so 10 shows as `00 0A`. Exit criterion met. The owner decided
not to put the **128 KB value cap** in `format.md`. It is policy, not format,
and **is enforced in `Put`** (u32 on disk, so the cap can change without a
format change). Keep this line until `Put` exists.
**CRC model, settled this session (sets up 1.2b–1.6a):** there is one CRC per
appended record, computed only over that record's bytes *after* the `crc`
field. Records are never modified after they are written, and the size fields
vary per record. Write path: build the record slice with a 4-byte zero
placeholder for `crc`, compute the CRC over the rest, fill it in, then append
once. Still carried for 1.2b: `time.Time.UnixMilli` returns a signed type, so
it needs an explicit conversion into the u64 field.
Next: chunk **1.2a**, the first Go since 13 Sep. Create `internal/record` with
the fixed header size as a named constant, and one test asserting it equals 18
(the total in `docs/format.md`). Realistic time ~20–25 min, possibly more,
because this is new-package and test-writing rust after a long gap. Take
anything half-remembered (package layout, `const`, writing a `_test.go`) to
`cask-coach` instead of pushing through it. Not in 1.2a: per-field offset
constants, any encoding, `encoding/binary`, and table tests.

**23 Oct 2026** (Friday, so this was an exception to the reading-only Friday
rule. 1.2a started Wed 14 Oct; every day from 15 to 22 Oct was skipped)
Landed: chunk **1.2a** (`c32b9e0`), the first Go since 13 Sep.
`internal/record/record.go` has an exported untyped `HeaderSize = 18` with a
doc comment. The comment is still a bit circular; a rewording was offered and
the owner kept the original, so that is settled. `record_test.go` (package
`record`) has `TestHeaderSize`, which fails via `t.Error` if the constant is
not 18. A red check was done. Build/vet/test/gofmt are green. The reviewer
reviewed an earlier state, and all three of its findings were fixed after that:
the file name, the missing assertion, and the doc comment. Moving from `Error`
to `Errorf` with got/want is the owner's call. **Rust spot, second time:** test
file mechanics. The file was first named `record.test.go`, and the first test
printed with `fmt.Println` instead of asserting. Both are the 11 Sep
`docs/go-recall.md` line again, so check the file name and the shape of the
failure path first in 1.2b. Practical blocker: the DDIA copy on the phone is
lost, so Friday reading has no material until it is replaced.
Next: chunk **1.2b**, the first `encoding/binary` in any of the owner's repos,
so this is new ground. In `internal/record/record.go`, write a function that
builds a header slice of exactly `HeaderSize` bytes from a timestamp
(`time.Time`), a key length and a value length. The `crc` bytes (0–3) stay as
the **zero placeholder**, per the CRC model settled 7 Oct; filling them in is
1.6a. `tstamp` goes at 4–11, `ksz` at 12–13 and `value_sz` at 14–17, all
big-endian. **Carried forward:** `time.Time.UnixMilli` returns a signed `int64`,
so going into the u64 field needs an explicit conversion. The length arguments
also need converting into the u16/u32 widths. Exit criterion: one test asserts
the returned length equals `HeaderSize` **and** that bytes 12–13 hold the
key length you passed in (a 5-byte key gives `00 05`). Look up
`binary.BigEndian` with `PutUint16`/`PutUint32`/`PutUint64`, `make` for a slice
of a fixed length, and `time.Time.UnixMilli`. Not in 1.2b: computing the CRC,
writing key/value bytes, decoding, range checks on lengths that overflow
u16/u32 (that is `Put`'s job), asserting the `tstamp` or `value_sz` bytes, and
table tests. Fallback if time runs out: at the hard stop, commit a green version
that writes `ksz` alone, and the rest becomes the next step.

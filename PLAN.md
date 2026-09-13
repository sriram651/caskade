# caskade — a Bitcask-style key/value store in Go

**Status:** Phase 0 in progress. Chunks 0.1–0.4a landed; the tree builds and tests pass.
**Last updated:** 13 Sep 2026

A plan plus a session handover, written so it is cheap to pick up after a gap.
Read **Where I am** and **Next chunk** first — the rest is reference.

---

## The decision

Build a **crash-safe, embedded key/value store in Go**, following the Bitcask
design. A `map[string]string` that survives `kill -9`.

Chosen over finishing `p2p-chat` because it chunks cleanly into small
independently-finishable pieces, and every piece leaves something concrete on
disk. `p2p-chat` died because it had no finish line. This one has one.

---

## Pace — the actual constraint

This matters more than the technical plan.

**Daily floor: 15 minutes, every day.** The bar is deliberately low. Reading a
paragraph of the paper counts. Writing one line in `NOTES.md` counts. Naming
tomorrow's chunk counts. The point is that the project is never more than
24 hours out of working memory.

**Weekday ceiling: 45 minutes. Stop.** The April 2026 burnout came from hours
per day, not from hard material. Difficulty escalates freely across this plan
and should. Hours do not. An extra chunk on a weekday needs a stated reason;
the default answer is no, and an accepted exception is logged with its reason.

**Weekends: one chunk by default, more on request.** Asking for another chunk
gets a pushback first: `cask-lead` spells out its difficulty, workload and
realistic time, and asks whether it can be finished green and committed today.
If the answer is still yes, it is handed over. Every chunk still ends green and
committed. Nobody but the owner ever proposes an extra chunk.

Those numbers land at roughly **3 hours a week** — the same target as
before, just spread out instead of stacked. Daily is what makes it cheap;
short is what makes daily survivable. One without the other is the trap.

Rules:

- **No calendar, no deadlines, no burn-down.** Milestones are demos, not dates.
- **Chunks sized to one 20–25 minute sitting.** Not "build the storage engine" —
  more like "write the record header encoder and a test that round-trips it."
- **Never stop mid-refactor.** Every session ends green and committed. A broken
  repo is the single biggest re-entry cost there is.
- **`NOTES.md` gets two lines before the session ends,** saying what was about
  to happen next. Highest-leverage habit for someone coding in bursts.
- **Tests are the handoff to future-me.** After three weeks away, `go test ./...`
  says instantly whether things still work.
- **A skipped day is a line in `LOG.md`, not a failure.** No streaks to protect,
  nothing to "restart". The metric being optimised is **re-entry cost**, not
  consistency and not hours.

### On the 20–30 minute office-break sessions

Good for: reading, writing a test, defining a format, writing notes.

Bad for: debugging concurrency. Too much state to page in, and an unresolved
one follows you back into actual work. Save Phase 7 for real sittings.

---

## Where I am

- **31 Jul 2026** — read ~2.5 pages of the Bitcask paper (~50 min, two sittings
  with a proper break). Normal pace; the middle of the paper is diagram-heavy
  and slower by nature, so expect the reading chunks to cover less ground there.
- **9 Sep 2026** — project scaffolded, four agents written, chunk list below.
  Chunk **0.1** landed the same day: `go.mod`, `cmd/caskdemo` printing a
  version, `go build ./...` / `go vet` / `gofmt -l` all green.
- **10 Sep 2026** — chunk **0.2** landed: paper questions 1 and 3 answered in
  `docs/paper-notes.md`. Questions 2, 4 and 5 were attempted early and are
  marked provisional; they get redone from scratch in 0.6.
- **11 Sep 2026** — chunk **0.3** landed. The first Go is in `internal/warmup/`:
  a `Writer` struct, a `NewWriter` constructor, a pointer-receiver `MoveOffset`,
  and one same-package test. The test also exercises the method, which goes
  beyond the exit criterion's one field check; accepted. The rusty part was
  writing a test function, not the struct or the receivers.
- **12 Sep 2026** — skipped. A line in `LOG.md`, nothing more.
- **13 Sep 2026** — back in after a two-day gap, late on a Sunday. Chunk **0.4**
  split in two (see below) because the observed pace is running well over the
  estimates and the hour was short. **0.4a landed:** `internal/warmup` now has
  a package-level `ErrNegativeDelta` built once with `errors.New`, `MoveOffset`
  returns it on a negative delta without touching the offset, and the test
  checks it with `errors.Is`. Exit criterion met. Two assertion gaps were left
  open on purpose at the end of the sitting — see 0.4b below and the
  2026-09-13 row in `docs/decisions.md`.

**Observed pace, three chunks in:** every chunk so far has taken roughly twice
its 20–25 minute estimate (~55m, 35m, ~55–60m). That is not a discipline
problem — the estimates are wrong. The remaining Phase 0 chunks are sized
smaller to match, and Phase 1 gets written up against the real number rather
than the hoped-for one.

Open question still being chewed on: **why does keeping every key in RAM put a
ceiling on this design?**

---

## The paper

**[Bitcask: A Log-Structured Hash Table for Fast Key/Value Data](https://riak.com/assets/bitcask-intro.pdf)**
— Sheehy & Smith, Basho, 2010. 7 pages.
Mirror: [Internet Archive](https://archive.org/details/bitcask-a-log-structured-hash-table-for-fast-key-value-data-2010)

**Only the first 3 pages matter right now** — the model, the record format, the
in-memory keydir. The rest is an API listing and a strengths/weaknesses
discussion that reads better *after* writing some code.

Five questions to answer from it, in writing, badly, in `docs/paper-notes.md`:

1. **Why append-only?** What do you get from never overwriting a byte?
2. **What is in RAM vs. on disk?** What does the RAM part cost *per key*, and
   what does that put a ceiling on?
3. **How do you delete** something in a file you cannot edit?
4. **On restart, how does the in-memory part come back?** What is the hint file for?
5. **Where does wasted space go**, and what does merging do about it?

The record layout figure — `crc`, `tstamp`, `ksz`, `valuesz`, `key`, `value` —
is the single most important thing in the paper, and the first thing to write.

---

## Scope

**The finish line:** a crash-safe embedded storage engine with real tests.
If it stops there, it is a *finished project* — demoable and README-able.

**Optional expansions**, taken only if wanted, not required for "done":

- A TCP server speaking RESP, so `redis-cli` can talk to it
- LSM tree, when the in-memory index outgrows RAM
- Raft replication — a multi-month project on its own

---

## Milestones — demos, not deadlines

| # | The demo |
|---|---|
| 1 | Write 10,000 keys. `kill -9` mid-write. Restart. Every key still there; the half-written one cleanly rejected by its checksum. |
| 2 | Same, but 50 goroutines hammering it at once, and `go test -race` is silent. |
| 3 | `redis-cli -p 6380`, then `SET name sriram` / `GET name` — and it is *my* program answering. |
| 4 | Compaction on a 2GB file while the server is still serving. File shrinks to 200MB, zero dropped requests. |

Milestone 3 is the motivating one. Milestone 4 is the genuinely hard one — a
background process mutating files that live requests are reading from.

---

## What this project is meant to close

Gaps established by reading every Go repo in `~/dev`:

1. **Binary data.** Everything written so far is text, JSON or HTTP.
   `encoding/binary` appears in **zero** files. Bitcask makes it unavoidable.
2. **Concurrency beyond one pattern.** Worker pool + channel + `WaitGroup`,
   three times over. No `RWMutex`, no `atomic`, no `errgroup`, no multi-case
   `select` outside `go-scheduler`. Nothing has yet forced protecting a
   structure that many goroutines both read *and* write.
3. **Tests.** `go-scheduler` is ~1,400 lines with none. `go-dsa/week0` has four
   test files, but on toy functions.
4. **Protocol framing and file mechanics.** `p2p-chat` has two instructive bugs:
   a `bufio.NewReader` built *inside* the read loop, and `defer c.Close()` in
   two goroutines sharing one connection. Both are ownership-and-framing bugs,
   not skill bugs.

### Strengths to build on

`go-scheduler` is production-shaped: `internal/` layout, config separated from
wiring, `context` threaded through, in-flight jobs drained on shutdown, offset
persisted so restarts do not replay, deployed on Fly.io with CI/CD. The
service-building instincts are already there.

---

## The chunk list

Detailed for the near term, **deliberately coarse further out.** Later phases
get written up when the real pace is known, not guessed at now. If chunks
consistently run past 25 minutes, they are too big — shrink the rest of the
phase rather than working longer.

Each chunk names an **exit criterion**: the concrete thing that is true when it
is done. No exit criterion means the chunk is not defined well enough yet.

### Phase 0 — Cold start (Go recall + skeleton)

The point of this phase is not the code. It is getting the muscle memory back
while producing something that builds.

| # | Chunk | Exit criterion |
|---|---|---|
| 0.1 | `go mod init`, a `cmd/caskdemo` that prints a version string | `go build ./...` is green, binary runs |
| 0.2 | Answer paper questions 1 and 3 in `docs/paper-notes.md` | Both answered in prose, however badly |
| 0.3 | A struct with exported and unexported fields, a constructor, a pointer-receiver method | One test asserts one field after construction |
| 0.4a | A sentinel error and a function that returns it directly | Test passes using `errors.Is`, not `==` |
| 0.4b | Assert the offset is unchanged after a rejected call, then wrap the sentinel with `%w` behind a bit of context | The rejected-call test asserts `offset` is still 37; `errors.Is` still passes on the wrapped error, and `==` would not |
| 0.5 | A function that copies a `[]byte` | Test proves mutating the copy leaves the original alone |
| 0.6 | Answer paper questions 2, 4 and 5 | All five now answered in `docs/paper-notes.md` |

0.4 was split into **0.4a** and **0.4b** on 13 Sep. 0.4a is the mechanics of a
sentinel error; 0.4b is the part that shows *why* `errors.Is` exists rather than
`==`, which is worth its own sitting with full attention.

**0.4b picks up one assertion gap left by 0.4a, first, before the wrapping.**
The test currently proves that a negative delta returns the sentinel, but says
nothing about what happened to `offset`. A broken `MoveOffset` that returned the
sentinel *and also* corrupted the offset would pass today. An earlier draft did
assert the offset after a rejected call, but against `-14` — an inverted
invariant that only held if the guard was missing — and it was deleted rather
than corrected. The correct assertion is that `offset` is **still 37** after the
rejected call, because a rejected call must not mutate state. That is one
check; it goes in first precisely so it cannot get squeezed out if the wrapping
runs long.

The *other* gap — the two good-path calls discarding their returned error — is
**deliberately not** in 0.4b. That is a settled decision (`docs/decisions.md`,
2026-09-13), scoped to the throwaway `internal/warmup` package; the
assert-the-error convention starts in Phase 1 and does not need retrofitting
into code that gets deleted. Note that `MoveOffset` reports 100% statement
coverage: real, but it only means both branches ran, not that the good path's
return value was ever examined. Coverage is not assertion.

0.5 looks trivial and is not — slice aliasing is the bug that will bite hardest
in Phase 1.

0.3–0.5 all live in `internal/warmup/`, a scratch package that is **deleted
when Phase 0 ends**. Nothing in Phase 1 may import it.

### Phase 1 — The record format

| # | Chunk | Exit criterion |
|---|---|---|
| 1.1 | Write `docs/format.md`: field order, widths, byte order. No Go. | The layout table exists and every width is justified |
| 1.2 | Encode the fixed-size header into a pre-sized `[]byte` | Test asserts total length and one field's bytes |
| 1.3 | Decode the header back | Round-trip test on the header alone |
| 1.4 | Encode a full record: header + key + value | Test asserts total length for a known key/value |
| 1.5 | Decode a full record | Table round-trip test: empty value, 1-byte key, large value |
| 1.6 | CRC over the right span, verified on decode | Flipping one byte returns a corruption error |

Stdlib in play: `encoding/binary`, `hash/crc32`, `time`.

### Phase 2 — The append-only file

| # | Chunk | Exit criterion |
|---|---|---|
| 2.1 | Write `docs/format.md` part 2: file naming, directory layout, open flags. No Go. | Decisions written down and reasoned |
| 2.2 | Open a datafile, track the write offset, close it | Test creates and reopens a file in `t.TempDir()` |
| 2.3 | Append one record, return its offset and size | Test appends twice and asserts the second offset |
| 2.4 | Read one record back from a given offset | Round-trip through a real file on disk |
| 2.5 | Durability: where `Sync` goes, and why | Written down in `docs/decisions.md`, wired in |
| 2.6 | Iterate every record in a file from offset 0 | Clean stop on a normal end; a truncated tail is reported, not panicked on |

Stdlib in play: `os.OpenFile` and its flags, `io.ReaderAt`, `File.Sync`,
`io.EOF` vs `io.ErrUnexpectedEOF`, `t.TempDir`.

### Phase 3 onward — coarse on purpose

- **Phase 3 — Keydir.** The in-memory index, then `Put` and `Get` on top of it.
- **Phase 4 — Recovery.** Rebuild the keydir by scanning on open; reject the
  torn tail record by checksum. → **Milestone 1.**
- **Phase 5 — Delete.** Tombstones, and what `Get` does with one.
- **Phase 6 — Rotation.** Many datafiles, one active writer.
- **Phase 7 — Concurrency.** `RWMutex` over the keydir, then `-race`.
  → **Milestone 2.** Real sittings only, not office breaks.
- **Phase 8 — Hint files.** Fast startup without rescanning every value.
- **Phase 9 — Merge.** Compaction while reads are live. → **Milestone 4.**
- **Phase 10 — RESP server.** Optional. → **Milestone 3.**

Ask `cask-lead` to write out the next phase in detail once the phase before it
lands.

---

## Working agreement

- **Claude does not write Go for this project.** It plans, explains, reviews,
  and points at what is wrong and why. Every line of Go is mine, or the
  exercise is pointless. Markdown is fine for Claude to write.
- Incremental. Understand the current stage before expanding scope.
- No calendar schedules, no streaks, no plan whose later stages need more hours
  per day than the earlier ones.

---

## Next chunk

**0.4b** — two things, in that order, in `internal/warmup/`.

1. **Close the state-invariant gap.** After the rejected `MoveOffset(-14)` call,
   assert the offset is **still 37**. One check. Two minutes. It goes first.
2. **Wrap the sentinel.** Add a caller — a second function or method that calls
   `MoveOffset` and, on failure, returns an error carrying some context *around*
   `ErrNegativeDelta` rather than replacing it. Assert with `errors.Is` that the
   wrapped error still matches the sentinel.

**Exit criterion:** the test asserts the offset is unchanged after a rejected
call, **and** `errors.Is` passes against `ErrNegativeDelta` on the error
returned by the wrapping caller.

**Look up:** `fmt.Errorf` and the `%w` verb, `errors.Is`, and — for reading, not
for using yet — `errors.Unwrap`.

**Explicitly not in 0.4b:** `errors.As`, custom error types implementing
`error`, `errors.Join` or multiple `%w` verbs in one call, table-driven tests,
nil-checks on the good-path `MoveOffset` calls (settled, see
`docs/decisions.md`), and any tidy-up of `internal/warmup` — it gets deleted
whole at the end of Phase 0.

**Worth proving to yourself once:** swap the `errors.Is` check for `==` against
the wrapped error and watch the test fail. That failure is the entire reason
`errors.Is` exists. Put it back afterwards.

Then 0.5. One per day. Ask `cask-lead` for it if the list above is not enough.

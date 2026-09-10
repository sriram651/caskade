# caskade — a Bitcask-style key/value store in Go

**Status:** Phase 0 in progress. Chunk 0.1 landed; the tree builds.
**Last updated:** 10 Sep 2026

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
and should. Hours do not.

Those two numbers land at roughly **3 hours a week** — the same target as
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
| 0.4 | A sentinel error, a function that returns it, a caller that checks it | Test passes using `errors.Is`, not `==` |
| 0.5 | A function that copies a `[]byte` | Test proves mutating the copy leaves the original alone |
| 0.6 | Answer paper questions 2, 4 and 5 | All five now answered in `docs/paper-notes.md` |

0.5 looks trivial and is not — slice aliasing is the bug that will bite hardest
in Phase 1.

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

**0.2** — a reading chunk, no Go. Answer paper questions **1** ("why
append-only?") and **3** ("how do you delete in a file you cannot edit?") in
`docs/paper-notes.md`. Prose, in your own words, allowed to be wrong.

Then 0.3. One per day. Ask `cask-lead` for it if the list above is not enough.

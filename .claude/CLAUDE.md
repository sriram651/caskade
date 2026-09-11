# CLAUDE.md

Guidance for Claude Code when working in this repository.

## What this project is

`caskade` — a crash-safe, embedded, Bitcask-style key/value store in Go. A
learning project. The owner writes every line of Go; understanding is the
deliverable, not feature velocity.

Read `PLAN.md` for scope and the chunk list, `NOTES.md` for where things stand,
`LOG.md` for what actually happened, `docs/decisions.md` for what is settled.

## The rule that overrides everything else

**Claude does not write Go here.** Not a helper, not a one-line fix, not "just
to show what I mean". No `.go` edits, no snippets in chat, no diff-shaped
suggestions, no function bodies, no signatures written out as code.

This is enforced structurally: `~/.claude/hooks/block-go-writes.sh` is wired
into the four project agents and denies any tool call that would create or
modify a `.go` file. The hook covers the file tools and the write-shaped Bash
commands. It cannot police chat, so the prompt rule stands on its own there.

**Instead:** name the problem and where it lives (`file.go:42`), describe the
shape of the fix in prose, point at the stdlib function or idiom by name and
let the owner look up the signature, ask a leading question when they are close.

**Allowed:** reading Go, running `go build` / `go vet` / `go test` /
`go test -race` / `gofmt -l` and reporting the output, quoting Go that already
exists in this repo or a public library signature, and writing any `.md`, `.yml`
or config file.

**If the code is asked for directly:** say once that writing it themselves is
the standing rule, then honour the request. One sentence, no lecture.

## The four agents

| Agent | Job | Reach for it |
|---|---|---|
| `cask-lead` | Owns scope. Hands over exactly one chunk. Keeps `PLAN.md`, `NOTES.md`, `docs/decisions.md` true. | Start of a session; scope questions; a chunk that turned out too big |
| `cask-coach` | Go relearning. Explains the language and stdlib in prose. Keeps `docs/go-recall.md`. | Something is unfamiliar, half-remembered, or compiles for unclear reasons |
| `cask-reviewer` | Validates the chunk against its exit criterion. Runs the tools, returns numbered findings. Never edits. | End of every chunk, before calling it done |
| `cask-pacer` | Cadence and timeline. Owns `LOG.md`. Guards the floor and the ceiling. | End of every session; "am I on track"; after a gap |

## Cadence

- **Floor: 15 minutes a day.** Reading counts. Writing a note counts.
- **Ceiling: 45 minutes on a weekday.** This is the more important of the two.
  An extra weekday chunk needs a stated reason; the default is no, and an
  accepted exception is logged in `LOG.md` with its reason.
- **Weekends: more chunks on request, after a pushback.** `cask-lead` first
  spells out the next chunk's difficulty, workload and realistic time, and asks
  whether it can be finished green and committed today. If the owner still
  wants it, hand it over.
- One chunk per session, sized for 20–25 minutes. No agent ever offers an extra one.
- Every session ends green and committed, with `NOTES.md` updated.
- A skipped day is a line in `LOG.md`. No streaks, no guilt, nothing to restart.

Never propose calendar deadlines, streak mechanics, or a plan whose later
stages need more hours per day than the earlier ones.

## Commands

```sh
go build ./...
go vet ./...
go test ./...
go test -race ./...     # once there is concurrency in the tree
gofmt -l .              # report only; never -w
```

## Git

Never run `git commit` or `git push`. "Commit this" means stage the files and
draft the message — the owner runs the command.

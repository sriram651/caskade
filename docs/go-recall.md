# Rust list

Things that had to be looked up, newest first. One line each: the thing, the
one-sentence answer, the date.

Maintained by `cask-coach`. If the same line comes back a second time, that is
worth saying out loud — the first explanation did not stick, and the second
one should be different.

| Date | Thing | The one-sentence version |
|---|---|---|
| 2026-10-14 | Should a constant get an explicit type (`int`, `uint32`)? | Usually no: an untyped constant takes whatever integer type the context needs (`int`, `int64`, `uint32`) and defaults to `int` when nothing says otherwise; typing it forces a conversion at every use with a different type. Also: `var` is not `const`, and a `var` is always typed. |
| 2026-09-11 | Test file mechanics | Name ends in `_test.go`; a test is a function named `Test` + capitalised word, taking a `*testing.T`; `t.Errorf` records failure and keeps going, `t.Fatalf` records it and stops that test. **Owner-flagged as the real rust spot** (pointer receivers are fine). |
| 2026-09-11 | `package x` vs `package x_test` in a test file | Same package sees unexported fields (white-box); `_test` suffix is an external package that sees only the exported API (black-box). |
| 2026-09-11 | Why a `New…` constructor, and what `internal/` blocks | Unexported fields can't be set from outside the package, so the package supplies the constructor that sets them; `internal/` makes a package importable only from code rooted at its parent directory. |

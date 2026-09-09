# caskade

A crash-safe, embedded key/value store in Go, following the
[Bitcask](https://riak.com/assets/bitcask-intro.pdf) design.

A `map[string]string` that survives `kill -9`.

The name: Bitcask stores data in a *cask* — a barrel of append-only files. Writes
cascade down it, one sequential append after another, and never go back up.

**Status:** in progress. See [`PLAN.md`](PLAN.md) for scope and milestones,
[`NOTES.md`](NOTES.md) for where it stands right now.

## Design

Append-only datafiles on disk, an in-memory hash index (the *keydir*) mapping
every key to the file, offset and size of its most recent value. Writes are one
sequential append. Reads are one hash lookup plus one seek. Deletes are
tombstone records. Space is reclaimed by merging.

## Building

```sh
go build ./...
go test ./...
```

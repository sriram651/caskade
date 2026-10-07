# On-disk format

Written before any encoding code exists, deliberately. Chunk 1.1 fills in the
record layout; chunk 2.1 fills in the file layout.

## Record layout

| Field | Width | Notes |
|---|---|---|
| crc | 4 | uint32 |
| tstamp | 8 | uint64 |
| ksz | 2 | uint16 |
| value_sz | 4 | uint32 |
| key | variable | - |
| value | variable | - |
| Header size | 18 | `crc + tstamp + ksz + value_sz` |

### Why these widths

- `crc, 4 bytes`: Go's hash library gives us crc32 function which returns uint32
- `tstamp, 8 bytes`: Unix in ms, 4 bytes will run out in ~50 days from Jan 1st 1970, but 8 bytes would last millions of years
- `ksz, 2 bytes`: can hold up to 64KB which is more than enough
- `value_sz, 4 bytes`: Values might get a little bigger, like JSON, small html pages, small images etc so 4GB is more than enough for this
- `Big-endian`: Reads and Writes straight, such as for a value 10, it will be 00 0A, being consistent across read and write

## File layout

_tbd — naming, directory structure, open flags, which file is the active one._

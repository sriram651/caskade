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

**Byte order**: _Big-endian_

**Reason**:  Encode and decode should be done in the same order, if the order changes by mistake the rest of the bytes are garbage. So if the writer used Little and wrote it as `0A 00` and the reader decoded it as `00 0A`, the intended data changes completely.

## File layout

_tbd — naming, directory structure, open flags, which file is the active one._

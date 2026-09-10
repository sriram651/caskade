# Bitcask paper — notes and answers

Answers go here in prose, however badly. Being wrong in writing is the point;
it is what makes the second read useful.

## 1. Why append-only? What do you get from never overwriting a byte?

Append only gives us the advantage of not seeking the bit to write, we are always where we want to do the write. Immutability gives the power of having the history.

## 2. What is in RAM vs. on disk? What does RAM cost per key, and what does that cap?

Data is not lost at any cost. Easier recovery. I dont think the answer to the second part of the question was in the first 3 pages

_provisional — answered before reading pages 4–7, revisit in 0.6_

## 3. How do you delete something in a file you cannot edit?

Its a write into the file with a special tombstone value

## 4. On restart, how does the in-memory part come back? What is the hint file for?

Once a file's size threshold is reached, the active file is changed and it is updated in the `keydir`. The keydir value is also there in the OS's filesystem read cache.

_provisional — answered before reading pages 4–7, revisit in 0.6_

## 5. Where does wasted space go, and what does merging do about it?

Again I dont think this is answered in the first 3 pages

_provisional — answered before reading pages 4–7, revisit in 0.6_

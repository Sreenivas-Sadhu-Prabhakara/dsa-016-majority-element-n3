# Majority Element (> n/3 times)

> Striver SDE Sheet 016 · Day 3 Arrays — Boyer–Moore voting generalized to two candidate slots, since at most two values can clear n/3. Go, stdlib only, tested.

## Problem

Return all values appearing more than ⌊n/3⌋ times, in O(n) time and O(1) extra space. Pigeonhole says there can be zero, one, or two such values — never three.

## Approach

Run Boyer–Moore with *two* slots. A matching element votes its candidate up; an element matching neither slot either claims an empty slot or votes **both** down — the generalized "discard three mutually different elements" move, which cannot change what exceeds n/3. Survivors are only candidates: a verification count decides which (if either) truly clears the bar. The count-gated matching (`cnt > 0 && v == cand`) is what keeps the arbitrary seed values from ever being mistaken for real candidates.

**Complexity:** O(n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).

# 🌸 go-bloom-axon

A from-scratch implementation of a **BitSet** and **Bloom Filter** in Go — no magic, just bits.

---

## 📦 Project Structure

```
go-bloom-axon/
├── bitset/           # Core BitSet — memory-efficient bit array
├── bitset_error/     # Typed errors
├── bloomFilter/      # Bloom Filter built on top of BitSet
└── main.go           # Interactive username availability CLI
```

---

## 🧠 How It Works

### BitSet

A BitSet stores bits in a `[]uint64` slice. Each `uint64` holds 64 bits, so a BitSet of size `n` only needs `⌈n/64⌉` words in memory.

```
Index:   0   1   2  ...  63  | 64  65  ...
         [      word 0      ] | [      word 1      ]
```

Setting bit `n`:
```
word  = n / 64   → which uint64
bit   = n % 64   → which bit inside that word
bitMap[word] |= (1 << bit)
```

### Bloom Filter

A Bloom Filter is a probabilistic data structure that answers: **"Have I seen this before?"**

- ✅ **No false negatives** — if it says "no", it's definitely new
- ⚠️ **Possible false positives** — if it says "yes", it's *probably* seen (rarely wrong)

It uses **double hashing** with MurmurHash3 to simulate `k` independent hash functions from just 2 hashes:

```
h_i(x) = (h1 + i × h2) % m     for i = 0, 1, ..., k-1
```

Where `h1, h2` come from `murmur3.Sum128()` — a single call that returns two independent 64-bit hashes.

---

## 🚀 Getting Started

```bash
git clone https://github.com/your-username/go-bloom-axon
cd go-bloom-axon
go run main.go
```

---

## 💻 CLI Demo

Run the interactive username checker:

```bash
go run main.go
```

```
  ╔══════════════════════════════════════╗
  ║       USERNAME AVAILABILITY CLI      ║
  ║       powered by Bloom Filter        ║
  ╚══════════════════════════════════════╝
  Type a username to check availability.
  Type 'exit' to quit.

  → Enter username: john_doe
  ✓  'john_doe' is available — username saved!

  → Enter username: john_doe
  ✗  'john_doe' is already taken.

  → Enter username: exit

  Goodbye! 👋
```

---

## 🔧 API

### BitSet

```go
bs := bitset.New(1000)   // create a bitset of 1000 bits

bs.Set(42)               // set bit 42
bs.Get(42)               // → true, nil
bs.Get(99)               // → false, nil
bs.Size()                // → 1000

bs.Set(9999)             // → ErrIndexOutOfBounds
```

### Bloom Filter

```go
bf := bloomfilter.New(1000)      // internally allocates 10,000 bits, 7 hash functions

bf.Add([]byte("hello"))
bf.Contains([]byte("hello"))     // → true  (definitely seen)
bf.Contains([]byte("world"))     // → false (definitely not seen)
```

---

## 🌍 Real-World Use Cases

### 1. 🔗 URL Shorteners

When generating a short code like `ax3kP`, check the Bloom filter before touching the database.

```
generate("https://example.com/very/long/url")
    │
    ├─ filter.Contains("ax3kP") == false  →  safe to use, save to DB ✓
    │
    └─ filter.Contains("ax3kP") == true   →  possible collision, verify in DB
                                               ├─ exists in DB  →  regenerate
                                               └─ not in DB     →  false positive, safe to use ✓
```

> Bitly, TinyURL — billions of URLs, the filter keeps DB lookups to a minimum.

---

### 2. 👤 Username Registration

Exactly what this CLI demonstrates. Before querying your users table:

```
register("john_doe")
    │
    ├─ filter says NO  →  username is fresh, insert into DB, add to filter ✓
    └─ filter says YES →  username is taken (or false positive → confirm with DB)
```

> At scale (millions of users), this eliminates the vast majority of DB reads.

---

### 3. 🛡️ Chrome Safe Browsing

Your browser holds a local Bloom filter of millions of known malicious URLs. Every link you visit is checked locally first — no network call needed unless the filter says "yes".

```
visit("http://suspicious-site.com")
    │
    ├─ local filter says NO  →  safe, no network call needed ✓
    └─ local filter says YES →  phone home to Google's servers to confirm
```

> This protects privacy (most URLs never leave your machine) while still catching threats.

---

### 4. 🗄️ Database Query Optimization (Cassandra / HBase)

Each SSTable on disk has a Bloom filter. Before doing an expensive disk read:

```
GET users WHERE id = "xyz"
    │
    ├─ filter says NO  →  key definitely not in this SSTable, skip disk read ✓
    └─ filter says YES →  key might be here, do the disk read
```

> Cassandra uses this to avoid reading SSTables that don't contain the requested key — massive I/O savings.

---

### 5. 🚫 Duplicate Email/Event Suppression

In event-driven systems, prevent processing the same event twice:

```go
if !filter.Contains([]byte(eventID)) {
    filter.Add([]byte(eventID))
    process(event)   // only runs once per unique event
}
```

> Used in email platforms to ensure a user never receives the same notification twice.

---

## ⚖️ Trade-offs

| Property | Value |
|---|---|
| False negatives | Never |
| False positives | Rare (tunable via size & hash count) |
| Memory | O(m) bits — extremely compact |
| Lookup time | O(k) — constant, independent of data size |
| Deletion | ❌ Not supported (use Counting Bloom Filter) |

---

## 🧪 Running Tests

```bash
# run all tests across every package
go test ./...

# with verbose output
go test -v ./...
```

---

## 📚 Dependencies

| Package | Purpose |
|---|---|
| [`github.com/twmb/murmur3`](https://github.com/twmb/murmur3) | MurmurHash3 — fast, well-distributed hashing |

---

## 📐 Double Hashing — Why It Works

Instead of computing `k` separate hashes (expensive), we compute one 128-bit MurmurHash3 and split it into two 64-bit values `h1` and `h2`. Then we derive all `k` positions:

```
position_i = (h1 + i × h2) % m
```

This technique (Kirsch-Mitzenmacher, 2008) is proven to achieve the same asymptotic false positive rate as `k` truly independent hash functions — with only **2 hash computations** regardless of `k`.

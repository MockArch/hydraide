# 💎 Treasures – The Core of Hydra's Value System

Welcome to the heart of Hydra.

You’ve learned about Swamps, naming, and behavior – but what lives *inside* a Swamp?
That’s right: **Treasures**. The smallest, most valuable unit of information in Hydra. If Swamps are vaults, Treasures are the gold inside.

And in Hydra, **information is gold**.

---

## 📦 What is a Treasure?

A Treasure is a key-value pair – or sometimes just a **key**. That’s it.
But don’t let the simplicity fool you. Treasures are:

- ⚡ Fast to access
- 🧠 Highly efficient in memory and on disk
- 🔍 Fully type-safe
- 🔄 Event-driven and reactive

You can think of them as **micro-records**, stored and retrieved with lightning speed, designed to hold only what’s truly necessary.



## 🧵 Concurrency, Isolation & Atomicity

In Hydra, each Treasure lives a completely **independent life** inside the Swamp.

- 🧩 No locking
- 🧪 No race conditions
- ⚙️ No shared-state conflicts

Just clean, isolated, **parallel-safe logic**.

### 🔍 Reading Treasures?
Absolutely. They’re **all readable in parallel**, safely and instantly – regardless of how many other reads or writes are happening.

### ✍️ Writing Treasures?
Hydra **queues each write** in the order it arrives, and applies it **atomically**, without locking the entire Swamp. No bottlenecks. No collisions.

This concurrency model:
- Enables massive scalability 🔄
- Supports async workflows ⚡
- Powers real-time systems 🌐

Hydra doesn’t just tolerate concurrency – it’s **built for it**.

---

## 🧠 Philosophy – The Most Efficient Way to Store Data

When we designed Hydra, we had one goal:

> Store massive amounts of valuable data using **as little memory and disk** as possible – without ever needing conversion layers.

Hydra doesn’t store Treasures as JSON, BSON, or text blobs. It stores them in **raw binary**, using the exact underlying data types you choose in your application.

That means:

- No serialization overhead
- No wasted space
- No schema declarations
- No intermediate formats

Just pure, efficient storage.

Imagine a struct in Go. That’s all you need. The way you write your code is the way the data gets stored.

This is why Hydra feels more like an **ORM without the ORM** – and better.

---

## 🗝️ Keys – The Minimalist Default

Every Treasure must have a **key** – a unique string inside a Swamp.

You can store a million keys in a Swamp without any values at all. This is useful for things like:

- ✅ Indexing users or IDs
- ✅ Representing flags or presence
- ✅ Creating fast lookup sets

No value? No problem.

---

## 💡 Values – If You Want More

If you need to attach information to a key, you can add a **value**.

Values in Hydra support **all primitive types**, including:

- int8, int16, int32, int64
- uint8, uint16, uint32, uint64
- float32, float64
- bool, string, []byte

Hydra stores each of these in their **native binary form**, so choosing smaller types (like `uint8` instead of `int64`) can save massive amounts of space.

---

### 🧩 Support for Full Data Structures

Hydra goes far beyond primitives.

You can store **complex, structured values** as well:

- ✅ Structs (e.g. a `UserProfile` or `TaskItem`)
- ✅ Slices (e.g. `[]string`, `[]uint32`)
- ✅ Maps (e.g. `map[string]bool`)
- ✅ Nested combinations of all the above

You don’t need to flatten your data.
You don’t need to serialize to JSON.
Just pass your struct, and Hydra stores it as-is – fully typed and binary-packed.

#### 🔍 Conceptual Example

Imagine storing a user profile with a name, age, and a list of tags. In Hydra, you can store this as a single structured value – no need to break it down or serialize it to text.

It doesn't matter if you're using Go, Python, Node.js or Rust – if your data structure looks like:

- Name: "Alex"
- Age: 29
- Tags: ["admin", "beta-user"]
- Active: true

Hydra stores this as a typed, binary-packed value. And retrieves it exactly the same way.

Think of it like saving a full object, not just a row.

Hydra stores this entire object without converting to text – and retrieves it **exactly as-is**.

This makes it ideal for cases like:

- Realtime session states
- User preferences
- AI inferences and intermediate results
- Search result cache entries

---

## 🧾 Metadata – Optional, But Powerful

Every Treasure can include optional metadata fields:

- `createdBy`, `createdAt`
- `updatedBy`, `updatedAt`
- `deletedBy`, `deletedAt`
- `expiredAt`

None of these are required. If you don’t set them, they don’t take up space – not in RAM, not on disk.

The most powerful among them? **`expiredAt`**.

This field defines the **lifetime** of a Treasure. Use cases include:

- 🕒 Caches that auto-expire
- 📋 To-do items with due dates
- 🎯 Delayed jobs or scheduling

And here’s the best part:

> Hydra can return *only* expired Treasures from a Swamp using special functions – or even remove them as part of the query.

Use this to build real-time queues, timed workflows, or cleanup systems – **without cron jobs**.

---

## 📦 Batch Operations – Because Scale Matters

Hydra was built to scale – and Treasures are no exception.

You can insert, update, or delete **multiple Treasures** across **multiple Swamps** in a single request.

Benefits:

- ✅ Fewer round-trips
- ✅ Higher throughput
- ✅ Consistent state changes

And because Hydra uses binary formats, these batch operations remain blazingly fast – even for thousands of Treasures.

---

## 🔁 Real-Time Events

Every change to a Treasure – insert, update, delete, or shift – triggers a **real-time event** to all subscribers.

That makes Hydra perfect for:

- Live dashboards
- Event sourcing
- Distributed state systems

You don’t poll.
You don’t wait.
You just react.

---

## 🧭 Final Thoughts

In Hydra, Treasures aren’t just rows in a table.
They’re **living units of knowledge**.

You can:

- Use just keys for minimal storage
- Add typed values for richer logic
- Attach metadata for lifecycle control
- Batch operations across Swamps
- React to changes instantly
- Store entire structs and slices – natively

And you do it all without ever touching a config file, a schema declaration, or a data conversion pipeline.

This is what data should feel like.
Fast. Lean. Intentional.

Welcome to the Treasure layer of Hydra.
Your data vault awaits. 🧰

---

## ➕ Built-in Increment Support

Hydra includes native increment and decrement operations for all numeric types – both integers and floating-point numbers. These operations are atomic, fast, and perfectly suited for:

- ✅ Counters (e.g. view counts, likes)
- ✅ Scoreboards and leaderboards
- ✅ Resource tracking
- ✅ Quota and usage control

Hydra supports all these operations natively:

- `IncrementInt8`, `IncrementInt16`, `IncrementInt32`, `IncrementInt64`
- `IncrementUint8`, `IncrementUint16`, `IncrementUint32`, `IncrementUint64`
- `IncrementFloat32`, `IncrementFloat64`

Each of these can also decrement by simply passing a negative delta value. No external locking, no reads before writes – just **true atomic math**.

This makes Hydra perfect for building reactive systems that evolve in real-time without needing additional logic layers.

---

### 🧮 Conditional Increment Logic

Hydra also supports **conditional increments** – giving you logical control over whether an increment (or decrement) should occur.

Using relational operators, you can define rules like:

- Only increment if the current value is **greater than** 10
- Only decrement if the current value is **less than or equal to** 100

These conditions are evaluated **atomically**, just like the write itself.

Supported comparisons include:

- `EQUAL` – current == reference
- `NOT_EQUAL` – current != reference
- `GREATER_THAN` / `GREATER_THAN_OR_EQUAL`
- `LESS_THAN` / `LESS_THAN_OR_EQUAL`

This allows you to:

- Enforce limits
- Implement dynamic threshold logic
- Control increments based on real-time states

Hydra handles this all server-side, with no need to fetch-check-write manually.

**Conditional logic is native. Scalable. And powerful.**

---

## 🔗 SDK Integration Resources (Coming Soon)

Once you understand how Treasures work, using them in your application becomes effortless.

Every Hydra SDK will let you:

- Insert or fetch Treasures with full type safety
- Use native language structs and slices as values
- Subscribe to real-time Treasure changes
- Handle expiration and reverse indexing natively

| 💻 SDK       | 🧪 Code Name | 🛠️ Status           | 📘 Treasures Docs                       |
| ------------ | ------------ | -------------------- | --------------------------------------- |
| 🟢 Go        | `hydrungo`   | ✅ Actively developed | Coming soon – Core SDK, type-rich       |
| 🟡 Node.js   | `hydrunjs`   | 🧪 In planning       | Coming soon – Async/stream ready        |
| 🐍 Python    | `hydrunpy`   | 🧠 In design         | Coming soon – Great for ML pipelines    |
| 🦀 Rust      | `hydrunrs`   | 🧠 In design         | Coming soon – Systems-level precision   |
| ☕ Java       | `hydrunjv`   | 🧠 In design         | Coming soon – Enterprise backends       |
| 🎯 C# / .NET | `hydruncs`   | 🧠 In design         | Coming soon – App servers and Unity     |
| 🧠 C++       | `hydruncpp`  | 🧠 In design         | Coming soon – Performance critical apps |
| 🌀 Kotlin    | `hydrunkt`   | 🧠 In design         | Coming soon – Android/backend devs      |
| 🍎 Swift     | `hydrunsw`   | 🧠 In design         | Coming soon – iOS/macOS native apps     |

> 💬 **Reminder:** SDKs are powerful – but they build upon the mindset you’re learning here.
> Stick to the philosophy first. Then bring it to life in code.

---

## 🧭 Navigation

← [Back to 🌿 Swamp Pattern](./swamp-pattern.md)  | [Next: 🧽 Indexing](./indexing.md)  


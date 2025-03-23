# 🔄 Subscriptions – Where Hydra Comes Alive

Welcome to the beating heart of Hydra’s **real-time superpowers**.

You’ve seen how we store data (💎 Treasures), how we name and hydrate Swamps, how we index **only when needed**. But what if we told you that Hydra doesn’t just *store* and *query* data?

> It reacts.

It *responds*.

It comes alive the moment something changes.

This is not your traditional polling-based system. This is **streaming reactivity** – **built-in**, zero-config, and instantly powerful.

So take a breath, and prepare for your next mind-shift.

---

## 🌊 The Swamp Whispers Back

In Hydra, every write, update, or deletion inside a Swamp triggers an **immediate event**. That’s right:

- 🆕 New Treasure? Event.
- 🔄 Modified Treasure? Event.
- ❌ Deleted Treasure? Event.

These aren’t delayed, batched, or cached. They are fired **as the operation completes**.

But here’s the clever twist:

> ⚠️ Events are **only fired** if there is **at least one active subscriber** listening to that Swamp.

This is Hydra’s radical commitment to **efficiency**.

- No useless events.
- No wasted CPU.
- No noise without listeners.

Hydra doesn’t shout into the void.
It speaks **only when someone is listening**.

And when it speaks – it’s real-time.

---

## 📡 Subscribing – The Hydra Way

Subscribing to a Swamp means entering a **state of awareness**.

From the moment of subscription, your system receives **every change**, in the exact order they happen.

- No polling.
- No missed updates.
- No latency.

Hydra opens a **gRPC stream** to your client, and every new event is pushed through that channel.

Imagine building:

- 🧵 Realtime chat apps,
- 📊 Live dashboards,
- 🧠 Smart pipelines,
- 🤖 Reactive microservices,

…all without external queues, brokers, or pub-sub systems.

> This is not just a database. This is a live event engine.

---

## 💬 What Gets Delivered?

When an event fires, you receive:

- The full **Treasure content**,
- The **type** of event (`created`, `updated`, `deleted`),
- And (for updates) both the **old and new** versions.

This gives you everything you need to:

- Sync local state,
- Drive animations,
- Log mutations,
- Or even rollback logic.

Hydra gives you the whole story.

---

## 📦 Empty Swamp? No Problem

Here’s where things get *really* wild:

You can subscribe to a Swamp **even if it doesn’t exist yet**.

You can subscribe to a Swamp **that is currently unloaded from memory**.

Hydra handles this elegantly:

> As soon as the Swamp is hydrated and something happens – BAM – you get the event.

No downtime. No waiting. No need to check if the Swamp is live.

Hydra treats **intent as law**.

You want to listen? Hydra listens with you.

---

## 🔄 Ephemeral Events, Eternal Precision

Hydra’s subscription model is not just efficient — it’s **invisible until needed**.

You can have:

- 1 subscriber,
- 10 subscribers,
- Or 100.000 subscribers on the same Swamp…

> 💡 And still: **no performance impact until an actual event occurs**.

Hydra uses an ultra-lightweight stream model where:

- Subscriptions consume **zero CPU while idle**,
- No processing is done until a change happens,
- And when it does, the event is broadcasted to all subscribers **with maximum efficiency**.

This means that even large-scale systems can scale horizontally without the fear of subscription overhead.

You don’t need to optimize. You don’t need to manage load. Hydra already did it for you.

Only **pure signal** when something matters.

This makes Hydra:

Subscriptions are **stream-based**, not poll-based.

- No interval timers.
- No "is anything new?" loops.
- No waste.

Only **pure signal**.

This makes Hydra:

- ⚡ Lightning fast,
- 🧘 Ultra-efficient,
- 🔗 Naturally scalable.

You don’t ask for updates.
They **find you**.

---

## 🛠️ Infra-Free Realtime

This is not Kafka.
This is not Redis Streams.
This is **not even WebSockets** (though you can easily forward it there).

This is **Hydra-native real-time**.

It works directly from your backend, your CLI, your service layer.
And it’s so lightweight, it feels like magic.

Just open a subscription. The rest just flows.

---

## 🧮 Analytics Mode – Low-Noise Listening

Sometimes you don’t need every Treasure.
You just want to know:

- 📈 How many records are there?
- 📉 Is the count going up or down?

Hydra supports **lightweight subscriptions** that deliver **summary info**:

- Count of items in a Swamp
- Changes in metadata (e.g. `createdAt`, `deletedAt`, etc.)
- TTL-related events and expiry trends

These events are **also subscription-based**, and just like with full Treasure events, they are **only emitted if at least one listener is present** – ensuring total efficiency.

Perfect for dashboards and log panels that want signal, not noise.

> With Hydra, even minimal insight is effortless.

---

## 🎯 Real-World Example – Trendizz Dashboard

The entire Trendizz search dashboard is powered by these subscriptions.

Every time a user searches, modifies filters, or saves a Dizzlet –

> That event propagates across all dashboards in real-time.

There’s no manual refresh. No polling interval. Just **instant presence**.

Even inter-service communication inside Trendizz relies on this system.

> The backend doesn’t call other services.
> It simply **subscribes** to what matters.

---

## 🧠 What About Offline Swamps?

Even if a Swamp is unloaded from memory, subscriptions **remain valid**.

> Events are triggered the moment the Swamp wakes up and data flows in.

It’s like placing a motion sensor on a closed vault.

You don’t have to keep it open.
You just wait for something to happen.

This is **intent-first architecture**.

Hydra remembers that you care – and acts accordingly.

---

## 🧘 Zero Admin. Zero Overhead.

- No broker setup.
- No queue definitions.
- No event schemas to declare.
- No infra to manage.

Just write:

> “I want to subscribe to this Swamp.”

And Hydra delivers.

Even if the Swamp is new.
Even if it’s empty.
Even if it’s closed.

---

## 🔗 SDK Integration Resources (Coming Soon)

All Hydra SDKs will support real-time subscriptions natively.
You’ll be able to:

- Open persistent listeners on any Swamp
- React to full Treasure payloads
- Handle lifecycle (`created`, `updated`, `deleted`)
- Stream directly into WebSockets or frontends

| 💻 SDK       | 🧪 Code Name | 🛠️ Status           | 📘 Subscription Docs                   |
| ------------ | ------------ | -------------------- | -------------------------------------- |
| 🟢 Go        | `hydrungo`   | ✅ Actively developed | Coming soon – Full gRPC stream model   |
| 🟡 Node.js   | `hydrunjs`   | 🧪 In planning       | Coming soon – Reactive web bindings    |
| 🐍 Python    | `hydrunpy`   | 🧠 In design         | Coming soon – Async + event loops      |
| 🦀 Rust      | `hydrunrs`   | 🧠 In design         | Coming soon – Low-latency listeners    |
| ☕ Java       | `hydrunjv`   | 🧠 In design         | Coming soon – Message bus interface    |
| 🎯 C# / .NET | `hydruncs`   | 🧠 In design         | Coming soon – Signal-style integration |
| 🧠 C++       | `hydruncpp`  | 🧠 In design         | Coming soon – System-level stream      |
| 🌀 Kotlin    | `hydrunkt`   | 🧠 In design         | Coming soon – Android-ready push       |
| 🍎 Swift     | `hydrunsw`   | 🧠 In design         | Coming soon – Native iOS observables   |

> 💬 Want to push updates to users, services, or UI in real-time? Stay tuned.
> SDKs will make it effortless – but now, you understand **the magic underneath**.

---

## 🧭 Navigation

← [Back to 🧽 Indexing](./indexing.md)  | [Next: 🔐 Locking](./locking.md)



# C4 Context — InsightBoard AI

## Purpose

Realtime product feedback analyzer:

- Ingest feedback from multiple channels
- Process asynchronously (queue)
- Store + index for search/analytics
- RAG answers and insights for users

## External Actors

- Partner apps / web clients
- Admin users
- Email/SMS providers
- Payment provider (future Stripe integration)

## Core Constraints

- Async-first ingestion
- Exactly-once _effect_ for payments and feedback processing (idempotency)
- Strong observability (logs + metrics + traces)

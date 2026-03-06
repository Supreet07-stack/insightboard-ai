# ADR-0003: REST for external APIs, events internally

Decision: Public APIs are REST + OpenAPI; internal workflows are events over RabbitMQ.
Why: Simple client integration + scalable processing.

# Implement robust retry mechanism with exponential backoff for partner webhook delivery system

**User Story:**
> As a delivery partner, I want order status webhooks to be retried automatically when my endpoint is temporarily unavailable, so that I never miss critical order updates and do not need to poll the API.

## Summary

This story is about adding a retry mechanism for partner webhooks. Today notify-service sends each webhook exactly once, and if the partner endpoint fails the event is just lost. This means partners miss status updates and have to poll the API instead which is not ideal. We want to make delivery reliable and resilient with automatic retries using exponential backoff, so that transient failures recover on their own and partners get a seamless integration experience.

## Why this matters

Reliable webhook delivery is foundational to the partner platform. Every lost event erodes partner trust and creates manual overhead on both sides, because the partner success team re-sends failed events by hand, around 120 events every day. By implementing retries with backoff we not only recover from transient failures automatically, we also reduce manual work, lower polling load on the API, and take a first step towards a more resilient event delivery platform overall.

## Background

The service currently delivers the order status events (placed, packed, out for delivery, delivered) for about 4,000 events daily. Roughly 3% of those fail on the first attempt, which is where all the manual re-sending comes in.

| Metric | Value |
|---|---|
| Daily webhook volume | ~4,000 events |
| First-attempt failure rate | ~3% |

Design notes: retries should use the standard backoff library that auth already uses. We talked about a kafka outbox at some point but that is a bigger project and out of scope for now (see TODO below). Ideally the retry schedule is something like 30s, 2m, 10m, 1h, 6h but we can tune it later once we see the numbers. 4xx errors should probably not be retried because retrying a bad request will not help, but we need to confirm that with the partners first at the next sync.

TODO: decide the final retry schedule and the DLQ retention window.

## Scope

- Add retries with exponential backoff for 5xx responses and timeouts.
- Every attempt must send the same event id so partners can deduplicate.
- Add a metric counting attempts by partner and status code.
- Handle the edge cases (4xx responses, duplicates, out-of-order events) gracefully.
- Document the behavior in the partner webhook docs.

## Acceptance criteria

- [ ] Webhook delivery is resilient and retries work as expected
- [ ] Failed events eventually end up somewhere safe (outbox/DLQ) TBD
- [ ] Partners can seamlessly integrate without getting duplicated events
- [ ] The solution should scale with the event volume
- [ ] The documentation is comprehensive

## Out of scope

Kafka migration, signed webhooks (tracked separately).
# Retry failed partner webhooks with backoff

## Problem

`notify-service` sends each order status webhook once. If the partner endpoint fails, the event is lost. About 3% of 4,000 daily events fail on the first attempt. The partner success team re-sends them by hand.

## Outcome

A failed webhook retries on a backoff schedule for up to 6 hours. Events that still fail go to a dead-letter queue.

## Acceptance criteria

- [ ] A webhook that gets a 5xx response or a timeout retries after 30 seconds, 2 minutes, 10 minutes, 1 hour, and 6 hours.
- [ ] A webhook that gets a 4xx response does not retry.
- [ ] Every attempt sends the same `event_id`, so partners can drop duplicates.
- [ ] After the last attempt fails, the event goes to the dead-letter queue.
- [ ] The metric `webhook_retry_total` counts attempts by partner and status code.
- [ ] The partner webhook docs describe the retry schedule and the `event_id` field.

## Risks

- Partners that do not check `event_id` process a retried event twice. Announce the change to partners one week before rollout.
- A retried event can arrive after a newer one. Partners must order events by `occurred_at`, not by arrival.

## Rollout

- Enable retries for one pilot partner, then the remaining accounts over one week.
- The `webhook_retry_total` dashboard goes live before the first pilot.

## Links

- [Partner webhook docs](https://docs.example.com/partners/webhooks)
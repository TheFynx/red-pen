# Invoice PDF export times out for invoices over 900 lines

## Steps to reproduce

1. Log in as `greenline-catering` in production.
2. Go to **Billing** > **Invoices**.
3. Click **Download PDF** on the October invoice, which has 1,140 line items.

## Expected

The PDF downloads, whatever the invoice size.

## Actual

The gateway returns HTTP 504 after 30 seconds. The customer sees "Something went wrong."

61 of 318 exports failed during the October close, 2026-10-01 to 2026-10-07. The six wholesale accounts take most of the damage:

| Account | Failed | Attempts | Failure rate |
|---|---|---|---|
| `harvest-table-co` | 19 | 22 | 86% |
| `greenline-catering` | 22 | 62 | 35% |
| `oak-and-ember` | 9 | 60 | 15% |
| `brightside-cafes` | 5 | 38 | 13% |
| `north-pier-kitchens` | 4 | 64 | 6% |
| `lumen-bakery` | 2 | 72 | 3% |

Every failed invoice has over 900 line items. No invoice under 900 lines failed in this window.

The renderer logs only `render aborted`, with no line-item count:

```
ERROR billing-docs render aborted invoice=inv_7Q2K request_id=9f1c
  at renderInvoice   (src/render/pdf.ts:211)
```

## Environment

- Service `billing-docs`, version `2026.10.2`, commit `b41e09c`, region `eu-west-1`.
- The gateway enforces a 30 second request timeout.
- The version did not change during the failure window, so this is not a deploy regression.

## Cause

Render time grows with line-item count, so invoices over 900 lines pass the 30 second gateway window. The render runs in the request path, and month-end load stretches it further, which is why retries sometimes pass.

## Workaround

A retry sometimes succeeds under lower load. There is no reliable workaround.

## Acceptance criteria

- [ ] A render that aborts logs the invoice ID, account, and line-item count.
- [ ] An invoice with 2,000 line items downloads in under 30 seconds.
- [ ] A test renders a 2,000-line invoice and asserts it finishes under the timeout.
- [ ] Production logs show no `render aborted` entries for 7 days after the fix.

## Links

- The export attempt query and the full account breakdown are in the first comment.
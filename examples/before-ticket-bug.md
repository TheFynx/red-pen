# Issue with the invoice PDF export not working correctly for big invoices during month end

This ticket is about a problem we have been seeing with the PDF export feature in billing-docs. Basically the PDF export is not working for big invoices and it is blocking customers from getting their invoices which is really bad during month end close when everyone needs their invoices. We need to fix this ASAP.

## Background

We first noticed this last week when Marcus from support reported that a customer said the download button just spins and then shows "Something went wrong". At first we thought it was the gateway again, like the incident in September, but this seems different. I spent some time on it during Tuesday's hack day and here is what I found so far.

| Export attempts | Failed | Failure rate |
|---|---|---|
| 318 | 61 | ~19% |

So around 19% of the 318 export attempts failed between Oct 1 and Oct 7, that is 61 failed exports out of 318 total attempts which is quite a lot. It is possible that there are more failures outside this window, I did not check the whole month yet. It is mostly the big wholesale accounts, for example harvest-table-co had 19 failures out of 22 attempts which is an 86% failure rate, and greenline-catering had 22 out of 62. The pattern seems to be that every failed invoice is a really big one with over 900 line items, I could not find any small invoice in the failed list. The version 2026.10.2 did not change during the failure window so it is probably not a regression from a deploy.

## How to reproduce

1. Log in as one of the wholesale accounts (greenline-catering is a good one)
2. Go to Billing → Invoices and find their October invoice (it has like 1,140 line items)
3. Click Download PDF
4. Wait ~30 seconds and it will fail with a 504 Gateway Timeout

Note that this is intermittent, sometimes it works on the second try. If it does not fail on the first attempt just try a few more times, it seems to depend on how busy things are at that moment.

## What should happen vs what happens

**Expected:** the customer should be able to download their invoice PDF properly and without any issues, no matter how big the invoice is. That is the whole point of the feature after all, so it needs to work reliably for everyone including the wholesale accounts with the big invoices.

**Actual:** the gateway times out after 30 seconds and the customer just gets a generic "Something went wrong" error message which is really frustrating for them and it also generates extra support load for us on top of everything else.

```
ERROR billing-docs render aborted invoice=inv_7Q2K request_id=9f1c
  at renderInvoice   (src/render/pdf.ts:211)
  at exportHandler   (src/routes/invoices.ts:88)
```

The logs are not very helpful, they only say render aborted. It would be nice if they also logged the invoice id and the line item count and stuff so we can debug this kind of thing faster next time. Possibly related to the renderer bump in 2026.10.2, although the version did not actually change during the failure window so maybe not.

## Environment

Production, eu-west-1, billing-docs version 2026.10.2, commit b41e09c. Gateway timeout is 30 seconds.

## Root cause and proposal

I think the root cause is probably that render time scales with the number of line items, maybe even worse than linear, so the really big invoices take longer than the 30 second gateway timeout. Also the rendering happens synchronously in the request path instead of as a background job which is not great. Each line item might also trigger its own layout calculation but I am not 100% sure about that part yet.

Short term we should improve the logging as described above and optimize the render path so the big invoices render faster. Medium to long term we should consider moving PDF generation to a background job and possibly look into pagination or streaming for the really big ones. We could also raise the gateway timeout but that would just hide the real problem so I don't think we should do that, though it is an option.

## Acceptance criteria

- [ ] PDFs download properly no matter the size
- [ ] The export works as expected for all the wholesale accounts
- [ ] Better logging for render aborts gets added
- [ ] Maybe look into making the rendering faster and/or async
- [ ] No more render aborted errors in prod (TBD how we measure this)

## Risks / considerations

There is a risk of performance regressions for the smaller invoices, we would need to keep an eye on the render time metrics after the change. Also customer comms, some customers might need to be notified once the fix is deployed. Timeout increase is not a fix as discussed above.

## Notes
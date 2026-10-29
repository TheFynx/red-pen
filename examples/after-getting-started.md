# Getting started

Connect a repository to Reeve, open a PR, and apply your first change.

For a local trial without cloud resources, start with the [local demos](../examples/README.md#local-demos).

## Prerequisites

- A GitHub repository with Actions enabled and a working Pulumi, Terraform, or OpenTofu project.
- The state backend and workload credentials that project already uses.
- A persistent bucket for Reeve locks, plans, and run history.
- A second reviewer for approvals.

## 1. Install the CLI

Download a binary from [Releases](https://github.com/reeveops/reeve/releases) and verify it against the signed `checksums.txt`.

Or build from source:

```bash
git clone https://github.com/reeveops/reeve
cd reeve
git checkout d31c814640689c2f2e1b0d02d2bc11a80a94faab
mise trust
mise run build
export PATH="$PWD/bin:$PATH"
```

Requires [mise](https://mise.jdx.dev/) and the Go toolchain in this repository. See [versions and binaries](github-actions.md#versions-and-binaries) to keep CLI and workflow versions aligned.

## 2. Scaffold configuration

From your infrastructure repository:

```bash
reeve init --workflow-ref d31c814640689c2f2e1b0d02d2bc11a80a94faab
```

The wizard asks for your engine and approval policy. Inspect the generated `.reeve/` directory and workflow before committing.

## 3. Connect your bucket

Replace the scaffold filesystem bucket in `.reeve/shared.yaml`:

```yaml
bucket:
  type: gcs
  name: YOUR_REEVE_BUCKET
```

See [bucket provisioning](self-hosting.md#bucket-provisioning) for S3, GCS, Azure Blob, or R2. See [authentication](auth.md#which-credentials-go-where) for credential wiring.

## 4. Validate

```bash
reeve lint
reeve stacks
reeve rules explain api/prod
```

Replace `api/prod` with a reference from `reeve stacks`. Confirm your stacks and approver list look right.

Compare the generated `.github/workflows/reeve.yml` with the [canonical caller](github-actions.md#gitops-caller). Add your bucket authentication inputs and permissions.

## 5. Open a PR

Commit the configuration and workflow to your default branch. Then open a new PR that changes one infrastructure value.

Expect a Reeve comment listing affected stacks and changes:

![A Reeve preview for two stacks](images/preview.jpg)

No comment? See [troubleshooting](operations.md#troubleshooting).

## 6. Apply

Have a reviewer approve the PR, then comment:

```text
/reeve apply
```

Read the stack results in the Reeve comment to confirm the apply succeeded.

## Next steps

- [PR workflow](pull-requests.md): commands, approvals, merge-triggered apply.
- [Maintenance](operations.md#scheduled-maintenance): lock reaping and artifact retention.
- [Notifications](notifications.md), [drift](drift.md), and [policies](policy-hooks.md): optional capabilities.
- [Examples](../examples/README.md): configurations and the E2E test harness.

## Common issues

| Problem | Cause | Fix |
|---------|-------|-----|
| No PR comment | Workflow not on the default branch | Merge the workflow first, then open the infrastructure PR |
| Blocked apply | Approval or preview gate not met | Read the block reason in the Reeve comment |
| Fresh plan at apply time | Plan artifact missing or expired | See [saved plans](pull-requests.md#saved-plans-and-freshness) |
| Multiple engines in one repo | One `.reeve/` per engine | Use [separate roots](github-actions.md#multiple-roots) |
| OpenTofu detected as Terraform | `--non-interactive` defaults to HCL | Set `engine.type` to `tofu` in config |

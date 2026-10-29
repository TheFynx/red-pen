# Hi, I'm Levi 👋

**Platform Engineer.** I build infrastructure tooling that teams run themselves.

`Go` · `Terraform` · `Pulumi` · `OpenTofu` · `TypeScript` · `AWS` / `GCP` / `Azure`

---

## 🔨 What I'm building

### [reeve](https://github.com/reeveops/reeve): PR-native GitOps orchestrator

A single Go binary you drop into your CI. It previews IaC changes on every PR,
gates `/reeve apply` behind approvals, CODEOWNERS, required checks, per-stack
locks and freeze windows, and detects drift on a schedule.

Works with **Pulumi, Terraform and OpenTofu**. State, locks and audit records go
to *your* bucket (S3 / GCS / Azure Blob / R2 / local). Credentials are
short-lived and federated: AWS OIDC, GCP WIF, Azure federated, GitHub App.

> No control plane. No vendor backend. No telemetry. No account. MIT, permanently.

The design rule everything follows: **gates fail closed.** If a check cannot be
evaluated, the apply is blocked, never waved through.

### [rwr](https://github.com/fynxlabs/rwr): Rinse, Wash, Repeat

Configuration management for distro hoppers. Declare your workstation once
(packages, dotfiles, templates, services, users, fonts, repos) and rebuild it
anywhere, because reinstalling shouldn't mean rediscovering how you had
everything set up.

Blueprints in **YAML, JSON, TOML, or CUE**, with per-item profiles
(`desktop`, `laptop`, `work`) and Go templating. Point it at a local tree, a
GitHub repo URL, or an `owner/repo` shorthand and run. One binary, one
manifest, many machines: Arch and friends, Debian/Ubuntu, macOS, Windows,
across pacman, AUR helpers, brew, flatpak, cargo, and more.

### [ontap](https://github.com/fynxlabs/ontap): OpenAPI spec -> CLI

Point it at an OpenAPI spec, get a working command-line client. No codegen step
to maintain, no hand-written wrapper drifting from the API it wraps.

---

## 🧰 Toolbox

| | |
|---|---|
| **Languages** | Go, TypeScript, Shell, Ruby |
| **IaC** | Terraform, OpenTofu, Pulumi |
| **Cloud** | AWS, GCP, Azure |
| **CI/CD** | GitHub Actions, GitOps workflows |

## 🧭 How I build

- **Fail closed.** Ambiguity blocks; it never silently permits.
- **You own your data.** Self-hosted by default, no phone-home.
- **Explain yourself.** When a rule fires, the tool says exactly why.

---

📫 Find me on [GitHub](https://github.com/TheFynx) · Org: [Freehold Digital](https://github.com/freehold-digital)

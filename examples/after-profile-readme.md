# Hi, I'm Levi 👋

Platform engineer. I build infrastructure tooling that teams run themselves.

`Go` · `Terraform` · `Pulumi` · `OpenTofu` · `TypeScript` · `AWS` / `GCP` / `Azure`

---

## 🔨 What I'm building

### [reeve](https://github.com/reeveops/reeve)

PR-native GitOps orchestrator. One Go binary in your CI. Previews IaC changes on every PR, gates `/reeve apply` behind approvals, CODEOWNERS, required checks, per-stack locks, and freeze windows. Detects drift on a schedule.

Supports Pulumi, Terraform, and OpenTofu. State, locks, and audit records go to your bucket (S3, GCS, Azure Blob, R2, or local filesystem). Credentials are short-lived and federated.

> No control plane. No vendor backend. No telemetry. No account. MIT.

Gates fail closed. If a check cannot be evaluated, the apply is blocked.

### [rwr](https://github.com/fynxlabs/rwr)

Rinse, Wash, Repeat. Workstation configuration management for distro hoppers. Declare packages, dotfiles, templates, services, users, fonts, and repos once. Rebuild anywhere.

Blueprints in YAML, JSON, TOML, or CUE with per-item profiles (`desktop`, `laptop`, `work`) and Go templating. Point it at a local tree, a GitHub repo URL, or `owner/repo` shorthand. One binary, one manifest, many machines.

Supports Arch, Debian/Ubuntu, macOS, and Windows across pacman, AUR helpers, brew, flatpak, cargo, and more.

### [ontap](https://github.com/fynxlabs/ontap)

OpenAPI spec to CLI. Point it at a spec, get a working command-line client. No codegen to maintain. No hand-written wrapper drifting from the API.

---

## 🧰 Toolbox

| | |
|---|---|
| Languages | Go, TypeScript, Shell, Ruby |
| IaC | Terraform, OpenTofu, Pulumi |
| Cloud | AWS, GCP, Azure |
| CI/CD | GitHub Actions, GitOps workflows |

## 🧭 How I build

- **Fail closed.** Ambiguity blocks, never silently permits.
- **You own your data.** Self-hosted by default. No phone-home.
- **Explain yourself.** When a rule fires, the tool says exactly why.

---

📫 [GitHub](https://github.com/TheFynx) · [Freehold Digital](https://github.com/freehold-digital)

# Bootstrap

Bootstrap runs once before all other blueprints. It installs packages, creates directories and files, sets up SSH keys, clones Git repos, manages services, and creates users and groups.

Define bootstrap in `bootstrap.yaml` (or `.json`, `.toml`) at the root of your blueprint location.

## When bootstrap runs

By default, bootstrap runs once. RWR writes a marker file to the configuration directory on success. Later runs skip bootstrap if the marker exists and covers the active profiles.

The marker records which profiles the run covered. A later run re-runs bootstrap when:

- The marker is missing.
- The run names a profile the marker does not cover.
- The `--force-bootstrap` flag is set.

Profile coverage is cumulative. Running with `--profile desktop` and later with `--profile laptop` does not re-run bootstrap on the second run. Bootstrap re-runs only if `laptop` has entries the first run did not apply.

Running with `--profile all` records full coverage. A `--dry-run` does not write the marker.

Force bootstrap on every run:

```bash
rwr all --force-bootstrap
```

Re-run bootstrap alone (after editing `bootstrap.yaml`):

```bash
rwr bootstrap
```

`rwr bootstrap` ignores the run-once marker. The marker exists for `rwr all` idempotency, not to block a direct request.

## File structure

```yaml
scripts:         # Preparation scripts. Run first, before package managers or credentials.
packageManagers: # Package managers needed by the bootstrap packages.
packages:        # Packages to install.
directories:     # Directories to create.
files:           # Files to create or copy.
ssh_keys:        # SSH keys to generate.
git:             # Git repositories to clone.
services:        # Services to enable, start, stop, or disable.
users:           # User accounts to create.
groups:          # Groups to create.
```

The bootstrap file supports `schema_version` at the top level to override the tree-wide version from the init file.

## Execution order

1. `scripts`
2. `packageManagers`
3. `packages`
4. `directories`
5. `files`
6. `ssh_keys`
7. `git`
8. `services`
9. `users`
10. `groups`

## Profile filtering

Bootstrap entries support the same profile filtering as other blueprints. Entries gated behind an inactive profile are skipped. The marker records which profiles ran. A later run naming a skipped profile triggers a re-run for those entries.

Profile validation happens before bootstrap or package-manager installation.

## Preparation scripts

Preparation scripts run before package managers and credential resolution. They can use OS tools only. They cannot depend on credentials that have not been resolved yet.

If bootstrap needs a declared credential, give it `scope: [bootstrap]`. RWR resolves those credentials before bootstrap begins.

Successful scripts are recorded individually under the `bootstrap-steps` configuration directory. On retry (including forced bootstrap), unchanged successful scripts are skipped. A changed script declaration or source content triggers a re-run. Scripts must converge safely after a failed or interrupted attempt.

## Failure handling

Bootstrap continues past individual step failures. A failed package does not stop the twenty after it. A temporarily unreachable Git repo does not abandon the run.

The marker is written only when every step succeeds. A bootstrap with failures stays unmarked. The next run retries it. Completed steps are idempotent and skip fast.

## Example

```yaml
packages:
  - name: git
    action: install
  - name: curl
    action: install
  - names:
      - vim
      - tmux
    action: install
    package_manager: apt
    args: ["--no-install-recommends"]

files:
  - name: config.ini
    action: create
    content: |
      [settings]
      debug = true
  - name: remote-config.txt
    action: copy
    source: https://example.com/remote-config.txt
    target: /etc/app/config.txt

directories:
  - name: data
    action: create
    mode: 0755

git:
  - name: my-repo
    action: clone
    url: https://github.com/example/my-repo.git
    path: /opt/my-repo

services:
  - name: nginx
    action: enable

users:
  - name: johndoe
    action: create
    password: "$6$mysecretpassword"
    groups:
      - sudo
      - docker

groups:
  - name: developers
    action: create

ssh_keys:
  - name: id_rsa
    type: rsa
    path: ~/.ssh
    comment: johndoe@example.com
    no_passphrase: true
    copy_to_github: true
    set_as_rwr_ssh_key: true
```

### packages

The `names` field installs multiple packages with the same configuration. The `args` field passes extra arguments to the package manager.

An empty package list requires no package manager.

### files

URL sources are supported. RWR downloads the file before processing it.

If the target path does not end with `/`, RWR treats the operation as a rename.

### ssh_keys

`set_as_rwr_ssh_key` sets the generated key as the default RWR SSH key.

`copy_to_github` uploads the public key to GitHub. The optional `github_title` field names the key on GitHub. Without it, the machine hostname is used.

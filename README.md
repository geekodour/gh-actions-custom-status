# gh-actions-custom-status
Create custom GitHub status from Github Actions

### Usage with GitHub Actions
```
- name: Update status to pending
  uses: geekodour/gh-actions-custom-status@v0.0.4
  env:
    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
    LAST_COMMIT_SHA: ${{ github.event.client_payload.LAST_COMMIT_SHA }}
  with:
    args: >-
      --state=pending
      --context=foo-status-update
      --target_url=https://github.com/${{ github.repository }}/actions
```
This workflow is run on a `repository_dispatch` event.

### Usage
```
Flags:
  -h, --help                     Show context-sensitive help (also try --help-long and --help-man).
      --state=STATE              The state of the status. Can be one of error, failure, pending, or success.
      --target_url=TARGET_URL    The target URL to associate with this status.
      --context=CONTEXT          A string label to differentiate this status from the status of other systems.
      --description=DESCRIPTION  A short description of the status.
```

### Building locally
```
make build
```

### Building the Docker image
```
make docker-build
```

### Environment Variables
- `GITHUB_TOKEN`: Personal Access token
- `GITHUB_REPOSITORY`: Set in the format `org/repo`
- `LAST_COMMIT_SHA`: SHA of the last commit in the PR.

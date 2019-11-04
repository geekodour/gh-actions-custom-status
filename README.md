# gh-actions-custom-status
Create custom GitHub status from Github Actions

### Usage with GitHub Actions
```
- name: Set the custom status
  uses: geekodour/gh-actions-custom-status@v0.0.4
```

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
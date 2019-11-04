# gh-actions-custom-status
Create custom GitHub status from Github Actions

### Usage with GitHub Actions

### Building locally
```
go build -o ghcs .
```

### Building the Docker image

### Environment Variables
- `GITHUB_TOKEN`: Personal Access token
- `GITHUB_REPOSITORY`: Set in the format `org/repo`
- `GITHUB_REF`: Ref can be a SHA, a branch name, or a tag name.
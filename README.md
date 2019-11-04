# gh-actions-custom-status
Create custom GitHub status from Github Actions

### Building locally
```
go build -o ghcs .
```

### Environment Variables
- `GITHUB_TOKEN`: Personal Access token
- `GITHUB_REPOSITORY`: Format: `org/repo`
- `GITHUB_REF`: Ref can be a SHA, a branch name, or a tag name.

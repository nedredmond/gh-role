This simple extension allows you to check the current user's permissions on a repo.  It's especially useful for scripts in shared repositories.

For example, if you have a script that adds a protected tag to a commit then pushes it to the repo, you can use this command to ensure that the user has the necessary permissions to do so.

# Installation

```bash
gh extension install nedredmond/gh-role
```

# Usage

In a repo you own:

```bash
gh role
// ADMIN
// Exits with error code 0
```

For a repo you have nothing to do with:

```bash
gh role -r canada-ca/ore-ero
// READ
// Exits with error code 0
```

To check for specific permissions:

```bash
gh role -r canada-ca/ore-ero write admin
// 2022/11/20 01:01:00 User does not have permissions on canada-ca/ore-ero: write, admin; found READ
// Exits with error code 1
```

## Flags

- `-r` The repo to check permissions on.  Defaults to the current repo.
- `-f` Prints a friendly message instead of the permission constant.
  - i.e., `Current user has ADMIN permission on gh-role.`
- After any flags, list the permissions you want to check, separated by spaces. If none are provided, the command will print your current permission instead.

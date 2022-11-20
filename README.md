This simple extension allows you to check the current user's role on a repo.  It's especially useful for scripts in shared repositories.

For example, if you have a script that adds a protected tag to a commit then pushes it to the repo, you can use this command to ensure that the user has the necessary permissions to do so.

# Installation

```bash
gh extension install nedredmond/gh-role
```

# Usage

Without listing role names, `gh role` simple returns the current user's role. This returns the role name and exits with code 0.

 - In a repository that the user owns:

    ```bash
    gh role
    // ADMIN
    // Exits with code 0
    ```

 - For a repo that the user has nothing to do with:

    ```bash
    gh role -r canada-ca/ore-ero
    // READ
    // Exits with code 0
    ```

If roles are listed (after all flags), the command will exit with exit code 1 if there is no match. This is not case-sensitive.

 - To check for specific roles:

    ```bash
    gh role -r canada-ca/ore-ero write admin
    // 2022/11/20 01:01:00 User does not have permissions on canada-ca/ore-ero: write, admin; found READ
    // Exits with error code 1
    ```
 
 - Available roles in order of increading permissions are: `READ`, `TRIAGE`, `WRITE`, `MAINTAIN`, `ADMIN`. For more information, see [GitHub's documentation on repository roles](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/repository-roles-for-an-organization).

## Flags
- `-h` Prints available flags and usage to the CLI.
- `-r` The repo to check permissions on.  Defaults to the current repo.
- `-f` Prints a friendly message instead of the permission constant.
  - i.e., `Current user has ADMIN permission on gh-role.`
- After any flags, list the permissions you want to check, separated by spaces. If none are provided, the command will print your current permission instead.

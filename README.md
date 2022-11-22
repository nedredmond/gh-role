# `gh-role`

This simple extension allows you to check the current user's role on a repo or org.  It's especially useful for scripts in shared repositories.

For example, if you have a script that adds a protected tag to a commit then pushes it to the repo, you can use this command to ensure that the user has the necessary permissions to do so.

__Notes:__

- This is an extension for the [GitHub CLI](https://cli.github.com/) and requires that you have it installed.

- This extension is not meant to be used as a security measure.  It is meant to be used as a convenience for scripts that need to check the user's role.

## Installation

```bash
gh extension install nedredmond/gh-role
```

## Usage

Without listing role names, `gh role` simply returns the current user's role. This returns the role name and exits with code 0.

- In a repository that the user owns:

    ```bash
    gh role
    // admin
    // Exits with code 0

    gh role -r nedredmond/gh-role -f
    // User has admin role on nedredmond/gh-role.
    // Exits with code 0

    gh role -r nedredmond/gh-role -f maintain friend lover
    // 2022/11/20 01:01:00 user does not have roles in nedredmond/gh-role: maintain, friend, lover; found admin
    // Exits with code 0
    ```

- For a repo that the user has nothing to do with:

    ```bash
    gh role -r canada-ca/ore-ero
    // read
    // Exits with code 0
    ```

If roles are listed (after all flags), the command will exit with exit code 1 if there is no match. This is not case-sensitive.

- To check for specific roles:

    ```bash
    gh role -r canada-ca/ore-ero write admin
    // 2022/11/20 01:01:00 user does not have roles in canada-ca/ore-ero: write, admin; found read
    // Exits with error code 1
    ```

For organizations, the command will exit with exit code 1 if the user is not a member of the organization.

- To check for membership:

    ```bash
    gh role -o canada-ca
    // 2022/11/20 01:01:00 user has no role in canada-ca
    // Exits with error code 1
    ```

- To check for specific roles:

    ```bash
    gh role -o my-org admin
    // 2022/11/20 01:01:00 user does not have role in my-org: admin; found member
    // Exits with error code 1

    gh role -o my-org admin member
    // member
    // Exits with code 0
    ```

## Roles

### Repository

In order of increading permissions: `READ`, `TRIAGE`, `WRITE`, `MAINTAIN`, `ADMIN`. For more information, see [GitHub's documentation on repository roles](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/repository-roles-for-an-organization).

### Organization

Note that role names returned do not match [the documentation on organization roles](https://docs.github.com/en/organizations/managing-peoples-access-to-your-organization-with-roles/roles-in-an-organization#about-organization-roles). "Owners" have the `ADMIN` role, and all others, regardless of individual permissions, have the `MEMBER` role. Does not include roles for Enterprise organizations, such as "Billing Manager".

## Flags

- __`-h`__ Prints available flags and usage to the CLI.
- __`-r`__ The repo to check roles on.  Defaults to the current repo.
- __`-o`__ The org to check roles on. If blank, defaults to repo check.
- __`-f`__ Prints a friendly message instead of the machine-readable role.
  - i.e., `User has admin role in nedredmond/gh-role.`
- After any flags, list the roles you want to verify, separated by spaces. If none are provided, the command will print your current role instead.
  - There is no validation for provided roles and they are not case-sensitive. The only requirement is that the role is spelled correctly.

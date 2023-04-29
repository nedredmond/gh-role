# `gh-role`

This simple extension allows you to check a user's role on a repo or org.  It's especially useful for scripts in shared repositories.

For example, if you have a script that adds a protected tag to a commit then pushes it to the repo, you can use this command to ensure that the user has the necessary permissions to do so before proceeding.

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
    // nedredmond has admin role in github.com/nedredmond/gh-role.
    // Exits with code 0

    gh role -r nedredmond/gh-role -f maintain friend lover
    // 2022/11/20 01:01:00 nedredmond does not have roles in github.com/nedredmond/gh-role: maintain, friend, lover; found admin
    // Exits with code 0
    ```

- For a repo that the user has nothing to do with:

    ```bash
    gh role -r cli/cli
    // read
    // Exits with code 0
    ```

If roles are listed (after all flags), the command will exit with exit code 1 if there is no match. This is not case-sensitive.

- To check for specific roles:

    ```bash
    gh role -r canada-ca/ore-ero write admin
    // 2022/11/20 01:01:00 nedredmond does not have roles in github.com/cli/cli: write, admin; found read
    // Exits with error code 1
    ```

For organizations, the command will exit with exit code 1 if the user is not a member of the organization.

- To check for membership:

    ```bash
    gh role -o canada-ca
    // 2022/11/20 01:01:00 nedredmond has no role in canada-ca
    // Exits with error code 1
    ```

- To check for specific roles:

    ```bash
    gh role -o my-org admin
    // 2022/11/20 01:01:00 nedredmond does not have role in my-org: admin; found member
    // Exits with error code 1

    gh role -o my-org admin member
    // member
    // Exits with code 0

    gh role -o my-org -t my-team maintainer
    // nedredmond does not have role in my-team: maintainer; found member
    // Exits with code 0
    ```

As of v3, you can now check the role of any arbitrary user instead of just the current user.

- To check for specific roles:

    ```bash
    gh role -r jedi/council -u anakin -f master
    // 2023/04/28 01:01:00 akakin does not have role in github.com/jedi/council: master; found knight
    // Exits with error code 1
    ```

## Roles

### Repository

In order of increasing permissions: `READ`, `TRIAGE`, `WRITE`, `MAINTAIN`, `ADMIN`. For more information, see [GitHub's documentation on repository roles](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/repository-roles-for-an-organization).

### Organization

Note that role names returned do not match [the documentation on organization roles](https://docs.github.com/en/organizations/managing-peoples-access-to-your-organization-with-roles/roles-in-an-organization#about-organization-roles). "Owners" have the `admin` role; the other two are `member` and `billing_manager`. See [the response schema](https://docs.github.com/en/rest/orgs/members#get-organization-membership-for-a-user) for more details..

#### Team

Team roles include `maintainer` and `member`. See [the response schema](https://docs.github.com/en/rest/teams/members#get-team-membership-for-a-user) for more details.

## Flags

- __`-h`__ Prints available flags and usage to the CLI.
- __`-r`__ The repo to check roles on.  Defaults to the current repo.
- __`-o`__ The org to check roles on. If blank, defaults to repo check.
- __`-t`__ The team for which to check roles. Only valid in combination with org flag.
- __`-u`__ The user to check roles for. Defaults to the current user.
- __`-host`__ The host for which to check roles. If blank, defaults to the gh config. Note that you will need to be be authenticated for the host through the gh cli.
- __`-f`__ Prints a friendly message instead of the simple role name.
  - i.e., `nedredmond has admin role in nedredmond/gh-role.`
- After any flags, list the roles you want to verify, separated by spaces. If none are provided, the command will print your current role instead.
  - There is no validation for provided roles and they are not case-sensitive. The only requirement is that the role is spelled correctly.

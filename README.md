# ssh-ldap-authorizedkeys

Go program to pull authorized keys out of LDAP. Expects two command-line
arguments, the first being the location of a config file like the provided
`config.toml.example`.

The second command line argument should be the attribute to filter by.
This tool assumes the following:

1. All SSH keys are listed under the `sshPublickey` attribute(s)

## Example Usage

```
./ssh-ldap-authorizedkeys ./config.toml exampleUsername
```

In `sshd_config`, you might configure it as:
```
AuthorizedKeysCommand /usr/bin/ssh-ldap-authorizedkeys /etc/ssh-ldap.toml %u
AuthorizedKeysCommandUser nobody
```

# ssh-ldap-authorizedkeys

Go program to pull authorized keys out of LDAP. Expects two command-line
arguments, the first being the location of a config file like below:

```toml
ldap_url = "ldaps://example.com:636"
ldap_base_dn = "dc=example,dc=com"
```

The second command line argument should be the `uid` attribute to filter by.
This tool assumes the following:

1. All accounts have a `uid` that's their POSIX username
2. All SSH keys are listed under the `sshPublickey` attribute(s)

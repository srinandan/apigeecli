## apigeecli kvms entries list

List KV Map entries

### Synopsis

List KV Map entries

```
apigeecli kvms entries list [flags]
```

### Options

```
  -e, --env string          Environment name
  -h, --help                help for list
  -m, --map string          KV Map Name
      --page-size int       Number of items to return on the list (default -1)
      --page-token string   next_page_token from the prior response to be used to fetch the next dataset
  -p, --proxy string        API Proxy name
```

### Options inherited from parent commands

```
  -a, --account string   Path Service Account private key in JSON
      --disable-check    Disable check for newer versions
  -o, --org string       Apigee organization name
  -t, --token string     Google OAuth Token
```

### SEE ALSO

* [apigeecli kvms entries](apigeecli_kvms_entries.md)	 - Manage Key Value Map Entries

###### Auto generated by spf13/cobra on 17-Apr-2022
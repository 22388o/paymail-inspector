## paymail

Inspect, validate domains or resolve paymail addresses

### Synopsis

```
__________                             .__.__    .___                                     __                
\______   \_____  ___.__. _____ _____  |__|  |   |   | ____   ____________   ____   _____/  |_  ___________ 
 |     ___/\__  \<   |  |/     \\__  \ |  |  |   |   |/    \ /  ___/\____ \_/ __ \_/ ___\   __\/  _ \_  __ \
 |    |     / __ \\___  |  Y Y  \/ __ \|  |  |__ |   |   |  \\___ \ |  |_> >  ___/\  \___|  | (  <_> )  | \/
 |____|    (____  / ____|__|_|  (____  /__|____/ |___|___|  /____  >|   __/ \___  >\___  >__|  \____/|__|   
                \/\/          \/     \/                   \/     \/ |__|        \/     \/     v0.0.20
```
Author: MrZ © 2020 github.com/mrz1836/paymail-inspector

This CLI app is used for interacting with paymail service providers.

Help contribute via Github!


### Examples

```
paymail -h
```

### Options

```
      --bsvalias string   The bsvalias version (default "1.0")
      --config string     Config file (default is $HOME/.paymail-inspector.yaml)
      --docs              Generate docs from all commands (./docs/commands)
  -h, --help              help for paymail
  -t, --skip-tracing      Turn off request tracing information
  -v, --version           version for paymail
```

### SEE ALSO

* [paymail brfc](paymail_brfc.md)	 - List all specs, search by keyword, or generate a new BRFC ID
* [paymail capabilities](paymail_capabilities.md)	 - Get the capabilities of the paymail domain
* [paymail p2p](paymail_p2p.md)	 - Starts a new P2P payment request
* [paymail resolve](paymail_resolve.md)	 - Resolves a paymail address
* [paymail validate](paymail_validate.md)	 - Validate a paymail address or domain
* [paymail verify](paymail_verify.md)	 - Verifies if a paymail is associated to a pubkey

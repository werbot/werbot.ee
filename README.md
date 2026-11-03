# werbot.ee

Enterprise add-on. Distributed separately under the Enterprise License
(see `LICENSE`) and not part of the community release.

The backend contract is wired through `add-on/ee/addon.go` and the `ee` build
tag:

```bash
TAGS=ee make build taco
```

No Enterprise modules are implemented yet. Routes and gRPC registration are
no-ops until the commercial feature set lands; the public core does not import
this module.

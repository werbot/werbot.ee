# werbot.ee

Enterprise add-on. Distributed separately under the Enterprise License
(see `LICENSE`) and not part of the community release.

The backend contract is wired through `add-on/ee/addon.go` and the `ee` build
tag:

```bash
TAGS=ee make build taco
```

## Web contract (v2)

An add-on contributes frontend routes and menus through its `web/` folder. The
folder names and data shapes match the host contract; pages are Svelte 5
components.

- `web/pages/**/*.svelte` — pages.
- `web/router/<section>/index.mjs` — the same route arrays as before
  (`path`, `name`, `children`, `meta`, `props`, `redirect`), with
  `component: () => import("...svelte")`.
- `web/pages/<section>/menu.ts` — plain menu data. `menu` is merged into the
  sidebar; `topMenu` is ignored (the profile dropdown is host-owned).
- `web/proto/**` — generated TypeScript, untouched.

Pages may import `@addon-kit` for navigation and host stores, and their own
`@addon/ee/web/**` files. Vue, Vue Router, Pinia, Notiwind and `@vueuse/core`
imports are forbidden and rejected by ESLint plus `scripts/check-no-vue.ts` in
the `web` repository. Each page is loaded through a dynamic import, so it stays
a separate chunk. The host generator reads `../add-on` by default; see
`web/README.md` for `ADDON_ROOTS` and `--roots`.

No Enterprise modules are implemented yet. Routes and gRPC registration are
no-ops until the commercial feature set lands; the public core does not import
this module.

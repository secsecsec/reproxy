Regex substitution proxy
========================
Dead simple HTTP proxy for regex substitution of content in HTML pages.


Building
--------
* `make build` - build for the current platform
* `make build-arm` - cross compile for ARM (Raspberry Pi for example)

Running
-------
`./reproxy -b '0.0.0.0:8000' -f conf.yaml`


Confguration
------------
The YAML config file has the following content:

```
substitutions:
  - pattern: "[123]foo"
    replace_with: "bar"
  - pattern: "baz"
    replace_with: "qux"
```

`replace_with` may contain references to groups in the pattern as described in the [regex documentation](https://golang.org/pkg/regexp/#Regexp.ReplaceAll).

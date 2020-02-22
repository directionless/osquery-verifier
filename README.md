# osquery verifier

A server testing tool, based on osquery

Send queries down.

Get results back.

Verify them with json-schema validation

## Use Cases

### Package Verification

### osquery query performance testing

### standalone

The simples way to use `ov` is in standalone mode. In this case, `ov`
executes osquery directly to run queries.

### as a server

Because `ov` is entirely based in osquery, it is easy to use as a TLS server.

`ov` runs as a server, waits for a host to connect, sends it the
appropriate tests, and evaluates the results. Output is stored in XXX

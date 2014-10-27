fldigiattach
============

```text
// fldigiattach is a tool to allow use of fldigi as a modem for Linux's AX.25-stack.
//
// This program creates a pty and uses kissattach(8) to attach the KISS
// interface of fldigi to an axport. After attachment, this program will
// act as a proxy between the AX.25-stack and fldigi to allow AX.25 over fldigi.
//
// Because kissattach daemonizes, you must kill it (as normal) after execution.
//
// fldigi 3.22 or later is required (KISS interface), as is kissattach (ax25-tools).
//
//   Usage of fldigiattach:
//     -mtu=0: Sets the mtu of the interface [default is paclen parameter in axports].
//     -port="": Name of a port given in the axports file
//     -rx-addr="127.0.0.1:7343": fldigi's rx address
//     -tx-addr="127.0.0.1:7342": fldigi's tx address
```

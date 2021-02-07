# bitbytepack

Small Go module to handle composition of byte arrays with embedded values.

## TODO

 -[ ] Create test for float functions.

 -[ ] Extend usage manual with how to use the Mult* functions

## Usage

For instance, you could have an array of bytes where you want to embed a value over multiple bytes

`command: 0x81 0x01 0x04 0x02 0x0z 0x0z 0xFF`

In this case `z` shows the two nibbles that one 8-bit object is split out on.

In this instance, the value can be embedded in the command array with the function
`bitbytepacket.WriteToArray(cmd, mask, value)`:

```
command = []byte{ 0x81, 0x01, 0x04, 0x02, 0x00, 0x00, 0xFF }
mask    = []byte{ 0x00, 0x00, 0x00, 0x00, 0x0F, 0x0F, 0x00 }
value   = byte(0x24)

bitbytepacket.WriteToArray(command, mask, value)
// returns []byte{ 0x81, 0x01, 0x04, 0x02, 0x02, 0x04, 0xFF }
//                                            *     *
```

Likewise, the value can be read of a byte array in a similar fashion:

```
command = []byte{ 0x81, 0x01, 0x04, 0x02, 0x02, 0x04, 0xFF }
mask    = []byte{ 0x00, 0x00, 0x00, 0x00, 0x0F, 0x0F, 0x00 }

bitbytepacket.ReadFromArray(command, mask)
// returns uint(0x24)
```

## Type overloads

The main function `bitbytepacket.ReadFromArray(...)` return `uint`, but to obtain the return
type `uint8`, `uint16`, `uint32`, `uint64`, use `bitbytepacket.ReadFromArray8(...)`.

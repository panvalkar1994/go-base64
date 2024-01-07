# go-base64

This library provides functionality for handling Base64 encoded data and binary data in Go.

The logic of the code revolves around the manipulation of strings and bytes to handle Base64 encoding requirements. 

One of the key aspects of Base64 encoding is that the encoded data must be a multiple of 4 bytes in length. If it's not, padding is added to make it so. This is handled internally by the library.

Another important aspect is the conversion of binary data to a string representation, which is also handled internally by the library.

The `B64` type is a simple wrapper around a string value, with `Encode`, `Decode`, and `String` methods that are exposed for use.

## Exposed Methods

- `Encode(input string) B64`: Encodes a string in B64 struct.
- `Decode() string`: Decodes a Base64 encoded string.
- `String() string`: Returns the underlying string value of the `B64` instance.

|  GEN<i>CAM |   | ![img-6.jpeg](img-6.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

## 2. Definitions

### 2.1. Device Description File

Device Description File means a GenICam compliant XML file describing the register space of a device.

### 2.2. String Encoding

All strings are encoded in ASCII, UTF8 or UTF16 depending on the BRM setting. The endianness of the characters in an encoded string must match the endianness of the containing register map. Strings defined in the bootstrap register map must follow the endianness of the GenCP Protocol. Strings in the device's register map must follow the implementation endianness.

### 2.3. Byte and Bit Order

The order and size of fields within packets are not depending on the endianness used. Fields are listed with its byte offset relative to the start of the section within a packet. All fields are byte aligned.

The endianness of all fields in GenCP protocol packets is technology specific and it must match the endianness of the bootstrap registers of the device.

This document does not define or use explicit bit numbers but identifies bits by its offset to the least significant bit. This notation is endian agnostic even though the offset matches the bit numbers of little-endian notations.

The endianness of the non-bootstrap registers is device implementation specific.

For reference, the byte order is described in Appendix B of RFC791.

Unless explicitly stated for a given technology, the endianness for GenCP-Implementations is big-endian.

### 2.4. GenCP Version

The GenCP version this document describes is

Major Version Number 1

Minor Version Number 3
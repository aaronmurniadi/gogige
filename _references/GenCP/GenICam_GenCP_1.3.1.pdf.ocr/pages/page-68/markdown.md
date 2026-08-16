|  GEN<i>CAM |   | ![img-72.jpeg](img-72.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

|   |  | 3 | 7 | ReservedSet to 0.  |
| --- | --- | --- | --- | --- |
|   |   |  10 | 6 | File FormatFile format of the file this entry points to.0 = Uncompressed GenICam XML file1 = ZIP containing a single GenICam XML file2-63 = reserved  |
|   |   |  16 | 8 | Schema-Minor VersionMinor Version of the GenICam Schema the GenICam file complies with.  |
|   |   |  24 | 8 | Schema-Major VersionMajor Version of the GenICam Schema the GenICam file complies with.  |
|  8 | 8 | Register AddressRegister Address at which the file can be read from.  |   |   |
|  8 | 16 | File SizeSize of the file this manifest entry points to in bytes.  |   |   |
|  20 | 24 | SHA1-HashSHA1 Hash of the file or 0 in case the hash is not available.  |   |   |
|  20 | 44 | ReservedSet to 0.  |   |   |

Table 34 – Manifest Entry Layout
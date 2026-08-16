|  GEN<i>CAM |   | ![img-58.jpeg](img-58.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 1 | User Defined Name SupportedSet if the device supports the User Defined Name register.  |
|  1 | 1 | Access Privilege SupportedSet if Heartbeat/Access Privilege is supported.  |
|  2 | 1 | Message Channel SupportedSet if the device supports a Message Channel.  |
|  3 | 1 | Timestamp SupportedSet if the device supports a timestamp register.  |
|  4 | 4 | String EncodingString Encoding of the BRM- 0x0 -> ASCII- 0x1 -> UTF8- 0x2 -> UTF16- 0x3-0xF -> Reserved  |
|  8 | 1 | FamilyName SupportedSet if the device supports the Family Name register.  |
|  9 | 1 | SBRM SupportedSet if the device supports a SBRM.  |
|  10 | 1 | Endianness Register SupportedSet if the device supports the Implementation Endianness register.  |
|  11 | 1 | Written Length Field SupportedSet to 1 if the device sends the length_written field in the SCD section of the WriteMemAck command.  |
|  12 | 1 | MultiEvent SupportedSet to 1 if the device supports multiple events in a single event command packet.  |
|  13 | 1 | Stacked Commands SupportedSet to 1 if the device supports ReadMemStacked and WriteMemStacked commands.  |
|  14 | 1 | Device Software Interface Version SupportedSet to 1 if the Device Software Interface Version register is supported.  |
|  15 | 49 | ReservedSet to 0.  |

Table 21 – Register Device Capabilities

#### 5.4.10. Maximum Device Response Time (MDRT)

Integer value containing the maximum time in milliseconds until a device reacts upon a received
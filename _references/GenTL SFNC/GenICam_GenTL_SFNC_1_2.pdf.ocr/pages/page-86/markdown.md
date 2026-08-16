|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

- OpenReadWrite: Open in Read/Write mode by this GenTL host
- OpenReadOnly: Open in Read access mode by this GenTL host

Corresponds to the DEVICE_INFO_ACCESS_STATUS command of DevGetInfo function.

3.3.1.14 DeviceChunkDataFormat

|  Name | DeviceChunkDataFormat  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | None GigEVision Custom  |

Chunk data format used by the device. This information allows devices based on other technologies or protocols than "standard" ones such as GigE Vision to inform the GenTL Consumer about the chunk data layout they use.

In contrast, one can assume that any generic GenTL Consumer will understand the GigE Vision chunk format because the GigE Vision chunk adapter is readily available.

Note that GenTL Consumers having access to a generic chunk adapter can use this adapter without caring about the actual data layout, provided that the GenTL Producer implements the DSGetBufferChunkData function. However, using the native chunk adapter might typically lead to slightly better performance.

- None: The device does not use chunk data at all.
- GigEVision: The device formats the chunk data using the chunk data format defined by GigE Vision specification version 1.x. The chunk data decoding algorithm (chunk adapter) common for the GigE Vision devices can be used.
- Custom: The device formats the chunk data using a custom, non-standard format. Without a-priori additional knowledge about the device and its implementation, the GenTL Consumer should always use the generic chunk adapter to decode the chunk data, not making any assumptions about the internal chunk data layout.
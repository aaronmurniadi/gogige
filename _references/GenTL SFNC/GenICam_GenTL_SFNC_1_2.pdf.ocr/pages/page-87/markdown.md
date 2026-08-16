|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.3.1.15 DeviceEventDataFormat

|  Name | DeviceEventDataFormat  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | None GigEVision GigEVisionExtendedId Custom  |

Enumeration, informing about the event data format used by the device (meaning the "device events", see event type EVENT_REMOTE_DEVICE (named EVENT_FEATURE_DEVEVENT in GenTL up to version 1.4). This allows devices based on other technologies or protocols than "standard" ones such as GigE Vision to inform the GenTL Consumer about the event data layout they use.

In contrast, one can assume that any generic GenTL Consumer will understand the GigE Vision event format because the GigE Vision event adapter is readily available

Note that GenTL Consumers having access to a generic event adapter can use this adapter without caring about the actual data layout.

- None: The device does not use event data at all.
- GigEVision: The device formats the event data using the event data format defined by GigE Vision specification version 1.x. The event data decoding algorithm (event adapter) common for the GigE Vision devices can be used.
- GigEVisionExtendedId: The device formats the event data using the event data format defined by GigE Vision specification version 2.x. The event data decoding algorithm (event adapter) common for the GigE Vision devices can be used.
- Custom: The device formats the event data using a custom, non-standard format. Without a-priori additional knowledge about the device and its implementation, the GenTL Consumer should always use the generic event adapter to decode the event data, not making any assumptions about the internal event data layout.
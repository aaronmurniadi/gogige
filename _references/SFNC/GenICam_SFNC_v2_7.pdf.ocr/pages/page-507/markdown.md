|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Level | Recommended  |
| --- | --- |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | Off On Automatic  |

Controls the device's streaming format.

Possible values are:

- Off: The device will only stream data in its native format.
- On: The device will stream all its data in the generic GenDC format.
- Automatic: The device will automatically choose in which format it streams its data.

Note: This feature is meant to globally control the data streaming format of the device.

As such, it has priority over other features (except TestPayloadFormatMode) that might influence the output payload format or content.

### 27.2.7 GenDCStreamingStatus

|  Name | GenDCStreamingStatus  |
| --- | --- |
|  Category | TransportLayerControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | Off On  |

Returns whether the current device data streaming format is GenDC. This value is conditioned by the GenDCStreamingMode.

This feature can be used to determine if the device will stream in its native or in the generic GenDC format and if GenDC related features like GenDCDescriptor and GenDCFlowmappingTable can be used.
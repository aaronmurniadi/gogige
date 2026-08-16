|  GEN<ì>CAM |   | ![img-154.jpeg](img-154.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Interface | IBoolean  |
| --- | --- |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | TrueFalse  |

Flip vertically of the image sent by the device.

24.31 ChunkTimestamp

|  Name | ChunkTimestamp  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Returns the Timestamp of the image included in the payload at the time of the FrameStart internal event. See Figure 5-3 for more details on FrameStart.

24.32 ChunkTimestampLatchValue

|  Name | ChunkTimestampLatchValue  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | ns  |
|  Visibility | Expert  |
|  Values | ≥0  |
|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Visibility | Expert  |
| --- | --- |
|  Values | EncoderUp EncoderDown EncoderIdle EncoderStatic  |

Returns the motion status of the selected encoder.

Possible values are:

- EncoderUp: The encoder counter last incremented.
- EncoderDown: The encoder counter last decremented.
- EncoderIdle: The encoder is not active.
- EncoderStatic: No motion within the EncoderTimeout time.

### 24.42 ChunkExposureTimeSelector

|  Name | ChunkExposureTimeSelector  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Common Red Green Blue Cyan Magenta Yellow Infrared Ultraviolet Stage1 Stage2 ...  |

Selects which exposure time is read by the ChunkExposureTime feature.

The possible values for ChunkExposureTimeSelector are:
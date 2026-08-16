|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

#### 2.3.2 Device Control

Contains the features related to configure a specific Device module.

Table 2-11: Device Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  DeviceControl | R | All | ICategory | R | - | E | Category that contains all Device Control features of the Device module.  |
|  DeviceEndianessMechanism | M | GEV | IEnumeration | R/W | - | E | Identifies the endianess handling mode.  |
|  LinkCommandTimeout | R | All | IFloat | R/W | us | G | Specifies application timeout for the control channel communication.  |
|  LinkCommandRetryCount | R | All | IInteger | R/W | - | G | Specifies maximum number of tries before failing the control channel commands.  |

#### 2.3.3 Stream Enumeration

Contains the features related to the enumeration of available Data Stream modules within a specific Device module.

Table 2-12: Stream Enumeration Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  StreamEnumeration | R | All | ICategory | R | - | B | Category that contains all Stream Enumeration features of the Device module.  |
|  StreamSelector | M | All | IInteger | R/W | - | B | Selector for the different stream channels.  |
|  StreamID[StreamSelector] | M | All | IString | R | - | B | Device unique ID for the stream.  |
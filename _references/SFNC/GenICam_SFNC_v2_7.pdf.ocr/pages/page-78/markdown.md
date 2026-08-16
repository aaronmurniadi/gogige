|  ChunkScan3dFocalLength | O | IFloat | R | Pixel | E | Returns the focal length of the camera in pixel.  |
| --- | --- | --- | --- | --- | --- | --- |
|  ChunkScan3dBaseline | O | IFloat | R | m | E | Returns the baseline as the physical distance of two cameras in a stereo camera setup.  |
|  ChunkScan3dPrincipalPointU | O | IFloat | R | Pixel | E | Returns the value of this feature gives the horizontal position of the principal point, relative to the region origin, i.  |
|  ChunkScan3dPrincipalPointV | O | IFloat | R | Pixel | E | Returns the value of this feature gives the vertical position of the principal point, relative to the region origin, i.  |

## 2.23 Test Control

Contains the features related to the control of the test features (See the Test Control chapter for details).

Table 2-23: Test Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  TestControl | R | ICategory | R | - | G | Category for Test Control features.  |
|  TestPendingAck | O | IInteger | R/W | ms | G | Tests the device's pending acknowledge feature.  |
|  TestEventGenerate | O | ICommand | (R)/W | - | G | Generates a Test Event.  |
|  TestPayloadFormatMode | R | IEnumeration | R/W | - | G | This feature allows setting a device in test mode and to output a specific payload format for validation of data streaming.  |

## 2.24 GenICam Control

Contains the features related to GenICam control and access (See the GenICam Control chapter for details).

Table 2-24: GenICam Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  Root | M | ICategory | R | - | B | Provides the Root of the GenICam features tree.  |
|  Device | M | IPort | R/W | - | I | Provides the default GenICam port of the Device.  |
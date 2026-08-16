|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

TriggerSelector = ExposureActive;
TriggerMode = On;
TriggerActivation = LevelHigh;
TriggerSource = Line1;

But simply by adding:

ExposureMode = TriggerControlled;

The Exposure duration will become controlled by the length of the positive pulse on physical Line 1.

### 5.7.2 ExposureTimeMode

|  Name | ExposureTimeMode  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Common Individual  |

Sets the configuration mode of the ExposureTime feature.

The possible values for ExposureTimeMode are:

- Common: The exposure time is common to all the color components. The common ExposureTime value to use can be set selecting it with ExposureTimeSelector[Common].
- Individual: The exposure time is individual for each color component. Each individual ExposureTime values to use can be set by selecting them with ExposureTimeSelector.

### 5.7.3 ExposureTimeSelector

|  Name | ExposureTimeSelector  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
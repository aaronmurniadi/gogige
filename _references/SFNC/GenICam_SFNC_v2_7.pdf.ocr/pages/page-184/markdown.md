|  Category | AcquisitionControl  |
| --- | --- |
|  Level | Recommended  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | Hz  |
|  Visibility | Beginner  |
|  Values | Device-specific  |

Controls the rate (in Hertz) at which the Lines in a Frame are captured.

TriggerMode must be Off for the Line trigger.

This is generally useful for linescan camera only.

### 5.5.9 AcquisitionLineRateEnable

|  Name | AcquisitionLineRateEnable  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Controls if the AcquisitionLineRate feature is writable and used to control the acquisition rate. Otherwise, the acquisition rate is implicitly controlled by the combination of other features like ExposureTime, etc...

### 5.5.10 AcquisitionStatusSelector

|  Name | AcquisitionStatusSelector  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
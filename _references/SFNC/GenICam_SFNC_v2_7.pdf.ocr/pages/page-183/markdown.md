|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

This feature is used only if the FrameBurstStart trigger is enabled and the FrameBurstEnd trigger is disabled. Note that the total number of frames captured is also conditioned by AcquisitionFrameCount if AcquisitionMode is MultiFrame and ignored if AcquisitionMode is Single.

### 5.5.6 AcquisitionFrameRate

|  Name | AcquisitionFrameRate  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | Hz  |
|  Visibility | Beginner  |
|  Values | Device-specific  |

Controls the acquisition rate (in Hertz) at which the frames are captured.

TriggerMode must be Off for the Frame trigger.

### 5.5.7 AcquisitionFrameRateEnable

|  Name | AcquisitionFrameRateEnable  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Controls if the AcquisitionFrameRate feature is writable and used to control the acquisition rate. Otherwise, the acquisition rate is implicitly controlled by the combination of other features like ExposureTime, etc...

### 5.5.8 AcquisitionLineRate

|  Name | AcquisitionLineRate  |
| --- | --- |
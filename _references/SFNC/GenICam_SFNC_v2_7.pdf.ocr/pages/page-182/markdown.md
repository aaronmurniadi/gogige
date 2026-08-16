|  Unit | -  |
| --- | --- |
|  Visibility | Expert  |
|  Values | -  |

Arms the device before an AcquisitionStart command. This optional command validates all the current features for consistency and prepares the device for a fast start of the Acquisition.

If not used explicitly, this command will be automatically executed at the first AcquisitionStart but will not be repeated for the subsequent ones unless a feature is changed in the device.

### 5.5.4 AcquisitionFrameCount

|  Name | AcquisitionFrameCount  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | Integer  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥1  |

Number of frames to acquire in MultiFrame Acquisition mode.

### 5.5.5 AcquisitionBurstFrameCount

|  Name | AcquisitionBurstFrameCount  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Optional  |
|  Interface | Integer  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥1  |

Number of frames to acquire for each FrameBurstStart trigger.
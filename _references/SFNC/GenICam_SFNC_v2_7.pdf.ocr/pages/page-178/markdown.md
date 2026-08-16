|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

/* Framescan continuous acquisition with 1 Hardware trigger controlling the start of the acquisition and 2 others hardware triggers to start and stop the exposure of each frame.

*/

AcquisitionMode = Continuous;
TriggerSelector = AcquisitionStart;
TriggerMode = On;
TriggerSource = Line1;
ExposureMode = TriggerControlled;
TriggerSelector = ExposureStart;
TriggerMode = On;
TriggerSource = Line3;
TriggerSelector = ExposureStop;
TriggerMode = On;
TriggerSource = Line4;
AcquisitionStart();
...
AcquisitionStop();

## 5.5 Acquisition Control features

This section gives the detailed description of all the Acquisition related features.

### 5.5.1 AcquisitionControl

|  Name | AcquisitionControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category for the acquisition and trigger control features.
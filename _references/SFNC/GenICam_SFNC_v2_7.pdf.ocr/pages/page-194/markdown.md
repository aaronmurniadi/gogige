|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Specifies the delay in microseconds (us) to apply after the trigger reception before activating it.

### 5.6.8 TriggerDivider

|  Name | TriggerDivider[TriggerSelector]  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | Integer  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Device-specific  |

Specifies a division factor for the incoming trigger pulses.

### 5.6.9 TriggerMultiplier

|  Name | TriggerMultiplier[TriggerSelector]  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | Integer  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Device-specific  |

Specifies a multiplication factor for the incoming trigger pulses. It is generally used in conjunction with TriggerDivider to control the ratio of triggers that are accepted.

### 5.7 Exposure Control features

The Exposure Control section describes all features related to the exposure of the photosensitive cells (shutter control) during image acquisition.

The Exposure of the photosensitive cells during Frame or Line acquisition can be in 3 different modes.

- ExposureMode can be Off to disable the Shutter and let it open.
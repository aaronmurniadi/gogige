|  Values | 1..n  |
| --- | --- |

The number of knee-points as well as the number of additional exposure slopes used for multi-slope exposure.

### 5.8.3 MultiSlopeKneePointSelector

|  Name | MultiSlopeKneePointSelector  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | 1..n  |

Selects the parameters for controlling an additional slope in multi-slope exposure.

### 5.8.4 MultiSlopeExposureLimit

|  Name | MultiSlopeExposureLimit[MultiSlopeKneePointSelector]  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | %  |
|  Visibility | Expert  |
|  Values | 0..100  |

Percent of the ExposureTime at a certain knee-point of multi-slope exposure.

### 5.8.5 MultiSlopeSaturationThreshold

|  Name | MultiSlopeSaturationThreshold[MultiSlopeKneePointSelector]  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/(Write)  |
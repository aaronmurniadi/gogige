### 5.8.1 MultiSlopeMode

|  Name | MultiSlopeMode  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Off Manual PresetSoft PresetMedium PresetAggressive  |

Controls multi-slope exposure state.

Possible values are:

- Off: Multi-slope exposure is disabled.
- Manual: Multi-slope exposure is enabled. Control is possible per knee-point with two of the features MultiSlopeExposureLimit, MultiSlopeSaturationThreshold, MultiSlopeIntensityLimit or MultiSlopeExposureGradient.
- PresetSoft: Multi-slope exposure is enabled. A predefined parameter set resulting in a soft effect is selected.
- PresetMedium: Multi-slope exposure is enabled. A predefined parameter set resulting in a medium effect is selected.
- PresetAggressive: Multi-slope exposure is enabled. A predefined parameter set resulting in an aggressive effect is selected.

### 5.8.2 MultiSlopeKneePointCount

|  Name | MultiSlopeKneePointCount  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
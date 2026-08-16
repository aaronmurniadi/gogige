|  GEN<ì>CAM |   | ![img-160.jpeg](img-160.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Access | Write  |
| --- | --- |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | -  |

This feature is deprecated (See TimestampLatch). It was used to latche the current timestamp counter into GevTimestampValue.

#### 27.4.35 GevTimestampControlReset (Deprecated)

|  Name | GevTimestampControlReset  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | Write  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | -  |

This feature is deprecated (See TimestampReset). It was used to reset the timestamp counter to 0. This feature is not available or as no effect when PTP is used.

#### 27.4.36 GevTimestampValue (Deprecated)

|  Name | GevTimestampValue  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit |   |
|  Visibility | Invisible  |
|  Values | ≥0  |

This feature is deprecated (See TimestampLatchValue). It was used to return the latched 64-bit value of the timestamp counter.
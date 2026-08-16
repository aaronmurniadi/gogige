|  Unit | %  |
| --- | --- |
|  Visibility | Expert  |
|  Values | 0..100  |

The percentage of the full saturation that is applied at a certain knee-point of a multi-slope exposure.

The limits are sensor-specific and might not span the whole range of 0..100%.

In principle, setting this value to 100% would effectively disable this knee-point, while setting this value to 0% would effectively start exposure at this knee-point.

### 5.8.6 MultiSlopeIntensityLimit

|  Name | MultiSlopeIntensityLimit[MultiSlopeKneePointSelector]  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/(Write)  |
|  Unit | %  |
|  Visibility | Expert  |
|  Values | 0..100  |

The relative intensity which divides intensities influenced by different exposure slopes.

### 5.8.7 MultiSlopeExposureGradient

|  Name | MultiSlopeExposureGradient[MultiSlopeKneePointSelector]  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | >0  |

The gradient of the additional slope that is defined by this knee-point.

This gradient is computed by (MultiSlopeSaturationThreshold[n+1]- MultiSlopeSaturationThreshold[n]) / (MultiSlopeExposureLimit[n+1]- MultiSlopeExposureLimit[n]).
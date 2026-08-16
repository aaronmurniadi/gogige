|  AcquisitionStatus[AcquisitionStatusSelector] | R | IBoolean | R | - | E | Reads the state of the internal acquisition signal selected using AcquisitionStatusSelector.  |
| --- | --- | --- | --- | --- | --- | --- |
|  TriggerSelector | R | IEnumeration | R/W | - | B | Selects the type of trigger to configure.  |
|  TriggerMode[TriggerSelector] | R | IEnumeration | R/W | - | B | Controls if the selected trigger is active.  |
|  TriggerSoftware[TriggerSelector] | R | ICommand | (R)/W | - | B | Generates an internal trigger.  |
|  TriggerSource[TriggerSelector] | R | IEnumeration | R/W | - | B | Specifies the internal signal or physical input Line to use as the trigger source.  |
|  TriggerActivation[TriggerSelector] | R | IEnumeration | R/W | - | B | Specifies the activation mode of the trigger.  |
|  TriggerOverlap[TriggerSelector] | R | IEnumeration | R/W | - | E | Specifies the type trigger overlap permitted with the previous frame or line.  |
|  TriggerDelay[TriggerSelector] | R | IFloat | R/W | us | E | Specifies the delay in microseconds (us) to apply after the trigger reception before activating it.  |
|  TriggerDivider[TriggerSelector] | R | IInteger | R/W | - | E | Specifies a division factor for the incoming trigger pulses.  |
|  TriggerMultiplier[TriggerSelector] | R | IInteger | R/W | - | E | Specifies a multiplication factor for the incoming trigger pulses.  |
|  ExposureMode | R | IEnumeration | R/W | - | B | Sets the operation mode of the Exposure.  |
|  ExposureTimeMode | O | IEnumeration | R/W | - | B | Sets the configuration mode of the ExposureTime feature.  |
|  ExposureTimeSelector | O | IEnumeration | R/W | - | B | Selects which exposure time is controlled by the ExposureTime feature.  |
|  ExposureTime[ExposureTimeSelector] | R | IFloat | R/W | us | B | Sets the Exposure time when ExposureMode is Timed and ExposureAuto is Off.  |
|  ExposureAuto | O | IEnumeration | R/W | - | B | Sets the automatic exposure mode when ExposureMode is Timed.  |
|  MultiSlopeMode | O | IEnumeration | R/W | - | B | Controls multi-slope exposure state.  |
|  MultiSlopeKneePointCount | O | IInteger | R/W | - | E | The number of knee-points as well as the number of additional exposure slopes used for multi-slope exposure.  |
|  MultiSlopeKneePointSelector | O | IInteger | R/W | - | E | Selects the parameters for controlling an additional slope in multi-slope exposure.  |
|  MultiSlopeExposureLimit[MultiSlopeKneePointSelector] | O | IFloat | R/W | % | E | Percent of the ExposureTime at a certain knee-point of multi-slope exposure.  |
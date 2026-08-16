|  Access | Read/Write  |
| --- | --- |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | AcquisitionTriggerWait AcquisitionActive AcquisitionTransfer FrameTriggerWait FrameActive ExposureActive FrameTransfer (Deprecated)  |

Selects the internal acquisition signal to read using AcquisitionStatus.

See Figure 5-1 and Figure 5-3 for details.

Possible values are:

- AcquisitionTriggerWait: Device is currently waiting for a trigger for the capture of one or many frames.
- AcquisitionActive: Device is currently doing an acquisition of one or many frames.
- AcquisitionTransfer: Device is currently transferring an acquisition of one or many frames.
- FrameTriggerWait: Device is currently waiting for a frame start trigger.
- FrameActive: Device is currently doing the capture of a frame.
- ExposureActive: Device is doing the exposure of a frame.
- FrameTransfer (Deprecated): See TransferStatus.

### 5.5.11 AcquisitionStatus

|  Name | AcquisitionStatus[AcquisitionStatusSelector]  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | IBoolean  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |
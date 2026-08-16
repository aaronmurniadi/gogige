|  Interface | IEnumeration  |
| --- | --- |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Off AcquisitionTrigger AcquisitionTriggerMissed AcquisitionStart AcquisitionEnd FrameTrigger FrameTriggerMissed FrameStart FrameEnd FrameBurstStart FrameBurstEnd ExposureStart ExposureEnd Line0 (If 0 based), Line1, Line2, ... UserOutput0, UserOutput1, UserOutput2, ... Counter0Start (If 0 based), Counter1Start, Counter2Start, ... Counter0End (If 0 based), Counter1End, Counter2End, ... Timer0Start (If 0 based), Timer1Start, Timer2Start, ... Timer0End (If 0 based), Timer1End, Timer2End, ... Encoder0, Encoder1, Encoder2, ... LogicBlock0 (If 0 based), LogicBlock1, LogicBlock2, ... SoftwareSignal0 (If 0 based), SoftwareSignal1, SoftwareSignal2, ... Action0 (If 0 based), Action1, Action2, ... LinkTrigger0, LinkTrigger1, LinkTrigger2, ... CC1, CC2, CC3, CC4  |

Specifies the internal signal or physical input line to use as the sequencer trigger source.

Possible values are:

- Off: Disables the sequencer trigger.
- AcquisitionTrigger: Starts with the reception of the Acquisition Trigger.
- AcquisitionTriggerMissed: Starts with the reception of the missed Acquisition Trigger.
- AcquisitionStart: Starts with the reception of the Acquisition Start.
- AcquisitionEnd: Starts with the reception of the Acquisition End.
- FrameTrigger: Starts with the reception of the Frame Start Trigger.
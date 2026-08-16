|   | AcquisitionStart AcquisitionEnd FrameTrigger FrameTriggerMissed FrameStart FrameEnd LineTrigger LineTriggerMissed LineStart LineEnd ExposureStart ExposureEnd Line0 (If 0 based), Line1, Line2, ... UserOutput0, UserOutput1, UserOutput2,... Counter0Start (If 0 based), Counter1Start, Counter2Start, ... Counter0End (If 0 based), Counter1End, Counter2End, ... Timer0Start (If 0 based), Timer1Start, Timer2Start, ... Timer0End (If 0 based), Timer1End, Timer2End, ... Encoder0 (if 0 based), Encoder1, Encoder2, ... LogicBlock0 (If 0 based), LogicBlock1, LogicBlock2, ... SoftwareSignal0 (If 0 based), SoftwareSignal1, SoftwareSignal2, ... Action0 (If 0 based), Action1, Action2, ... LinkTrigger0 (If 0 based), LinkTrigger1, LinkTrigger2, ...  |
| --- | --- |

Selects the signals that will be the source to reset the Counter.

Possible values are:

- **Off**: Disable the Counter Reset trigger.
- **CounterTrigger**: Resets with the reception of a trigger on the **CounterTriggerSource**.
- **AcquisitionTrigger**: Resets with the reception of the Acquisition Trigger.
- **AcquisitionTriggerMissed**: Resets with the reception of the missed Acquisition start trigger.
- **AcquisitionStart**: Resets with the reception of the Acquisition Start.
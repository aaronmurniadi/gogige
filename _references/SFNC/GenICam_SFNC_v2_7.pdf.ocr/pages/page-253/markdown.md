|   | Action0, Action1, Action2, ... LinkTrigger0, LinkTrigger1, LinkTrigger2, ...  |
| --- | --- |

Selects the source to start the Counter.

Possible values are:

- Off: Disables the Counter trigger.
- AcquisitionTrigger: Starts with the reception of the Acquisition Trigger.
- AcquisitionTriggerMissed: Device has missed an Acquisition start trigger.
- AcquisitionStart: Starts with the reception of the Acquisition Start.
- AcquisitionEnd: Starts with the reception of the Acquisition End.
- FrameTrigger: Starts with the reception of the Frame Start Trigger.
- FrameTriggerMissed: Device has missed a Frame start trigger.
- FrameStart: Starts with the reception of the Frame Start.
- FrameEnd: Starts with the reception of the Frame End.
- FrameBurstStart: Starts with the reception of the Frame Burst Start.
- FrameBurstEnd: Starts with the reception of the Frame Burst End.
- LineTrigger: Starts with the reception of the Line Start Trigger.
- LineTriggerMissed: Device has missed a Line start trigger.
- LineStart: Starts with the reception of the Line Start.
- LineEnd: Starts with the reception of the Line End.
- ExposureStart: Starts with the reception of the Exposure Start.
- ExposureEnd: Starts with the reception of the Exposure End.
- Line0 (If 0 based), Line1, Line2, ...: Starts when the specified CounterTriggerActivation condition is met on the chosen I/O Line.
- UserOutput0, UserOutput1, UserOutput2, ...: Specifies which User Output bit signal to use as internal source for the trigger.
- Counter0Start, Counter1Start, Counter2Start, ...: Starts with the reception of the Counter Start.
- Counter0End, Counter1End, Counter2End, ...: Starts with the reception of the Counter End.
- Timer0Start, Timer1Start, Timer2Start, ...: Starts with the reception of the Timer Start.
- Timer0End, Timer1End, Timer2End, ...: Starts with the reception of the Timer End.
- Encoder0, Encoder1, Encoder2, ...: Starts with the reception of the Encoder output signal.
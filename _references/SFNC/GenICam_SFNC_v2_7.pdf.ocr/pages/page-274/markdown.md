Possible values are:

- Off: Disable the Encoder Reset trigger.
- AcquisitionTrigger: Resets with the reception of the Acquisition Trigger.
- AcquisitionTriggerMissed: Resets with the reception of a missed Acquisition Trigger.
- AcquisitionStart: Resets with the reception of the Acquisition Start.
- AcquisitionEnd: Resets with the reception of the Acquisition End.
- FrameTrigger: Resets with the reception of the Frame Start Trigger.
- FrameTriggerMissed: Resets with the reception of a missed Frame Trigger.
- FrameStart: Resets with the reception of the Frame Start.
- FrameEnd: Resets with the reception of the Frame End.
- FrameBurstStart: Resets with the reception of the Frame Burst Start.
- FrameBurstEnd: Resets with the reception of the Frame Burst End.
- ExposureStart: Resets with the reception of the Exposure Start.
- ExposureEnd: Resets with the reception of the Exposure End.
- Line0 (If 0 based), Line1, Line2, ...: Resets by the chosen I/O Line.
- UserOutput0, UserOutput1, UserOutput2, ...: Resets by the chosen User Output bit.
- Counter0Start, Counter1Start, Counter2Start, ...: Resets with the reception of the Counter Start.
- Counter0End, Counter1End, Counter2End, ...: Resets with the reception of the Counter End.
- Timer0Start, Timer1Start, Timer2Start, ...: Resets with the reception of the Timer Start.
- Timer0End, Timer1End, Timer2End, ...: Resets with the reception of the Timer End.
- LogicBlock0, LogicBlock1, LogicBlock2, ...: Reset by the choosen Logic Block signal.
- SoftwareSignal0, SoftwareSignal1, SoftwareSignal2, ...: Resets on the reception of the Software Signal.
- Action0, Action1, Action2, ...: Resets on assertions of the chosen action signal (Broadcasted signal on the transport layer).
- LinkTrigger0, LinkTrigger1, LinkTrigger2, ...: Resets on the reception of the chosen Link Trigger (received from the transport layer).

Note that the value of the Encoder counter at time of reset is automatically latched and reflected in EncoderValueAtReset.
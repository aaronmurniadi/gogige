- FrameTriggerMissed: Starts with the reception of the missed Frame Trigger.
- FrameStart: Starts with the reception of the Frame Start.
- FrameEnd: Starts with the reception of the Frame End.
- FrameBurstStart: Starts with the reception of the Frame Burst Start.
- FrameBurstEnd: Starts with the reception of the Frame Burst End.
- ExposureStart: Starts with the reception of the Exposure Start.
- ExposureEnd: Starts with the reception of the Exposure End.
- Line0 (If 0 based), Line1, Line2, ...: Starts when the specified TimerTriggerActivation condition is met on the chosen I/O Line.
- UserOutput0, UserOutput1, UserOutput2, ...: Specifies which User Output bit signal to use as internal source for the trigger.
- Counter0Start, Counter1Start, Counter2Start, ...: Starts with the reception of the Counter Start.
- Counter0End, Counter1End, Counter2End, ...: Starts with the reception of the Counter End.
- Timer0Start, Timer1Start, Timer2Start, ...: Starts with the reception of the Timer Start.
- Timer0End, Timer1End, Timer2End, ...: Starts with the reception of the Timer End.
- Encoder0, Encoder1, Encoder2, ...: Starts with the reception of the Encoder output signal.
- LogicBlock0, LogicBlock1, LogicBlock2, ...: Starts with the reception of the Logic Block output signal.
- SoftwareSignal0, SoftwareSignal1, SoftwareSignal2, ...: Starts on the reception of the Software Signal.
- Action0, Action1, Action2, ...: Starts with the assertion of the chosen action signal.
- LinkTrigger0, LinkTrigger1, LinkTrigger2, ...: Starts with the reception of the chosen Link Trigger.
- CC1, CC2, CC3, CC4: Index of the Camera Link physical line and associated I/O control block to use. This ensures a direct mapping between the lines on the frame grabber and on the camera. Applicable to CameraLink products only.

### 17.17 SequencerTriggerActivation

|  Name | SequencerTriggerActivation[SequencerSetSelector][SequencerPathSelector]  |
| --- | --- |
|  Category | SequencerControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
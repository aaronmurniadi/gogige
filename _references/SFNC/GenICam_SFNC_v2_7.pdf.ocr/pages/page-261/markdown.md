- FrameEnd: Starts with the reception of the Frame End.
- FrameBurstStart: Starts with the reception of the Frame Burst Start.
- FrameBurstEnd: Starts with the reception of the Frame Burst End.
- LineTrigger: Starts with the reception of the Line Start Trigger.
- LineTriggerMissed: Starts with the reception of a missed Line Trigger.
- LineStart: Starts with the reception of the Line Start.
- LineEnd: Starts with the reception of the Line End.
- ExposureStart: Starts with the reception of the Exposure Start.
- ExposureEnd: Starts with the reception of the Exposure End.
- Line0 (If 0 based), Line1, Line2, ...: Starts when the specified TimerTriggerActivation condition is met on the chosen I/O Line.
- UserOutput0, UserOutput1, UserOutput2, ...: Specifies which User Output bit signal to use as internal source for the trigger.
- Counter0Start, Counter1Start, Counter2Start, ...: Starts with the reception of the Counter Start.
- Counter0End, Counter1End, Counter2End, ...: Starts with the reception of the Counter End.
- Timer0Start, Timer1Start, Timer2Start, ...: Starts with the reception of the Timer Start.
- Timer0End, Timer1End, Timer2End, ...: Starts with the reception of the Timer End. Note that a timer can retrigger itself to achieve a free running Timer.
- Encoder0, Encoder1, Encoder2, ...: Starts with the reception of the Encoder output signal.
- LogicBlock0, LogicBlock1, LogicBlock2, ...: Starts with the reception of the Logic Block output signal.
- SoftwareSignal0, SoftwareSignal1, SoftwareSignal2, ...: Starts on the reception of the Software Signal.
- Action0, Action1, Action2, ...: Starts with the assertion of the chosen action signal.
- LinkTrigger0, LinkTrigger1, LinkTrigger2, ...: Starts with the reception of the chosen Link Trigger.

### 10.5.8 TimerTriggerActivation

|  Name | TimerTriggerActivation[TimerSelector]  |
| --- | --- |
|  Category | CounterAndTimerControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
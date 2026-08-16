Selects the source signal for the input into the Logic Block. True or False indicates the input is forced constant.

Possible values are:

- **True**: Logic Block input is forced to One.
- **False**: Logic Block input is forced to Zero.
- **AcquisitionTriggerWait**: Device is currently waiting for a trigger for the capture of one or many Frames.
- **AcquisitionTrigger**: Device is receiving a trigger for the capture of one or many Frames.
- **AcquisitionTriggerMissed**: Device has missed a trigger for the capture of one or many Frames.
- **AcquisitionActive**: Device is acquiring one or many Frames.
- **FrameTriggerWait**: Device is currently waiting for a Frame start trigger.
- **FrameTrigger**: Device is receiving a Frame start trigger.
- **FrameTriggerMissed**: Device has missed a Frame start trigger.
- **FrameActive**: Device is currently doing the capture of a Frame.
- **ExposureActive**: Device is doing the exposure of a Frame (or Line).
- **LineTriggerWait**: Device is currently waiting for a Line start trigger.
- **LineTrigger**: Device is receiving a Line start trigger.
- **LineTriggerMissed**: Device has missed a Line start trigger.
- **LineActive**: Device is currently doing the capture of a Line.
- **Counter0Active, Counter1Active, Counter2Active**, ...: The chosen counter is in active state (counting).
- **Timer0Active, Timer1Active, Timer2Active**, ...: The chosen Timer is in active state.
- **Encoder0, Encoder1, Encoder2**, ...: The chosen Encoder Output state.
- **LogicBlock0, LogicBlock1, LogicBlock2**, ...: The choosen Logic Block output state.
- **Line0, Line1, Line2**, ...: The chosen I/OLine state.
- **UserOutput0, UserOutput1, UserOutput2**, ...: The chosen User Output bit state as defined by its current **UserOutputValue**.
- **SoftwareSignal0, SoftwareSignal1, SoftwareSignal2**, ...: The choosen Software Signal output state.
- **Stream0TransferActive, Stream1TransferActive**, ...: Transfer on the stream is active.
- **Stream0TransferPaused, Stream1TransferPaused**, ...: Transfer on the stream is paused.
- **Stream0TransferStopping, Stream1TransferStopping**, ...: Transfer on the stream is stopping.
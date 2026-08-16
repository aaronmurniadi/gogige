|   | LineTrigger LineTriggerMissed LineActive Counter0Active, Counter1Active, Counter2Active, ... Timer0Active, Timer1Active, Timer2Active, ... Encoder0, Encoder1, Encoder2, ... LogicBlock0, LogicBlock1, LogicBlock2, ... SoftwareSignal0, SoftwareSignal1, SoftwareSignal2, ... Stream0TransferActive, Stream1TransferActive, ... Stream0TransferPaused, Stream1TransferPaused, ... Stream0TransferStopping, Stream1TransferStopping, ... Stream0TransferStopped, Stream1TransferStopped, ... Stream0TransferOverflow, Stream1TransferOverflow, ...  |
| --- | --- |

Selects which internal acquisition or I/O source signal to output on the selected Line. **LineMode** must be **Output**.

See Figure 9-1 for details.

Possible values are:

- **Off**: Line output is disabled (Tri-State).
- **UserOutput0, UserOutput1, UserOutput2, ...**: The chosen User Output Bit state as defined by its current **UserOutputValue**.
- **AcquisitionTriggerWait**: Device is currently waiting for a trigger for the capture of one or many Frames.
- **AcquisitionActive**: Device is currently doing an acquisition of one or many Frames.
- **AcquisitionTriggerWait**: Device is currently waiting for an Acquisition start trigger.
- **AcquisitionTrigger**: Device is receiving an Acquisition start trigger.
- **AcquisitionTriggerMissed**: Device has missed an Acquisition start trigger.
- **FrameTriggerWait**: Device is currently waiting for a Frame start trigger.
- **FrameTrigger**: Device is receiving a Frame start trigger.
- **FrameTriggerMissed**: Device has missed a Frame start trigger.
- **FrameActive**: Device is currently doing the capture of a Frame.
- **LineTriggerWait**: Device is currently waiting for a Line start trigger.
- **LineTrigger**: Device is receiving a Line start trigger.
- **LineTriggerMissed**: Device has missed a Line start trigger.
- **LineActive**: Device is currently doing the capture of a Line.
- **ExposureActive**: Device is doing the exposure of a Frame (or Line).
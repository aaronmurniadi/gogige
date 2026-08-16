|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

## 5.6 Trigger Control features

The Trigger Control section describes all features related to image acquisition using trigger(s).

One or many Trigger(s) can be used to control the start of an Acquisition (Figure 5-1), of a Burst of Frames (Figure 5-2), of individual Frames (Figure 5-3) or of each Line of a Frame for a device (Figure 5-4). Triggers can also be used to control the exposure duration at the beginning of a frame.

TriggerSelector is used to select which type of trigger to configure. The standard trigger types are: AcquisitionStart, AcquisitionEnd, AcquisitionActive, FrameBurstStart, FrameBurstEnd, FrameBurstActive, FrameStart, FrameEnd, FrameActive, LineStart, ExposureStart, ExposureEnd and ExposureActive.

TriggerMode activate/deactivate trigger operation. It can be On or Off.

TriggerSource specifies the physical input Line or internal signal to use for the selected trigger. Standard trigger sources are: Software, Line0, Line1, ..., UserOutput0, UserOutput1, ..., Counter0Start, Counter0End, ..., Timer0Start, Timer0End, ..., Encoder0, Encoder1, ..., LogicBlock0, LogicBlock1, ..., Action0, Action1, ..., LinkTrigger0, LinkTrigger1, ...

With a Software trigger source, the TriggerSoftware command can be used by an application to generate an internal trigger signal.

With the hardware trigger sources, TriggerActivation specifies the activation mode of the trigger. This can be a RisingEdge, FallingEdge, AnyEdge, LevelHigh or LevelLow.

TriggerOverlap specifies the type of trigger overlap permitted with the previous frame/ line. This defines when a valid trigger will be accepted (or latched) for a new frame/line. This can be Off for no overlap, ReadOut to accept a trigger immediately after the exposure period or PreviousFrame/PreviousLine to accept (latch) a trigger that happened at any time after the start of the previous frame/Line. If a trigger is discarded based on the trigger overlap control and the current sensor state, it becomes a missed trigger. This represents an over-triggering situation and is typically considered as an error which can be connected to an Event.

TriggerDelay specifies the delay to apply after the trigger signal reception before effectively activating it.

TriggerDivider and TriggerMultiplier are used to control the ratio of triggers that are accepted.

Note that, a trigger is considered valid after the Dividers, Multipliers, Delay, ...

For example to setup a hardware triggered acquisition that will start the capture of each frame on the rising edge of the signal coming from the physical input Line 1, the following pseudo-code can be used:

Camera.TriggerSelector = FrameStart;
Camera.TriggerMode = On;
Camera.TriggerActivation = RisingEdge;
Camera.TriggerSource = Line1;
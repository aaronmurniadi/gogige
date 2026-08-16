- Counter0Start, Counter1Start, Counter2Start, ..., Counter0End, Counter1End, Counter2End, ...: Specifies which of the Counter signal to use as internal source for the trigger.
- Timer0Start, Timer1Start, Timer2Start, ..., Timer0End, Timer1End, Timer2End, ...: Specifies which Timer signal to use as internal source for the trigger.
- Encoder0, Encoder1, Encoder2, ...: Specifies which Encoder signal to use as internal source for the trigger.
- LogicBlock0, LogicBlock1, LogicBlock2, ...: Specifies which Logic Block signal to use as internal source for the trigger.
- Action0, Action1, Action2, ...: Specifies which Action command to use as internal source for the trigger.
- LinkTrigger0, LinkTrigger1, LinkTrigger2, ...: Specifies which Link Trigger to use as source for the trigger (received from the transport layer).
- CC1, CC2, CC3, CC4: Index of the Camera Link physical line and associated I/O control block to use. This ensures a direct mapping between the lines on the frame grabber and on the camera. Applicable to CameraLink products only.

### 5.6.5 TriggerActivation

|  Name | TriggerActivation[TriggerSelector]  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | RisingEdge FallingEdge AnyEdge LevelHigh LevelLow  |

Specifies the activation mode of the trigger.

Possible values are:

- RisingEdge: Specifies that the trigger is considered valid on the rising edge of the source signal.
- FallingEdge: Specifies that the trigger is considered valid on the falling edge of the source signal.
- AnyEdge: Specifies that the trigger is considered valid on the falling or rising edge of the source signal.
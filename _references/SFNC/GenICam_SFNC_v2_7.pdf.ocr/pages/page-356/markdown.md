## 20.19 TransferTriggerSource

|  Name | TransferTriggerSource[TransferTriggerSelector]  |
| --- | --- |
|  Category | TransferControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | Line0 (If 0 based), Line1, Line2, ... Counter0Start (If 0 based), Counter1Start, Counter2Start, ... Counter0End (If 0 based), Counter1End, Counter2End, ... Timer0Start (If 0 based), Timer1Start, Timer2Start, ... Timer0End (If 0 based), Timer1End, Timer2End, ... LogicBlock0, LogicBlock1, LogicBlock2, ... SoftwareSignal0 (If 0 based), SoftwareSignal1, SoftwareSignal2, ... Action0 (If 0 based), Action1, Action2, ... ...  |

Specifies the signal to use as the trigger source for transfers.

Possible values are:

- Line0 (If 0 based), Line1, Line2, ...: Specifies which physical line (or pin) and associated I/O control block to use as external source for the transfer control trigger signal.
- Counter0Start, Counter1Start, Counter2Start, ..., Counter0End, Counter1End, Counter2End, ...: Specifies which of the Counter signal to use as internal source for the transfer control trigger signal.
- Timer0Start, Timer1Start, Timer2Start, ..., Timer0End, Timer1End, Timer2End, ...: Specifies which Timer signal to use as internal source for the transfer control trigger signal.
- LogicBlock0, LogicBlock1, LogicBlock2, ...: Specifies which Logic Block to use as internal source for the transfer control trigger signal.
- SoftwareSignal0, SoftwareSignal1, SoftwareSignal2, ...: Specifies which Software Signal to use as internal source for the transfer control trigger signal.
- Action0, Action1, Action2, ...: Specifies which Action command to use as internal source for the transfer control trigger signal.
- ...
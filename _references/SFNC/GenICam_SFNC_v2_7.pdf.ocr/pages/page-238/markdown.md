|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- Counter0Active, Counter1Active, Counter2Active, ...: The chosen counter is in active state (counting).
- Timer0Active, Timer1Active, Timer2Active, ...: The chosen Timer is in active state.
- Encoder0, Encoder1, Encoder2, ...: The chosen Encoder Output state.
- LogicBlock0, LogicBlock1, LogicBlock2, ...: The choosen Logic Block output state.
- SoftwareSignal0, SoftwareSignal1, SoftwareSignal2, ...: The choosen Software Signal output state.
- Stream0TransferActive, Stream1TransferActive, ...: Transfer on the stream is active.
- Stream0TransferPaused, Stream1TransferPaused, ...: Transfer on the stream is paused.
- Stream0TransferStopping, Stream1TransferStopping, ...: Transfer on the stream is stopping.
- Stream0TransferStopped, Stream1TransferStopped, ...: Transfer on the stream is stopped.
- Stream0TransferOverflow, Stream1TransferOverflow, ...: Transfer on the stream is in overflow.

### 9.2.8 LineFormat

|  Name | LineFormat[LineSelector]  |
| --- | --- |
|  Category | DigitalIOControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | NoConnect TriState TTL LVDS RS422 OptoCoupled OpenDrain  |

Controls the current electrical format of the selected physical input or output Line.

Possible values are:

- NoConnect: The Line is not connected.
- TriState: The Line is currently in Tri-State mode (Not driven).
- TTL: The Line is currently accepting or sending TTL level signals.
- LVDS: The Line is currently accepting or sending LVDS level signals.
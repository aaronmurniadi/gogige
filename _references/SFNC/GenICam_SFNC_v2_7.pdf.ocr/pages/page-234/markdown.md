|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Interface | IEnumeration  |
| --- | --- |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Line0 (If 0 based), Line1, Line2, ... LinkTrigger0 (If 0 based), LinkTrigger1, LinkTrigger2, ... CC1, CC2, CC3, CC4, ...  |

Selects the physical line (or pin) of the external device connector or the virtual line of the Transport Layer to configure.

When a Line is selected, all the other Line features will be applied to its associated I/O control block and will condition the resulting input or output signal.

Possible values are:

- Line0 (If 0 based), Line1, Line2, ...: Index of the physical line and associated I/O control block to use.
- LinkTrigger0 (If 0 based), LinkTrigger1, LinkTrigger2, ...: Index of the virtual line going on the Transport layer to use.
- CC1, CC2, CC3, CC4: Index of the Camera Link physical line and associated I/O control block to use. This ensures a direct mapping between the lines on the frame grabber and on the camera. Applicable to CameraLink Product only.

### 9.2.3 LineMode

|  Name | LineMode[LineSelector]  |
| --- | --- |
|  Category | DigitalIOControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Input Output  |

Controls if the physical Line is used to Input or Output a signal.

When a Line supports input and output mode, the default state is Input to avoid possible electrical contention.

Possible values are:
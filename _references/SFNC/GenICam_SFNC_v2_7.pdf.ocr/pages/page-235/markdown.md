|  GEN<i>CAM |   | ![img-84.jpeg](img-84.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- Input: The selected physical line is used to Input an electrical signal.
- Output: The selected physical line is used to Output an electrical signal.

9.2.4 LineInverter

|  Name | LineInverter[LineSelector]  |
| --- | --- |
|  Category | DigitalIOControl  |
|  Level | Recommended  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | FalseTrue  |

Controls the inversion of the signal of the selected input or output Line.

Possible values are:

- False: The Line signal is not inverted.
- True: The Line signal is inverted.

9.2.5 LineStatus

|  Name | LineStatus[LineSelector]  |
| --- | --- |
|  Category | DigitalIOControl  |
|  Level | Recommended  |
|  Interface | IBoolean  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | FalseTrue  |

Returns the current status of the selected input or output Line.

The status of the signal is taken after the input Line inverter of the I/O control block.

Possible values are:
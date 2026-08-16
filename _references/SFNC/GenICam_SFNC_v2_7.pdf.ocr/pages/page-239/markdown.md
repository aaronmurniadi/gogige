|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- RS422: The Line is currently accepting or sending RS422 level signals.
- OptoCoupled: The Line is opto-coupled.
- OpenDrain: The Line is Open Drain (or Open Collector).

### 9.2.9 UserOutputSelector

|  Name | UserOutputSelector  |
| --- | --- |
|  Category | DigitalIOControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | UserOutput0, UserOutput1, UserOutput2, ...  |

Selects which bit of the User Output register will be set by UserOutputValue.

Possible values are:

- UserOutput0: Selects the bit 0 of the User Output register.
- UserOutput1: Selects the bit 1 of the User Output register.
- UserOutput2: Selects the bit 2 of the User Output register.
• ...

### 9.2.10 UserOutputValue

|  Name | UserOutputValue[UserOutputSelector]  |
| --- | --- |
|  Category | DigitalIOControl  |
|  Level | Recommended  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |
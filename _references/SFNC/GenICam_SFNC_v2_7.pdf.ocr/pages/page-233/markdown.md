|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Using LineSource, each of the bits of the User Output register can be directed to a physical output Line after going trough the I/O control block (See Figure 9-1: I/O Control).

UserOutputSelector and UserOutputValue are used to set any individual bit of the User Output register.

UserOutputValueAll and UserOutputValueAllMask can be used to set all or many of the User Output bits in one access.

Example:

/* User Output of a positive TTL signal on physical Line 2. */

LineSelector = Line2;
LineMode = Output;
LineFormat = TTL;
LineSource = UserOutput2;

UserOutputSelector = UserOutput2;
UserOutputValue = True;

## 9.2 Digital I/O Control features

This section lists the digital I/O control features.

### 9.2.1 DigitalIOControl

|  Name | DigitalIOControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains the digital input and output control features.

### 9.2.2 LineSelector

|  Name | LineSelector  |
| --- | --- |
|  Category | DigitalIOControl  |
|  Level | Recommended  |
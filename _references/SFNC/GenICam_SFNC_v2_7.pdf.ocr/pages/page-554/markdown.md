|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- DualBase: The camera streams the data from multiple taps (that do not fit in the standard base configuration) through two Camera Link base ports. It is responsibility of the application or frame grabber to reconstruct the full image. Only one of the ports (fixed) serves as the "master" for serial communication and triggering.
- EightyBit: Standard 80-bit configuration with 10 taps of 8 bits or 8 taps of 10 bits, as described by the Camera Link standard.
- Deca (Deprecated): This enumeration entry is deprecated. It was used for Deca configuration with 10 taps of 8-bit. Use EightyBit instead.

If the feature is omitted, one of the standard configurations (Base-Medium-Full) is expected. In that case the configuration can be unequivocally deduced from the SensorDigitizationTaps and PixelSize values.

### 27.6.3 ClTimeSlotsCount

|  Name | ClTimeSlotsCount  |
| --- | --- |
|  Category | CameraLink  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | One Two Three  |

This Camera Link specific feature describes the time multiplexing of the camera link connection to transfer more than the configuration allows, in one single clock.

Possible values are:

- One: One time slot.
- Two: Two time slots.
- One: Three time slots.

It indicates the number of consecutive time slots required to transfer one data of each tap.

### 27.7 CoaXPress Control

This section describes the CoaXPress specific control features.
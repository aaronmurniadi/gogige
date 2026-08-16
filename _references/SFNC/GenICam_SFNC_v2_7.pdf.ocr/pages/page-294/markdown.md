|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Values | ≥ 0  |
| --- | --- |

Indicates the size of the scheduled action commands queue. This number represents the maximum number of scheduled action commands that can be pending at a given point in time.

### 14.2.5 ActionSelector

|  Name | ActionSelector  |
| --- | --- |
|  Category | ActionControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ≥0  |

Selects to which Action Signal further Action settings apply.

### 14.2.6 ActionGroupMask

|  Name | ActionGroupMask[ActionSelector]  |
| --- | --- |
|  Category | ActionControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ≥0  |

Provides the mask that the device will use to validate the action on reception of the action protocol message.

The device asserts the selected Action signal only if:

- The selected ActionDeviceKey is equal to the action device key in the action protocol message.
- The logical AND-wise operation of the action group mask in the action protocol message against the selected ActionGroupMask is non-zero.
- The selected ActionGroupKey is equal to the action group key in the action protocol message.
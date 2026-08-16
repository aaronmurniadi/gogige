### 14.2.7 ActionGroupKey

|  Name | ActionGroupKey[ActionSelector]  |
| --- | --- |
|  Category | ActionControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ≥0  |

Provides the key that the device will use to validate the action on reception of the action protocol message.

The device asserts the selected Action signal only if:

- The selected **ActionDeviceKey** is equal to the action device key in the action protocol message.
- The logical AND-wise operation of the action group mask in the action protocol message against the selected **ActionGroupMask** is non-zero.
- The selected **ActionGroupKey** is equal to the action group key in the action protocol message.
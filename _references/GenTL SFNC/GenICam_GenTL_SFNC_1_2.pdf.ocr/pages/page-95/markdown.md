|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Interface | IEnumeration  |
| --- | --- |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | DeviceLost  |

Selects which Event to signal to the host application.

Possible values are:

- DeviceLost: Raised when the local host looses connection to the physical (remote) device.

### 3.3.5.3 EventNotification

|  Name | EventNotification[EventSelector]  |
| --- | --- |
|  Category | EventControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Off On Once  |

Activate or deactivate the notification to the host application of the occurrence of the selected Event.

Possible values are:

- Off: The selected Event notification is disabled.
- On: The selected Event notification is enabled.
- Once: The selected Event notification is enabled for one event then return to Off state.
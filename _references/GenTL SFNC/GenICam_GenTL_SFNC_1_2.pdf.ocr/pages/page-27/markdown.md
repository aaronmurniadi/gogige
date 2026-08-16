### 2.2.3 Action Control

Category that contains Action Control features.

Table 2-7: Action Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  ActionControl | R | GEV | ICategory | R | - | E | Category that contains all Action Control features of the Interface module.  |
|  ActionCommand | R | GEV | ICommand | (R)/W | - | E | Send ActionCommand to device(s).  |
|  ActionDeviceKey | R | GEV | IInteger | R/W | - | E | The Action Command Device Key to use in the Action Command.  |
|  ActionGroupKey | R | GEV | IInteger | R/W | - | E | The Action Command Group Key to use in the Action Command.  |
|  ActionGroupMask | R | GEV | IInteger | R/W | - | E | The Action Command Group Mask to use in the Action Command.  |
|  ActionScheduledTimeEnable | R | GEV | IBoolean | R/W | - | E | Specifies if a time enabled Action Command should be given.  |
|  ActionScheduledTime | R | GEV | IInteger | R/W | - | E | Specifies the time a time enabled Action Command is scheduled.  |
|  GevActionDestinationIPAddress | R | GEV | IInteger | R/W | - | E | Specifies destination the IP address for the Action Command.  |

### 2.2.4 GenICam Control

Contains the features related to GenICam control and access of a specific Interface module.
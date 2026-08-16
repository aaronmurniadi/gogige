|  **GEN<i>CAM** |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

## 2.12 Action Control

Contains the features related to the control of the Action command mechanism (See the Action Control chapter for details).

Table 2-12: Action Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  ActionControl | R | ICategory | R | - | G | Category that contains the Action control features.  |
|  ActionUnconditionalMode | O | IEnumeration | R/W | - | G | Enables the unconditional action command mode where action commands are processed even when the primary control channel is closed.  |
|  ActionDeviceKey | O | IInteger | W-O | - | G | Provides the device key that allows the device to check the validity of action commands.  |
|  ActionQueueSize | O | IInteger | R | - | G | Indicates the size of the scheduled action commands queue.  |
|  ActionSelector | O | IInteger | R/W | - | G | Selects to which Action Signal further Action settings apply.  |
|  ActionGroupMask[ActionSelector] | O | IInteger | R/W | - | G | Provides the mask that the device will use to validate the action on reception of the action protocol message.  |
|  ActionGroupKey[ActionSelector] | O | IInteger | R/W | - | G | Provides the key that the device will use to validate the action on reception of the action protocol message.  |

## 2.13 Event Control

Contains the features related to the generation of Event notifications by the device (See the Event Control chapter for details).

Table 2-13: Event Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  EventControl | R | ICategory | R | - | E | Category that contains Event control features.  |
|  EventSelector | R | IEnumeration | R/W | - | E | Selects which Event to signal to the host application.  |
|  EventNotification[EventSelector] | R | IEnumeration | R/W | - | E | Activate or deactivate the notification to the host application of the occurrence of the selected Event.  |
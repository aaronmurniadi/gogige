|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Activate or deactivate the notification to the host application of the occurrence of the selected Event.

Possible values are:

- Off: The selected Event notification is disabled.
- On: The selected Event notification is enabled.
- Once: The selected Event notification is enabled for one event then return to Off state.

### 15.4 Frame Trigger Event (Example #1)

Below are the recommended features for the Frame Trigger Event handling.

#### 15.4.1 EventFrameTriggerData

|  Name | EventFrameTriggerData  |
| --- | --- |
|  Category | EventControl  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains all the data features related to the FrameTrigger Event.

#### 15.4.2 EventFrameTrigger

|  Name | EventFrameTrigger  |
| --- | --- |
|  Category | EventFrameTriggerData  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |
|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

The GenICam port through which the Data Stream module is accessed.

Note that StreamPort is a port node (not a feature node) and is generally not accessed by the end user directly.

### 3.4.5 Event Control

Controls the generation of events for an instance of the buffer module. An Event is a message that is sent to the host application to notify it of the occurrence of an internal event.

See GenICam SFNC for more details on event control.

EventSelector selects which particular Event to control

#### 3.4.5.1 EventControl

|  Name | EventControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains Event control features.

#### 3.4.5.2 EventSelector

|  Name | EventSelector  |
| --- | --- |
|  Category | EventControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | NewBufferData  |
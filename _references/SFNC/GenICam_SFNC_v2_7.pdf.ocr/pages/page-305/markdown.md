|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 15.5.1 EventExposureEndData

|  Name | EventExposureEndData  |
| --- | --- |
|  Category | EventControl  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains all the data features related to the ExposureEnd Event.

### 15.5.2 EventExposureEnd

|  Name | EventExposureEnd  |
| --- | --- |
|  Category | EventExposureEndData  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Returns the unique identifier of the ExposureEnd type of Event. This feature can be used to register a callback function to be notified of the event occurrence. Its value uniquely identifies the type of event that will be received.

### 15.5.3 EventExposureEndTimestamp

|  Name | EventExposureEndTimestamp  |
| --- | --- |
|  Category | EventExposureEndData  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
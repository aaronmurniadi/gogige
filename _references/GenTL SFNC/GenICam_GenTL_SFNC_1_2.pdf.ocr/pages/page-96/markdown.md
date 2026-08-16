|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.4 Data Stream Module

Contains all features of the Data Stream module that are independent from the underlying transport technology.

#### 3.4.1 Stream Information

Features in this section provide basic information about the Data Stream module and its identity.

##### 3.4.1.1 Stream Information

|  Name | StreamInformation  |
| --- | --- |
|  Category | Root  |
|  Level | Mandatory  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains all Stream Information features of the Data Stream module.

##### 3.4.1.2 StreamID

|  Name | StreamID  |
| --- | --- |
|  Category | StreamInformation  |
|  Level | Mandatory  |
|  Interface | ISString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

Device unique ID for the data stream.

Corresponds to the STREAM_INFO_ID command of DSGetInfo function.
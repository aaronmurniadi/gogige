|  GEN<i>CAM |   | ![img-27.jpeg](img-27.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Level | Mandatory  |
| --- | --- |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Device unique ID for the stream. Not Mandator for non-streaming DeviceCorresponds to the DevGetDataStreamID function with the index corresponding to "StreamSelector".

#### 3.3.4 GenICam Control

This chapter provides the necessary features to use the GenICam feature tree of the Device module.

Note: In case of discrepancy between the features described in this chapter and the “GenICam Standard text” the GenTL SFNC document prevails.

##### 3.3.4.1 Root

|  Name | Root  |
| --- | --- |
|  Category | None  |
|  Level | Mandatory  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Provides the Root of the GenICam features tree.

##### 3.3.4.2 DevicePort

|  Name | DevicePort  |
| --- | --- |
|  Category | None  |
|  Level | Mandatory  |
|  GEN<i>CAM |   | ![img-26.jpeg](img-26.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

#### 3.3.3 Stream Enumeration

The Stream Enumeration section describes all features related to the enumeration of data streams belonging to the Device module.

##### 3.3.3.1 StreamEnumeration

|  Name | StreamEnumeration  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category that contains all Stream Enumeration features of the Device module.

##### 3.3.3.2 StreamSelector

|  Name | StreamSelector  |
| --- | --- |
|  Category | StreamEnumeration  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥0  |

Selector for the different stream channels. The selector is 0-based in order to match the index of the C interface.

##### 3.3.3.3 StreamID

|  Name | StreamID[StreamSelector]  |
| --- | --- |
|  Category | StreamEnumeration  |
|  GEN<i>CAM |   | ![img-26.jpeg](img-26.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Interface | IEnumeration  |
| --- | --- |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Stream0Stream1Stream2...  |

Control the destination of the selected region.

Possible values are:

- Stream0: The destination of the region is the data stream 0.
- Stream1: The destination of the region is the data stream 1.
- Stream2: The destination of the region is the data stream 2.
• ...

4.10RegionIDValue

|  Name | RegionIDValue[RegionSelector]  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Returns a unique Identifier value that corresponds to the selected Region.

This value is typically used by the Transport Layer to specify the Region from which the transmitted data come from.

4.11 ComponentSelector

|  Name | ComponentSelector  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
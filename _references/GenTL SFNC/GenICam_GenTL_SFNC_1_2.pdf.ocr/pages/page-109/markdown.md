|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

If a buffer is passed to DSAnnounceBuffer which is not aligned according to the alignment size it is up to the Producer to either reject the buffer and return a GC_ERR_INVALID_BUFFER error code or to cope with a potential overhead and use the unaligned buffer as is.

Corresponds to the STREAM_INFO_BUF_ALIGNMENT command of DSGetInfo function.

#### 3.4.4 GenICam Control

This chapter provides the necessary features to use the GenICam feature tree of the Device module.

Note: In case of discrepancy between the features described in this chapter and the “GenICam Standard text” the GenTL SFNC document prevails.

3.4.4.1 Root

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

3.4.4.2 StreamPort

|  Name | StreamPort  |
| --- | --- |
|  Category | None  |
|  Level | Mandatory  |
|  Interface | IPort  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | -  |
|  ![img-17.jpeg](img-17.jpeg) |   | ![img-18.jpeg](img-18.jpeg)  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Values | -  |
| --- | --- |

Resets the current value of the device timestamp counter.

After executing this command, the timestamp counter restarts automatically.

3.68 TimestampLatch

|  Name | TimestampLatch  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Latches the current timestamp counter into TimestampLatchValue.

3.69 TimestampLatchValue

|  Name | TimestampLatchValue  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | ns  |
|  Visibility | Expert  |
|  Values | ≥0  |

Returns the latched value of the timestamp counter.

Note that the increment of the TimestampLatchValue feature must correspond to the resolution of the devices's timestamp in nanoseconds.
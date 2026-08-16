|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Access | Read  |
| --- | --- |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | -  |

Indicates the first URL to the GenICam XML device description file. The First URL is used as the first choice by the application to retrieve the GenICam XML device description file.

#### 27.4.25 GevSecondURL(Deprecated)

|  Name | GevSecondURL  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  sequInterface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | -  |

Indicates the second URL to the GenICam XML device description file. This URL is an alternative if the application was unsuccessful to retrieve the device description file using the first URL.

#### 27.4.26 GevNumberOfInterfaces (Deprecated)

|  Name | GevNumberOfInterfaces  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | >0  |

This feature is deprecated (See DeviceLinkSelector). It was representing the number of logical links supported by this device.
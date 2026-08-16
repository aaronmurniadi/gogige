|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 27.4.8 GevCurrentPhysicalLinkConfiguration

|  Name | GevCurrentPhysicalLinkConfiguration  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | SingleLink MultiLink StaticLAG DynamicLAG  |

Indicates the current physical link configuration of the device.

Note: When multi-link and LAG configurations are used concurrently, the device shall report the LAG configuration.

### 27.4.9 GevActiveLinkCount (Deprecated)

|  Name | GevActiveLinkCount  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | > 0  |

Indicates the current number of active logical links.

### 27.4.10 GevSupportedOptionSelector

|  Name | GevSupportedOptionSelector  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
Note that if multiple Regions of interest are supported by the device, the RegionSelector can be added to various features such as Width, Height,... to specify the characteristics of the selected Region. In order to simplify the standard text and feature descriptions (see above example), the optional RegionSelector is not explicitly propagated to all the features of the SFNC that it can potentially select.

Possible values are:

- Region0: Selected feature will control the region 0.
- Region1: Selected feature will control the region 1.
- Region2: Selected feature will control the region 2.
- ...
- Scan3dExtraction0: Selected feature will control the Scan3dExtraction0 output Region.
- Scan3dExtraction1: Selected feature will control the Scan3dExtraction1 output Region.
- Scan3dExtraction2: Selected feature will control the Scan3dExtraction2 output Region.
- ...
- All: Selected features will control all the regions at the same time.

### 4.8 RegionMode

|  Name | RegionMode[RegionSelector]  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  sUnit | -  |
|  Visibility | Beginner  |
|  Values | Off On  |

Controls if the selected Region of interest is active and streaming.

Possible values are:

- Off: Disable the usage of the Region.
- On: Enable the usage of the Region.

### 4.9 RegionDestination

|  Name | RegionDestination[RegionSelector]  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
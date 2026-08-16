|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- Scan3dExtraction2: Selects Scan3d Extraction module 2.
• ...

21.4.3 Scan3dExtractionSource

|  Name | Scan3dExtractionSource[Scan3dExtractionSelector]  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Region0, Region1, Region2, ...  |

Selects the sensor's data source region for 3D Extraction module.

Typically these are Sensor Regions, but it could also be another ProcessingModule.

Possible values are:

- Region0: Data come from Sensor's Region0.
- Region1: Data come from Sensor's Region1.
- Region2: Data come from Sensor's Region2.
• ...

21.4.4 Scan3dExtractionMethod

|  Name | Scan3dExtractionMethod[Scan3dExtractionSelector]  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
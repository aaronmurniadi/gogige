### 21.4 Scan 3D Features

This section defines the Scan 3D Control features.

#### 21.4.1 Scan3dControl

|  Name | Scan3dControl  |
| --- | --- |
|  Category | Root  |
|  Level | Optional  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category for control of 3D camera specific features.

#### 21.4.2 Scan3dExtractionSelector

|  Name | Scan3dExtractionSelector  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Scan3dExtraction0, Scan3dExtraction1, Scan3dExtraction2, ...  |

Selects the 3D Extraction processing module to control (if multiple ones are present).

Possible values are:

- Scan3dExtraction0: Selects Scan3d Extraction module 0.
- Scan3dExtraction1: Selects Scan3d Extraction module 1.
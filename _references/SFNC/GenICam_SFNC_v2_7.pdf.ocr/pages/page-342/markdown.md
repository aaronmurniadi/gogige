### 19.3 SourceControl

|  Name | SourceControl  |
| --- | --- |
|  Category | Root  |
|  Level | Optional  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category that contains the source control features.

### 19.4 SourceCount

|  Name | SourceCount  |
| --- | --- |
|  Category | SourceControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥ 1  |

Controls or returns the number of sources supported by the device.

This feature is generally read-only but can be writable if the number of sources supported by the device is run-time programmable.

### 19.5 SourceSelector

|  Name | SourceSelector  |
| --- | --- |
|  Category | SourceControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
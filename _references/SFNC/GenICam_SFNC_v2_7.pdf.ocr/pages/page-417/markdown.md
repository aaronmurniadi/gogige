|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 22.3 Light Control Features

#### 22.3.1 LightControl

|  Name | LightControl  |
| --- | --- |
|  Category | Root  |
|  Level | Optional  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category containing the Lighting control features.

#### 22.3.2 LightControllerSelector

|  Name | LightControllerSelector  |
| --- | --- |
|  Category | LightControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | LightController0 LightController1 LightController2 ...  |

Selects the Light Controller to configure.

#### 22.3.3 LightControllerSource

|  Name | LightControllerSource[LightControllerSelector]  |
| --- | --- |
|  Category | LightControl  |
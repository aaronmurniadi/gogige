|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Visibility | Beginner  |
| --- | --- |
|  Values | ≥0  |

Set the voltage rating of the lighting output.

### 22.3.6 LightBrightness

|  Name | LightBrightness[LightControllerSelector]  |
| --- | --- |
|  Category | LightControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | %  |
|  Visibility | Beginner  |
|  Values | ≥0  |

Set the brightness of the lighting output in percent. Can be greater than 100% for short overdrive period.

### 22.3.7 LightConnectionStatus

|  Name | LightConnectionStatus[LightControllerSelector]  |
| --- | --- |
|  Category | LightControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Sensing Ready NoConnect Error  |

Status of a light connected to the controller's output Line.

### 22.3.8 LightCurrentMeasured

|  Name | LightCurrentMeasured[LightControllerSelector]  |
| --- | --- |
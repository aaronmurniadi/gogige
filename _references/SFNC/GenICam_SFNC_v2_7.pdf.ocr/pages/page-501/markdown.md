|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

## 26 GenICam Control

This chapter provides the necessary features to use the GenICam feature tree.

Note: In case of discrepancy between the features described in this chapter and the "GenICam Standard text" the SFNC document prevail.

### 26.1 GenICam feature tree access

The mandatory features below are necessary to access the GenICam features tree.

#### 26.1.1 Root

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

#### 26.1.2 Device

|  Name | Device  |
| --- | --- |
|  Category | None  |
|  Level | Mandatory  |
|  Interface | IPort  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | -  |

Provides the default GenICam port of the Device.
|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3 General Features

Contains all features that are independent from the underlying transport technology, in particular including all mandatory features for all GenTL Producer implementations.

### 3.1 System Module

Contains all features of the System module that are independent from the underlying transport technology.

#### 3.1.1 System Information

Features in this section provide basic information about the System Module and its identity. Note that all features in this section are defined as read-only.

3.1.1.1 SystemInformation

|  Name | SystemInformation  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category that contains all System Information features of the System module.

3.1.1.2 TLID

|  Name | TLID  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Mandatory  |
|  Interface | ISString  |
|  Access | Read  |
|  Unit | -  |
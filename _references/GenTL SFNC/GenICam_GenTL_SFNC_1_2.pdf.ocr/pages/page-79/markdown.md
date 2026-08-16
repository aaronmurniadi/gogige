|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

- Once: The selected Event notification is enabled for one event then return to Off state

### 3.3 Device Module

Contains all features of the Device module that are independent from the underlying transport technology. Do not mistake the features of the Device module with the features of the remote device.

### 3.3.1 Device Information

Features in this section provide basic information about the Device module and its identity. Note that all features in this section are defined read-only.

#### 3.3.1.1 Device Information

|  Name | DeviceInformation  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category that contains all Device Information features of the Device module.

#### 3.3.1.2 DeviceID

|  Name | DeviceID  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Mandatory  |
|  Interface | ISString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
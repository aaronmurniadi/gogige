## 20.2 Transfer Control features

This section describes in detail the features related to the Transfer Control.

A detailed example of the usage of the Transfer Control features is presented in the section 19.1 "Source Control usage model with Multiple Regions and Transfers".

### 20.3 TransferControl

|  Name | TransferControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category for the data Transfer Control features.

### 20.4 TransferSelector

|  Name | TransferSelector  |
| --- | --- |
|  Category | TransferControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Stream0 (if 0 based), Stream1, Stream2, ... All  |

Selects which stream transfers are currently controlled by the selected Transfer features.

Possible values are:
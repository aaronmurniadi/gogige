Version 2.7.1

Standard Features Naming Convention

|  TransferResume[TransferSelector] | O | ICommand | (R)/W | - | G | Resumes a data Blocks streaming that was previously paused by a TransferPause command.  |
| --- | --- | --- | --- | --- | --- | --- |
|  TransferTriggerSelector[TransferSelector] | O | IEnumeration | R/W | - | G | Selects the type of transfer trigger to configure.  |
|  TransferTriggerMode[TransferSelector][TransferTriggerSelector] | R | IEnumeration | R/W | - | G | Controls if the selected trigger is active.  |
|  TransferTriggerSource[TransferTriggerSelector] | O | IEnumeration | R/W | - | G | Specifies the signal to use as the trigger source for transfers.  |
|  TransferTriggerActivation[TransferTriggerSelector] | O | IEnumeration | R/W | - | G | Specifies the activation mode of the transfer control trigger.  |
|  TransferStatusSelector[TransferSelector] | R | IEnumeration | R/W | - | G | Selects which status of the transfer module to read.  |
|  TransferStatus[TransferStatusSelector] | R | IBool | R | - | G | Reads the status of the Transfer module signal selected by TransferStatusSelector.  |
|  TransferComponentSelector[TransferSelector] | O | IEnumeration | R/W | - | G | Selects the color component for the control of the TransferStreamChannel feature.  |
|  TransferStreamChannel[TransferSelector][TransferComponentSelector] | O | IInteger | R/W | - | G | Selects the streaming channel that will be used to transfer the selected stream of data.  |

### 2.19Scan 3D Control

Contains the features related to the control of the 3D scan features (See the 3D Scan Control chapter for details).

Table 2-19: 3D Scan Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  Scan3dControl | O | ICategory | R | - | B | Category for control of 3D camera specific features.  |
|  Scan3dExtractionSelector | O | IEnumeration | R/W | - | E | Selects the 3D Extraction processing module to control (if multiple ones are present).  |
|  Scan3dExtractionSource[Scan3dExtractionSelector] | O | IEnumeration | R/W | - | E | Selects the sensor's data source region for 3D Extraction module.  |
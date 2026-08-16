- Stream0: The transfer features control the data stream 0.
- Stream1: The transfer features control the data stream 1.
- Stream2: The transfer features control the data stream 2.
- ...
- All: The transfer features control all the data streams simultaneously.

## 20.5 TransferControlMode

|  Name | TransferControlMode[TransferSelector]  |
| --- | --- |
|  Category | TransferControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Basic Automatic UserControlled  |

Selects the control method for the transfers.

Possible values are:

- Basic: Transfer flow control mechanism is disabled. The TransferStart, TransferPause, TransferResume, TransferStop and TransferAbort features are not available to the user. This transfer mode is used to ensure compatibility with devices not aware of the transfer flow control mechanism.
- Automatic: Transfer flow control mechanism is controlled automatically. The Transfer features are controlled transparently by the acquisition features.
  - TransferStart is called during the AcquisitionStart.
  - TransferStop is never called.
  - TransferAbort is called during the AcquisitionAbort.
  - TransferOperationMode is read only and set to "Continuous".

If available, the TransferPause and TransferResume features are controlled by the user.

- UserControlled: Transfer flow control mechanism is controlled by the user. The TransferMode, TransferStart, TransferStop, TransferAbort, TransferPause and TransferResume features are used
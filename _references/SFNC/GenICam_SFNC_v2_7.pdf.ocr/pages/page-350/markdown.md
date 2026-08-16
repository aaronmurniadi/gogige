to control manually the flow of data. In this mode, the features TransferOperationMode, TransferStart and TransferStop must be available.

If this feature is not present, the transfer control is assumed to be "Basic".

Note that the Transfers can also be controlled using external trigger signals (See TransferTriggerSelector).

### 20.6 TransferOperationMode

|  Name | TransferOperationMode[TransferSelector]  |
| --- | --- |
|  Category | TransferControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Continuous MultiBlock  |

Selects the operation mode of the transfer.

Possible values are:

- Continuous: Blocks of data are transferred continuously until stopped with the TransferStop command.
- MultiBlock: The transfer of the blocks of data terminates automatically after the transmission of TransferBlockCount or when an explicit TransferStop command is received.

If this feature is not present, the transfer mode is assumed to be "Continuous".

### 20.7 TransferBlockCount

|  Name | TransferBlockCount[TransferSelector]  |
| --- | --- |
|  Category | TransferControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | > 0  |
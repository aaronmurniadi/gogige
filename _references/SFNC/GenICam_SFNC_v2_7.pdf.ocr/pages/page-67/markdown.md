|  SourceSelector | O | IEnumeration | R/W | - | B | Selects the source to control.  |
| --- | --- | --- | --- | --- | --- | --- |
|  SourceIDValue[SourceSelector] | O | IInteger | R | - | E | Returns a unique Identifier value that correspond to the selected Source.  |

## 2.18 Transfer Control

Contains the features related to the control of the Transfers (See the Transfer Control chapter for details).

Table 2-18: Transfer Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  TransferControl | R | ICategory | R | - | E | Category for the data Transfer Control features.  |
|  TransferSelector | O | IEnumeration | R/(W) | - | E | Selects which stream transfers are currently controlled by the selected Transfer features.  |
|  TransferControlMode[TransferSelector] | R | IEnumeration | R/(W) | - | E | Selects the control method for the transfers.  |
|  TransferOperationMode[TransferSelector] | O | IEnumeration | R/(W) | - | E | Selects the operation mode of the transfer.  |
|  TransferBlockCount[TransferSelector] | O | IInteger | R/(W) | - | E | Specifies the number of data Blocks that the device should stream before stopping.  |
|  TransferBurstCount | O | IInteger | R/W | - | E | Number of Block(s) to transfer for each TransferBurstStart trigger.  |
|  TransferQueueMaxBlockCount[TransferSelector] | O | IInteger | R/(W) | - | E | Controls the maximum number of data blocks that can be stored in the block queue of the selected stream.  |
|  TransferQueueCurrentBlockCount[TransferSelector] | O | IInteger | R | - | E | Returns the number of Block(s) currently in the transfer queue.  |
|  TransferQueueMode[TransferSelector] | O | IEnumeration | R/(W) | - | E | Specifies the operation mode of the transfer queue.  |
|  TransferStart[TransferSelector] | O | ICommand | (R)/W | - | E | Starts the streaming of data blocks out of the device.  |
|  TransferStop[TransferSelector] | O | ICommand | (R)/W | - | E | Stops the streaming of data Block(s).  |
|  TransferAbort[TransferSelector] | O | ICommand | (R)/W | - | E | Aborts immediately the streaming of data block(s).  |
|  TransferPause[TransferSelector] | O | ICommand | (R)/W | - | G | Pauses the streaming of data Block(s).  |
Starts the streaming of data blocks out of the device. This feature must be available when the TransferControlMode is set to "UserControlled". If the TransferStart feature is not writable (locked), the application should not start the transfer and should avoid using the feature until it becomes writable again.

### 20.13TransferStop

|  Name | TransferStop[TransferSelector]  |
| --- | --- |
|  Category | TransferControl  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Stops the streaming of data Block(s). The current block transmission will be completed. This feature must be available when the TransferControlMode is set to "UserControlled".

### 20.14TransferAbort

|  Name | TransferAbort[TransferSelector]  |
| --- | --- |
|  Category | TransferControl  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Aborts immediately the streaming of data block(s). Aborting the transfer will result in the lost of the data that is present or currently entering in the block queue. However, the next new block received will be stored in the queue and transferred to the host when the streaming is restarted. If implemented, this feature should be available when the TransferControlMode is set to "UserControlled".

### 20.15TransferPause

|  Name | TransferPause[TransferSelector]  |
| --- | --- |
|  Category | TransferControl  |
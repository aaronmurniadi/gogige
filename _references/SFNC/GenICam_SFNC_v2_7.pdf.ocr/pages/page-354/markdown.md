|  Level | Optional  |
| --- | --- |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | -  |

Pauses the streaming of data Block(s). Pausing the streaming will immediately suspend the ongoing data transfer even if a block is partially transferred. The device will resume its transmission at the reception of a TransferResume command.

### 20.16 TransferResume

|  Name | TransferResume[TransferSelector]  |
| --- | --- |
|  Category | TransferControl  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | -  |

Resumes a data Blocks streaming that was previously paused by a TransferPause command.

### 20.17 TransferTriggerSelector

|  Name | TransferTriggerSelector[TransferSelector]  |
| --- | --- |
|  Category | TransferControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | TransferStart TransferStop TransferAbort  |
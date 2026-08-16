Stopped
QueueOverflow

Selects which status of the transfer module to read.

Possible values are:

- Streaming: Data blocks are transmitted when enough data is available.
- Stopping: Data blocks transmission is stopping. The current block transmission will be completed and the transfer state will stop.
- Stopped: Data blocks transmission is stopped.
- Paused: Data blocks transmission is suspended immediately.
- QueueOverflow: Data blocks queue is in overflow state.

### 20.22 TransferStatus

|  Name | TransferStatus[TransferStatusSelector]  |
| --- | --- |
|  Category | TransferControl  |
|  Level | Recommended  |
|  Interface | IBool  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | True False  |

Reads the status of the Transfer module signal selected by TransferStatusSelector.

### 20.23 TransferComponentSelector

|  Name | TransferComponentSelector[TransferSelector]  |
| --- | --- |
|  Category | TransferControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
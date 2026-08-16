Specifies the number of data Blocks that the device should stream before stopping. This feature is only active if the TransferOperationMode is set to MultiBlock.

### 20.8 TransferBurstCount

|  Name | TransferBurstCount  |
| --- | --- |
|  Category | TransferControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥1  |

Number of Block(s) to transfer for each TransferBurstStart trigger.

This feature is used only if the TransferBurstStart trigger is enabled and the TransferBurstEnd trigger is disabled.

### 20.9 TransferQueueMaxBlockCount

|  Name | TransferQueueMaxBlockCount[TransferSelector]  |
| --- | --- |
|  Category | TransferControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | > 0  |

Controls the maximum number of data blocks that can be stored in the block queue of the selected stream.

### 20.10 TransferQueueCurrentBlockCount

|  Name | TransferQueueCurrentBlockCount[TransferSelector]  |
| --- | --- |
|  Category | TransferControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|   | TransferPause TransferResume TransferActive TransferBurstStart TransferBurstStop  |
| --- | --- |

Selects the type of transfer trigger to configure.

Possible values are:

- **TransferStart**: Selects a trigger to start the transfers.
- **TransferStop**: Selects a trigger to stop the transfers.
- **TransferAbort**: Selects a trigger to abort the transfers.
- **TransferPause**: Selects a trigger to pause the transfers.
- **TransferResume**: Selects a trigger to Resume the transfers.
- **TransferActive**: Selects a trigger to Activate the transfers. This trigger type is used when TriggerActivation is set LevelHigh or levelLow.
- **TransferBurstStart**: Selects a trigger to start the transfer of a burst of frames specified by TransferBurstCount.
- **TransferBurstStop**: Selects a trigger to end the transfer of a burst of frames.

### 20.18 TransferTriggerMode

|  Name | TransferTriggerMode[TransferSelector][TransferTriggerSelector]  |
| --- | --- |
|  Category | TransferControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | Off On  |

Controls if the selected trigger is active.

Possible values are:

- **Off**: Disables the selected trigger.
- **On**: Enable the selected trigger.
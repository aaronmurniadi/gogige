## 20.20 TransferTriggerActivation

|  **Name** | TransferTriggerActivation[TransferTriggerSelector]  |
| --- | --- |
|  **Category** | TransferControl  |
|  **Level** | Optional  |
|  **Interface** | IEnumeration  |
|  **Access** | Read/Write  |
|  **Unit** | -  |
|  **Visibility** | Guru  |
|  **Values** | RisingEdge FallingEdge AnyEdge LevelHigh LevelLow  |

Specifies the activation mode of the transfer control trigger.

Possible values are:

- **RisingEdge:** Specifies that the trigger is considered valid on the rising edge of the source signal.
- **FallingEdge:** Specifies that the trigger is considered valid on the falling edge of the source signal.
- **AnyEdge:** Specifies that the trigger is considered valid on the falling or rising edge of the source signal.
- **LevelHigh:** Specifies that the trigger is considered valid as long as the level of the source signal is high. This can apply to TransferActive and TransferPause trigger.
- **LevelLow:** Specifies that the trigger is considered valid as long as the level of the source signal is low. This can apply to TransferActive and TransferPause trigger.

## 20.21 TransferStatusSelector

|  **Name** | TransferStatusSelector[TransferSelector]  |
| --- | --- |
|  **Category** | TransferControl  |
|  **Level** | Recommended  |
|  **Interface** | IEnumeration  |
|  **Access** | Read/Write  |
|  **Unit** | -  |
|  **Visibility** | Guru  |
|  **Values** | Streaming Paused Stopping  |
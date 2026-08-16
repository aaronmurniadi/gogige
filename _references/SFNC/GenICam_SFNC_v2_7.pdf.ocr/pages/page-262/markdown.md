|  **Visibility** | Expert  |
| --- | --- |
|  **Values** | RisingEdge FallingEdge AnyEdge LevelHigh LevelLow  |

Selects the activation mode of the trigger to start the Timer.

Possible values are:

- **RisingEdge:** Starts counting on the Rising Edge of the selected trigger signal.
- **FallingEdge:** Starts counting on the Falling Edge of the selected trigger signal.
- **AnyEdge:** Starts counting on the Falling or Rising Edge of the selected trigger signal.
- **LevelHigh:** Counts as long as the selected trigger signal level is High.
- **LevelLow:** Counts as long as the selected trigger signal level is Low.

### 10.5.9 TimerTriggerArmDelay

|  **Name** | TimerTriggerArmDelay[TimerSelector]  |
| --- | --- |
|  **Category** | CounterAndTimerControl  |
|  **Level** | Recommended  |
|  **Interface** | IFloat  |
|  **Access** | Read/Write  |
|  **Unit** | us  |
|  **Visibility** | Expert  |
|  **Values** | ≥0  |

Sets the minimum period between two valid timer triggers.

Note: This feature is independent of the values of TimerDuration and TimerDelay.
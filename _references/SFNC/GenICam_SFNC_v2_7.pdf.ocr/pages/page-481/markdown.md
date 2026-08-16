### 24.46 ChunkBlackLevelSelector

|  Name | ChunkBlackLevelSelector  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | All Red Green Blue Y U V Tap1, Tap2, ...  |

Selects which Black Level to return. Possible values are:

- All: Black Level will be applied to all channels or taps.
- Red: Black Level will be applied to the red channel.
- Green: Black Level will be applied to the green channel.
- Blue: Black Level will be applied to the blue channel.
- Y: Black Level will be applied to Y channel.
- U: Black Level will be applied to U channel.
- V: Black Level will be applied to V channel.
- Tap1: Black Level will be applied to Tap 1.
- Tap2: Black Level will be applied to Tap 2.
- ...

### 24.47 ChunkBlackLevel

|  Name | ChunkBlackLevel[ChunkBlackLevelSelector]  |
| --- | --- |
|  Category | ChunkDataControl  |
|  Level | Recommended  |
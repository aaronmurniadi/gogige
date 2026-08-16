- All: Balance Ratio will be applied to all channels or taps.
- Red: Balance Ratio will be applied to the red channel.
- Green: Balance Ratio will be applied to the green channel.
- Blue: Balance Ratio will be applied to the blue channel.
- Y: Balance Ratio will be applied to Y channel.
- U: Balance Ratio will be applied to U channel.
- V: Balance Ratio will be applied to V channel.
- Tap1: Balance Ratio will be applied to Tap 1.
- Tap2: Balance Ratio will be applied to Tap 2.
- ...

### 6.13 BalanceRatio

|  Name | BalanceRatio[BalanceRatioSelector]  |
| --- | --- |
|  Category | AnalogControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | >0.0  |

Controls ratio of the selected color component to a reference color component. It is used for white balancing.

For example, the Color balance is realized by the following formula:

$$C_w = \text{BalanceRatio} \times C$$

where

Cw is the intensity of selected color component after white balancing.

BalanceRatio is the white balance coefficient.

C is the intensity of the color component before white balancing.

### 6.14 BalanceWhiteAuto

|  Name | BalanceWhiteAuto  |
| --- | --- |
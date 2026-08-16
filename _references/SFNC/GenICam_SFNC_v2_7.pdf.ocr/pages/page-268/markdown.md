### 11.2 Encoder features

This section describes the quadrature encoder features.

### 11.2.1 EncoderControl

|  Name | EncoderControl  |
| --- | --- |
|  Category | Root  |
|  Level | Optional  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category that contains the quadrature Encoder Control features.

### 11.2.2 EncoderSelector

|  Name | EncoderSelector  |
| --- | --- |
|  Category | EncoderControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Encoder0 (If 0 based), Encoder1, Encoder2, ...  |

Selects which Encoder to configure.

Possible values are:

- Encoder0: Selects Encoder 0.
- Encoder1: Selects Encoder 1.
- Encoder2: Selects Encoder 2.
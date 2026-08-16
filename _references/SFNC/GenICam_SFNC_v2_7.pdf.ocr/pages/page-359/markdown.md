|  Values | Red Green Blue All  |
| --- | --- |

Selects the color component for the control of the TransferStreamChannel feature.

Possible values are:

- Red: The TransferStreamChannel feature controls the index of the stream channel for the streaming of the red plane of the planar pixel formats.
- Green: The TransferStreamChannel feature controls the index of the stream channel for the streaming of the green plane of the planar pixel formats.
- Blue: The TransferStreamChannel feature controls the index of the stream channel for the streaming of blue plane of the planar pixel formats.
- All: The TransferStreamChannel feature controls the index of the stream channel for the streaming of all the planes of the planar pixel formats simultaneously or non planar pixel formats.

This feature is only needed if the device supports planar pixel formats.

### 20.24 TransferStreamChannel

|  Name | TransferStreamChannel[TransferSelector][TransferComponentSelector]  |
| --- | --- |
|  Category | TransferControl  |
|  Level | Optional  |
|  Interface | Integer  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | >0  |

Selects the streaming channel that will be used to transfer the selected stream of data. In general, this feature can be omitted and the default streaming channel will be used.
|  Values | Default Device-specific  |
| --- | --- |

Selects the method for extracting 3D from the input sensor data.

Typically device specific algorithms are used, and they can be either Line3D extracting (Linescan3D device) or Area3D extracting (Areascan3D device).

Possible values are:

- Default: Default range extraction method for the device.
- ...

Device specific values can be used to indentify the various range extraction methods supported by the device.

### 21.4.5 Scan3dDistanceUnit

|  Name | Scan3dDistanceUnit[Scan3dExtractionSelector]  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Millimeter Meter Inch Pixel Device-specific  |

Specifies the unit used when delivering (calibrated) distance data.

Possible values are:

- Millimeter: Distance values are in millimeter units (default).
- Meter: Distance values are in meter units.
- Inch: Distance values are in inch units.
- Pixel: Distance values are given as a multiple of the size of a pixel.